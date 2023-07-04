package service

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/dto"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/adapter"
	sqlpkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/timepkg"
)

type Service struct {
	repository sqlpkg.Repository[entity.User]
}

func NewService(r sqlpkg.Repository[entity.User]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	n := timepkg.NowUnixMilli()
	v.CreatedAt = n
	v.UpdatedAt = n
	err = s.repository.Create(v)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), w)
}

func (s *Service) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	offset := limit * (page - 1)
	rows, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		// log err
		response.Respond(http.StatusOK, make([]any, 0), w)
		return
	}
	var e entity.User
	d, err := sqlpkg.GetEntities(
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
	m := make(map[string]any)
	m["items"] = d
	m["limit"] = limit
	m["page"] = page
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (s *Service) ReadOneInternal(id string) *sql.Row {
	return s.repository.ReadOne(id)
}

func (s *Service) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	row := s.ReadOneInternal(id)
	if row == nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	e := new(entity.User)
	d, err := sqlpkg.GetEntity(
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

func (s *Service) Update(id string, d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	v.UpdatedAt = timepkg.NowUnixMilli()
	rowsAffected, err := s.repository.Update(
		id,
		v,
	)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), w)
}

func (s *Service) Delete(id string, w http.ResponseWriter, r *http.Request) {
	rowsAffected, err := s.repository.Delete(id)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(map[string]bool{"success": true}), w)
}
