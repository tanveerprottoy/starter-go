package repository

import (
	"database/sql"
	"fmt"
	"log"
	"txp/restapistarter/internal/app/module/user/entity"
	sqlUtil "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/data/sql/postgres"
)

type UserRepository[T entity.User] struct {
}

func (r *UserRepository[T]) Create(e *entity.User) error {
	_, err := postgres.DB.Exec(
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

func (r *UserRepository[T]) ReadMany() (*sql.Rows, error) {
	rows, err := postgres.DB.Query(
		"SELECT * FROM users", // WHERE id IS NOT NULL
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *UserRepository[T]) ReadOne(id string) *sql.Row {
	row := postgres.DB.QueryRow(
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return row
}

func (r *UserRepository[T]) Update(id string, e *entity.User) (int64, error) {
	q := "UPDATE users SET name = $2 WHERE id = $1"
	res, err := postgres.DB.Exec(
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

func (r *UserRepository[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM users WHERE id = $1"
	res, err := postgres.DB.Exec(
		q,
		id,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return sqlUtil.GetRowsAffected(res), nil
}
