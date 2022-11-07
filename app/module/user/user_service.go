package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"txp/restapistarter/app/module/user/dto"
	"txp/restapistarter/app/module/user/entity"
	"txp/restapistarter/app/module/user/repository"
	"txp/restapistarter/app/util"
	"txp/restapistarter/pkg/coreutil"
	sqlUtil "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/responseutil"
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
		responseutil.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		&entity.User{
			Name: b.Name,
		},
	)
	if err != nil {
		responseutil.RespondError(http.StatusInternalServerError, errors.New(util.InternalServerError), w)
		return
	}
	responseutil.Respond(http.StatusCreated, responseutil.BuildData(b), w)
}

func (s *UserService) ReadMany(w http.ResponseWriter, r *http.Request) {
	rows, err := s.repository.ReadMany()
	if err != nil {
		responseutil.RespondError(http.StatusInternalServerError, err, w)
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
		responseutil.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	responseutil.Respond(http.StatusOK, responseutil.BuildData(d), w)
}

func (s *UserService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := coreutil.GetURLParam(r, util.UrlKeyId)
	row := s.repository.ReadOne(userId)
	if row == nil {
		responseutil.RespondError(http.StatusInternalServerError, errors.New(util.InternalServerError), w)
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
		responseutil.RespondError(http.StatusInternalServerError, err, w)
		return
	}
	responseutil.Respond(http.StatusOK, responseutil.BuildData(d), w)
}

func (s *UserService) Update(w http.ResponseWriter, r *http.Request) {
	userId := coreutil.GetURLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		responseutil.RespondError(http.StatusBadRequest, err, w)
		return
	}
	rowsAffected, err := s.repository.Update(
		userId,
		&entity.User{
			Name: b.Name,
		},
	)
	if err != nil || rowsAffected <= 0 {
		responseutil.RespondError(http.StatusInternalServerError, errors.New(util.InternalServerError), w)
		return
	}
	responseutil.Respond(http.StatusOK, responseutil.BuildData(b), w)
}

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := coreutil.GetURLParam(r, util.UrlKeyId)
	rowsAffected, err := s.repository.Delete(userId)
	if err != nil || rowsAffected <= 0 {
		responseutil.RespondError(http.StatusInternalServerError, errors.New(util.InternalServerError), w)
		return
	}
	responseutil.Respond(http.StatusOK, responseutil.BuildData(map[string]bool{"success": true}), w)
}
