# Go Huberi App

## Description

This is go Point of Sale (POS) app.

## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate

## Configuration

All configuration is in `config.yaml` file.

## API Spec

All API Spec is in `api-spec.yaml` file.

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root@tcp(localhost:3306)/pos" -path db/migrations up
```

## Run Application

### Run web server

```shell
go run cmd/web/main.go
```