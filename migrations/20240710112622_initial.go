package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInitial, downInitial)
}

func upInitial(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE users
	(
	    id            	serial not null unique,
	    created_at    	timestamptz not null,
	    updated_at    	timestamptz,
	    deleted_at    	timestamptz,
	    passport_number varchar(15) not null unique,
	    surname      	varchar(255) not null,
	    name          	varchar(255) not null,
	    patronymic      varchar(255) not null,
	    address         varchar(255) not null
	);
	CREATE TABLE tasks
	(
	    id            	serial not null unique,
	    created_at    	timestamptz not null,
	    updated_at    	timestamptz,
	    deleted_at    	timestamptz,
	    title      		varchar(255) not null,
	    description     varchar
	);
	CREATE TABLE time_costs
	(
	    id            	serial not null unique,
	    created_at    	timestamptz not null,
	    updated_at    	timestamptz,
	    deleted_at    	timestamptz,
	    task_id 		int references tasks (id) on delete cascade not null,
        user_id 		int references users (id) on delete cascade      not null,
	    start_at        timestamptz,
	    end_at      	timestamptz
	);
	`)
	return err
}

func downInitial(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
DROP TABLE time_costs;
DROP TABLE tasts;
DROP TABLE users;
`)
	return err
}
