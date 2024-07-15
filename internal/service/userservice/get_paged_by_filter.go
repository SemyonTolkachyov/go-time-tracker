package userservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/entity"
	"go-time-tracker/internal/model/filter"
)

func (u Service) GetPagedByFilter(ctx context.Context, size, number int, filter filter.User) (*[]entity.User, error) {
	log.Infof("Getting users paged by filter %s, size=%d, pageNumber=%d", filter, size, number)
	offset := (number - 1) * size
	users, err := u.userStorage.GetPagedByFilter(ctx, offset, size, filter)
	if err != nil {
		log.Errorf("Error getting page with users number %d with size %d by filter %s: %v", number, size, filter, err)
		return nil, err
	}
	return users, nil
}
