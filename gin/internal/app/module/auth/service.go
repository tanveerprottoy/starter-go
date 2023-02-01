package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/entity"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/service"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/adapter"
	_http "github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/jwt"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
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
	splits, err := _http.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	tokenBody := splits[1]
	claims, err := jwt.VerifyToken(tokenBody)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	// find user
	row := s.userService.ReadOneInternal(claims.Payload.Id)
	if row == nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	d, err := adapter.RowToUserEntity(row)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	return d
}
