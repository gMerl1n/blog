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

func (h *Handler) CreatePost(c *gin.Context) {

	var input CreatePostRequest

	if err := c.BindJSON(&input); err != nil {
		return
	}

	postID, err := h.Services.ServicePost.CreatePost(input.Title, input.Body)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, postID)

}

func (h *Handler) GetPostByID(c *gin.Context) {

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	post, err := h.Services.ServicePost.GetPostByID(postID)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, post)

}
