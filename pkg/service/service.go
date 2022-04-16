package service

import (
	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"
)

type Model interface {
	Create(model models.Model) (models.Model, error)
}

type Service struct {
	Model
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Model: NewModelService(repos.Model),
	}
}