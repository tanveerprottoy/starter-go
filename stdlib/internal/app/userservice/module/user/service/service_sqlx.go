package service

import (
	"errors"
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/dto"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql/sqlxpkg"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/response"
)

type ServiceSqlx struct {
	repository sqlxpkg.Repository[entity.User]
}

func NewServiceSqlx(r sqlxpkg.Repository[entity.User]) *ServiceSqlx {
	s := new(ServiceSqlx)
	s.repository = r
	return s
}

func (s *ServiceSqlx) readOneInternal(id string, w http.ResponseWriter) (entity.User, error) {
	return s.repository.ReadOne(id)
}

func (s *ServiceSqlx) Create(d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	// convert dto to entity
	b := entity.User{}
	b.Name = d.Name
	b. = d.Author
	b.PublicationYear = d.PublicationYear
	n := time.Now().UnixMilli()
	b.CreatedAt = n
	b.UpdatedAt = n
	err := s.repository.Create(&b)
	if err != nil {
		s.handleError(err, w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), w)
}

func (s *ServiceSqlx) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		s.handleError(err, w)
		return
	}
	m := make(map[string]any)
	m["items"] = d
	m["limit"] = limit
	m["page"] = page
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (s *ServiceSqlx) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		s.handleError(err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(b), w)
}

func (s *ServiceSqlx) Update(id string, d *dto.CreateUpdateUserDto, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		s.handleError(err, w)
		return
	}
	b.Title = d.Title
	b.Author = d.Author
	b.PublicationYear = d.PublicationYear
	b.UpdatedAt = time.Now().UnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		s.handleError(err, w)
		return
	}
	if rows > 0 {
		response.Respond(http.StatusOK, response.BuildData(b), w)
		return
	}
	response.RespondError(http.StatusBadRequest, errors.New("operation was not successful"), w)
}

func (s *ServiceSqlx) Delete(id string, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		s.handleError(err, w)
		return
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		s.handleError(err, w)
		return
	}
	if rows > 0 {
		response.Respond(http.StatusOK, response.BuildData(b), w)
		return
	}
	response.RespondError(http.StatusBadRequest, errors.New("operation was not successful"), w)
}
