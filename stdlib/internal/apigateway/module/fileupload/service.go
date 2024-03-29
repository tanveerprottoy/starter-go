package fileupload

import (
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/s3pkg"
)

type Service struct {
	s3client *s3pkg.Client
}

func NewService(s3client *s3pkg.Client) *Service {
	s := new(Service)
	s.s3client = s3client
	return s
}

func (s *Service) UploadOne(p []byte, w http.ResponseWriter, r *http.Request) {
	/* d, err := adapter.BytesToValue[entity.Content](p)
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
	response.Respond(http.StatusCreated, d, w) */
}

func (s *Service) UploadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	/* rows, err := s.repository.ReadMany()
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
	response.Respond(http.StatusOK, d, w) */
}
