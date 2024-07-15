package userservice

import (
	"go-time-tracker/internal/adapter/http"
	"go-time-tracker/internal/adapter/pgsql"
)

type Service struct {
	userStorage pgsql.User
	httpSource  http.User
}

func NewUserService(storage pgsql.User, httpSource http.User) *Service {
	return &Service{storage, httpSource}
}
