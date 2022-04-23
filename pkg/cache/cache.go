package cache

import models "github.com/MIHAIL33/Service-Nats-streaming/model"

type Model interface {
	AddAll(models *[]models.Model) error
	AddOne(model *models.Model) error
	GetAll() (*[]models.Model, error)
	GetOneById(id string) (*models.Model, error)
}

type Cache struct {
	Model
}

func NewCache() *Cache {
	return &Cache{
		Model: NewCacheModel(),
	}
}