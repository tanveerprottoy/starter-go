package content

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/gin/module/content/entity"
	"github.com/tanveerprottoy/starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/starter-go/gin/pkg/adapter"
	sqlpkg "github.com/tanveerprottoy/starter-go/gin/pkg/data/sql"
	"github.com/tanveerprottoy/starter-go/gin/pkg/httppkg"
	"github.com/tanveerprottoy/starter-go/gin/pkg/response"
	"github.com/tanveerprottoy/starter-go/gin/pkg/timepkg"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	s := new(Service)
	s.repository = repository
	return s
}

func (s *Service) Create(p []byte, ctx *gin.Context) {
	d, err := adapter.BytesToType[entity.Content](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	d.CreatedAt = timepkg.Now()
	d.UpdatedAt = timepkg.Now()
	err = s.repository.Create(d)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	response.Respond(http.StatusCreated, d, ctx)
}

func (s *Service) ReadMany(limit, page int, ctx *gin.Context) {
	rows, err := s.repository.ReadMany()
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	var e entity.Content
	d, err := sqlpkg.GetEntities(
		rows,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	response.Respond(http.StatusOK, d, ctx)
}

func (s *Service) ReadOne(id string, ctx *gin.Context) {
	row := s.repository.ReadOne(id)
	if row == nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	e := new(entity.Content)
	d, err := sqlpkg.GetEntity(
		row,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}
	response.Respond(http.StatusOK, d, ctx)
}

func (s *Service) Update(id string, p []byte, ctx *gin.Context) {
	userId := httppkg.GetURLParam(ctx, constant.KeyId)
	d, err := adapter.BytesToType[entity.Content](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	d.CreatedAt = timepkg.Now()
	d.UpdatedAt = timepkg.Now()
	rowsAffected, err := s.repository.Update(userId, d)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	response.Respond(http.StatusOK, d, ctx)
}

func (s *Service) Delete(id string, ctx *gin.Context) {
	rowsAffected, err := s.repository.Delete(id)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	response.Respond(http.StatusOK, map[string]bool{"success": true}, ctx)
}
