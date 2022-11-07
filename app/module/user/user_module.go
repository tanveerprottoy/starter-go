package user

import (
	"txp/restapistarter/app/module/user/repository"
	"txp/restapistarter/pkg/router"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModule struct {
	Router          *UserRouter
	Handler         *UserHandler
	Service         *UserService
	Repository      *repository.UserRepository
	MongoRepository *repository.UserMongoRepository
}

func NewUserModule(db *mongo.Database, router *router.Router) *UserModule {
	m := new(UserModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = new(repository.UserRepository)
	m.MongoRepository = repository.NewUserMongoRepository(db)
	m.Service = NewUserService(m.Repository)
	m.Handler = NewUserHandler(m.Service)
	m.Router = NewUserRouter(router, m)
	return m
}
