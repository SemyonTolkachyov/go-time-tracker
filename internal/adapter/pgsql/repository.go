package pgsql

import (
	"context"
	"go-time-tracker/internal/entity"
	"go-time-tracker/internal/model"
	"go-time-tracker/internal/model/filter"
	"go-time-tracker/internal/model/input"
)

type User interface {
	Create(ctx context.Context, user model.UserData) (int, error)
	Update(ctx context.Context, input input.UpdateUser) (int, error)
	Delete(ctx context.Context, userId int) (int, error)
	Get(ctx context.Context, userId int) (*entity.User, error)
	GetByFilter(ctx context.Context, filter filter.User) (*[]entity.User, error)
	GetPagedByFilter(ctx context.Context, offset, limit int, filter filter.User) (*[]entity.User, error)
	Transactor
}

type TimeCost interface {
	Create(ctx context.Context, input model.TimeCostData) (int, error)
	Get(ctx context.Context, timeCostId int) (*entity.TimeCost, error)
	GetByFilter(ctx context.Context, filter filter.TimeCost) (*[]entity.TimeCost, error)
	Delete(ctx context.Context, timeCostId int) (int, error)
	Update(ctx context.Context, input input.UpdateTimeCost) (int, error)
	UpdateByFilter(ctx context.Context, input input.UpdateTimeCost, filter filter.TimeCost) (int, error)
	Transactor
}
