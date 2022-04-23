package handler

import (
	"github.com/MIHAIL33/Service-Nats-streaming/pkg/service"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/MIHAIL33/Service-Nats-streaming/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		models := api.Group("models")
		{

			cache := models.Group("cache")
			{
				cache.GET("", h.getAllModelsFromCache)
				cache.GET(":id", h.getModelFromCacheById)
			}

			models.POST("", h.createModel)
			models.GET("", h.getAllModels)
			models.GET(":id", h.getModelById)
			models.DELETE(":id", h.deleteModel)
		}
	}

	return router
}