package content

import (
	"errors"
	"net/http"
	"time"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/module/content/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/adapter"
	sqlPkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"

	"github.com/go-chi/chi"
)

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {
	s := new(Service)
	s.repository = repository
	return s
}

func (s *Service) Create(p []byte, w http.ResponseWriter, r *http.Request) {
	d, err := adapter.BytesToType[entity.Content](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	err = s.repository.Create(d)
	if err != nil {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	response.Respond(http.StatusCreated, d, w)
}

func (s *Service) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	rows, err := s.repository.ReadMany()
	if err != nil {
		response.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	var e entity.Content
	d, err := sqlPkg.GetEntities(
		rows,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		response.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	response.Respond(http.StatusOK, d, w)
}

func (s *Service) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	row := s.repository.ReadOne(id)
	if row == nil {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	e := new(entity.Content)
	d, err := sqlPkg.GetEntity(
		row,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		response.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	response.Respond(http.StatusOK, d, w)
}

func (s *Service) Update(id string, p []byte, w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, constant.KeyId)
	d, err := adapter.BytesToType[entity.Content](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	rowsAffected, err := s.repository.Update(userId, d)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	response.Respond(http.StatusOK, d, w)
}

func (s *Service) Delete(id string, w http.ResponseWriter, r *http.Request) {
	rowsAffected, err := s.repository.Delete(id)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	response.Respond(
		http.StatusOK,
		map[string]bool{"success": true},
		w,
	)
}
