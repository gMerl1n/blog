package handlers

import (
	"net/http"

	"github.com/gMerl1n/blog/internal/services"
	"github.com/gMerl1n/blog/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	TokenManager jwt.ITokenManager
	Services     *services.Service
	logger       *logrus.Logger
}

func NewHandler(service *services.Service, tokenManager jwt.ITokenManager, logger *logrus.Logger) Handler {
	return Handler{
		Services:     service,
		TokenManager: tokenManager,
		logger:       logger,
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
			posts.POST("/", h.TokenAuthMiddleware, h.CreatePost)
			posts.PATCH("/", h.TokenAuthMiddleware, h.UpdatePost)
			posts.GET("/:id", h.GetPostByID)
			posts.GET("/", h.GetPosts)
		}
		users := api.Group("/users")
		{
			users.POST("/", h.CreateUser)
			users.POST("/login", h.LoginUser)

		}
	}

	return router

}
