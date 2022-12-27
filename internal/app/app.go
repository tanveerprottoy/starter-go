package app

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"txp/restapistarter/internal/app/module/auth"
	"txp/restapistarter/internal/app/module/content"
	"txp/restapistarter/internal/app/module/user"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/internal/pkg/middleware"
	routerPkg "txp/restapistarter/internal/pkg/router"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/data/sql/postgres"
	"txp/restapistarter/pkg/router"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	validatorPkg "txp/restapistarter/pkg/validator"
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
	Validate         *validator.Validate
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
	a.UserModule = user.NewUserModule(a.MongoDBClient.DB, a.PostgresDBClient.DB, a.Validate)
	a.AuthModule = auth.NewAuthModule(a.UserModule.Service)
	a.ContentModule = content.NewContentModule(a.PostgresDBClient.DB)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	routerPkg.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
	routerPkg.RegisterContentRoutes(a.router, constant.V1, a.ContentModule)
}

func (a *App) initValidators() {
	a.Validate = validator.New()
	_ = a.Validate.RegisterValidation("notempty", validatorPkg.NotEmpty)
}

func (a *App) initLogger() {
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
}

// Init app
func (a *App) InitComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
	a.initValidators()
	a.initLogger()
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

// Run app
func (a *App) RunTLSSimpleConfig() {
	err := http.ListenAndServeTLS(":443", "cert.crt", "key.key", a.router.Mux)
	if err != nil {
		log.Fatal(err)
	}
}

// use mutual TLS and not just one-way TLS,
// we must instruct it to require client authentication to ensure clients present a certificate from our CA when they connect
func (a *App) RunTLSMutual() {
	caCert, _ := ioutil.ReadFile("ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
		Handler:   a.router.Mux,
	}
	server.ListenAndServeTLS("cert.crt", "key.key")
}
