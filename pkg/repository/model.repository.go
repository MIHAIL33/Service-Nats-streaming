package repository

import (
	"encoding/json"
	"fmt"

	models "github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/jmoiron/sqlx"
)

type ModelRepository struct {
	db *sqlx.DB
}

func NewModelRepository(db *sqlx.DB) *ModelRepository {
	return &ModelRepository{ db: db }
}

func (r *ModelRepository) Create(model models.Model) (*models.Model, error) {
	createModelQuery := fmt.Sprintf("INSERT INTO %s (model) VALUES ($1) RETURNING model", modelsTable)
	jsonModel, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	var res models.Model
	err = r.db.QueryRow(createModelQuery, jsonModel).Scan(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ModelRepository) GetById(id string) (*models.Model, error) {
	getModelQuery := fmt.Sprintf("SELECT * FROM %s WHERE model->>'order_uid' = $1", modelsTable)
	var res models.Model
	err := r.db.QueryRow(getModelQuery, id).Scan(&res)
	if err != nil {
		return nil, err 
	}

	return &res, nil
}

func (r *ModelRepository) GetAll() (*[]models.Model, error) {
	getAllModelQuery := fmt.Sprintf("SELECT * FROM %s", modelsTable)
	var res []models.Model
	err := r.db.Select(&res, getAllModelQuery)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r* ModelRepository) Delete(id string) (*models.Model, error) {
	deleteModelQuery := fmt.Sprintf("DELETE FROM %s WHERE model->>'order_uid' = $1 RETURNING *", modelsTable)
	var res models.Model
	err := r.db.QueryRow(deleteModelQuery, id).Scan(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}