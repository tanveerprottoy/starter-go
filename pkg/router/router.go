package router

import (
	"txp/restapistarter/pkg/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	)
}
