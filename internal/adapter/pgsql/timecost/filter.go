package timecost

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql/common"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model/filter"
	"strings"
)

// getDefaultFilter add default filter for time costs
func getDefaultFilter() string {
	return fmt.Sprintf("%s IS NULL", names.TimeCostColDeletedAt)
}

// getWhereStatementByFilter get where expression by filter for time cost
func (r Repository) getWhereStatementByFilter(filter filter.TimeCost, startIdx int) (string, []interface{}) {
	log.Debugf("Getting where expression for time cost by filter=%s with start param index=%d", filter, startIdx)
	expressions := make([]string, 0)
	values := make([]interface{}, 0)
	argId := startIdx

	if filter.Id != nil {
		common.GetIntFilterExpression(&expressions, &values, &argId, names.TimeCostColId, *filter.Id)
	}

	if filter.CreatedAt != nil {
		common.GetTimeFilterExpression(&expressions, &values, &argId, names.TimeCostColCreatedAt, *filter.CreatedAt)
	}

	if filter.UpdatedAt != nil {
		common.GetTimeFilterExpression(&expressions, &values, &argId, names.TimeCostColUpdatedAt, *filter.UpdatedAt)
	}

	if filter.TaskId != nil {
		common.GetIntFilterExpression(&expressions, &values, &argId, names.TimeCostColTaskId, *filter.TaskId)
	}

	if filter.UserId != nil {
		common.GetIntFilterExpression(&expressions, &values, &argId, names.TimeCostColUserId, *filter.UserId)
	}

	if filter.StartAt != nil {
		common.GetTimeFilterExpression(&expressions, &values, &argId, names.TimeCostColStartAt, *filter.StartAt)
	}

	if filter.EndAt != nil {
		common.GetTimeFilterExpression(&expressions, &values, &argId, names.TimeCostColEndAt, *filter.EndAt)
	}

	return strings.Join(expressions, " AND "), values
}
