package trackerservice

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/model"
	"go-time-tracker/internal/model/filter"
	"go-time-tracker/internal/model/input"
	"sort"
	"time"
)

func (s Service) GetUserCostsByPeriod(ctx context.Context, input input.TimeCostsByPeriod) (*[]model.CostData, error) {
	log.Infof("Getting user costs by period %s", input)
	start := filter.TimeFilter{
		Op:   filter.GreaterOp,
		Time: &input.Start,
	}
	end := filter.TimeFilter{
		Op:   filter.LessOp,
		Time: &input.End,
	}
	periodFilter := filter.TimeCost{
		UserId:  &input.UserId,
		StartAt: &start,
		EndAt:   &end,
	}

	tCosts, err := s.timeCostStorage.GetByFilter(ctx, periodFilter)
	if err != nil {
		log.Errorf("Error getting time costs from db by filter%s for input %s: %v", periodFilter, input, err)
		return nil, err
	}
	timeCosts := *tCosts

	sort.Slice(timeCosts, func(i, j int) bool { return timeCosts[i].TaskId < timeCosts[j].TaskId })
	log.Debugf("Creating cost data from time costs for input %s", input)
	costData := make([]model.CostData, 0)
	for i, timeCost := range timeCosts {
		var cost time.Duration

		if timeCost.StartAt.Before(input.Start) {
			cost = timeCost.EndAt.Sub(input.Start)
		} else if timeCost.EndAt.After(input.End) {
			cost = input.End.Sub(*timeCost.StartAt)
		} else {
			cost = timeCost.EndAt.Sub(*timeCost.StartAt)
		}

		if i > 0 && timeCosts[i-1].TaskId == timeCost.TaskId {
			j := len(costData) - 1
			costData[j].Duration += cost
			costData[j].Hours = int(costData[j].Duration.Hours())
			costData[j].Minutes = int(costData[j].Duration.Minutes()) - costData[j].Hours*60
		} else {
			costData = append(costData, model.CostData{
				TaskId:   timeCost.TaskId,
				Duration: cost,
			})
		}
	}
	log.Debugf("Sorting cost data for input %s", input)
	sort.Slice(costData, func(i, j int) bool { return costData[i].Duration > costData[j].Duration })

	return &costData, nil
}
