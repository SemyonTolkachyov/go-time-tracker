package timecost

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/entity"
)

// Get return time cost from pgsql db
func (r Repository) Get(ctx context.Context, timeCostId int) (*entity.TimeCost, error) {
	log.Debugf("Get time cost data from database by timeCostId=%d", timeCostId)
	var timeCost entity.TimeCost
	q := fmt.Sprintf(`SELECT * FROM %s WHERE %s = $1 AND %s`, names.TimeCostTable, names.TimeCostColId, getDefaultFilter())
	if err := pgsql.FromCtxOrData(ctx, r.db).GetContext(ctx, timeCost, q, timeCostId); err != nil {
		return nil, err
	}

	return &timeCost, nil
}
