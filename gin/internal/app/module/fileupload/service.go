package content

import "github.com/gin-gonic/gin"

type Service struct {
}

func NewService() *Service {
	s := new(Service)
	return s
}

func (s *Service) UploadOne(p []byte, ctx *gin.Context) {
	/* d, err := adapter.BytesToValue[entity.Content](p)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err)
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

func (s *Service) UploadMany(limit, page int, ctx *gin.Context) {
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
