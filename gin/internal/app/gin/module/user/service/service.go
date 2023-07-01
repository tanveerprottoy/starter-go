package service

import (
	"github.com/tanveerprottoy/starter-go/gin/internal/app/module/user/dto"
	"github.com/tanveerprottoy/starter-go/gin/internal/app/module/user/entity"
	"github.com/tanveerprottoy/starter-go/gin/pkg/data/sql/sqlxpkg"
)

type Service struct {
	repository sqlxpkg.Repository[entity.User]
}

func NewService(r sqlxpkg.Repository[entity.User]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(d *dto.CreateUpdateUserDto, ctx *gin.Context) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	v.CreatedAt = timepkg.Now()
	v.UpdatedAt = timepkg.Now()
	err = s.repository.Create(v)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), ctx)
}

func (s *Service) ReadMany(limit, page int, ctx *gin.Context) {
	offset := limit * (page - 1)
	rows, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		// log err
		response.Respond(http.StatusOK, make([]any, 0), ctx)
		return
	}
	var e entity.User
	d, err := sqlPkg.GetEntities(
		rows,
		&e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
	if err != nil {
		// log err
		response.Respond(http.StatusOK, make([]any, 0), ctx)
		return
	}
	m := make(map[string]any)
	m["items"] = d
	m["limit"] = limit
	m["page"] = page
	response.Respond(http.StatusOK, response.BuildData(m), ctx)
}

func (s *Service) ReadOneInternal(id string) *sql.Row {
	return s.repository.ReadOne(id)
}

func (s *Service) ReadOne(id string, ctx *gin.Context) {
	row := s.ReadOneInternal(id)
	if row == nil {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	e := new(entity.User)
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
	response.Respond(http.StatusOK, response.BuildData(d), ctx)
}

func (s *Service) Update(id string, d *dto.CreateUpdateUserDto, ctx *gin.Context) {
	v, err := adapter.AnyToType[entity.User](d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	v.UpdatedAt = timepkg.Now()
	rowsAffected, err := s.repository.Update(id, v)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(d), ctx)
}

func (s *Service) Delete(id string, ctx *gin.Context) {
	rowsAffected, err := s.repository.Delete(id)
	if err != nil || rowsAffected <= 0 {
		response.RespondError(http.StatusInternalServerError, errors.New(constant.InternalServerError), ctx)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(map[string]bool{"success": true}), ctx)
}
