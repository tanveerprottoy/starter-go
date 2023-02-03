package app

import (
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/auth"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/content"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/middleware"
	routerPkg "github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/router"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/data/nosql/mongodb"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/data/sql/postgres"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/router"

	"github.com/go-playground/validator/v10"
	// "go.uber.org/zap"
)

// App struct
type App struct {
	MongoDBClient    *mongodb.Client
	PostgresDBClient *postgres.Client
	router           *router.Router
	Middlewares      []any
	AuthModule       *auth.Module
	UserModule       *user.Module
	ContentModule    *content.Module
	Validate         *validator.Validate
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
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
	a.UserModule = user.NewModule(a.MongoDBClient.DB, a.PostgresDBClient.DB, a.Validate)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
	a.ContentModule = content.NewModule(a.PostgresDBClient.DB)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	routerPkg.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
	routerPkg.RegisterContentRoutes(a.router, constant.V1, a.ContentModule)
}

/* func (a *App) initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"proxy.log",
	}
	cfg.Build()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	task := "taskName"
	logger.Info("failed to do task",
		// Structured context as strongly typed Field values.
		zap.String("url", task),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
} */

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
	// a.initLogger()
}

// Run app
func (a *App) Run() {
	a.router.Engine.Run(":8080")
}

// Run app
func (a *App) RunTLS() {
	a.router.Engine.Run(":443", "cert.crt", "key.key")
}
