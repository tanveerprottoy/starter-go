package errorpkg

import (
	"errors"
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/constant"
)

func NewError(m string) error {
	return errors.New(m)
}

func HandleDBError(err error) *HTTPError {
	httpErr := &HTTPError{Code: http.StatusBadRequest, Err: err}
	if err.Error() == "sql: no rows in result set" {
		httpErr.Code = http.StatusNotFound
		httpErr.Err = NewError(constant.NotFound)
	}
	return httpErr
}
