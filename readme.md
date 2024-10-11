для скачивания
```bash
go get github.com/zongrade/testeffective

go mod tidy
```

`.env` файл имеет следующие переменные
* `DB_USER`
* `DB_NAME`
* `SERVER_PORT`
* `DB_PORT`
* `HOST`
* `PASSWORD`
* `SERVER_READ_TIMEOUT`
* `SERVER_WRITE_TIMEOUT`
* `SERVER_IDLE_TIMEOUT`
* `SERVER_READ_HEADER_TIMEOUT`
* `SERVER_SHUTDOWN_TIMEOUT`

для миграции используется `goose`
[goose.github](https://github.com/pressly/goose)

для запуска
```bash
go run ./cmd/app/main.go