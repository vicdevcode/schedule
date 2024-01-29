PHONY: start

start:
	@go build -o bin/app.exe cmd/app/main.go
	@./bin/app.exe

PHONY: migrate

migrate:
	@go build -o bin/migrate.exe cmd/migrate/main.go
	@./bin/migrate.exe
