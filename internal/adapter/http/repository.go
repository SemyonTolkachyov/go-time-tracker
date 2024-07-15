package http

import (
	"go-time-tracker/internal/model"
	"net/http"
	"net/url"
)

// User http repository interface
type User interface {
	GetUserInfo(passportNumber string) (*model.UserInfo, error)
}

// Client with base url struct
type Client struct {
	*http.Client
	BaseURL *url.URL
}
