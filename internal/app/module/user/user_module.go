package user

import (
	"txp/restapistarter/internal/app/module/user/entity"
	"txp/restapistarter/internal/app/module/user/repository"
	"txp/restapistarter/internal/pkg/middleware"
	data "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/router"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModule struct {
	Router          *UserRouter
	Handler         *UserHandler
	Service         *UserService
	Repository      data.Repository[entity.User]
	MongoRepository *repository.UserMongoRepository
}

func NewUserModule(db *mongo.Database, router *router.Router, authMiddleware *middleware.AuthMiddleware) *UserModule {
	m := new(UserModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = new(repository.UserRepository[entity.User])
	m.MongoRepository = repository.NewUserMongoRepository(db)
	m.Service = NewUserService(m.Repository)
	m.Handler = NewUserHandler(m.Service)
	m.Router = NewUserRouter(router, m, authMiddleware)
	return m
}
