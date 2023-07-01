package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/entity"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql/sqlxpkg"
)

const TableName = "users"

type RepositorySqlx[T entity.User] struct {
	db *sqlx.DB
}

func NewRepositorySqlx(db *sqlx.DB) *RepositorySqlx[entity.User] {
	r := new(RepositorySqlx[entity.User])
	r.db = db
	return r
}

func (r *RepositorySqlx[T]) Create(e *entity.User) error {
	var lastId string
	err := r.db.QueryRow("INSERT INTO "+TableName+" (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastId)
	if err != nil {
		return err
	}
	e.Id = lastId
	return nil
}

func (r *RepositorySqlx[T]) ReadMany(limit, offset int) ([]entity.User, error) {
	d := []entity.User{}
	err := r.db.Select(&d, "SELECT * FROM "+TableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *RepositorySqlx[T]) ReadOne(id string) (entity.User, error) {
	b := entity.User{}
	err := r.db.Get(&b, "SELECT * FROM "+TableName+" WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r *RepositorySqlx[T]) Update(id string, e *entity.User) (int64, error) {
	q := "UPDATE " + TableName + " SET name = $2, updated_at = $3 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Name, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}

func (r *RepositorySqlx[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM " + TableName + " WHERE id = $1"
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}
