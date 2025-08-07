package dmp

import v1 "github.com/KubeOperator/kubepi/internal/model/v1"

type Binding struct {
	v1.BaseModel `storm:"inline"`
	v1.Metadata  `storm:"inline"`
	UserRef      string `json:"UserRef" storm:"inline"`
	DmpRef       string `json:"dmpRef" storm:"index"`
	DmpUser      string `json:"dmpUser" storm:"inline"`
	DmpPassword  string `json:"dmpPassword" strom:"inline"`
}
