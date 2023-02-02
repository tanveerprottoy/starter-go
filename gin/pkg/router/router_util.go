package router

import "github.com/gin-gonic/gin"

func GetURLParam(ctx *gin.Context, key string) string {
	return ctx.Param(key)
}

func GetQueryParam(ctx *gin.Context, key string) string {
	return ctx.Request.URL.GetQuery(key)
}
