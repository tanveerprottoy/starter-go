package user

import "txp/restapistarter/app/module/user/repository"

type UserModule struct {
	UserHandler    *UserHandler
	UserService    *UserService
	UserRepository *repository.UserRepository
}

func (m *UserModule) InitComponents() {
	m.UserRepository = new(repository.UserRepository)
	m.UserService = NewUserService(
		m.UserRepository,
	)
	m.UserHandler = NewUserHandler(
		m.UserService,
	)
}
