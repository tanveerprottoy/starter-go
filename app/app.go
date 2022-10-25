package app

import (
	"log"
	"net/http"
	"txp/restapistarter/app/module/content"
	"txp/restapistarter/app/module/user"
	"txp/restapistarter/pkg/data/sql/postgres"
)

// global var
var (
	// configs
	Configs       map[string]interface{}
	UserModule    *user.UserModule
	ContentModule *content.ContentModule
)

// App struct
type App struct {
	router *Router
}

func (a *App) initDB() {
	postgres.InitDB()
}

func (a *App) initModules() {
	UserModule = new(user.UserModule)
	UserModule.InitComponents()
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
