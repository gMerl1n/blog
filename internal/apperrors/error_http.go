package apperrors

import (
	"github.com/gin-gonic/gin"
)

func BadResponse(ctx *gin.Context, err error) {

	basicError1 := err.(*basicError)

	ctx.AbortWithStatusJSON(basicError1.httpStatusCode, basicError1)

}
