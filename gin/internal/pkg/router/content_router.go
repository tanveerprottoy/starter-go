package router

import (
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/content"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/router"
)

func RegisterContentRoutes(router *router.Router, version string, module *content.Module) {
	routes := router.Engine.Group(constant.ApiPattern + version + constant.ContentsPattern)
	routes.Get(constant.RootPattern, module.Handler.ReadMany)
	routes.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
	routes.Post(constant.RootPattern, module.Handler.Create)
	routes.Patch(constant.RootPattern+"{id}", module.Handler.Update)
	routes.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
}
