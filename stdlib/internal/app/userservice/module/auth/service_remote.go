package auth

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/auth/dto"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/config"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
)

type ServiceRemote struct {
	HTTPClient *httppkg.HTTPClient
}

func NewServiceRemote(c *httppkg.HTTPClient) *ServiceRemote {
	s := new(ServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *ServiceRemote) Authorize(w http.ResponseWriter, r *http.Request) any {
	_, err := httppkg.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	u, err := httppkg.Request[dto.AuthUserDto](
		http.MethodPost,
		fmt.Sprintf("%s%s", config.GetEnvValue("USER_SERVICE_BASE_URL"), constant.UserServiceAuthEndpoint),
		r.Header,
		nil,
		s.HTTPClient,
	)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	return u
}
