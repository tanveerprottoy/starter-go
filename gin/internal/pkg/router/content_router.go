package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/content"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
)

func RegisterContentRoutes(e *gin.Engine, version string, module *content.Module) {
	routes := e.Group(constant.ApiPattern + version + constant.ContentsPattern)
	routes.GET(constant.RootPattern, module.Handler.ReadMany)
	routes.GET(constant.RootPattern+"{id}", module.Handler.ReadOne)
	routes.POST(constant.RootPattern, module.Handler.Create)
	routes.PATCH(constant.RootPattern+"{id}", module.Handler.Update)
	routes.DELETE(constant.RootPattern+"{id}", module.Handler.Delete)
}
