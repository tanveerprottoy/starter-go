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
	UserMongoRepository *repository.UserMongoRepository
}

func NewUserModule(db *mongo.Database, router *router.Router) *UserModule {
	m := new(UserModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.UserRepository = new(repository.UserRepository)
	m.UserMongoRepository = repository.NewUserMongoRepository(db)
	m.UserService = NewUserService(m.UserRepository)
	m.UserHandler = NewUserHandler(m.UserService)
	m.UserRouter = NewUserRouter(router, m)
	return m
}
