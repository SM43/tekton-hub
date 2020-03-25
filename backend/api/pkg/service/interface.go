package service

import (
	"github.com/jinzhu/gorm"
	"github.com/redhat-developer/tekton-hub/backend/api/pkg/app"
	"go.uber.org/zap"
)

type Service interface {
	Resource() *Resource
}

type ServiceImpl struct {
	app app.Base
	log *zap.SugaredLogger
	db  *gorm.DB
}

func New(base app.Base) *ServiceImpl {
	return &ServiceImpl{
		app: base,
		log: base.Logger().With("name", "db"),
		db:  base.DB(),
	}
}

func (s *ServiceImpl) Resource() *Resource {
	return &Resource{s.db, s.log}
}
