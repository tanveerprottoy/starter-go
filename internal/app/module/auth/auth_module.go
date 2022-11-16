package auth

import "txp/restapistarter/internal/app/module/user"

type AuthModule struct {
	Service *AuthService
}

func NewAuthModule(userService *user.UserService) *AuthModule {
	m := new(AuthModule)
	m.Service = NewAuthService(userService)
	return m
}
