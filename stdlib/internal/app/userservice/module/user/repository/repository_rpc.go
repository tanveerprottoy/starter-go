package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/proto"
	sqlpkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql/mysqlpkg"
)

type RepositoryRPC struct {
}

func (r *RepositoryRPC) Create(e *proto.User) (string, error) {
	var lastId string = ""
	result, err := mysqlpkg.GetInstance().DB.Exec(
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
	rows, err := mysqlpkg.GetInstance().DB.Query(
		"SELECT * FROM users", // WHERE id IS NOT NULL
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *RepositoryRPC) ReadOne(id string) *sql.Row {
	row := mysqlpkg.GetInstance().DB.QueryRow(
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
	res, err := mysqlpkg.GetInstance().DB.Exec(
		q,
		id,
		e.Name,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return sqlpkg.GetRowsAffected(res), nil
}

func (r *RepositoryRPC) Delete(id string) (int64, error) {
	q := "DELETE FROM users WHERE id = $1"
	res, err := mysqlpkg.GetInstance().DB.Exec(
		q,
		id,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return sqlpkg.GetRowsAffected(res), nil
}
