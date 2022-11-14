package middleware

import (
	"context"
	"net/http"
	"txp/restapistarter/internal/app/module/auth"
	"txp/restapistarter/internal/pkg/constant"
)

type AuthMiddleWare struct {
	Service *auth.AuthService
}

func NewAuthMiddleWare(s *auth.AuthService) *AuthMiddleWare {
	m := new(AuthMiddleWare)
	m.Service = s
	return m
}

// AuthUserMiddleWare auth user
func (m *AuthMiddleWare) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := m.Service.Authorize(w, r)
		if payload == nil {
			return
		}
		ctx := context.WithValue(r.Context(), constant.ContextKey, payload)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
