package user

import (
	"database/sql"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/handler"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/repository"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/apigateway/module/user/service"
	sqlpkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	Handler         *handler.Handler
	Service         *service.Service
	Repository      sqlpkg.Repository[entity.User]
	MongoRepository *repository.RepositoryAlt
}

func NewModule(db *mongo.Database, dbSql *sql.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = repository.NewRepository(dbSql)
	m.MongoRepository = repository.NewRepositoryAlt(db)
	m.Service = service.NewService(m.Repository)
	m.Handler = handler.NewHandler(m.Service, validate)
	return m
}
