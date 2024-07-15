package trackerservice

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/apperror"
	"go-time-tracker/internal/model"
	"go-time-tracker/internal/model/filter"
	"time"
)

func (s Service) Start(ctx context.Context, userId, taskId int) error {
	log.Infof("Starting tracker for userId %d by taskId %d", userId, taskId)
	err := s.timeCostStorage.TxWithOpts(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	}, func(ctx pgsql.TxContext) error {
		stopFilter := filter.TimeFilter{Op: filter.IsNullOp}
		f := filter.TimeCost{UserId: &userId, TaskId: &taskId, EndAt: &stopFilter}
		costs, err := s.timeCostStorage.GetByFilter(ctx, f)
		if err != nil {
			log.Errorf("Error getting costs for userId %d by taskId %d: %v", userId, taskId, err)
			return err
		}
		if len(*costs) > 0 {
			return apperror.NewExistsError("already started")
		}

		startTime := time.Now().UTC()
		_, err = s.timeCostStorage.Create(ctx, model.TimeCostData{
			TaskId:  taskId,
			UserId:  userId,
			StartAt: &startTime,
			EndAt:   nil,
		})
		if err != nil {
			log.Errorf("Error creating time cost for userId %d by taskId %d: %v", userId, taskId, err)
			return err
		}

		return nil
	})
	if err != nil {
		log.Errorf("Error starting tracker for userId %d by taskId %d: %v", userId, taskId, err)
		return err
	}
	return nil
}
