package pgsql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// Tx represents transaction with context as inner object.
type Tx struct {
	context.Context
	Tx *sqlx.Tx
}

func (t *Tx) WithValue(key, value any) TxContext {
	return &Tx{
		Context: context.WithValue(t.Context, key, value),
		Tx:      t.Tx,
	}
}

// Prepare query.
func (t *Tx) Prepare(query string) (*sql.Stmt, error) {
	return t.Tx.PrepareContext(t.Context, query)
}

// Exec executes query with args.
func (t *Tx) Exec(query string, args ...any) (sql.Result, error) {
	return t.Tx.ExecContext(t.Context, query, args...)
}

// Query loads data from db.
func (t *Tx) Query(query string, args ...any) (*sql.Rows, error) {
	return t.Tx.QueryContext(t.Context, query, args...)
}

// QueryRow loads single row from db.
func (t *Tx) QueryRow(query string, args ...any) *sql.Row {
	return t.Tx.QueryRowContext(t.Context, query, args...)
}

// Select loads data in dest from db
func (t *Tx) Select(dest interface{}, query string, args ...any) error {
	return t.Tx.SelectContext(t.Context, dest, query, args...)
}

// Commit this transaction.
func (t *Tx) Commit() error {
	return t.Tx.Commit()
}

// Rollback cancel this transaction.
func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}
