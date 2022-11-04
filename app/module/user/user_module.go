package user

import (
	"txp/restapistarter/app/module/user/repository"
	"txp/restapistarter/pkg/router"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModule struct {
	UserRouter     *UserRouter
	UserHandler    *UserHandler
	UserService    *UserService
	UserRepository *repository.UserRepository
}

func NewUserModule(db *mongo.Database, router *router.Router) *UserModule {
	m := new(UserModule)
	m.UserRouter = NewUserRouter(router, m)
	m.UserRepository = new(repository.UserRepository)
	m.UserService = NewUserService(m.UserRepository)
	m.UserHandler = NewUserHandler(m.UserService)
	return m
}
