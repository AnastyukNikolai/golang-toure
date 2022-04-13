package handler

import (
	"github.com/gin-gonic/gin"
	"golang-ture/internal/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		items := api.Group("items")
		{
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.POST("/", h.createItem)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
