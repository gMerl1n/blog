package handlers

import (
	"net/http"

	"github.com/gMerl1n/blog/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Services *services.Service
	logger   *logrus.Logger
}

func NewHandler(service *services.Service, logger *logrus.Logger) Handler {
	return Handler{
		Services: service,
		logger:   logger,
	}
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
			posts.POST("/", h.CreatePost)
			posts.GET("/:id", h.GetPostByID)
			posts.GET("/", h.GetPosts)
		}
	}

	return router

}
