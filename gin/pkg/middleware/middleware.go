package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/tanveerprottoy/starter-go/pkg/constant"
	"github.com/tanveerprottoy/starter-go/pkg/jwt"
	"github.com/tanveerprottoy/starter-go/pkg/response"
)

// JSONContentTypeMiddleWare content type json setter middleware
func JSONContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// CORSEnableMiddleWare enable cors
func CORSEnableMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		next.ServeHTTP(w, r)
	})
}

// JWTMiddleWare checks auth of the request
func JWTMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			// Token is missing
			response.RespondError(http.StatusForbidden, errors.New("auth token is missing"), w)
			return
		}
		split := strings.Split(tokenHeader, " ")
		// token format is `Bearer {tokenBody}`
		if len(split) != 2 {
			response.RespondError(http.StatusForbidden, errors.New("token format is invalid"), w)
			return
		}
		tokenBody := split[1]
		claims, err := jwt.VerifyToken(tokenBody)
		if err != nil {
			response.RespondError(http.StatusForbidden, err, w)
			return
		}
		ctx := context.WithValue(r.Context(), constant.ContextPayloadKey, claims.Payload)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
