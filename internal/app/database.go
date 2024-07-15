package app

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/config"
	_ "go-time-tracker/migrations"
	"strings"
)

// newPgSqlxConnect create and return sqlx pg connection
func (a *App) newPgSqlxConnect(cfg config.SQLConfig) (*sqlx.DB, error) {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("host=%s port=%s ", cfg.Host, cfg.Port))
	builder.WriteString(fmt.Sprintf("user=%s password=%s ", cfg.User, cfg.Password))
	builder.WriteString(fmt.Sprintf("dbname=%s ", cfg.DBName))
	builder.WriteString(fmt.Sprintf("sslmode=%s ", cfg.SslMode))

	params := builder.String()

	log.Debugf("Connecting to postgresql with params=%s", params)

	db, err := sqlx.Open("postgres", params)
	if err != nil {
		log.Errorf("Error connecting to postgresql with params=%s", params)
		return nil, err
	}

	return db, err
}

// createIfNotExistsDbPg init database if does not exist
func (a *App) createIfNotExistsDbPg(cfg config.SQLConfig) error {
	_, err := a.pgsqlx.Exec("SELECT VERSION()")
	if err != nil {
		var pgErr *pq.Error
		ok := errors.As(err, &pgErr)
		if ok {
			if pgErr.Code == "3D000" {
				log.Info("Creating postgresql database")
				err := a.initPgDb(cfg)
				if err != nil {
					log.Errorf("Error initializing postgresql with err: %s", err)
					return err
				}
			}
		}
	}
	return nil
}

// initPgDb create new db
func (a *App) initPgDb(cfg config.SQLConfig) error {
	dbName := cfg.DBName
	cfg.DBName = "postgres"
	db, err := a.newPgSqlxConnect(cfg)
	if err != nil {
		return err
	}
	defer db.Close()
	q := fmt.Sprintf(`CREATE DATABASE %s`, dbName)
	_, err = db.Exec(q)
	return err
}

// upMigrationPgSqlx up all available migrations
func (a *App) upMigrationPgSqlx() error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(a.pgsqlx.DB, "."); err != nil {
		return err
	}
	return nil
}
