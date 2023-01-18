package errors

import "errors"

func NewError(m string) error {
	return errors.New(m)
}
