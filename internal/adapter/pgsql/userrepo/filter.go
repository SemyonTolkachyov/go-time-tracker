package userrepo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql/common"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model/filter"
	"strings"
)

// getDefaultFilter add default filter for users
func getDefaultFilter() string {
	return fmt.Sprintf("%s IS NULL", names.UserColDeletedAt)
}

// getWhereStatementByFilter get where expression by filter for user
func (r Repository) getWhereStatementByFilter(filter filter.User, startIdx int) (string, []interface{}) {
	log.Debugf("Getting where expression for user by filter=%s with start param index=%d", filter, startIdx)
	expressions := make([]string, 0)
	values := make([]interface{}, 0)
	argId := startIdx

	if filter.Id != nil {
		common.GetIntFilterExpression(&expressions, &values, &argId, names.UserColId, *filter.Id)
	}

	if filter.CreatedAt != nil {
		common.GetTimeFilterExpression(&expressions, &values, &argId, names.UserColCreatedAt, *filter.CreatedAt)
	}

	if filter.UpdatedAt != nil {
		common.GetTimeFilterExpression(&expressions, &values, &argId, names.UserColUpdatedAt, *filter.UpdatedAt)
	}

	if filter.PassportNumber != nil {
		common.GetStringFilterExpression(&expressions, &values, &argId, names.UserColPassportNumber, *filter.PassportNumber)
	}

	if filter.Surname != nil {
		common.GetStringFilterExpression(&expressions, &values, &argId, names.UserColSurname, *filter.Surname)
	}

	if filter.Name != nil {
		common.GetStringFilterExpression(&expressions, &values, &argId, names.UserColName, *filter.Name)
	}

	if filter.Patronymic != nil {
		common.GetStringFilterExpression(&expressions, &values, &argId, names.UserColPatronymic, *filter.Patronymic)
	}

	if filter.Address != nil {
		common.GetStringFilterExpression(&expressions, &values, &argId, names.UserColAddress, *filter.Address)
	}

	return strings.Join(expressions, " AND "), values
}
