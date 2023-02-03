package content

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/app/module/content/entity"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/internal/pkg/constant"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/adapter"
	sqlPkg "github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/data/sql"
	httpPkg "github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/http"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/response"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/time"
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
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
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
	d, err := sqlPkg.GetEntities(
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
	d, err := sqlPkg.GetEntity(
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
	userId := httpPkg.GetURLParam(ctx, constant.KeyId)
	d, err := adapter.BytesToType[entity.Content](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
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
