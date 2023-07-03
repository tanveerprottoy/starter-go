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

// Hanlder is responsible for extracting data
// from request body and building and seding response
type HandlerSqlx struct {
	service  *service.ServiceSqlx
	validate *validator.Validate
}

func NewHandlerSqlx(s *service.ServiceSqlx, v *validator.Validate) *HandlerSqlx {
	h := new(HandlerSqlx)
	h.service = s
	h.validate = v
	return h
}

func (h *HandlerSqlx) parseValidateRequestBody(ctx *gin.Context) (dto.CreateUpdateUserDto, error) {
	var d dto.CreateUpdateUserDto
	err := jsonpkg.Decode(ctx.Request.Body, &d)
	if err != nil {
		return d, err
	}
	// validate request body
	err = h.validate.Struct(d)
	if err != nil {
		// Struct is invalid
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	}
	return d, err
}

func (h *HandlerSqlx) Create(ctx *gin.Context) {
	d, err := h.parseValidateRequestBody(ctx)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	e, httpErr := h.service.Create(&d, ctx)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, ctx)
		return
	}
	response.Respond(http.StatusCreated, e, ctx)
}

func (h *HandlerSqlx) ReadMany(ctx *gin.Context) {
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
	e, httpErr := h.service.ReadMany(limit, page, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, ctx)
	}
	response.Respond(http.StatusOK, e, ctx)
}

func (h *HandlerSqlx) ReadOne(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	e, httpErr := h.service.ReadOne(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, ctx)
	}
	response.Respond(http.StatusOK, e, ctx)
}

func (h *HandlerSqlx) Update(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	d, err := h.parseValidateRequestBody(ctx)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	e, httpErr := h.service.Update(id, &d, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, ctx)
	}
	response.Respond(http.StatusOK, e, ctx)
}

func (h *HandlerSqlx) Delete(ctx *gin.Context) {
	id := httppkg.GetURLParam(ctx, constant.KeyId)
	e, httpErr := h.service.Delete(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, ctx)
	}
	response.Respond(http.StatusOK, e, ctx)
}
