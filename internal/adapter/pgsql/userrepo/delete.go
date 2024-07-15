package userrepo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
)

// Delete remove user from pgsql db
func (r Repository) Delete(ctx context.Context, userId int) (int, error) {
	log.Debugf("Removing user from database by id=%d", userId)
	q := fmt.Sprintf(`UPDATE %s SET %s = NOW() AT TIME ZONE 'UTC' WHERE %s = $1 AND %s`, names.UsersTable, names.UserColDeletedAt, names.UserColId, getDefaultFilter())
	i, err := pgsql.FromCtxOrData(ctx, r.db).ExecContext(ctx, q, userId)
	if err != nil {
		return 0, err
	}
	affected, _ := i.RowsAffected()
	return int(affected), nil
}
