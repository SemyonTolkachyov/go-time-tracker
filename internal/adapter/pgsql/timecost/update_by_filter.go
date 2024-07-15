package timecost

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model/filter"
	"go-time-tracker/internal/model/input"
	"strings"
	"time"
)

// UpdateByFilter time costs in pgsql db by filter
func (r Repository) UpdateByFilter(ctx context.Context, input input.UpdateTimeCost, filter filter.TimeCost) (int, error) {
	log.Debugf("Updating time costs data in database by filter:%s new data:%s", filter, input)
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

	whereStatement, filterValues := r.getWhereStatementByFilter(filter, argId+1)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s AND %s", names.TimeCostTable, setQuery, whereStatement, getDefaultFilter())
	args = append(args, filterValues...)
	i, err := pgsql.FromCtxOrData(ctx, r.db).ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	affected, _ := i.RowsAffected()
	return int(affected), nil
}
