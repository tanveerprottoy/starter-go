package adapter

import (
	"database/sql"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/module/user/entity"
	_sql "github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql"
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
