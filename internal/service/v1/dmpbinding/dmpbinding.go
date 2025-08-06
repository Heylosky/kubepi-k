package dmpbinding

import (
	"errors"
	"time"

	v1Dmp "github.com/KubeOperator/kubepi/internal/model/v1/dmp"
	"github.com/KubeOperator/kubepi/internal/service/v1/common"
	"github.com/asdine/storm/v3/q"
	"github.com/google/uuid"
)

type Service interface {
	common.DBService
	CreateDmpBinding(binding *v1Dmp.Binding, options common.DBOptions) error
	Delete(name string, options common.DBOptions) error
	GetDmpBindingByDmpName(DmpName string, options common.DBOptions) ([]v1Dmp.Binding, error)
	GetBindingByDmpNameAndUserName(dmpName string, userName string, options common.DBOptions) (*v1Dmp.Binding, error)
}

func NewService() Service {
	return &service{}
}

type service struct {
	common.DefaultDBService
}

func (s *service) CreateDmpBinding(binding *v1Dmp.Binding, options common.DBOptions) error {
	db := s.GetDB(options)
	binding.UUID = uuid.New().String()
	binding.CreateAt = time.Now()
	binding.UpdateAt = time.Now()
	return db.Save(binding)
}

func (s *service) Delete(name string, options common.DBOptions) error {
	db := s.GetDB(options)
	var binding v1Dmp.Binding
	if err := db.One("Name", name, &binding); err != nil {
		return err
	}
	if binding.BuiltIn {
		return errors.New("can not delete this resource,because it created by system")
	}
	return db.DeleteStruct(&binding)
}

func (s *service) GetDmpBindingByDmpName(dmpName string, options common.DBOptions) ([]v1Dmp.Binding, error) {
	db := s.GetDB(options)
	query := db.Select(q.Eq("DmpRef", dmpName))
	var rbs []v1Dmp.Binding
	if err := query.Find(&rbs); err != nil {
		return rbs, err
	}
	return rbs, nil
}

func (s *service) GetBindingByDmpNameAndUserName(dmpName string, userName string, options common.DBOptions) (*v1Dmp.Binding, error) {
	db := s.GetDB(options)
	query := db.Select(q.And(q.Eq("DmpRef", dmpName), q.Eq("UserRef", userName)))
	var rb v1Dmp.Binding
	if err := query.First(&rb); err != nil {
		return nil, err
	}
	return &rb, nil
}
