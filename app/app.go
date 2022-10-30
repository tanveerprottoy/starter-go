package app

import (
	"log"
	"net/http"
	"txp/restapistarter/app/module/content"
	"txp/restapistarter/app/module/user"
	"txp/restapistarter/pkg/data/nosql/mongodb"
	"txp/restapistarter/pkg/data/sql/postgres"
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
	router   *Router
}

func (a *App) initDB() {
	postgres.InitDBClient()
	a.DBClient = mongodb.NewDBClient()
}

func (a *App) initModules() {
	UserModule = new(user.UserModule)
	UserModule.InitComponents(a.DBClient.DB)
	ContentModule = new(content.ContentModule)
	ContentModule.InitComponents()
}

// Init app
func (a *App) InitComponents() {
	a.initDB()
	a.initModules()
	a.router = NewRouter()
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
