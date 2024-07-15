package trackerservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/apperror"
	"go-time-tracker/internal/model/filter"
	"go-time-tracker/internal/model/input"
	"time"
)

func (s Service) Stop(ctx context.Context, userId, taskId int) error {
	log.Infof("Stopping tracker for userId=%d by taskId=%d", userId, taskId)
	stopTime := time.Now().UTC()
	newData := input.UpdateTimeCost{EndAt: &stopTime}
	stopFilter := filter.TimeFilter{Op: filter.IsNullOp}
	f := filter.TimeCost{UserId: &userId, TaskId: &taskId, EndAt: &stopFilter}
	r, err := s.timeCostStorage.UpdateByFilter(ctx, newData, f)
	if err != nil {
		log.Errorf("Error updating time cost for userId=%d by taskId=%d: %v", userId, taskId, err)
		return err
	}
	if r == 0 {
		log.Warnf("No time cost for userId=%d by taskId=%d", userId, taskId)
		return apperror.NewNotFoundError("time cost not found")
	}
	return nil
}
