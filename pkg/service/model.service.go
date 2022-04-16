package service

import (
	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"
)

type ModelService struct {
	repo repository.Model
}

func NewModelService(repo repository.Model) *ModelService {
	return &ModelService{ repo: repo }
}

func (s *ModelService) Create(model models.Model) (models.Model, error) {
	return s.repo.Create(model)
}