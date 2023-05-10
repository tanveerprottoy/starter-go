package user

import (
	"database/sql"
	"fmt"
	"log"
	"txp/userservice/pkg/data"
	"txp/userservice/app/module/user/proto"
)

type RepositoryRPC struct {
}

func (r *RepositoryRPC) Create(e *proto.User) (string, error) {
	var lastId string = ""
	result, err := data.DB.Exec(
		"INSERT INTO users (name)"+
			"VALUES ($1) RETURNING id",
		e.Name,
	)
	if err != nil {
		log.Println(err)
		return lastId, err
	}
	if err != nil {
		log.Println(err)
	}
	temp, _ := result.LastInsertId()
	lastId = fmt.Sprintf("%d", temp)
	return lastId, nil
}

func (r *RepositoryRPC) ReadMany() (*sql.Rows, error) {
	rows, err := data.DB.Query(
		"SELECT * FROM users", // WHERE id IS NOT NULL
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *RepositoryRPC) ReadOne(id string) *sql.Row {
	row := data.DB.QueryRow(
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return row
}

func (r *RepositoryRPC) Update(
	id string,
	e *proto.User,
) (int64, error) {
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

func (r *RepositoryRPC) Delete(id string) (int64, error) {
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
