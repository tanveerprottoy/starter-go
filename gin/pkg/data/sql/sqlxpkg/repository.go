package sqlxpkg

import (
	"github.com/jmoiron/sqlx"
)

type Repository[T any] interface {
	Create(e *T) error

	ReadMany() (*sqlx.Rows, error)

	ReadOne(id string) *sqlx.Row

	Update(id string, e *T) (int64, error)

	Delete(id string) (int64, error)
}
