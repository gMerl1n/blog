package handlers

import (
	"strings"

	er "github.com/gMerl1n/blog/internal/apperrors"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) TokenAuthMiddleware(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		er.BadResponse(c, er.Unauthorized.SetCause("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		er.BadResponse(c, er.Unauthorized.SetCause("invalid auth header"))
		return
	}

	if len(headerParts[1]) == 0 {
		er.BadResponse(c, er.Unauthorized.SetCause("token is empty"))
		return
	}

	userId, err := h.TokenManager.Parse(headerParts[1])
	if err != nil {
		er.BadResponse(c, er.Unauthorized.SetCause(err.Error()))
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) getUserID(ctx *gin.Context) (int, error) {

	id, ok := ctx.Get(userCtx)
	if !ok {
		return 0, er.NotFound.SetCause("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, er.IncorrectData.SetCause("user id is of invalid type")
	}

	return idInt, nil

}
