package content

import "net/http"

type ContentHandler struct {
	service *ContentService
}

func NewContentHandler(
	service *ContentService,
) *ContentHandler {
	h := new(ContentHandler)
	h.service = service
	return h
}

func (h *ContentHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.Create(
		w,
		r,
	)
}

func (h *ContentHandler) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.ReadMany(
		w,
		r,
	)
}

func (h *ContentHandler) ReadOne(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.ReadOne(
		w,
		r,
	)
}

func (h *ContentHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.Update(
		w,
		r,
	)
}

func (h *ContentHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.Delete(
		w,
		r,
	)
}
