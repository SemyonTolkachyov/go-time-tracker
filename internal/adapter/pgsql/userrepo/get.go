package userrepo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/entity"
)

// Get return user from pgsql db
func (r Repository) Get(ctx context.Context, userId int) (*entity.User, error) {
	log.Debugf("Getting user from database by userId=%d", userId)
	var user entity.User
	q := fmt.Sprintf(`SELECT * FROM %s WHERE %s = $1 AND %s`, names.UsersTable, names.UserColId, getDefaultFilter())
	if err := pgsql.FromCtxOrData(ctx, r.db).GetContext(ctx, &user, q, userId); err != nil {
		return nil, err
	}

	return &user, nil
}
