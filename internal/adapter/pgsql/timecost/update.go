package timecost

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model/input"
	"strings"
	"time"
)

// Update time cost in pgsql db
func (r Repository) Update(ctx context.Context, input input.UpdateTimeCost) (int, error) {
	log.Debugf("Updating time cost data in database by id=%d new data=%s", input.Id, input)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.TimeCostColUserId, argId))
		args = append(args, *input.UserId)
		argId++
	}

	if input.TaskId != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.TimeCostColTaskId, argId))
		args = append(args, *input.TaskId)
		argId++
	}

	if input.StartAt != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.TimeCostColStartAt, argId))
		args = append(args, *input.StartAt)
		argId++
	}

	if input.EndAt != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.TimeCostColEndAt, argId))
		args = append(args, *input.EndAt)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("%s=$%v", names.TimeCostColUpdatedAt, argId))
	args = append(args, time.Now().UTC())

	setQuery := strings.Join(setValues, ", ")
	args = append(args, input.Id)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d AND %s", names.TimeCostTable, setQuery, names.TimeCostColId, argId+1, getDefaultFilter())

	i, err := pgsql.FromCtxOrData(ctx, r.db).ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	affected, _ := i.RowsAffected()
	return int(affected), nil
}
