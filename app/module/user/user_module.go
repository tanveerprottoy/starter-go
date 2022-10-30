package user

import (
	"txp/restapistarter/app/module/user/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModule struct {
	DB             *mongo.Database
	UserHandler    *UserHandler
	UserService    *UserService
	UserRepository *repository.UserRepository
}

func (m *UserModule) InitComponents(db *mongo.Database) {
	m.UserRepository = new(repository.UserRepository)
	m.UserService = NewUserService(
		m.UserRepository,
	)
	m.UserHandler = NewUserHandler(
		m.UserService,
	)
}
