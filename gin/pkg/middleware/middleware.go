package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/pkg/jwtpkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/response"
)

// JWTMiddleWare checks auth of the request
func JWTMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := ctx.Request.Header["Authorization"]
		if h == nil && len(h) == 0 {
			response.RespondError(http.StatusForbidden, errors.New("auth token is missing"), ctx)
			return
		}
		tkHeader := h[0]
		if tkHeader == "" {
			// Token is missing
			response.RespondError(http.StatusForbidden, errors.New("auth token is missing"), ctx)
			return
		}
		split := strings.Split(tkHeader, " ")
		// token format is `Bearer {tokenBody}`
		if len(split) != 2 {
			response.RespondError(http.StatusForbidden, errors.New("token format is invalid"), ctx)
			return
		}
		tokenBody := split[1]
		claims, err := jwtpkg.VerifyToken(tokenBody)
		if err != nil {
			response.RespondError(http.StatusForbidden, err, ctx)
			return
		}
		ctx.Set(constant.ContextPayloadKey, claims.Payload)
		ctx.Next()
	}
}

// JSONContentTypeMiddleWare content type json setter middleware
/* func JSONContentTypeMiddleWare(next http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header().Set("Content-Type", "application/json")
		ctx.ServeHTTP(w, r)
	}
} */

// CORSEnableMiddleWare enable cors
/* func CORSEnableMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		next.ServeHTTP(w, r)
	})
} */
