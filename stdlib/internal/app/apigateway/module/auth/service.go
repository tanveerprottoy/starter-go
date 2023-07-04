package auth

import (
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/service"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/adapter"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/jwtpkg"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
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
	splits, err := httppkg.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	tokenBody := splits[1]
	claims, err := jwtpkg.VerifyToken(tokenBody)
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
