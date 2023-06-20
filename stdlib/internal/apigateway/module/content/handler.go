package content

import (
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/adapter"
	httpPkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/http"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	h := new(Handler)
	h.service = service
	return h
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r)
}

func (h *Handler) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httpPkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httpPkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	h.service.ReadMany(limit, page, w, r)
}

func (h *Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httpPkg.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httpPkg.GetURLParam(r, constant.KeyId)
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, p, w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httpPkg.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
