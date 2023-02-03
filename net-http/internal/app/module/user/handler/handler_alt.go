package handler

import (
	"net/http"

	"github.com/tanveerprottoy/rest-api-starter-go/net-http/internal/app/module/user/dto"
	"github.com/tanveerprottoy/rest-api-starter-go/net-http/internal/app/module/user/service"
	"github.com/tanveerprottoy/rest-api-starter-go/net-http/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/net-http/pkg/adapter"
	httpPkg "github.com/tanveerprottoy/rest-api-starter-go/net-http/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/net-http/pkg/response"
)

type HandlerAlt struct {
	service *service.ServiceAlt
}

func NewHandlerAlt(s *service.ServiceAlt) *HandlerAlt {
	h := new(HandlerAlt)
	h.service = s
	return h
}

func (h *HandlerAlt) Create(w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(d, w, r)
}

func (h *HandlerAlt) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httpPkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httpPkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	h.service.ReadMany(limit, page, w, r)
}

func (h *HandlerAlt) ReadManyWithNestedDocQuery(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httpPkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httpPkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	k0 := httpPkg.GetQueryParam(r, "key0")
	k1 := httpPkg.GetQueryParam(r, "key1")
	h.service.ReadManyWithNestedDocQuery(limit, page, k0, k1, w, r)
}

func (h *HandlerAlt) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httpPkg.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *HandlerAlt) Update(w http.ResponseWriter, r *http.Request) {
	id := httpPkg.GetURLParam(r, constant.KeyId)
	d, err := adapter.BodyToType[dto.CreateUpdateUserDto](r.Body)
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

func (h *HandlerAlt) Delete(w http.ResponseWriter, r *http.Request) {
	id := httpPkg.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
