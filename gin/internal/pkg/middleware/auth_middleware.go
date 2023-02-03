package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/auth"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
)

type AuthMiddleware struct {
	Service *auth.Service
}

func NewAuthMiddleware(s *auth.Service) *AuthMiddleware {
	m := new(AuthMiddleware)
	m.Service = s
	return m
}

// AuthUserMiddleWare
func (m *AuthMiddleware) AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := m.Service.Authorize(ctx)
		if payload == nil {
			return
		}
		ctx.Set(constant.KeyAuthUser, payload)
		ctx.Next()
	}
}
