package router

import (
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/content"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/router"
)

func RegisterContentRoutes(router *router.Router, version string, module *content.Module) {
	routes := router.Engine.Group(constant.ApiPattern + version + constant.ContentsPattern)
	routes.GET(constant.RootPattern, module.Handler.ReadMany)
	routes.GET(constant.RootPattern+"{id}", module.Handler.ReadOne)
	routes.POST(constant.RootPattern, module.Handler.Create)
	routes.PATCH(constant.RootPattern+"{id}", module.Handler.Update)
	routes.DELETE(constant.RootPattern+"{id}", module.Handler.Delete)
}
