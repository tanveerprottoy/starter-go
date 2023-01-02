package auth

import (
	"net/http"

	"github.com/tanveerprottoy/rest-api-starter-go/internal/app/module/user/entity"
	"github.com/tanveerprottoy/rest-api-starter-go/internal/app/module/user/service"
	"github.com/tanveerprottoy/rest-api-starter-go/internal/pkg/adapter"
	_http "github.com/tanveerprottoy/rest-api-starter-go/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/pkg/jwt"
	"github.com/tanveerprottoy/rest-api-starter-go/pkg/response"
)

type Service struct {
	userService *service.Service
}

func NewService(userService *service.Service) *Service {
	s := new(Service)
	s.userService = userService
	return s
}

func (s *Service) Authorize(w http.ResponseWriter, r *http.Request) *entity.User {
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
