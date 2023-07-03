package gin

import (
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/auth"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/content"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/user"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/middleware"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/router"
	"github.com/tanveerprottoy/starter-go/gin/pkg/data/nosql/mongodb"
	"github.com/tanveerprottoy/starter-go/gin/pkg/data/sql/postgres"

	"github.com/go-playground/validator/v10"
	// "go.uber.org/zap"
)

// App struct
type App struct {
	MongoDBClient    *mongodb.Client
	PostgresDBClient *postgres.Client
	gin              *pkg.Gin
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
	router.RegisterUserRoutes(a.gin.Engine, constant.V1, a.UserModule, m)
	router.RegisterContentRoutes(a.gin.Engine, constant.V1, a.ContentModule)
}

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.gin = pkg.NewGin()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	a.gin.Engine.Run(":8080")
}

// Run app
func (a *App) RunTLS() {
	a.gin.Engine.Run(":443", "cert.crt", "key.key")
}
