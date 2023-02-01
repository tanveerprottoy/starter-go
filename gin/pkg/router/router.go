package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/middleware"
)

// Router struct
type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	r := new(Router)
	r.Engine = gin.Default()
	r.registerGlobalMiddlewares()
	return r
}

func (r *Router) registerGlobalMiddlewares() {
	r.Engine.Use(
		middleware.JSONContentTypeMiddleWare(),
		middleware.CORSEnableMiddleWare(),
	)
}
