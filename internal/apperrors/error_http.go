package apperrors

import (
	"github.com/gin-gonic/gin"
)

func BadResponse(ctx *gin.Context, err error) {

	basicErr := err.(*basicError)

	ctx.AbortWithStatusJSON(basicErr.httpStatusCode, basicErr)

}
