package userrepo

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-time-tracker/internal/adapter/pgsql"
	"go-time-tracker/internal/adapter/pgsql/common/names"
	"go-time-tracker/internal/model"
	"time"
)

// Create add user in pgsql db
func (r Repository) Create(ctx context.Context, user model.UserData) (int, error) {
	log.Debugf("Adding user to database %s", user)
	var id int
	q := fmt.Sprintf(`INSERT INTO %s (%s, %s, %s, %s, %s, %s)
	VALUES ($1, $2, $3, $4, $5, $6 ) RETURNING id;`,
		names.UsersTable, names.UserColCreatedAt, names.UserColPassportNumber, names.UserColSurname, names.UserColName, names.UserColPatronymic, names.UserColAddress)
	row := pgsql.FromCtxOrData(ctx, r.db).QueryRowContext(ctx, q, time.Now().UTC(), user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
