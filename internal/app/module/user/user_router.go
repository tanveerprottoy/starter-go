package user

import (
	"txp/restapistarter/internal/app/pkg/constant"
	"txp/restapistarter/pkg/router"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	router *router.Router
}

func NewUserRouter(router *router.Router, module *UserModule) *UserRouter {
	r := new(UserRouter)
	r.router = router
	r.registerRoutes(constant.V1, module)
	return r
}

func (r *UserRouter) registerRoutes(version string, module *UserModule) {
	r.router.Mux.Route(
		constant.ApiPattern+version+constant.UsersPattern,
		func(r chi.Router) {
			r.Get(
				constant.RootPattern,
				module.Handler.ReadMany,
			)
			r.Get(
				constant.RootPattern+"{id}",
				module.Handler.ReadOne,
			)
			r.Post(
				constant.RootPattern,
				module.Handler.Create,
			)
			r.Patch(
				constant.RootPattern+"{id}",
				module.Handler.Update,
			)
			r.Delete(
				constant.RootPattern+"{id}",
				module.Handler.Delete,
			)
		},
	)
}
