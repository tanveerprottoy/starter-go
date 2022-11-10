package content

import (
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/router"

	"github.com/go-chi/chi"
)

type ContentRouter struct {
	router *router.Router
}

func NewContentRouter(router *router.Router, module *ContentModule) *ContentRouter {
	r := new(ContentRouter)
	r.router = router
	r.registerRoutes(constant.V1, module)
	return r
}

func (r *ContentRouter) registerRoutes(version string, module *ContentModule) {
	r.router.Mux.Route(
		constant.ApiPattern+version+constant.ContentsPattern,
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
