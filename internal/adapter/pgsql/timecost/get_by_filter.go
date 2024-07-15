package timecost

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/entity"
	"go-time-tracker/internal/model/filter"
)

// GetByFilter return time costs from pgsql db by filter
func (r Repository) GetByFilter(ctx context.Context, filter filter.TimeCost) (*[]entity.TimeCost, error) {
	log.Debugf("Getting time cost data from database by filter: %s", filter)
	var timeCosts []entity.TimeCost
	filterStatement, args := r.getWhereStatementByFilter(filter, 1)

	selectStatement := fmt.Sprintf("SELECT * FROM %s ", names.TimeCostTable)

	whereStatement := fmt.Sprintf("WHERE %s ", getDefaultFilter())
	if len(args) > 0 {
		whereStatement = whereStatement + fmt.Sprintf("AND %s", filterStatement)
	}

	q := selectStatement + whereStatement

	err := pgsql.FromCtxOrData(ctx, r.db).SelectContext(ctx, &timeCosts, q, args...)

	if err != nil {
		return nil, err
	}
	return &timeCosts, nil
}
