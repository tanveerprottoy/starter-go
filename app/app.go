package app

import (
	"log"
	"net/http"
	"txp/restapistarter/app/module/content"
	"txp/restapistarter/app/module/user"
)

// global var
var (
	UserModule    *user.UserModule
	ContentModule *content.ContentModule
)

// App struct
type App struct {
	router *Router
}

func (a *App) initModules() {
	UserModule = new(user.UserModule)
	UserModule.InitComponents()
	ContentModule = new(content.ContentModule)
	ContentModule.InitComponents()
}

// Init app
func (a *App) InitComponents() {
	a.initModules()
	a.router = NewRouter()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		"0.0.0.0:3000",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
