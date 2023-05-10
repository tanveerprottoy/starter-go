package errorpkg

import "errors"

func NewError(m string) error {
	return errors.New(m)
}
