package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (h *Handler) CreatePost(ctx *gin.Context) {

	var input CreatePostRequest

	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode post request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	postID, err := h.Services.ServicePost.CreatePost(ctx, input.Title, input.Body)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to create user. Error: %s ", err))
		er.BadResponse(ctx, err)
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
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	post, err := h.Services.ServicePost.GetPostByID(ctx, postID)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to get post from db %s ", err))
		er.BadResponse(ctx, err)
		return
	}

	h.logger.Info(fmt.Sprintf("post with id %d got successfully", post.ID))

	ctx.JSON(http.StatusCreated, post)

}

func (h *Handler) GetPosts(ctx *gin.Context) {

	h.logger.Info("getting list of posts")

	listPosts, err := h.Services.ServicePost.GetPosts(ctx)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to get all posts: %s", err))
		er.BadResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, listPosts)

}
