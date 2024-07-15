package app

import (
	"github.com/jmoiron/sqlx"
	"go-time-tracker/internal/adapter/http"
	"go-time-tracker/internal/adapter/http/userhttprepo"
	"go-time-tracker/internal/adapter/pgsql/timecost"
	"go-time-tracker/internal/adapter/pgsql/userrepo"
	"go-time-tracker/internal/service"
	"go-time-tracker/internal/service/trackerservice"
	"go-time-tracker/internal/service/userservice"
)

type Container struct {
	pgsql *sqlx.DB
	http  *http.Client
}

func NewContainer(pgSqlxConn *sqlx.DB, client *http.Client) *Container {

	return &Container{
		pgsql: pgSqlxConn,
		http:  client,
	}
}

func (c *Container) GetService() *service.Service {
	return service.NewService(c.getUserService(), c.getTrackerService())
}

func (c *Container) getPgsql() *sqlx.DB {
	return c.pgsql
}

func (c *Container) getHttpClient() *http.Client {
	return c.http
}

func (c *Container) getUserService() *userservice.Service {
	return userservice.NewUserService(c.getUserRepo(), c.getUserHttpRepo())
}

func (c *Container) getTrackerService() *trackerservice.Service {
	return trackerservice.NewTrackerService(c.getTimeCostRepo())
}

func (c *Container) getUserHttpRepo() *userhttprepo.Repository {
	return userhttprepo.NewUserRepository(c.getHttpClient())
}

func (c *Container) getUserRepo() *userrepo.Repository {
	return userrepo.NewRepository(c.getPgsql())
}

func (c *Container) getTimeCostRepo() *timecost.Repository {
	return timecost.NewRepository(c.getPgsql())
}
