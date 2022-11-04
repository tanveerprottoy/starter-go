package app

import (
	"log"
	"net/http"
	"txp/restapistarter/app/module/content"
	"txp/restapistarter/app/module/user"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/data/sql/postgres"
	"txp/restapistarter/pkg/router"
)

// global var
var (
	Configs       map[string]interface{}
	UserModule    *user.UserModule
	ContentModule *content.ContentModule
	DBClient      *mongodb.DBClient
)

// App struct
type App struct {
	DBClient *mongodb.DBClient
	router   *router.Router
}

func (a *App) initDB() {
	postgres.InitDBClient()
	a.DBClient = mongodb.NewDBClient()
}

func (a *App) initModules() {
	UserModule = user.NewUserModule(a.DBClient.DB, a.router)
	ContentModule = content.NewContentModule(a.DBClient.DB, a.router)
}

// Init app
func (a *App) InitComponents() {
	a.initDB()
	a.initModules()
	a.router = router.NewRouter()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		"127.0.0.1:8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
