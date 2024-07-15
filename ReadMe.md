## **Go time tracker**


**Описание**


Приложение позволяет вести затраты времени по задачам.
Созданные миграции применяются при старте приложения.
Конфигурационные данные вынесены в .env-файл.
- Бэкэнд golang (основные библиотеки: gin, goose, sqlx)
- Бд postgresql

Архитектура приложения разделена на уровни:

- Domain (services, entities)
- Adapters (db repository, third-party api)
- Handlers (app api)

**Эндпоинты**


GET /swagger/*any  - *свагер*
POST /api/v1/users/create - *создание пользователя*
GET /api/v1/users/get - *получение пользователя*
PUT /api/v1/users/update - *обновление пользователя*
DELETE /api/v1/users/delete - *удаление пользователя (используется мягкое удаление)*
GET /api/v1/users/get-by-filter - *получение пользователей по фильтру*
GET /api/v1/users/get-by-filter-paged - *получение страницы с пользователями по фильтру*
GET /api/v1/tracker/get-user-costs-by-period - *получение затрат по задачам за период*
POST /api/v1/tracker/start - *начать отсчет по задаче*
PUT /api/v1/tracker/stop - *завершить отсчет по задаче*

Более подробная информация доступна в свагере

**Основные команды**

Создание миграции

    goose -dir ./migrations create migration_name go

Генерация документации для свагера

    swag init -o api -g cmd/http.go
