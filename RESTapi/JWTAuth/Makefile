.PHONY: build

build:
	go build -v ./cmd/api/

run-postgres:
	docker run --name local-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=restapi -d -p 5432:5432 postgres

stop-postgres:
	docker stop local-postgres && docker rm local-postgres

migrate-up:
	migrate -path migrations -database "postgresql://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up  2>&1 | tee -a migrate-up.log

migrate-down:
	migrate -path migrations -database "postgresql://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down  2>&1 | tee -a migrate-down.log

.DEFAULT_GOAL := build