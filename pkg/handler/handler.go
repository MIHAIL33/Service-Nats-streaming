package handler

import (
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		models := api.Group("/models")
		{
			models.POST("/", h.createModel)
			models.GET("/", h.getAllModels)
			models.GET("/:id", h.getModelById)
			models.DELETE("/:id", h.deleteModel)
		}
	}

	return router
}