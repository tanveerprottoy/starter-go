package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"txp/restapistarter/app/module/user/dto"
	"txp/restapistarter/internal/app/module/user/entity"
	"txp/restapistarter/internal/app/module/user/repository"
	"txp/restapistarter/internal/app/pkg/constant"
	sqlUtil "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/response"
	"txp/restapistarter/pkg/router"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	s := new(UserService)
	s.repository = repository
	return s
}

func (s *UserService) Create(w http.ResponseWriter, r *http.Request) {
	var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		&entity.User{
			Name: b.Name,
		},
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
	d, err := sqlUtil.GetEntities(
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

func (s *UserService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := router.GetURLParam(r, constant.UrlKeyId)
	row := s.repository.ReadOne(userId)
	if row == nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	e := new(entity.User)
	d, err := sqlUtil.GetEntity(
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
	userId := router.GetURLParam(r, constant.UrlKeyId)
	var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	rowsAffected, err := s.repository.Update(
		userId,
		&entity.User{
			Name: b.Name,
		},
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