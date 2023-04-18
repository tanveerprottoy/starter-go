package router

import (
	"github.com/tanveerprottoy/starter-go/internal/app/module/user"
	"github.com/tanveerprottoy/starter-go/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/internal/pkg/middleware"
	"github.com/tanveerprottoy/starter-go/pkg/router"

	"github.com/go-chi/chi"
)

func RegisterUserRoutes(router *router.Router, version string, module *user.Module, authMiddleWare *middleware.AuthMiddleware) {
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
