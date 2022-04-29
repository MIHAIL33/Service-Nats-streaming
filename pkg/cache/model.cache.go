package cache

import (
	"errors"
	"fmt"

	models "github.com/MIHAIL33/Service-Nats-streaming/model"
)

type CacheModel struct {
	models *[]models.Model
}

func NewCacheModel() *CacheModel {
	return &CacheModel{
		models: new([]models.Model),
	}
}

func (cm *CacheModel) AddAll(models *[]models.Model) error {
	*cm.models = *models
	return nil
}

func (cm *CacheModel) AddOne(model *models.Model) error {
	*cm.models = append(*cm.models, *model)
	return nil
}

func (cm *CacheModel) GetAll() (*[]models.Model, error) {
	return cm.models, nil
}

func (cm *CacheModel) GetOneById(id string) (*models.Model, error) {
	for _, model := range *cm.models {
		if model.Order_uid == id {
			return &model, nil
		}
	}
	err := fmt.Sprintf("Not found model with order_uid = %s", id)
	return nil, errors.New(err)
}
