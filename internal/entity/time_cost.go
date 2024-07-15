package entity

import (
	"fmt"
	"time"
)

type TimeCost struct {
	Id        int        `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	TaskId    int        `db:"task_id"`
	UserId    int        `db:"user_id"`
	StartAt   *time.Time `db:"start_at"`
	EndAt     *time.Time `db:"end_at"`
}

func (t TimeCost) String() string {
	return fmt.Sprintf("{Id: %d, CreatedAt: %s, UpdatedAt: %s, DeletedAt: %s, TaskId: %d, UserId: %d, StartAt: %s, EndAt: %s}", t.Id, t.CreatedAt, t.UpdatedAt, t.DeletedAt, t.TaskId, t.UserId, t.StartAt, t.EndAt)
}
