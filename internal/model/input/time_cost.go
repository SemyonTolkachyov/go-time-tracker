package input

import (
	"fmt"
	"go-time-tracker/internal/apperror"
	"time"
)

type TimeCostsByPeriod struct {
	UserId int       `json:"user_id" binding:"required"`
	Start  time.Time `json:"start" binding:"required"`
	End    time.Time `json:"end" binding:"required"`
}

type UpdateTimeCost struct {
	Id      int        `json:"id" binding:"required"`
	TaskId  *int       `json:"task_id"`
	UserId  *int       `json:"user_id"`
	StartAt *time.Time `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}

func (t TimeCostsByPeriod) String() string {
	return fmt.Sprintf("{UserId: %d, Start: %s, End: %s}", t.UserId, t.Start, t.End)
}

func (u UpdateTimeCost) String() string {
	return fmt.Sprintf("{Id: %d, TaskId: %d, UserId: %d, StartAt: %s, EndAt: %s}", u.Id, u.TaskId, u.UserId, u.StartAt, u.EndAt)
}

func (u UpdateTimeCost) Validate() error {
	if u.TaskId == nil && u.UserId == nil && u.StartAt == nil && u.EndAt == nil {
		return apperror.NewValidationError("update structure has no values")
	}
	return nil
}
