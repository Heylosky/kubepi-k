package dmp

import (
	v1Cluster "github.com/KubeOperator/kubepi/internal/model/v1/cluster"
	"time"
)

type Dmp struct {
	v1Cluster.Cluster
	Accessable   bool         `json:"accessable"`
	MemberCount  int          `json:"memberCount"`
	DmpUser      string       `json:"dmpUser"`
	DmpPassword  string       `json:"dmpPassword"`
	ExtraDmpInfo ExtraDmpInfo `json:"extraDmpInfo"`
}

type ExtraDmpInfo struct {
	TotalNodeNum      int     `json:"totalNodeNum"`
	ReadyNodeNum      int     `json:"readyNodeNum"`
	CPUAllocatable    float64 `json:"cpuAllocatable"`
	CPURequested      float64 `json:"cpuRequested"`
	MemoryAllocatable float64 `json:"memoryAllocatable"`
	MemoryRequested   float64 `json:"memoryRequested"`
	Health            bool    `json:"health"`
	Message           string  `json:"message"`
}

type Member struct {
	Name        string    `json:"name"`
	DmpUser     string    `json:"dmpUser"`
	DmpPassword string    `json:"dmpPassword"`
	BindingName string    `json:"bindingName"`
	CreateAt    time.Time `json:"createAt"`
}
