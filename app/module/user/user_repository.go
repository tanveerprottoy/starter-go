package user

import (
	"database/sql"
	"fmt"
	"log"
	"txp/restapistarter/app/module/user/entity"
	"txp/restapistarter/pkg/data"
)

type UserRepository struct {
}

func (r *UserRepository) Create(e *entity.User) error {
	_, err := data.DB.Exec(
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

func (r *UserRepository) ReadMany() (*sql.Rows, error) {
	rows, err := data.DB.Query(
		"SELECT * FROM users", // WHERE id IS NOT NULL
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *UserRepository) ReadOne(id string) *sql.Row {
	row := data.DB.QueryRow(
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return row
}

func (r *UserRepository) Update(id string, e *entity.User) (int64, error) {
	q := "UPDATE users SET name = $2 WHERE id = $1"
	res, err := data.DB.Exec(
		q,
		id,
		e.Name,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}

func (r *UserRepository) Delete(id string) (int64, error) {
	q := "DELETE FROM users WHERE id = $1"
	res, err := data.DB.Exec(
		q,
		id,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}
