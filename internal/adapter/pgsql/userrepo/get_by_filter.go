package userrepo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/entity"
	"go-time-tracker/internal/model/filter"
)

// GetByFilter return users from pgsql db by filter
func (r Repository) GetByFilter(ctx context.Context, filter filter.User) (*[]entity.User, error) {
	log.Debugf("Getting users from database by filter=%s", filter)
	var users []entity.User
	filterStatement, args := r.getWhereStatementByFilter(filter, 1)

	selectStatement := fmt.Sprintf("SELECT * FROM %s ", names.UsersTable)

	whereStatement := fmt.Sprintf("WHERE %s ", getDefaultFilter())
	if len(args) > 0 {
		whereStatement = whereStatement + fmt.Sprintf("AND %s", filterStatement)
	}

	q := selectStatement + whereStatement

	err := pgsql.FromCtxOrData(ctx, r.db).SelectContext(ctx, &users, q, args...)
	if err != nil {
		return nil, err
	}
	return &users, nil
}
