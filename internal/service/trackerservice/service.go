package trackerservice

import (
	"go-time-tracker/internal/adapter/pgsql"
)

type Service struct {
	timeCostStorage pgsql.TimeCost
}

func NewTrackerService(storage pgsql.TimeCost) *Service {
	return &Service{storage}
}
