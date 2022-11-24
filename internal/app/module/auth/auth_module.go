package auth

import "txp/restapistarter/internal/app/module/user/service"

type AuthModule struct {
	Service *AuthService
}

func NewAuthModule(userService *service.UserService) *AuthModule {
	m := new(AuthModule)
	m.Service = NewAuthService(userService)
	return m
}
