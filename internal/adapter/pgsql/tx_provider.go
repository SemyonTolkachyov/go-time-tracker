package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/apperror"
)

// Connector for sqlx database.
type Connector interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

// TxProvider for manage transactions
type TxProvider struct {
	conn Connector
}

func NewTxProvider(conn Connector) *TxProvider {
	return &TxProvider{
		conn: conn,
	}
}

// AcquireWithOpts transaction from db with options.
func (t *TxProvider) AcquireWithOpts(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	log.Infof("Acquire transaction from db with opts=%v", opts)
	tx, err := t.conn.BeginTxx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Tx{
		Context: context.WithValue(ctx, txKeyType{}, Access(tx)),
		Tx:      tx,
	}, nil
}

// Acquire transaction from db.
func (t *TxProvider) Acquire(ctx context.Context) (*Tx, error) {
	return t.AcquireWithOpts(ctx, &DefaultTxOpts)
}

// TxWithOpts runs fn in transaction with options.
func (t *TxProvider) TxWithOpts(ctx context.Context, opts *sql.TxOptions, fn func(TxContext) error) error {
	log.Infof("Run sql tx with opts=%v", opts)
	tx, err := t.AcquireWithOpts(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			log.Warnf("Recovering from panic in TxWithOpts error is: %v", r)
			_ = tx.Rollback()
			err, _ = r.(error)
		} else if err != nil {
			err = tx.Rollback()
			var txErr *pq.Error
			ok := errors.As(err, &txErr)
			if ok && txErr.Code == "40001" {
				err = apperror.NewBadRequestError("transaction in progress")
			}
		} else {
			err = tx.Commit()
		}

		if ctx.Err() != nil && errors.Is(err, context.DeadlineExceeded) {
			log.Errorf("query response time exceeded the configured timeout")
		}
	}()

	err = fn(tx)

	return err
}

// Tx runs fn in transaction.
func (t *TxProvider) Tx(ctx context.Context, fn func(TxContext) error) error {
	return t.TxWithOpts(ctx, &DefaultTxOpts, fn)
}
