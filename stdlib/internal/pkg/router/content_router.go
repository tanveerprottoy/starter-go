package router

import (
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/apigateway/module/content"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"

	"github.com/go-chi/chi"
)

func RegisterContentRoutes(router *Router, version string, module *content.Module) {
	router.Mux.Route(
		constant.ApiPattern+version+constant.ContentsPattern,
		func(r chi.Router) {
			r.Get(constant.RootPattern, module.Handler.ReadMany)
			r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
			r.Post(constant.RootPattern, module.Handler.Create)
			r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
			r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
		},
	)
}
