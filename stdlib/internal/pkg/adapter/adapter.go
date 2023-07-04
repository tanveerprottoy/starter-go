package adapter

import (
	"database/sql"

	"github.com/tanveerprottoy/starter-go/stdlib/internal/app/userservice/module/user/entity"
	sqlpkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/data/sql"
)

func RowToUserEntity(row *sql.Row) (*entity.User, error) {
	e := new(entity.User)
	return sqlpkg.GetEntity(
		row,
		e,
		&e.Id,
		&e.Name,
		&e.CreatedAt,
		&e.UpdatedAt,
	)
}
