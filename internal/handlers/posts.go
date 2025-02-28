package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodPost {
		h.logger.Warn("http method should be post")
		er.BadResponse(ctx, er.NotAllowed.SetCause("http method should be post"))
		return
	}

	var input requests.CreatePostRequest

	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode post request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	postID, err := h.Services.ServicePost.CreatePost(ctx.Request.Context(), input.Title, input.Body)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to create user. Error: %s ", err))
		er.BadResponse(ctx, err)
		return
	}

	h.logger.Info(fmt.Sprintf("post with id %d created successfully", postID))

	ctx.JSON(http.StatusOK, postID)

}

func (h *Handler) GetPostByID(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodGet {
		h.logger.Warn("http method should be get")
		er.BadResponse(ctx, er.NotAllowed.SetCause("http method should be get"))
		return
	}

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

	if ctx.Request.Method != http.MethodPost {
		h.logger.Warn("http method should be post")
		er.BadResponse(ctx, er.NotAllowed.SetCause("http method should be post"))
		return
	}

	h.logger.Info("getting list of posts")

	listPosts, err := h.Services.ServicePost.GetPosts(ctx)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to get all posts: %s", err))
		er.BadResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, listPosts)

}

func (h *Handler) UpdatePost(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodPatch {
		h.logger.Warn("http method should be patch")
		er.BadResponse(ctx, er.NotAllowed.SetCause("http method should be patch"))
		return
	}

	h.logger.Info("Updating post")

	var input requests.UpdatePostRequest

	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode post request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	postUpdatedID, err := h.Services.ServicePost.UpdatePost(ctx, input)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode post request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, postUpdatedID)

}
