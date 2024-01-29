# Это чистая архитектура

Шаблон для инициализации проектов под чистую архитектуру. Имеется здесь:
`config`, `logger`, `postgres`, `gorm`, `gin`.

## Как встроить это в свой проект?

Сначала инициализируем свой проект через

```shell
go mod init <repository>
```

Скопировать всё, кроме: `go.mod`, `go.sum`, `example.env`, `example.config.yaml`

Добавить свои `.env, config/config.yaml` файлы, исходя из их примеров.

Заменить import на свой репозиторий, а также просто исправить все ошибки.

список библиотек:

```shell
go get github.com/fatih/color
go get github.com/gin-gonic/gin
go get github.com/google/uuid
go get github.com/ilyakaznacheev/cleanenv
go get github.com/joho/godotenv
go get gorm.io/gorm
go get gorm.io/driver/postgres
```
