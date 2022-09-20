package app

import (
	"log"
	"net/http"
	"os"
	"txp/restapistarter/app/module/content"
	"txp/restapistarter/app/module/user"
	"txp/restapistarter/pkg/util"
)

// global var
var (
	// configs
	Configs map[string]interface{}
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

func (a *App) initConfigs() {
	fileBytes, _ := os.ReadFile("../config/dev.json")
	_ = util.Unmarshal(fileBytes, &Configs)
	log.Print(Configs)
}

// Init app
func (a *App) InitComponents() {
	a.initConfigs()
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
