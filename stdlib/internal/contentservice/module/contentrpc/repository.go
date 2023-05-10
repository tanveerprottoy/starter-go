package content

import (
	"database/sql"
	"fmt"
	"log"
	"txp/contentservice/app/module/content/proto"
	"txp/contentservice/pkg/data"
)

type Repository struct {
}

func (r *Repository) Create(
	e *proto.Content,
) (string, error) {
	var lastId string = ""
	result, err := data.DB.Exec(
		"INSERT INTO contents (name)"+
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

func (r *Repository) ReadMany() (*sql.Rows, error) {
	rows, err := data.DB.Query(
		"SELECT * FROM contents", // WHERE id IS NOT NULL
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *Repository) ReadOne(id string) *sql.Row {
	row := data.DB.QueryRow(
		"SELECT * FROM contents WHERE id = $1 LIMIT 1",
		id,
	)
	return row
}

func (r *Repository) Update(
	id string,
	e *proto.Content,
) (int64, error) {
	q := "UPDATE contents SET name = $2 WHERE id = $1"
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

func (r *Repository) Delete(id string) (int64, error) {
	q := "DELETE FROM contents WHERE id = $1"
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
