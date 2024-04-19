DB_URL=postgres://postgres:welcome1@localhost:5432/bwa?sslmode=disable

ifeq ($(OS),Windows_NT)
	DETECTED_OS := Windows
else
	DETECTED_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

.SILENT: help
help:
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo " migration-create name={name}  Create migration"
	@echo " migration-up                  Up migrations"
	@echo " migration-down                Down last migration"
	@echo " go-run                        Run Project"

# Build

.SILENT: migration-create
migration-create:
	@migrate create -ext sql -dir ./migrations -seq $(name)

# Up migration

.SILENT: migration-up
migration-up:
	@migrate -database $(DB_URL) -path ./migrations up

# Down migration

.SILENT: migration-down
migration-down:
	@migrate -database $(DB_URL) -path ./migrations down 1

.SILENT: go-run
go-run:
	@go run main.go -config=properties/bwa-startup.prod.yaml

.DEFAULT_GOAL := help

