package mysql

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

// Exec executes a query without returning any rows.
func Exec(query string, args ...any) (sql.Result, error) {
	return DB.Exec(query, args...)
}

func ExecContext(
	ctx context.Context,
	query string,
	args ...any,
) (sql.Result, error) {
	return DB.ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func Query(query string, args ...any) (*sql.Rows, error) {
	return DB.Query(query, args)
}

func QueryContext(
	ctx context.Context,
	query string,
	args ...any,
) (*sql.Rows, error) {
	return DB.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func QueryRow(query string, args ...any) *sql.Row {
	return DB.QueryRow(query, args)
}

func QueryRowContext(
	ctx context.Context,
	query string,
	args ...any,
) *sql.Row {
	return DB.QueryRowContext(ctx, query, args...)
}

// Prepare creates a prepared statement for later queries or executions.
func Prepare(query string) (*sql.Stmt, error) {
	return DB.Prepare(query)
}

func PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return DB.PrepareContext(ctx, query)
}
