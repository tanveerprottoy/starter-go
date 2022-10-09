package app

import (
	"txp/restapistarter/app/util"
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
	r.registerUserRoutes(
		util.V1,
	)
	r.registerContentRoutes(
		util.V1,
	)
	return r
}

func (r *Router) registerMiddlewares() {
	r.Mux.Use(
		middleware.JSONContentTypeMiddleWare,
	)
}

func (r *Router) registerUserRoutes(
	version string,
) {
	r.Mux.Route(
		util.ApiPattern+version+util.UsersPattern,
		func(r chi.Router) {
			r.Get(
				util.RootPattern,
				UserModule.UserHandler.ReadMany,
			)
			r.Get(
				util.RootPattern+"{id}",
				UserModule.UserHandler.ReadOne,
			)
			r.Post(
				util.RootPattern,
				UserModule.UserHandler.Create,
			)
			r.Patch(
				util.RootPattern+"{id}",
				UserModule.UserHandler.Update,
			)
			r.Delete(
				util.RootPattern+"{id}",
				UserModule.UserHandler.Delete,
			)
		},
	)
}

func (r *Router) registerContentRoutes(
	version string,
) {
	r.Mux.Route(
		util.ApiPattern+version+util.ContentsPattern,
		func(r chi.Router) {
			r.Get(
				util.RootPattern,
				ContentModule.ContentHandler.ReadMany,
			)
			r.Get(
				util.RootPattern+"{id}",
				ContentModule.ContentHandler.ReadOne,
			)
			r.Post(
				util.RootPattern,
				ContentModule.ContentHandler.Create,
			)
			r.Patch(
				util.RootPattern+"{id}",
				ContentModule.ContentHandler.Update,
			)
			r.Delete(
				util.RootPattern+"{id}",
				ContentModule.ContentHandler.Delete,
			)
		},
	)
}
