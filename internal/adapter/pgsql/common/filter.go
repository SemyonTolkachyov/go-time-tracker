package common

import (
	"fmt"
	"go-time-tracker/internal/model/filter"
)

// GetIntFilterExpression for pgsql
func GetIntFilterExpression(expressions *[]string, values *[]interface{}, argId *int, col string, filterValue int) {
	*expressions = append(*expressions, fmt.Sprintf("%s=$%d", col, *argId))
	*values = append(*values, filterValue)
	*argId++
}

// GetStringFilterExpression for pgsql
func GetStringFilterExpression(expressions *[]string, values *[]interface{}, argId *int, col string, filterValue string) {
	*expressions = append(*expressions, fmt.Sprintf("%s LIKE $%d", col, *argId))
	*values = append(*values, "%"+filterValue+"%")
	*argId++
}

// GetTimeFilterExpression for pgsql
func GetTimeFilterExpression(expressions *[]string, values *[]interface{}, argId *int, col string, filterValue filter.TimeFilter) {
	if filterValue.Op == filter.IsNotNullOp || filterValue.Op == filter.IsNullOp {
		*expressions = append(*expressions, fmt.Sprintf("%s %s", col, filterValue.Op))
	} else {
		*expressions = append(*expressions, fmt.Sprintf("DATE_TRUNC($%d,%s)%s$%d", *argId, col, filterValue.Op, *argId+1))
		*values = append(*values, "SECOND")
		*values = append(*values, filterValue.Time)
		*argId += 2
	}
}
