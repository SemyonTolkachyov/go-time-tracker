package filter

import "fmt"

type TimeCost struct {
	Id        *int
	CreatedAt *TimeFilter
	UpdatedAt *TimeFilter
	TaskId    *int
	UserId    *int
	StartAt   *TimeFilter
	EndAt     *TimeFilter
}

func (t TimeCost) String() string {
	return fmt.Sprintf("{Id: %d, CreatedAt: %s, UpdatedAt: %s, TaskId: %d, UserId: %d, StartAt: %s, EndAt: %s}", t.Id, t.CreatedAt, t.UpdatedAt, t.TaskId, t.UserId, t.StartAt, t.EndAt)
}
