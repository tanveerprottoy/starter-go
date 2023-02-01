package response

import (
	"github.com/gin-gonic/gin"
)

func BuildData[T any](p T) *Response[T] {
	return &Response[T]{Data: p}
}

func Respond(code int, data any, ctx *gin.Context) {
	ctx.JSON(code, data)
}

func RespondError(code int, err error, ctx *gin.Context) {
	ctx.AbortWithError(code, err)
}
