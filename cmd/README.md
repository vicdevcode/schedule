# CMD

Entrypoint файлы. Здесь можно инициализировать нужные скрипты, к примеру в `app/main.go` инициализируем config, берем оттуда данные и передаем уже в скрипт `app.Run`. Также можно здесь запускать миграции `migrate/main.go`, тесты и так далее.
