package content

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {
	h := new(Handler)
	h.service = service
	return h
}

func (h *Handler) UploadOne(ctx *gin.Context) {
	/* p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r) */
}

func (h *Handler) UploadMany(ctx *gin.Context) {
	/* p, err := adapter.IOReaderToBytes(r.Body)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(p, w, r) */
}
