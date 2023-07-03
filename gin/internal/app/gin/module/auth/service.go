package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/user/entity"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/user/service"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/adapter"
	"github.com/tanveerprottoy/starter-go/gin/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/jwtpkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/response"
)

type Service struct {
	userService *service.Service
}

func NewService(userService *service.Service) *Service {
	s := new(Service)
	s.userService = userService
	return s
}

func (s *Service) Authorize(ctx *gin.Context) *entity.User {
	splits, err := httppkg.ParseAuthToken(ctx)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, ctx)
		return nil
	}
	tokenBody := splits[1]
	claims, err := jwtpkg.VerifyToken(tokenBody)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, ctx)
		return nil
	}
	// find user
	row := s.userService.ReadOneInternal(claims.Payload.Id)
	if row == nil {
		response.RespondError(http.StatusForbidden, err, ctx)
		return nil
	}
	d, err := adapter.RowToUserEntity(row)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, ctx)
		return nil
	}
	return d
}
