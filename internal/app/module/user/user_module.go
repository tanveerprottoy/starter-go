package user

import (
	"txp/restapistarter/internal/app/module/user/entity"
	"txp/restapistarter/internal/app/module/user/handler"
	"txp/restapistarter/internal/app/module/user/repository"
	"txp/restapistarter/internal/app/module/user/service"
	data "txp/restapistarter/pkg/data/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModule struct {
	Handler         *handler.UserHandler
	Service         *service.UserService
	Repository      data.Repository[entity.User]
	MongoRepository *repository.UserMongoRepository
}

func NewUserModule(db *mongo.Database) *UserModule {
	m := new(UserModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = new(repository.UserRepository[entity.User])
	m.MongoRepository = repository.NewUserMongoRepository(db)
	m.Service = service.NewUserService(m.Repository)
	m.Handler = handler.NewUserHandler(m.Service)
	return m
}
