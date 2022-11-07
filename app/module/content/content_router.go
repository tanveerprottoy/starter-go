package content

import (
	"txp/restapistarter/app/util"
	"txp/restapistarter/pkg/router"

	"github.com/go-chi/chi"
)

type ContentRouter struct {
	router *router.Router
}

func NewContentRouter(router *router.Router, module *ContentModule) *ContentRouter {
	r := new(ContentRouter)
	r.router = router
	r.registerRoutes(util.V1, module)
	return r
}

func (r *ContentRouter) registerRoutes(version string, module *ContentModule) {
	r.router.Mux.Route(
		util.ApiPattern+version+util.ContentsPattern,
		func(r chi.Router) {
			r.Get(
				util.RootPattern,
				module.Handler.ReadMany,
			)
			r.Get(
				util.RootPattern+"{id}",
				module.Handler.ReadOne,
			)
			r.Post(
				util.RootPattern,
				module.Handler.Create,
			)
			r.Patch(
				util.RootPattern+"{id}",
				module.Handler.Update,
			)
			r.Delete(
				util.RootPattern+"{id}",
				module.Handler.Delete,
			)
		},
	)
}
