package model

import (
	"fmt"
	"time"
)

type TimeCostData struct {
	Id      int
	TaskId  int
	UserId  int
	StartAt *time.Time
	EndAt   *time.Time
}

type CostData struct {
	TaskId   int
	Duration time.Duration
	Hours    int
	Minutes  int
}

func (t TimeCostData) String() string {
	return fmt.Sprintf("{Id: %d, TaskId: %d, UserId: %d, StartAt: %s, EndAt: %s}", t.Id, t.TaskId, t.UserId, t.StartAt, t.EndAt)
}

func (c CostData) String() string {
	return fmt.Sprintf("{TaskId:%d, Duration: %s, Hours: %d, Minutes: %d}", c.TaskId, c.Duration.String(), c.Hours, c.Minutes)
}
