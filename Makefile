# Variáveis
BINARY_NAME=myapp
MAIN_PATH=cmd/api/main.go

# Go comandos
.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

.PHONY: run
run:
	go run $(MAIN_PATH)

.PHONY: test
test:
	go test ./... -v

.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: mock
mock:
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mock/repository_mock.go

.PHONY: swagger
swagger:
	swag init -g $(MAIN_PATH) -o api/docs

.PHONY: clean
clean:
	rm -f bin/$(BINARY_NAME)
	rm -f coverage.out

.PHONY: deps
deps:
	go mod tidy
	go mod verify

.PHONY: createMigration
createMigration:
	migrate create -ext=sql -dir=scripts/migrations -seq init

.PHONY: migrate
migrate:
	migrate -path=scripts/migrations -database "postgres://postgres:postgres@localhost:5432/myapp?sslmode=disable" -verbose up

.PHONY: migrateDown
migrateDown:
	migrate -path=scripts/migrations -database "postgres://postgres:postgres@localhost:5432/myapp?sslmode=disable" -verbose down