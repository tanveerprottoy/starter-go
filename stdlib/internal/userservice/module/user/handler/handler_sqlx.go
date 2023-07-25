package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/userservice/module/user/dto"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/userservice/module/user/service"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/adapter"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/jsonpkg"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
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

func (h *HandlerSqlx) parseValidateRequestBody(r *http.Request) (dto.CreateUpdateUserDto, error) {
	var d dto.CreateUpdateUserDto
	err := jsonpkg.Decode(r.Body, &d)
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

func (h *HandlerSqlx) Create(w http.ResponseWriter, r *http.Request) {
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	ctx := r.Context()
	e, httpErr := h.service.Create(&d, ctx)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
		return
	}
	response.Respond(http.StatusCreated, e, w)
}

func (h *HandlerSqlx) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httppkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httppkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	e, httpErr := h.service.ReadMany(limit, page, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *HandlerSqlx) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.ReadOne(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *HandlerSqlx) Update(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	e, httpErr := h.service.Update(id, &d, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *HandlerSqlx) Delete(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.Delete(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}
