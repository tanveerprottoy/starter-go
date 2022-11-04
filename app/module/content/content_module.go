package content

import (
	"txp/restapistarter/pkg/router"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContentModule struct {
	ContentRouter     *ContentRouter
	ContentHandler    *ContentHandler
	ContentService    *ContentService
	ContentRepository *ContentRepository
}

func NewContentModule(db *mongo.Database, router *router.Router) *ContentModule {
	m := new(ContentModule)
	m.ContentRouter = NewContentRouter(router, m)
	m.ContentRepository = new(ContentRepository)
	m.ContentService = NewContentService(m.ContentRepository)
	m.ContentHandler = NewContentHandler(m.ContentService)
	return m
}
