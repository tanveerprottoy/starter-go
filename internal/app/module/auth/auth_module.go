package auth

type AuthModule struct {
	Service *AuthService
}

func NewAuthModule() *AuthModule {
	m := new(AuthModule)
	m.Service = NewAuthService()
	return m
}
