PHONY: run_db create_migration make_migration_up make_migration_down go_get run_server


run_db:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

create_migration:
	migrate create -ext sql -dir ./schema -seq init

make_migration_up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

make_migration_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

go_get:
	go get -u github.com/$(package)

run_server:
	go run cmd/main.go
