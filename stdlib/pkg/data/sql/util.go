package sql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/adapter"
)

func GetRowsAffected(result sql.Result) int64 {
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return rows
}

func GetEntities[T any](rows *sql.Rows, obj *T, params ...interface{}) ([]*T, error) {
	var entities []*T
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(params...); err != nil {
			return nil, fmt.Errorf("GetEntities %v", err)
		}
		adapter.ValuesToStruct(params, obj)
		entities = append(entities, obj)
	}
	return entities, nil
}

func GetEntity[T any](row *sql.Row, obj *T, params ...interface{}) (*T, error) {
	if err := row.Scan(params...); err != nil {
		return nil, fmt.Errorf("GetEntity %v", err)
	}
	adapter.ValuesToStruct(params, obj)
	return obj, nil
}
