package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetURLParam(ctx *gin.Context, key string) string {
	return ctx.Param(key)
}

func GetQueryParam(ctx *gin.Context, key string) string {
	return ctx.Request.URL.GetQuery(key)
}

func ParseAuthToken(ctx *gin.Context) ([]string, error) {
	tokenHeader := ctx.Header.Get("Authorization")
	if tokenHeader == "" {
		// Token is missing
		return nil, errors.New("auth token is missing")
	}
	splits := strings.Split(tokenHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		return nil, errors.New("token format is invalid")
	}
	return splits, nil
}

func BindGin[T any](ctx *gin.Context, v T) error {
	return ctx.ShouldBind(&v)
}
