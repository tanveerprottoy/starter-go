package user

import (
	"txp/restapistarter/app/util"
	"txp/restapistarter/pkg/router"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	router *router.Router
}

func NewUserRouter(router *router.Router, module *UserModule) *UserRouter {
	r := new(UserRouter)
	r.router = router
	r.registerRoutes(util.V1, module)
	return r
}

func (r *UserRouter) registerRoutes(version string, module *UserModule) {
	r.router.Mux.Route(
		util.ApiPattern+version+util.UsersPattern,
		func(r chi.Router) {
			r.Get(
				util.RootPattern,
				module.UserHandler.ReadMany,
			)
			r.Get(
				util.RootPattern+"{id}",
				module.UserHandler.ReadOne,
			)
			r.Post(
				util.RootPattern,
				module.UserHandler.Create,
			)
			r.Patch(
				util.RootPattern+"{id}",
				module.UserHandler.Update,
			)
			r.Delete(
				util.RootPattern+"{id}",
				module.UserHandler.Delete,
			)
		},
	)
}
