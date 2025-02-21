package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	UserID int
	Title  string
	Body   string
}

func (h *Handler) CreatePost(ctx *gin.Context) {

	var input CreatePostRequest

	if err := ctx.BindJSON(&input); err != nil {
		return
	}

	postID, err := h.Services.ServicePost.CreatePost(ctx, input.Title, input.Body)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, postID)

}

func (h *Handler) GetPostByID(ctx *gin.Context) {

	postID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	post, err := h.Services.ServicePost.GetPostByID(ctx, postID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, post)

}
