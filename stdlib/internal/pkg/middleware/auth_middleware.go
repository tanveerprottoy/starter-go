package middleware

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/starter-go/internal/app/module/auth"
	"github.com/tanveerprottoy/starter-go/internal/pkg/constant"
)

type AuthMiddleware struct {
	Service *auth.Service
}

func NewAuthMiddleware(s *auth.Service) *AuthMiddleware {
	m := new(AuthMiddleware)
	m.Service = s
	return m
}

// AuthUserMiddleWare auth user
func (m *AuthMiddleware) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := m.Service.Authorize(w, r)
		if payload == nil {
			return
		}
		ctx := context.WithValue(r.Context(), constant.KeyAuthUser, payload)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
