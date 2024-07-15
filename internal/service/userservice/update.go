package userservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/model/input"
)

func (u Service) Update(ctx context.Context, input input.UpdateUser) (int, error) {
	log.Infof("Updating user %s", input)
	err := input.Validate()
	if err != nil {
		log.Errorf("User validation error for input %s: %v", input, err)
		return 0, err
	}
	i, err := u.userStorage.Update(ctx, input)
	if err != nil {
		log.Errorf("User update error with user id=%d: %v", input.Id, err)
	}
	return i, err
}
