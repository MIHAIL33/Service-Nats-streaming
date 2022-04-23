package service

import (
	"errors"

	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/cache"
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/repository"
	"github.com/sirupsen/logrus"
)

type ModelService struct {
	repo repository.Model
	ch cache.Model
}

func NewModelService(repo repository.Model, ch cache.Model) *ModelService {
	return &ModelService{ 
		repo: repo,
		ch: ch,
	}
}

func (s *ModelService) Create(model models.Model) (*models.Model, error) {
	oldModel, _ := s.GetById(model.Order_uid)
	// if err != nil {
	// 	logrus.Errorln(err.Error())
	// }
	if oldModel != nil {
		return nil, errors.New("model with this order_uid already exist")
	}
	newModel, err := s.repo.Create(model)
	if err != nil {
		return nil, err
	}
	s.ch.AddOne(newModel)
	return newModel, nil
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

func (s *ModelService) AddAllInCache() error {
	models, err := s.GetAll()
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	err = s.ch.AddAll(models)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

func (s *ModelService) GetAllFromCache() (*[]models.Model, error) {
	return s.ch.GetAll()
}

func (s *ModelService) GetModelFromCacheById(id string) (*models.Model, error) {
	return s.ch.GetOneById(id)
}