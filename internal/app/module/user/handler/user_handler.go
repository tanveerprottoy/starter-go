package handler

import (
	"net/http"
	"txp/restapistarter/internal/app/module/user/dto"
	"txp/restapistarter/internal/app/module/user/service"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/router"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(s *service.UserService, v *validator.Validate) *UserHandler {
	h := new(UserHandler)
	h.service = s
	h.validate = v
	return h
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BodyToValue[dto.CreateUpdateUserDto](r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(d, w, r)
}

func (h *UserHandler) ReadMany(w http.ResponseWriter, r *http.Request) {
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

func (h *UserHandler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	d, err := adapter.BodyToValue[dto.CreateUpdateUserDto](r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, d, w, r)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
