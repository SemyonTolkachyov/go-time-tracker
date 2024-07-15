package service

import (
	"context"
	"go-time-tracker/internal/entity"
	"go-time-tracker/internal/model"
	"go-time-tracker/internal/model/filter"
	"go-time-tracker/internal/model/input"
)

type UserService interface {
	Create(ctx context.Context, input input.NewUser) (int, error)
	Update(ctx context.Context, input input.UpdateUser) (int, error)
	Delete(ctx context.Context, userId int) (int, error)
	Get(ctx context.Context, userId int) (*entity.User, error)
	GetByFilter(ctx context.Context, filter filter.User) (*[]entity.User, error)
	GetPagedByFilter(ctx context.Context, size, number int, filter filter.User) (*[]entity.User, error)
}

type TrackerService interface {
	Start(ctx context.Context, userId, taskId int) error
	Stop(ctx context.Context, userId, taskId int) error
	GetUserCostsByPeriod(ctx context.Context, input input.TimeCostsByPeriod) (*[]model.CostData, error)
}

type Service struct {
	UserService
	TrackerService
}

func NewService(userService UserService, trackerService TrackerService) *Service {
	return &Service{userService, trackerService}
}
