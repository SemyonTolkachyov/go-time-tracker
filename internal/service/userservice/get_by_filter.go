package userservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/entity"
	"go-time-tracker/internal/model/filter"
)

func (u Service) GetByFilter(ctx context.Context, filter filter.User) (*[]entity.User, error) {
	log.Infof("getting user by filter: %s", filter)
	users, err := u.userStorage.GetByFilter(ctx, filter)
	if err != nil {
		log.Errorf("Error getting users by filter %s: %v", filter, err)
		return nil, err
	}
	return users, nil
}
