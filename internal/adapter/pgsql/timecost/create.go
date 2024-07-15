package timecost

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model"
	"time"
)

// Create add time cost in pgsql db
func (r Repository) Create(ctx context.Context, input model.TimeCostData) (int, error) {
	log.Debugf("Adding time cost data to database %s", input)
	var id int
	q := fmt.Sprintf(`INSERT INTO %s (%s, %s, %s, %s, %s)
	VALUES ($1, $2, $3, $4, $5 ) RETURNING id;`,
		names.TimeCostTable, names.TimeCostColCreatedAt, names.TimeCostColTaskId, names.TimeCostColUserId, names.TimeCostColStartAt, names.TimeCostColEndAt)

	row := pgsql.FromCtxOrData(ctx, r.db).QueryRowContext(ctx, q, time.Now().UTC(), input.TaskId, input.UserId, input.StartAt, input.EndAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
