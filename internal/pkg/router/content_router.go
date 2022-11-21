package router

import (
	"txp/restapistarter/internal/app/module/content"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/router"

	"github.com/go-chi/chi"
)

func RegisterContentRoutes(router *router.Router, version string, module *content.ContentModule) {
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
