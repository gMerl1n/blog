package handlers

import (
	"net/http"

	"github.com/gMerl1n/blog/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *services.Service
}

func NewHandler(service *services.Service) Handler {
	return Handler{Services: service}
}

func (h *Handler) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Ok")
}

func (h *Handler) InitRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", h.TestHandler)

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
