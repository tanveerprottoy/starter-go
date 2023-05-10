package content

import (
	"database/sql"
)

type Module struct {
	RPC        *RPC
	Service    *Service
	Repository *Repository
}

func NewModule(db *sql.DB) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewRPC(m.Service)
	return m
}
