.SILENT:

run:
	go run cmd/main.go

build:
	go build cmd/main.go

database:
	docker run --name=todo_db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

migrate up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down