package content

import (
	"txp/restapistarter/pkg/router"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContentModule struct {
	Router     *ContentRouter
	Handler    *ContentHandler
	Service    *ContentService
	Repository *ContentRepository
}

func NewContentModule(db *mongo.Database, router *router.Router) *ContentModule {
	m := new(ContentModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = new(ContentRepository)
	m.Service = NewContentService(m.Repository)
	m.Handler = NewContentHandler(m.Service)
	m.Router = NewContentRouter(router, m)
	return m
}
