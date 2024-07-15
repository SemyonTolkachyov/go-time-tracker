package userrepo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model/input"
	"strings"
	"time"
)

// Update user in pgsql db
func (r Repository) Update(ctx context.Context, input input.UpdateUser) (int, error) {
	log.Debugf("Updating user data in database by id=%d new data=%s", input.Id, input)
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.PassportNumber != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.UserColPassportNumber, argId))
		args = append(args, *input.PassportNumber)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.UserColSurname, argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.UserColName, argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.UserColPatronymic, argId))
		args = append(args, *input.Patronymic)
		argId++
	}

	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", names.UserColAddress, argId))
		args = append(args, *input.Address)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("%s=$%v", names.UserColUpdatedAt, argId))
	args = append(args, time.Now().UTC())

	setQuery := strings.Join(setValues, ", ")
	args = append(args, input.Id)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d AND %s", names.UsersTable, setQuery, names.UserColId, argId+1, getDefaultFilter())

	i, err := pgsql.FromCtxOrData(ctx, r.db).ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	affected, _ := i.RowsAffected()
	return int(affected), nil
}
