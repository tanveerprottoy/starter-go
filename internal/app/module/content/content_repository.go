package content

import (
	"database/sql"
	"fmt"
	"log"
	"txp/restapistarter/app/module/content/entity"
	sqlUtil "txp/restapistarter/pkg/data/sql"
	"txp/restapistarter/pkg/data/sql/postgres"
)

type ContentRepository struct {
}

func (r *ContentRepository) Create(e *entity.Content) error {
	_, err := postgres.DB.Exec(
		"INSERT INTO contents (name)"+
			"VALUES ($1)",
		e.Name,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *ContentRepository) ReadMany() (*sql.Rows, error) {
	rows, err := postgres.DB.Query(
		"SELECT * FROM contents", // WHERE id IS NOT NULL
	)
	if err != nil {
		return nil, fmt.Errorf("ReadMany %v", err)
	}
	return rows, nil
}

func (r *ContentRepository) ReadOne(id string) *sql.Row {
	row := postgres.DB.QueryRow(
		"SELECT * FROM contents WHERE id = $1 LIMIT 1",
		id,
	)
	return row
}

func (r *ContentRepository) Update(id string, e *entity.Content) (int64, error) {
	q := "UPDATE contents SET name = $2 WHERE id = $1"
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

func (r *ContentRepository) Delete(id string) (int64, error) {
	q := "DELETE FROM contents WHERE id = $1"
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