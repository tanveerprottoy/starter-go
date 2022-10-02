package content

import (
	"encoding/json"
	"errors"
	"net/http"
	"txp/restapistarter/app/module/content/dto"
	"txp/restapistarter/app/module/content/entity"
	sqlUtil "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/util"

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
	var b *dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		&entity.Content{
			Name: b.Name,
		},
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	util.Respond(http.StatusCreated, b, w)
}

func (s *ContentService) ReadMany(w http.ResponseWriter, r *http.Request) {
	rows, err := s.repository.ReadMany()
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	var e entity.Content
	d, err := sqlUtil.GetEntities(
		rows,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(http.StatusOK, d, w)
}

func (s *ContentService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	row := s.repository.ReadOne(userId)
	if row == nil {
		util.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	e := new(entity.Content)
	d, err := sqlUtil.GetEntity(
		row,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(http.StatusOK, d, w)
}

func (s *ContentService) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(
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
		util.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	util.Respond(http.StatusOK, b, w)
}

func (s *ContentService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	rowsAffected, err := s.repository.Delete(userId)
	if err != nil || rowsAffected <= 0 {
		util.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		map[string]bool{"success": true},
		w,
	)
}
