package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/user/dto"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/user/service"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/pkg/adapter"
	"github.com/tanveerprottoy/starter-go/gin/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/jsonpkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/response"
)

type HandlerAlt struct {
	service *service.ServiceAlt
	validate *validator.Validate
}

func NewHandlerAlt(s *service.ServiceAlt, v *validator.Validate) *HandlerAlt {
	h := new(HandlerAlt)
	h.service = s
	h.validate = v
	return h
}

func (h *HandlerAlt) parseValidateRequestBody(r *http.Request) (dto.CreateUpdateUserDto, error) {
	var d dto.CreateUpdateUserDto
	err := jsonpkg.Decode(r.Body, &d)
	if err != nil {
		return d, err
	}
	// validate request body
	err = h.validate.Struct(d)
	if err != nil {
		// Struct is invalid
		// for checking only
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	}
	return d, err
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

func (h *HandlerAlt) ReadManyWithNestedDocQuery(ctx *gin.Context) {
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
	k0 := httppkg.GetQueryParam(ctx, "key0")
	k1 := httppkg.GetQueryParam(ctx, "key1")
	h.service.ReadManyWithNestedDocQuery(limit, page, k0, k1, ctx)
}

func (h *HandlerAlt) ReadOne(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	h.service.ReadOne(id, ctx)
}

func (h *HandlerAlt) Update(ctx *gin.Context) {
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

func (h *HandlerAlt) Delete(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	h.service.Delete(id, ctx)
}
