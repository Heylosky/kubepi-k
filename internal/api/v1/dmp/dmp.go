package dmp

import (
	goContext "context"
	"errors"
	"github.com/KubeOperator/kubepi/internal/api/v1/commons"
	"github.com/KubeOperator/kubepi/internal/api/v1/session"
	"github.com/KubeOperator/kubepi/internal/service/v1/cluster"
	"github.com/KubeOperator/kubepi/internal/service/v1/clusterapp"
	"github.com/KubeOperator/kubepi/internal/service/v1/clusterrepo"
	"github.com/KubeOperator/kubepi/internal/service/v1/common"
	"github.com/KubeOperator/kubepi/internal/service/v1/dmpbinding"
	"github.com/KubeOperator/kubepi/internal/service/v1/imagerepo"
	pkgV1 "github.com/KubeOperator/kubepi/pkg/api/v1"
	"github.com/KubeOperator/kubepi/pkg/kubernetes"
	"github.com/asdine/storm/v3"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
	"time"
)

type Handler struct {
	clusterService     cluster.Service
	dmpBindingService  dmpbinding.Service
	clusterRepoService clusterrepo.Service
	imageRepoService   imagerepo.Service
	clusterAppService  clusterapp.Service
}

func NewHandler() *Handler {
	return &Handler{
		clusterService:     cluster.NewService(),
		dmpBindingService:  dmpbinding.NewService(),
		clusterRepoService: clusterrepo.NewService(),
		imageRepoService:   imagerepo.NewService(),
		clusterAppService:  clusterapp.NewService(),
	}
}

func (h *Handler) SearchDmps() iris.Handler {
	return func(ctx *context.Context) {
		// 搜索条件
		pageNum, _ := ctx.Values().GetInt(pkgV1.PageNum)
		pageSize, _ := ctx.Values().GetInt(pkgV1.PageSize)
		showExtra := ctx.URLParamExists("showExtra")
		var conditions commons.SearchConditions
		if err := ctx.ReadJSON(&conditions); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Values().Set("message", err.Error())
			return
		}

		// 获取集群数据
		clusters, total, err := h.clusterService.Search(pageNum, pageSize, conditions.Conditions, common.DBOptions{})
		if err != nil && err != storm.ErrNotFound {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Values().Set("message", err.Error())
			return
		}
		result := make([]Dmp, 0)

		// 根据用户配置Accessable字段
		u := ctx.Values().Get("profile")
		profile := u.(session.UserProfile)
		for i := range clusters {
			c := Dmp{Cluster: clusters[i]}
			bs, err := h.dmpBindingService.GetDmpBindingByDmpName(c.Name, common.DBOptions{})
			if err != nil && !errors.Is(err, storm.ErrNotFound) {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Values().Set("message", err.Error())
				return
			}
			c.MemberCount = len(bs)
			for j := range bs {
				if bs[j].UserRef == profile.Name {
					c.Accessable = true
					c.DmpUser = bs[j].DmpUser
					c.DmpPassword = bs[j].DmpPassword
				}
			}
			result = append(result, c)
		}

		// Extra字段信息
		if showExtra {
			ctx1, cancel := goContext.WithTimeout(goContext.Background(), 10*time.Second)
			defer cancel()

			wg := sync.WaitGroup{}
			go func(result []Dmp) {
				for i := range result {
					wg.Add(1)
					c := kubernetes.NewKubernetes(&result[i].Cluster)
					go func(i int, ctx1 goContext.Context) {
						defer wg.Done()
						info, _ := getExtraClusterInfo(ctx1, c)
						result[i].ExtraDmpInfo = info
					}(i, ctx1)
				}
				wg.Wait()
				cancel()
			}(result)

			<-ctx1.Done()
		}

		// 设置返回值
		ctx.Values().Set("data", pkgV1.Page{Items: result, Total: total})
	}
}

func getExtraClusterInfo(context goContext.Context, client kubernetes.Interface) (ExtraDmpInfo, error) {
	err := client.Ping()
	if err != nil {
		return ExtraDmpInfo{Health: false, Message: err.Error()}, err
	}
	c, err := client.Client()
	if err != nil {
		return ExtraDmpInfo{Health: false, Message: err.Error()}, err
	}
	nodesList, err := c.CoreV1().Nodes().List(context, metav1.ListOptions{})
	if err != nil {
		return ExtraDmpInfo{Health: true, Message: err.Error()}, err
	}
	nodes := nodesList.Items

	totalCpu := float64(0)
	totalMemory := float64(0)
	usedCpu := float64(0)
	usedMemory := float64(0)
	readyNodes := 0
	for i := range nodes {
		conditions := nodes[i].Status.Conditions
		for i := range conditions {
			if conditions[i].Type == "Ready" {
				if conditions[i].Status == "True" {
					readyNodes += 1
				}
			}
		}
		cpu := nodes[i].Status.Allocatable.Cpu().AsApproximateFloat64()
		totalCpu += cpu
		memory := nodes[i].Status.Allocatable.Memory().AsApproximateFloat64()
		totalMemory += memory
	}
	podsList, err := c.CoreV1().Pods("").List(goContext.TODO(), metav1.ListOptions{})
	if err != nil {
		return ExtraDmpInfo{Health: true, Message: err.Error()}, err
	}
	pods := podsList.Items
	for i := range pods {
		for j := range pods[i].Spec.Containers {
			cpu := pods[i].Spec.Containers[j].Resources.Requests.Cpu().AsApproximateFloat64()
			usedCpu += cpu
			memory := pods[i].Spec.Containers[j].Resources.Requests.Memory().AsApproximateFloat64()
			usedMemory += memory

		}
	}
	result := ExtraDmpInfo{
		Health:            true,
		TotalNodeNum:      len(nodes),
		ReadyNodeNum:      readyNodes,
		CPUAllocatable:    totalCpu,
		CPURequested:      usedCpu,
		MemoryAllocatable: totalMemory,
		MemoryRequested:   usedMemory,
	}
	return result, nil
}

func Install(parent iris.Party) {
	handler := NewHandler()
	sp := parent.Party("/dmps")
	sp.Post("/search", handler.SearchDmps())
	sp.Post("/:name/members", handler.CreateDmpMember())
	sp.Delete("/:name/members/:member", handler.DeleteDmpMember())
}
