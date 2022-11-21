package content

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ContentModule struct {
	Handler    *ContentHandler
	Service    *ContentService
	Repository *ContentRepository
}

func NewContentModule(db *mongo.Database) *ContentModule {
	m := new(ContentModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = new(ContentRepository)
	m.Service = NewContentService(m.Repository)
	m.Handler = NewContentHandler(m.Service)
	return m
}
