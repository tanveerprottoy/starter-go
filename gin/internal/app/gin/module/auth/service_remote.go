package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/auth/dto"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/pkg/config"
	"github.com/tanveerprottoy/starter-go/gin/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/response"
)

type ServiceRemote struct {
	HTTPClient *httppkg.HTTPClient
}

func NewServiceRemote(c *httppkg.HTTPClient) *ServiceRemote {
	s := new(ServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *ServiceRemote) Authorize(ctx *gin.Context) any {
	_, err := httppkg.ParseAuthToken(ctx)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, ctx)
		return nil
	}
	u, err := httppkg.Request[dto.AuthUserDto](
		http.MethodPost,
		fmt.Sprintf("%s%s", config.GetEnvValue("USER_SERVICE_BASE_URL"), constant.UserServiceAuthEndpoint),
		ctx.Request.Header,
		nil,
		s.HTTPClient,
	)
	if err != nil {
		response.RespondError(http.StatusForbidden, err, ctx)
		return nil
	}
	return u
}
