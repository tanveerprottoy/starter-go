package content

import (
	"database/sql"
)

type Module struct {
	Handler    *Handler
	Service    *Service
}

func NewModule(db *sql.DB) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Service = NewService()
	m.Handler = NewHandler(m.Service)
	return m
}
