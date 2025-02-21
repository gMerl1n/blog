package handlers

import (
	"fmt"
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

	h.logger.Info(fmt.Sprintf("creating post by user %d", input.UserID))

	postID, err := h.Services.ServicePost.CreatePost(ctx, input.Title, input.Body)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to create user. Error: %s ", err))
		return
	}

	h.logger.Info(fmt.Sprintf("post with id %d created successfully", postID))

	ctx.JSON(http.StatusOK, postID)

}

func (h *Handler) GetPostByID(ctx *gin.Context) {

	h.logger.Info("gettings post by id...")

	postID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to get post id from query params %s ", err))
		return
	}

	post, err := h.Services.ServicePost.GetPostByID(ctx, postID)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to get post from db %s ", err))
		return
	}

	h.logger.Info(fmt.Sprintf("post with id %d got successfully", post.ID))

	ctx.JSON(http.StatusCreated, post)

}
