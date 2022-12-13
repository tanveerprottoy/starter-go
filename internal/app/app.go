package app

import (
	"log"
	"net/http"
	"txp/restapistarter/internal/app/module/auth"
	"txp/restapistarter/internal/app/module/content"
	"txp/restapistarter/internal/app/module/user"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/internal/pkg/middleware"
	_routerModule "txp/restapistarter/internal/pkg/router"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/data/sql/postgres"
	"txp/restapistarter/pkg/router"
)

// App struct
type App struct {
	MongoDBClient    *mongodb.DBClient
	PostgresDBClient *postgres.DBClient
	router           *router.Router
	Middlewares      []any
	AuthModule       *auth.AuthModule
	UserModule       *user.UserModule
	ContentModule    *content.ContentModule
}

func (a *App) initDB() {
	a.MongoDBClient = mongodb.GetInstance()
	a.PostgresDBClient = postgres.GetInstance()
}

func (a *App) initMiddlewares() {
	authMiddleWare := middleware.NewAuthMiddleware(a.AuthModule.Service)
	a.Middlewares = append(a.Middlewares, authMiddleWare)
}

func (a *App) initModules() {
	a.UserModule = user.NewUserModule(a.MongoDBClient.DB, a.PostgresDBClient.DB)
	a.AuthModule = auth.NewAuthModule(a.UserModule.Service)
	a.ContentModule = content.NewContentModule(a.PostgresDBClient.DB)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	_routerModule.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
	_routerModule.RegisterContentRoutes(a.router, constant.V1, a.ContentModule)
}

// Init app
func (a *App) InitComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
