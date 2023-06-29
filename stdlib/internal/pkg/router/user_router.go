package router

import (
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/apigateway/module/user"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/middleware"

	"github.com/go-chi/chi"
)

func RegisterUserRoutes(router *Router, version string, module *user.Module, authMiddleWare *middleware.AuthMiddleware) {
	router.Mux.Group(
		func(r chi.Router) {
			r.Use(authMiddleWare.AuthUser)
			r.Route(
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
