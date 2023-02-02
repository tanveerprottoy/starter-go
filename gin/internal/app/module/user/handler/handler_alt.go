package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/dto"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/service"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/adapter"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/router"
)

type HandlerAlt struct {
	service *service.ServiceAlt
}

func NewHandlerAlt(s *service.ServiceAlt) *HandlerAlt {
	h := new(HandlerAlt)
	h.service = s
	return h
}

func (h *HandlerAlt) Create(ctx *gin.Context) {
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](ctx.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err)
		return
	}
	h.service.Create(d, ctx)
}

func (h *HandlerAlt) ReadMany(ctx *gin.Context) {
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

func (h *HandlerAlt) ReadManyWithNestedDocQuery(ctx *gin.Context) {
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
	k0 := router.GetQueryParam(ctx, "key0")
	k1 := router.GetQueryParam(ctx, "key1")
	h.service.ReadManyWithNestedDocQuery(limit, page, k0, k1, ctx)
}

func (h *HandlerAlt) ReadOne(ctx *gin.Context) {
	id := router.GetURLParam(ctx, constant.KeyId)
	h.service.ReadOne(id, ctx)
}

func (h *HandlerAlt) Update(ctx *gin.Context) {
	id := router.GetURLParam(ctx, constant.KeyId)
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](ctx.Request.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, err)
		return
	}
	h.service.Update(id, d, ctx)
}

func (h *HandlerAlt) Delete(ctx *gin.Context) {
	id := router.GetURLParam(ctx, constant.KeyId)
	h.service.Delete(id, ctx)
}
