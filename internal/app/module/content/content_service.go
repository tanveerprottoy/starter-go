package content

import (
	"encoding/json"
	"errors"
	"net/http"
	"txp/restapistarter/internal/app/module/content/dto"
	"txp/restapistarter/internal/app/module/content/entity"
	"txp/restapistarter/internal/pkg/constant"
	sqlcore "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/response"

	"github.com/go-chi/chi"
)

type ContentService struct {
	repository *ContentRepository
}

func NewContentService(
	repository *ContentRepository,
) *ContentService {
	s := new(ContentService)
	s.repository = repository
	return s
}

func (s *ContentService) Create(w http.ResponseWriter, r *http.Request) {
	var b dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		&entity.Content{
			Name: b.Name,
		},
	)
	if err != nil {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	response.Respond(http.StatusCreated, b, w)
}

func (s *ContentService) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
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
	d, err := sqlcore.GetEntities(
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

func (s *ContentService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, constant.UrlKeyId)
	row := s.repository.ReadOne(userId)
	if row == nil {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	e := new(entity.Content)
	d, err := sqlcore.GetEntity(
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

func (s *ContentService) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, constant.UrlKeyId)
	var b dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response.RespondError(
			http.StatusBadRequest,
			err,
			w,
		)
		return
	}
	rowsAffected, err := s.repository.Update(
		userId,
		&entity.Content{
			Name: b.Name,
		},
	)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(
			http.StatusInternalServerError,
			errors.New(constant.InternalServerError),
			w,
		)
		return
	}
	response.Respond(http.StatusOK, b, w)
}

func (s *ContentService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, constant.UrlKeyId)
	rowsAffected, err := s.repository.Delete(userId)
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
