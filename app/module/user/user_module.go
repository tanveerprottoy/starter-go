package user

type UserModule struct {
	UserHandler    *UserHandler
	UserService    *UserService
	UserRepository *UserRepository
}

func (m *UserModule) InitComponents() {
	m.UserRepository = new(UserRepository)
	m.UserService = NewUserService(
		m.UserRepository,
	)
	m.UserHandler = NewUserHandler(
		m.UserService,
	)
}
