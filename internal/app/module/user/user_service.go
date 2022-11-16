package user

import (
	"errors"
	"net/http"
	_sql "database/sql"
	"txp/restapistarter/internal/app/module/user/entity"
	"txp/restapistarter/internal/pkg/constant"
	"txp/restapistarter/pkg/adapter"
	"txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/router"
)

type UserService struct {
	repository sql.Repository[entity.User] // repository.UserRepository[entity.User]
}

func NewUserService(r sql.Repository[entity.User]) *UserService {
	s := new(UserService)
	s.repository = r
	return s
}

func (s *UserService) Create(w http.ResponseWriter, r *http.Request) {
	/*var b dto.CreateUpdateUserDto
	err := json.Decode(r.Body, &b)
	d, err := adapter.AnyToValue[schema.UserSchema](b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	} */
	defer r.Body.Close()
	b, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d, err := adapter.BytesToValue[entity.User](b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		d,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(b), w)
}

func (s *UserService) ReadMany(w http.ResponseWriter, r *http.Request) {
	rows, err := s.repository.ReadMany()
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
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
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (s *UserService) ReadOneInternal(id string) *_sql.Row {
	return s.repository.ReadOne(id)
}

func (s *UserService) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.UrlKeyId)
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

func (s *UserService) Update(w http.ResponseWriter, r *http.Request) {
	id := router.GetURLParam(r, constant.UrlKeyId)
	defer r.Body.Close()
	b, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d, err := adapter.BytesToValue[entity.User](b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	rowsAffected, err := s.repository.Update(
		id,
		d,
	)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(b), w)
}

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := router.GetURLParam(r, constant.UrlKeyId)
	rowsAffected, err := s.repository.Delete(userId)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(map[string]bool{"success": true}), w)
}
