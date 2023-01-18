package sqlx

import "database/sql"

type Repository[T any] interface {
	Create(e *T) error

	ReadMany() (*sql.Rows, error)

	ReadOne(id string) *sql.Row

	Update(id string, e *T) (int64, error)

	Delete(id string) (int64, error)
}
