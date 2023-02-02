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
	p, err := adapter.IOReaderToBytes(ctx.Request.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err)
		return
	}
	h.service.Create(p, ctx)
}

func (h *Handler) ReadMany(ctx *gin.Context) {
	limit := 10
	page := 1
	var err error
	limitStr := router.GetQueryParam(ctx, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err)
			return
		}
	}
	pageStr := router.GetQueryParam(ctx, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err)
			return
		}
	}
	h.service.ReadMany(limit, page, ctx)
}

func (h *Handler) ReadOne(ctx *gin.Context) {
	id := router.GetURLParam(ctx, constant.KeyId)
	h.service.ReadOne(id, ctx)
}

func (h *Handler) Update(ctx *gin.Context) {
	id := router.GetURLParam(ctx, constant.KeyId)
	p, err := adapter.IOReaderToBytes(ctx.Request.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err)
		return
	}
	h.service.Update(id, p, ctx)
}

func (h *Handler) Delete(ctx *gin.Context) {
	id := router.GetURLParam(ctx, constant.KeyId)
	h.service.Delete(id, ctx)
}
