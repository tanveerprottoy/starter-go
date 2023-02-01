package errors

import (
	"errors"
)

func NewError(m string) error {
	return errors.New(m)
}

/* func AbortWithError(code int, message string, ctx *gin.Context) {
	ctx.AbortWithError(
		code,
		errors.New(
			message,
		),
	)
} */
