package userservice

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (u Service) Delete(ctx context.Context, userId int) (int, error) {
	log.Infof("Deleting user with id %d", userId)
	i, err := u.userStorage.Delete(ctx, userId)
	if err != nil {
		log.Errorf("Error deleting user with id %d: %v", userId, err)
		return 0, err
	}
	return i, nil
}
