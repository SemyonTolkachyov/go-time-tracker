package pgsql

import (
	"context"
	"database/sql"
)

// Transactor interface
type Transactor interface {
	Tx(ctx context.Context, fn func(TxContext) error) error
	TxWithOpts(ctx context.Context, opts *sql.TxOptions, fn func(TxContext) error) error
	Acquire(ctx context.Context) (*Tx, error)
	AcquireWithOpts(ctx context.Context, opts *sql.TxOptions) (*Tx, error)
}

// TxContext interface for DAO operations with context.
type TxContext interface {
	context.Context
	WithValue(key, value any) TxContext
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

// Access interface for simple DML operations.
type Access interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

// DefaultTxOpts is package variable with default transaction level
var DefaultTxOpts = sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  false,
}

type txKeyType struct{}

// FromCtxOrData returns access interface from context or data arg.
func FromCtxOrData(ctx context.Context, data Access) Access {
	value, ok := ctx.Value(txKeyType{}).(Access)
	if ok {
		return value
	}
	return data
}
