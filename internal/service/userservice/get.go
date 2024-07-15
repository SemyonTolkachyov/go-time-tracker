package userservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/entity"
)

func (u Service) Get(ctx context.Context, userId int) (*entity.User, error) {
	log.Infof("Getting user with id %d", userId)
	user, err := u.userStorage.Get(ctx, userId)
	if err != nil {
		log.Errorf("Error getting user with id %d: %v", userId, err)
		return nil, err
	}
	return user, nil
}
