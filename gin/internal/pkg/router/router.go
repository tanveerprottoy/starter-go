package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterGlobalMiddlewares(e *gin.Engine) {
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	e.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	e.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	// e.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := e.Group("/", AuthRequired())
	// exactly the same as:
	// authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	// authorized.Use(AuthRequired())
}
