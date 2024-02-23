# ==================================================================================== # 
# HELPERS
# ==================================================================================== #

# Run the app
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== # 
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build 
build:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #

## run: runs api
.PHONY: run
run:
	@go run ./cmd/api -db-dsn=${POSTGRES_DSN}

## psql: connects to database using psql
.PHONY: psql
psql:
	migrate create -seq -ext .sql -dir ./migrations $(name)

## migration: create a new database migration (requires name argument)
.PHONY: migration
migration:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## up: apply all up database migrations
.PHONY: up
up: confirm
	@echo "Running up migrations.."
	migrate -path=./migrations -database=$(POSTGRES_DSN) up


# ==================================================================================== # 
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit 
audit:
	@echo 'Tidying and verifying module dependencies...' 
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

