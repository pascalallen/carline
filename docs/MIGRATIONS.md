# Migrations

This project uses a library called [migrate](https://github.com/golang-migrate/migrate) for database migration 
tooling.

## Creating a migration

```bash
bin/migrate create -ext sql -seq create_permissions_table
```

2 SQL files will be generated at `internal/carline/infrastructure/database/migrations`.

## Run migrations

```bash
bin/migrate -database "postgres://pascalallen:password@postgres:5432/carline?sslmode=disable" -path . up
```

## Roll back migrations

```bash
bin/migrate -database "postgres://pascalallen:password@postgres:5432/carline?sslmode=disable" -path . down
```
