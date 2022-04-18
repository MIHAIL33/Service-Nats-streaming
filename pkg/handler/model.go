package handler

import (
	"net/http"

	"github.com/MIHAIL33/Service-Nats-streaming/model"
	"github.com/gin-gonic/gin"
)

type getAllModelResponse struct {
	Data *[]models.Model `json:"data"`
}


// @Summary Create
// @Tags API
// @Description Create new model
// @Accept json
// @Produce json
// @Param model body models.Model true "model info"
// @Success 200 {object} models.Model
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/models [post]
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

// @Summary GetAll
// @Tags API
// @Description Get all model
// @Accept json
// @Produce json
// @Success 200 {object} []models.Model
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/models [get]
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

// @Summary GetById
// @Tags API
// @Description Get model by order_uid
// @Accept json
// @Produce json
// @Param id path string true "order_uid"
// @Success 200 {object} models.Model
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/models/{id} [get]
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

// @Summary Delete
// @Tags API
// @Description delete model by order_uid
// @Accept json
// @Produce json
// @Param id path string true "order_uid"
// @Success 200 {object} models.Model
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/models/{id} [delete]
func (h *Handler) deleteModel(c *gin.Context) {
	modelId := c.Param("id")
	if modelId == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	model, err := h.services.Delete(modelId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model)
}
