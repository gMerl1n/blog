package handlers

import (
	"fmt"
	"net/http"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodPost {
		h.logger.Warn("http method should be post to create user")
		er.BadResponse(ctx, er.NotAllowed.SetCause("http method should be post to create user"))
		return
	}

	var input requests.CreateUserRequest

	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode post request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	tokens, err := h.Services.ServiceUser.CreateUser(ctx, input.Name, input.Email, input.Password, input.RepeatPassword)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to create user. Error: %s", err))
		er.BadResponse(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, tokens)

}

func (h *Handler) LoginUser(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodPost {
		h.logger.Warn("http method should be post to login user")
		er.BadResponse(ctx, er.NotAllowed.SetCause("http method should be post to login user"))
		return
	}

	var input requests.LoginUserRequest

	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode login user request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	tokens, err := h.Services.ServiceUser.LoginUser(ctx, input.Email, input.Password)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to login user and get tokens %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, tokens)

}
