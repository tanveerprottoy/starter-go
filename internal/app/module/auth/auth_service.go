package auth

import (
	"errors"
	"net/http"
	"txp/restapistarter/pkg/jwt"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/strings"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	s := new(AuthService)
	return s
}

func (s *AuthService) Authorize(w http.ResponseWriter, r *http.Request) any {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		// Token is missing
		response.RespondError(http.StatusForbidden, errors.New("auth token is missing"), w)
		return nil
	}
	splits := strings.Split(tokenHeader, " ")
	// token format is `Bearer {tokenBody}`
	if len(splits) != 2 {
		response.RespondError(http.StatusForbidden, errors.New("token format is invalid"), w)
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
