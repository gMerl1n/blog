package handlers

import (
	"fmt"
	"net/http"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gMerl1n/blog/internal/entities/requests"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {

	var input requests.CreateUserRequest

	if err := ctx.BindJSON(&input); err != nil {
		h.logger.Warn(fmt.Sprintf("failed to decode post request data. Error: %s", err))
		er.BadResponse(ctx, er.IncorrectRequestParams.SetCause(err.Error()))
		return
	}

	userID, err := h.Services.ServiceUser.CreateUser(ctx, input.Name, input.Email, input.Password, input.RepeatPassword)
	if err != nil {
		h.logger.Warn(fmt.Sprintf("failed to create user. Error: %s", err))
		er.BadResponse(ctx, err)
	}

	ctx.JSON(http.StatusCreated, userID)

}
