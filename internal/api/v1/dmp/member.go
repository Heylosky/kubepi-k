package dmp

import (
	"fmt"

	"github.com/KubeOperator/kubepi/internal/api/v1/session"
	v1 "github.com/KubeOperator/kubepi/internal/model/v1"
	v1Dmp "github.com/KubeOperator/kubepi/internal/model/v1/dmp"
	"github.com/KubeOperator/kubepi/internal/server"
	"github.com/KubeOperator/kubepi/internal/service/v1/common"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// Create Dmp Member
// @Tags dmps
// @Summary Create Dmp Member
// @Description Create Dmp Member
// @Accept  json
// @Produce  json
// @Param dmp path string true "集群名称"
// @Param request body Member true "request"
// @Success 200 {object} Member
// @Security ApiKeyAuth
// @Router /dmps/{dmp}/members [post]
func (h *Handler) CreateDmpMember() iris.Handler {
	return func(ctx *context.Context) {
		name := ctx.Params().GetString("name")
		var req Member
		err := ctx.ReadJSON(&req)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Values().Set("message", err.Error())
			return
		}
		if req.Name == "" {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Values().Set("message", "username can not be none")
			return
		}
		if req.DmpUser == "" {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Values().Set("message", "dmpUser can not be none")
			return
		}
		if req.DmpPassword == "" {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Values().Set("message", "dmpPassword can not be none")
			return
		}
		u := ctx.Values().Get("profile")
		profile := u.(session.UserProfile)
		binding := v1Dmp.Binding{
			BaseModel: v1.BaseModel{
				Kind:      "DmpBinding",
				CreatedBy: profile.Name,
			},
			Metadata: v1.Metadata{
				Name: fmt.Sprintf("%s-%s-dmp-binding", name, req.Name),
			},
			UserRef:     req.Name,
			DmpRef:      name,
			DmpUser:     req.DmpUser,
			DmpPassword: req.DmpPassword,
		}

		tx, _ := server.DB().Begin(true)

		if err := h.dmpBindingService.CreateDmpBinding(&binding, common.DBOptions{DB: tx}); err != nil {
			_ = tx.Rollback()
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Values().Set("message", "unable to complete authorization")
			return
		}

		_ = tx.Commit()
		ctx.Values().Set("data", req)
	}
}

// Delete Dmp Member
// @Tags dmps
// @Summary Delete dmp Member by name
// @Description Delete dmp Member by name
// @Accept  json
// @Produce  json
// @Param dmp path string true "集群名称"
// @Param members path string true "成员名称"
// @Success 200 {number} 200
// @Security ApiKeyAuth
// @Router /dmps/{dmp}/members/{member} [delete]
func (h *Handler) DeleteDmpMember() iris.Handler {
	return func(ctx *context.Context) {
		name := ctx.Params().GetString("name")
		memberName := ctx.Params().GetString("member")
		u := ctx.Values().Get("profile")
		profile := u.(session.UserProfile)
		c, err := h.clusterService.Get(name, common.DBOptions{})
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Values().Set("message", fmt.Sprintf("get cluster failed: %s", err.Error()))
			return
		}
		if c.CreatedBy == memberName {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.Values().Set("message", fmt.Sprintf("can not delete or update cluster importer %s", profile.Name))
			return
		}

		binding, err := h.dmpBindingService.GetBindingByDmpNameAndUserName(c.Name, memberName, common.DBOptions{})
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Values().Set("message", fmt.Sprintf("get dmp failed: %s", err.Error()))
			return
		}
		tx, err := server.DB().Begin(true)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Values().Set("message", fmt.Sprintf("get dmp failed: %s", err.Error()))
			return
		}
		if err := h.dmpBindingService.Delete(binding.Name, common.DBOptions{DB: tx}); err != nil {
			_ = tx.Rollback()
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Values().Set("message", fmt.Sprintf("delete dmp binding failed: %s", err.Error()))
			return
		}
		_ = tx.Commit()
	}
}
