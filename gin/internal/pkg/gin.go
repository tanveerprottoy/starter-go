package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/router"
)

// Gin Engine container struct
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
type Gin struct {
	Engine *gin.Engine
}

func NewGin() *Gin {
	g := &Gin{}
	g.Engine = gin.Default()
	router.RegisterGlobalMiddlewares(g.Engine)
	return g
}
