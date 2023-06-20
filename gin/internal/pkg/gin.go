package pkg

import "github.com/gin-gonic/gin"

// Gin Engine container struct
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
type Gin struct {
	Engine *gin.Engine
}

func NewGin() *Gin {
	g := &Gin{}
	g.Engine = gin.Default()
	return g
}
