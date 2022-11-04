package router

import (
	"txp/restapistarter/pkg/middleware"

	"github.com/go-chi/chi"
)

// Router struct
type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	r := &Router{}
	r.Mux = chi.NewRouter()
	r.registerMiddlewares()
	return r
}

func (r *Router) registerMiddlewares() {
	r.Mux.Use(
		middleware.JSONContentTypeMiddleWare,
	)
}
