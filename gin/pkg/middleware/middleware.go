package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/jwt"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
)

// JSONContentTypeMiddleWare content type json setter middleware
func JSONContentTypeMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}

// CORSEnableMiddleWare enable cors
func CORSEnableMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
		// ctx.Request.Header.Set("Access-Control-Allow-Credentials", "true")
		// ctx.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Request.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}

// JWTMiddleWare checks auth of the request
func JWTMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			// Token is missing
			response.RespondError(http.StatusForbidden, errors.New("auth token is missing"), ctx)
			return
		}
		split := strings.Split(tokenHeader, " ")
		// token format is `Bearer {tokenBody}`
		if len(split) != 2 {
			response.RespondError(http.StatusForbidden, errors.New("token format is invalid"), ctx)
			return
		}
		tokenBody := split[1]
		claims, err := jwt.VerifyToken(tokenBody)
		if err != nil {
			response.RespondError(http.StatusForbidden, err, ctx)
			return
		}
		ctx.Set(constant.ContextPayloadKey, claims.Payload)
		ctx.Next()
	}
}
