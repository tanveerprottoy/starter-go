package content

import (
	"encoding/json"
	"errors"
	"net/http"
	"txp/restapistarter/app/module/content/dto"
	"txp/restapistarter/app/module/content/entity"
	"txp/restapistarter/app/util"
	"txp/restapistarter/pkg/coreutil"
	sqlcoreutil "txp/restapistarter/pkg/data/sql"

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
		coreutil.RespondError(http.StatusBadRequest, err, w)
		return
	}
	err = s.repository.Create(
		&entity.Content{
			Name: b.Name,
		},
	)
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	coreutil.Respond(http.StatusCreated, b, w)
}

func (s *ContentService) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
	rows, err := s.repository.ReadMany()
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	var e entity.Content
	d, err := sqlcoreutil.GetEntities(
		rows,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	coreutil.Respond(http.StatusOK, d, w)
}

func (s *ContentService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	row := s.repository.ReadOne(userId)
	if row == nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	e := new(entity.Content)
	d, err := sqlcoreutil.GetEntity(
		row,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		coreutil.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	coreutil.Respond(http.StatusOK, d, w)
}

func (s *ContentService) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		coreutil.RespondError(
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
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	coreutil.Respond(http.StatusOK, b, w)
}

func (s *ContentService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	rowsAffected, err := s.repository.Delete(userId)
	if err != nil || rowsAffected <= 0 {
		coreutil.RespondError(
			http.StatusInternalServerError,
			errors.New(util.InternalServerError),
			w,
		)
		return
	}
	coreutil.Respond(
		http.StatusOK,
		map[string]bool{"success": true},
		w,
	)
}
