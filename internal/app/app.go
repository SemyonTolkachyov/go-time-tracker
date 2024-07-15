package app

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/http"
	"go-time-tracker/internal/config"
	"sync"
)

type App struct {
	cfg config.Config

	c     *Container
	cOnce *sync.Once

	mysql  *sql.DB
	pgsqlx *sqlx.DB
	http   *http.Client
}

var a *App

func NewApp() (*App, error) {
	log.Debug("Creating new app")
	log.Debug("Loading config")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Errorf("Error loading config: %s", err)
		return nil, err
	}

	app := &App{
		cOnce: &sync.Once{},
		cfg:   cfg,
	}
	log.Debug("Config is loaded")
	log.Debug("Creating pgsql connection")
	pgSqlxConn, err := app.newPgSqlxConnect(cfg.Db)
	if err != nil {
		log.Errorf("Error creating pgsql connection: %s", err)
		return nil, err
	}
	app.pgsqlx = pgSqlxConn
	log.Debug("Created pgsql connection")

	log.Debug("Creating if database not exists")
	err = app.createIfNotExistsDbPg(cfg.Db)
	if err != nil {
		log.Errorf("Error creating database: %s", err)
		return nil, err
	}
	log.Debug("Database exist")

	log.Debug("Creating http client")
	httpClient, err := app.newHttpClient(cfg.UserApi)
	if err != nil {
		log.Errorf("Error creating http client: %s", err)
		return nil, err
	}
	app.http = httpClient
	log.Debug("Created http client")

	log.Debug("Creating di container")
	app.c = NewContainer(app.pgsqlx, app.http)
	log.Debug("Created container")

	return app, nil
}

func SetGlobalApp(app *App) {
	a = app
}

func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}

	return a, nil
}

func InitApp() error {
	initLogger()
	log.Info("Init app ...")
	app, err := NewApp()
	if err != nil {
		log.Fatalf("Fail to create app: %s", err)
		return err
	}
	log.Info("Init app success")
	log.Info("Up available migrations")
	err = app.upMigrationPgSqlx()
	if err != nil {
		log.Fatalf("Fail to up migrations: %s", err)
		return err
	}
	log.Info("Up available migrations success")
	SetGlobalApp(app)

	return nil
}
