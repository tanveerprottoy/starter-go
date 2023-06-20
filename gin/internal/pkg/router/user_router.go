package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/module/user"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/middleware"
)

func RegisterUserRoutes(e *gin.Engine, version string, module *user.Module, authMiddleWare *middleware.AuthMiddleware) {
	routes := e.Group(constant.ApiPattern + version + constant.UsersPattern)
	routes.Use(authMiddleWare.AuthUser())
	routes.GET(constant.RootPattern, module.Handler.ReadMany)
	routes.GET(constant.RootPattern+"{id}", module.Handler.ReadOne)
	routes.POST(constant.RootPattern, module.Handler.Create)
	routes.PATCH(constant.RootPattern+"{id}", module.Handler.Update)
	routes.DELETE(constant.RootPattern+"{id}", module.Handler.Delete)
}
