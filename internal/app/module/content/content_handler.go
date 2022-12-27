package content

import (
	"net/http"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/router"
)

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

func (h *ContentHandler) Create(w http.ResponseWriter, r *http.Request) {
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r)
}

func (h *ContentHandler) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := router.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := router.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	h.service.ReadMany(limit, page, w, r)
}

func (h *ContentHandler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *ContentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, p, w, r)
}

func (h *ContentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
