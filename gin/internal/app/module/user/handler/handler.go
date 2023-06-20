package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/module/user/dto"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/module/user/service"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/pkg/adapter"
	"github.com/tanveerprottoy/starter-go/gin/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/response"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(s *service.Service, v *validator.Validate) *Handler {
	h := new(Handler)
	h.service = s
	h.validate = v
	return h
}

func (h *Handler) Create(ctx *gin.Context) {
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](ctx.Request.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	h.service.Create(d, ctx)
}

func (h *Handler) ReadMany(ctx *gin.Context) {
	limit := 10
	page := 1
	var err error
	limitStr := httppkg.GetQueryParam(ctx, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, ctx)
			return
		}
	}
	pageStr := httppkg.GetQueryParam(ctx, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, ctx)
			return
		}
	}
	h.service.ReadMany(limit, page, ctx)
}

func (h *Handler) ReadOne(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	h.service.ReadOne(id, ctx)
}

func (h *Handler) Update(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](ctx.Request.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	h.service.Update(id, d, ctx)
}

func (h *Handler) Delete(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	h.service.Delete(id, ctx)
}
