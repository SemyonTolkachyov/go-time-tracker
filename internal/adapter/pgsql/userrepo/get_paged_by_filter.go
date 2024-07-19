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

// GetPagedByFilter return users from pgsql db by filter paged
func (r Repository) GetPagedByFilter(ctx context.Context, offset, limit int, filter filter.User) (*[]entity.User, error) {
	log.Debugf("Getting users paged from database by filter=%s limit=%d offset=%d", filter, limit, offset)
	var users []entity.User
	filterStatement, args := r.getWhereStatementByFilter(filter, 1)

	selectStatement := fmt.Sprintf(
		"SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s ",
		names.UserColId,
		names.UserColCreatedAt,
		names.UserColUpdatedAt,
		names.UserColDeletedAt,
		names.UserColPassportNumber,
		names.UserColName,
		names.UserColSurname,
		names.UserColPatronymic,
		names.UserColAddress,
		names.UsersTable,
	)

	whereStatement := fmt.Sprintf("WHERE %s", getDefaultFilter())
	if len(args) > 0 {
		whereStatement = whereStatement + fmt.Sprintf(" AND %s", filterStatement)
	}

	args = append(args, limit, offset)
	limitStatement := fmt.Sprintf(" ORDER BY %s LIMIT $%d OFFSET $%d", names.UserColId, len(args)-1, len(args))

	q := selectStatement + whereStatement + limitStatement

	err := pgsql.FromCtxOrData(ctx, r.db).SelectContext(ctx, &users, q, args...)
	if err != nil {
		return nil, err
	}
	return &users, nil
}
