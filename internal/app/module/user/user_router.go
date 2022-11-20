package user

import (
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/internal/pkg/middleware"
	"txp/restapistarter/pkg/router"

	"github.com/go-chi/chi"
)

type UserRouter struct {
	router         *router.Router
	authMiddleWare *middleware.AuthMiddleware
}

func NewUserRouter(
	router *router.Router,
	module *UserModule,
	authMiddleWare *middleware.AuthMiddleware,
) *UserRouter {
	r := new(UserRouter)
	r.router = router
	r.authMiddleWare = authMiddleWare
	r.registerRoutes(constant.V1, module)
	return r
}

func (r *UserRouter) registerRoutes(version string, module *UserModule) {
	r.router.Mux.Group(
		func(router chi.Router) {
			router.Use(r.authMiddleWare.AuthUser)
			r.router.Mux.Route(
				constant.ApiPattern+version+constant.UsersPattern,
				func(r chi.Router) {
					r.Get(constant.RootPattern, module.Handler.ReadMany)
					r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
					r.Post(constant.RootPattern, module.Handler.Create)
					r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
					r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
				},
			)
		},
	)
}
