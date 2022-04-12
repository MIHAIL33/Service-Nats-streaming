package repository

import "github.com/jmoiron/sqlx"

type Model interface {

}

type Repository struct {
	Model
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}