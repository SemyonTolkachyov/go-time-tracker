package timecost

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
)

// Delete remove time cost from pgsql db
func (r Repository) Delete(ctx context.Context, timeCostId int) (int, error) {
	log.Debugf("Removing time cost data from database by id=%d", timeCostId)
	q := fmt.Sprintf(`UPDATE %s SET %s = NOW() AT TIME ZONE 'UTC' WHERE %s = $1 AND %s`, names.TimeCostTable, names.TimeCostColDeletedAt, names.TimeCostColId, getDefaultFilter())
	i, err := pgsql.FromCtxOrData(ctx, r.db).ExecContext(ctx, q, timeCostId)
	if err != nil {
		return 0, err
	}
	affected, _ := i.RowsAffected()
	return int(affected), nil
}
