package repository

import (
	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/jmoiron/sqlx"
)

type Model interface {
	Create(model models.Model) (*models.Model, error)
	GetById(id string) (*models.Model, error)
	GetAll() (*[]models.Model, error)
	Delete(id string) (*models.Model, error)
}

type Repository struct {
	Model
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Model: NewModelRepository(db),
	}
}