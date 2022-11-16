package auth

import (
	"net/http"

	_http "txp/restapistarter/pkg/http"
	"txp/restapistarter/pkg/jwt"
	"txp/restapistarter/pkg/response"
)

type AuthServiceRemote struct {
	HTTPClient *_http.HTTPClient
}

func NewServiceRemote(c *_http.HTTPClient) *AuthServiceRemote {
	s := new(AuthServiceRemote)
	s.HTTPClient = c
	return s
}

func (s *AuthServiceRemote) Authorize(w http.ResponseWriter, r *http.Request) any {
	splits, err := _http.ParseToken(r)
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
	return claims
}
