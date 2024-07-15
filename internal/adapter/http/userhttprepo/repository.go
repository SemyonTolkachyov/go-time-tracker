package userhttprepo

import (
	"go-time-tracker/internal/adapter/http"
)

// Repository http user repo
type Repository struct {
	*http.Client
}

func NewUserRepository(client *http.Client) *Repository {
	return &Repository{
		client,
	}
}
