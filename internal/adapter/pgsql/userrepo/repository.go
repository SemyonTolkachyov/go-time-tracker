package userrepo

import (
	"github.com/jmoiron/sqlx"
	"go-time-tracker/internal/adapter/pgsql"
)

// Repository of users for pgsql
type Repository struct {
	db *sqlx.DB
	*pgsql.TxProvider
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db, TxProvider: pgsql.NewTxProvider(db)}
}