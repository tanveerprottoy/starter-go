package sqlx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Exec executes a query without returning any rows.
func Exec(db *sqlx.DB, query string, args ...any) (sql.Result, error) {
	return db.Exec(query, args...)
}

func ExecContext(
	db *sqlx.DB,
	ctx context.Context,
	query string,
	args ...any,
) (sql.Result, error) {
	return db.ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func Query(db *sqlx.DB, query string, args ...any) (*sql.Rows, error) {
	return db.Query(query, args)
}

func QueryContext(
	db *sqlx.DB,
	ctx context.Context,
	query string,
	args ...any,
) (*sql.Rows, error) {
	return db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func QueryRow(db *sqlx.DB, query string, args ...any) *sql.Row {
	return db.QueryRow(query, args)
}

func QueryRowContext(
	db *sqlx.DB,
	ctx context.Context,
	query string,
	args ...any,
) *sql.Row {
	return db.QueryRowContext(ctx, query, args...)
}

// Prepare creates a prepared statement for later queries or executions.
func Prepare(db *sqlx.DB, query string) (*sql.Stmt, error) {
	return db.Prepare(query)
}

func PrepareContext(db *sqlx.DB, ctx context.Context, query string) (*sql.Stmt, error) {
	return db.PrepareContext(ctx, query)
}
