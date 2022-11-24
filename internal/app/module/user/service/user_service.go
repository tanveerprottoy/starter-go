package service

import (
	_sql "database/sql"
	"errors"
	"net/http"
	"txp/restapistarter/internal/app/module/user/entity"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/time"
)

type UserService struct {
	repository sql.Repository[entity.User]
}

func NewUserService(r sql.Repository[entity.User]) *UserService {
	s := new(UserService)
	s.repository = r
	return s
}

func (s *UserService) Create(p []byte, w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BytesToValue[entity.User](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	err = s.repository.Create(d)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), w)
}

func (s *UserService) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	offset := limit * (page - 1)
	rows, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		// log err
		response.Respond(http.StatusOK, make([]any, 0), w)
		return
	}
	var e entity.User
	d, err := sql.GetEntities(
		rows,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		// log err
		response.Respond(http.StatusOK, make([]any, 0), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (s *UserService) ReadOneInternal(id string) *_sql.Row {
	return s.repository.ReadOne(id)
}

func (s *UserService) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	row := s.ReadOneInternal(id)
	if row == nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	e := new(entity.User)
	d, err := sql.GetEntity(
		row,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (s *UserService) Update(id string, p []byte, w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BytesToValue[entity.User](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d.UpdatedAt = time.Now()
	rowsAffected, err := s.repository.Update(
		id,
		d,
	)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (s *UserService) Delete(id string, w http.ResponseWriter, r *http.Request) {
	rowsAffected, err := s.repository.Delete(id)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(map[string]bool{"success": true}), w)
}
