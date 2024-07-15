package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upSeed, downSeed)
}

func upSeed(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
INSERT INTO public.tasks (id, created_at, title, description) VALUES (1, NOW() AT TIME ZONE 'UTC', 'Task1', 'Test');
INSERT INTO public.tasks (id, created_at, title, description) VALUES (2, NOW() AT TIME ZONE 'UTC', 'Task2', 'Test');
INSERT INTO public.tasks (id, created_at, title, description) VALUES (3, NOW() AT TIME ZONE 'UTC', 'Task3', 'Test');
INSERT INTO public.tasks (id, created_at, title, description) VALUES (4, NOW() AT TIME ZONE 'UTC', 'Task4', 'Test');
INSERT INTO public.tasks (id, created_at, title, description) VALUES (5, NOW() AT TIME ZONE 'UTC', 'Task5', 'Test');

INSERT INTO public.users (id, created_at, passport_number, surname, name, patronymic, address) VALUES (1, NOW() AT TIME ZONE 'UTC', '4444 444444', 'Тест', 'Иван', 'Иванович', 'Москва');
INSERT INTO public.users (id, created_at, passport_number, surname, name, patronymic, address) VALUES (2, NOW() AT TIME ZONE 'UTC', '5555 555555', 'Тест', 'Олег', 'Олегович', 'Тула');
INSERT INTO public.users (id, created_at, passport_number, surname, name, patronymic, address) VALUES (3, NOW() AT TIME ZONE 'UTC', '6666 666666', 'Тест', 'Сергей', 'Сергеевич', 'Воронеж');
INSERT INTO public.users (id, created_at, passport_number, surname, name, patronymic, address) VALUES (4, NOW() AT TIME ZONE 'UTC', '7777 777777', 'Тест', 'Андрей', 'Андреевич', 'Новосибирск');
INSERT INTO public.users (id, created_at, passport_number, surname, name, patronymic, address) VALUES (5, NOW() AT TIME ZONE 'UTC', '8888 888888', 'Тест', 'Семён', 'Семёнович', 'Красноярск');

INSERT INTO public.time_costs (id, created_at, task_id, user_id, start_at, end_at) VALUES (1, NOW() AT TIME ZONE 'UTC', 1, 1, '2024-07-05 00:00:00.000000 +00:00', '2024-07-05 10:00:00.000000 +00:00');
INSERT INTO public.time_costs (id, created_at, task_id, user_id, start_at, end_at) VALUES (2, NOW() AT TIME ZONE 'UTC', 2, 2, '2024-07-05 00:00:00.000000 +00:00', null);
INSERT INTO public.time_costs (id, created_at, task_id, user_id, start_at, end_at) VALUES (3, NOW() AT TIME ZONE 'UTC', 3, 1, '2024-07-05 00:00:00.000000 +00:00', '2024-07-06 00:30:00.000000 +00:00');
INSERT INTO public.time_costs (id, created_at, task_id, user_id, start_at, end_at) VALUES (4, NOW() AT TIME ZONE 'UTC', 4, 3, '2024-07-05 00:00:00.000000 +00:00', null);
INSERT INTO public.time_costs (id, created_at, task_id, user_id, start_at, end_at) VALUES (5, NOW() AT TIME ZONE 'UTC', 5, 1, '2024-07-05 00:00:00.000000 +00:00', '2024-07-05 01:00:00.000000 +00:00');`)
	return err
}

func downSeed(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
DELETE FROM time_costs WHERE id IN (1,2,3,4,5);
DELETE FROM tasks WHERE id IN (1,2,3,4,5);
DELETE FROM users WHERE id IN (1,2,3,4,5);`)
	return err
}
