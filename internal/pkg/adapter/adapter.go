package adapter

import (
	"database/sql"

	"txp/restapistarter/internal/app/module/user/entity"
	_sql "txp/restapistarter/pkg/data/sql"
)

func RowToUserEntity(row *sql.Row) (*entity.User, error) {
	e := new(entity.User)
	return _sql.GetEntity(
		row,
		e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
}
