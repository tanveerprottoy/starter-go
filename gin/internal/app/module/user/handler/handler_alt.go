package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/dto"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/user/service"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/adapter"
	httpPkg "github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
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
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](ctx.Request.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	h.service.Create(d, ctx)
}

func (h *HandlerAlt) ReadMany(ctx *gin.Context) {
	limit := 10
	page := 1
	var err error
	limitStr := httpPkg.GetQueryParam(ctx, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, ctx)
			return
		}
	}
	pageStr := httpPkg.GetQueryParam(ctx, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, ctx)
			return
		}
	}
	h.service.ReadMany(limit, page, ctx)
}

func (h *HandlerAlt) ReadManyWithNestedDocQuery(ctx *gin.Context) {
	limit := 10
	page := 1
	var err error
	limitStr := httpPkg.GetQueryParam(ctx, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, ctx)
			return
		}
	}
	pageStr := httpPkg.GetQueryParam(ctx, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, ctx)
			return
		}
	}
	k0 := httpPkg.GetQueryParam(ctx, "key0")
	k1 := httpPkg.GetQueryParam(ctx, "key1")
	h.service.ReadManyWithNestedDocQuery(limit, page, k0, k1, ctx)
}

func (h *HandlerAlt) ReadOne(ctx *gin.Context) {
	id := httpPkg.GetURLParam(ctx, constant.KeyId)
	h.service.ReadOne(id, ctx)
}

func (h *HandlerAlt) Update(ctx *gin.Context) {
	id := httpPkg.GetURLParam(ctx, constant.KeyId)
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

func (h *HandlerAlt) Delete(ctx *gin.Context) {
	id := httpPkg.GetURLParam(ctx, constant.KeyId)
	h.service.Delete(id, ctx)
}
