package content

import (
	"database/sql"
)

type ContentModule struct {
	Handler    *ContentHandler
	Service    *ContentService
	Repository *ContentRepository
}

func NewContentModule(db *sql.DB) *ContentModule {
	m := new(ContentModule)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewContentRepository(db)
	m.Service = NewContentService(m.Repository)
	m.Handler = NewContentHandler(m.Service)
	return m
}
