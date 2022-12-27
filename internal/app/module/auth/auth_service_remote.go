package auth

import (
	"fmt"
	"net/http"

	"txp/restapistarter/internal/app/module/auth/dto"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/config"
	httpinternal "txp/restapistarter/pkg/http"
	"txp/restapistarter/pkg/response"
)

type AuthServiceRemote struct {
	HTTPClient *httpinternal.HTTPClient
}

func NewServiceRemote(c *httpinternal.HTTPClient) *AuthServiceRemote {
	s := new(AuthServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *AuthServiceRemote) Authorize(w http.ResponseWriter, r *http.Request) any {
	_, err := httpinternal.ParseAuthToken(r)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, w)
		return nil
	}
	u, err := httpinternal.Request[dto.AuthUserDto](
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
