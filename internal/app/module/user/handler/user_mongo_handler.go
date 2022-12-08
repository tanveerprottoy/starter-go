package handler

import (
	"net/http"
	"txp/restapistarter/internal/app/module/user/service"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/router"
)

type UserMongoHandler struct {
	service *service.UserMongoService
}

func NewUserMongoHandler(s *service.UserMongoService) *UserMongoHandler {
	h := new(UserMongoHandler)
	h.service = s
	return h
}

func (h *UserMongoHandler) Create(w http.ResponseWriter, r *http.Request) {
	/*var b dto.CreateUpdateUserDto
	err := json.Decode(r.Body, &b)
	d, err := adapter.AnyToValue[schema.UserSchema](b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	} */
	defer r.Body.Close()
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r)
}

func (h *UserMongoHandler) ReadMany(w http.ResponseWriter, r *http.Request) {
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

func (h *UserMongoHandler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *UserMongoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	defer r.Body.Close()
	p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, p, w, r)
}

func (h *UserMongoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}