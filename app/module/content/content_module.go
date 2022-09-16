package content

type ContentModule struct {
	ContentHandler    *ContentHandler
	ContentService    *ContentService
	ContentRepository *ContentRepository
}

func (m *ContentModule) InitComponents() {
	m.ContentRepository = new(ContentRepository)
	m.ContentService = NewContentService(
		m.ContentRepository,
	)
	m.ContentHandler = NewContentHandler(
		m.ContentService,
	)
}
