package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tanveerprottoy/starter-go/internal/app/module/user/entity"
	sqlUtil "github.com/tanveerprottoy/starter-go/pkg/data/sql"
)

type Repository[T entity.User] struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository[entity.User] {
	r := new(Repository[entity.User])
	r.db = db
	return r
}

func (r *Repository[T]) Create(e *entity.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (name)"+
			"VALUES ($1)",
		e.Name,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *Repository[T]) ReadMany(limit, offset int) (*sql.Rows, error) {
	rows, err := r.db.Query(
		"SELECT * FROM users LIMIT $1 OFFSET $2", // WHERE id IS NOT NULL
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *Repository[T]) ReadOne(id string) *sql.Row {
	row := r.db.QueryRow(
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return row
}

func (r *Repository[T]) Update(id string, e *entity.User) (int64, error) {
	q := "UPDATE users SET name = $2 WHERE id = $1"
	res, err := r.db.Exec(
		q,
		id,
		e.Name,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return sqlUtil.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM users WHERE id = $1"
	res, err := r.db.Exec(
		q,
		id,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return sqlUtil.GetRowsAffected(res), nil
}
