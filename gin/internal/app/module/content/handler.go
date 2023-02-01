package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/adapter"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/router"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {
	h := new(Handler)
	h.service = service
	return h
}

func (h *Handler) Create(ctx *gin.Context) {
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r)
}

func (h *Handler) ReadMany(ctx *gin.Context) {
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

func (h *Handler) ReadOne(ctx *gin.Context) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *Handler) Update(ctx *gin.Context) {
	id := router.GetURLParam(r, constant.KeyId)
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, p, w, r)
}

func (h *Handler) Delete(ctx *gin.Context) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
