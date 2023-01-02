package auth

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/rest-api-starter-go/internal/app/module/auth/dto"
	"github.com/tanveerprottoy/rest-api-starter-go/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/pkg/config"
	httpPkg "github.com/tanveerprottoy/rest-api-starter-go/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/pkg/response"
)

type ServiceRemote struct {
	HTTPClient *httpPkg.HTTPClient
}

func NewServiceRemote(c *httpPkg.HTTPClient) *ServiceRemote {
	s := new(ServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *ServiceRemote) Authorize(w http.ResponseWriter, r *http.Request) any {
	_, err := httpPkg.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	u, err := httpPkg.Request[dto.AuthUserDto](
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
