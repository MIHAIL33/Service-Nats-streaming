package handler

import (
	"net/http"

	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/gin-gonic/gin"
)

type getAllModelResponse struct {
	Data *[]models.Model `json:"data"`
}

func (h *Handler) createModel(c *gin.Context) {
	var input models.Model
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	model, err := h.services.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}

func (h *Handler) getAllModels(c *gin.Context) {
	models, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllModelResponse{
		Data: models,
	})
}

func (h *Handler) getModelById(c *gin.Context) {
	modelId := c.Param("id")
	if modelId == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	model, err := h.services.Model.GetById(modelId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, model)
}

func (h *Handler) deleteModel(c *gin.Context) {

}
