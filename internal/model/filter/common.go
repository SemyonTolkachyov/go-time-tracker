package filter

import (
	"fmt"
	"go-time-tracker/internal/apperror"
	"time"
)

const (
	EqualOp     = "="
	LessOp      = "<"
	GreaterOp   = ">"
	IsNullOp    = "IS NULL"
	IsNotNullOp = "IS NOT NULL"
)

const (
	EqualFilterOp   = "eq"
	LessFilterOp    = "le"
	GreaterFilterOp = "gt"
)

type TimeFilter struct {
	Op   string
	Time *time.Time
}

func (tf TimeFilter) String() string {
	return fmt.Sprintf("{Op: %s, Time: %s}", tf.Op, tf.Time.Format(time.DateTime))
}

func ParseTimeFilter(filter string) (*TimeFilter, error) {
	opStr := filter[0:2]
	var op string
	switch opStr {
	case EqualFilterOp:
		op = EqualOp
		break
	case LessFilterOp:
		op = LessOp
		break
	case GreaterFilterOp:
		op = GreaterOp
		break
	default:
		return nil, apperror.NewUnknownOperatorError(fmt.Sprintf("got unknown operator: %s", opStr))
	}
	datetime, err := time.Parse(time.DateTime, filter[2:])
	if err != nil {
		return nil, err
	}
	return &TimeFilter{Op: op, Time: &datetime}, nil
}
