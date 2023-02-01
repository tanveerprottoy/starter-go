package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/auth/dto"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/config"
	httpPkg "github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
)

type ServiceRemote struct {
	HTTPClient *httpPkg.HTTPClient
}

func NewServiceRemote(c *httpPkg.HTTPClient) *ServiceRemote {
	s := new(ServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *ServiceRemote) Authorize(ctx *gin.Context) any {
	_, err := httpPkg.ParseAuthToken(ctx)
	if err != nil {
		response.RespondError(http.StatusForbidden, err)
		return nil
	}
	u, err := httpPkg.Request[dto.AuthUserDto](
		http.MethodPost,
		fmt.Sprintf("%s%s", config.GetEnvValue("USER_SERVICE_BASE_URL"), constant.UserServiceAuthEndpoint),
		ctx.Header,
		nil,
		s.HTTPClient,
	)
	if err != nil {
		response.RespondError(http.StatusForbidden, err)
		return nil
	}
	return u
}
