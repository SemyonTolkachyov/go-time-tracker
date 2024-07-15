package userservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/model"
	"go-time-tracker/internal/model/input"
)

func (u Service) Create(ctx context.Context, input input.NewUser) (int, error) {
	log.Infof("Creating user %s", input)
	userInfo, err := u.httpSource.GetUserInfo(input.PassportNumber)
	if err != nil {
		log.Errorf("Error getting user info by passport number %s: %v", input.PassportNumber, err)
		return 0, err
	}
	log.Debugf("Get user info %s with passport number: %s", userInfo, input.PassportNumber)
	userData := model.UserData{
		Name:           userInfo.Name,
		Surname:        userInfo.Surname,
		Patronymic:     userInfo.Patronymic,
		Address:        userInfo.Address,
		PassportNumber: input.PassportNumber,
	}
	i, err := u.userStorage.Create(ctx, userData)
	if err != nil {
		log.Errorf("Creating user error by passport number %s: %v", input.PassportNumber, err)
	}
	return i, err
}
