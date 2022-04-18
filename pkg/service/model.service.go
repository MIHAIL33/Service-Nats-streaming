package service

import (
	"errors"

	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"
	"github.com/sirupsen/logrus"
)

type ModelService struct {
	repo repository.Model
}

func NewModelService(repo repository.Model) *ModelService {
	return &ModelService{ repo: repo }
}

func (s *ModelService) Create(model models.Model) (*models.Model, error) {
	oldModel, err := s.GetById(model.Order_uid)
	if err != nil {
		logrus.Errorln(err.Error())
	}
	if oldModel != nil {
		return nil, errors.New("model with this order_uid already exist")
	}
	return s.repo.Create(model)
}

func (s *ModelService) GetById(id string) (*models.Model, error) {
	return s.repo.GetById(id)
}

func (s *ModelService) GetAll() (*[]models.Model, error) {
	return s.repo.GetAll()
}

func (s *ModelService) Delete(id string) (*models.Model, error) {
	return s.repo.Delete(id)
} 