package handlers

import (
	"github.com/gMerl1n/blog/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{Services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.POST("/")
			posts.GET("/:id")
		}
	}

	return router

}
