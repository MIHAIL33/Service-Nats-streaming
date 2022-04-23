package service

import (
	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/cache"
)

type Model interface {
	Create(model models.Model) (*models.Model, error)
	GetById(id string) (*models.Model, error)
	GetAll() (*[]models.Model, error)
	Delete(id string) (*models.Model, error)

	AddAllInCache() error
	GetAllFromCache() (*[]models.Model, error)
	GetModelFromCacheById(id string) (*models.Model, error)
}

type Service struct {
	Model
}

func NewService(repos *repository.Repository, ch *cache.Cache) *Service {
	return &Service{
		Model: NewModelService(repos.Model, ch.Model),
	}
}