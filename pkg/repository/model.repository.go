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

type Mod struct {
	Model models.Model
}

func NewModelRepository(db *sqlx.DB) *ModelRepository {
	return &ModelRepository{ db: db }
}

func (r *ModelRepository) Create(model models.Model) (models.Model, error) {
	createModelQuery := fmt.Sprintf("INSERT INTO %s (model) VALUES ($1) RETURNING *", modelsTable)
	jsonModel, err := json.Marshal(model)
	if err != nil {
		return model, err
	}
	
	var res models.Model
	err = r.db.QueryRow(createModelQuery, jsonModel).Scan(&res)
	if err != nil {
		return model, err
	}

	return res, nil
}