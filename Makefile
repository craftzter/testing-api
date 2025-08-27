.PHONY: gen, up, down, start, stop,
gen:
	sqlc generate

up:
	migrate -path migrations -database "postgres://root:secret@localhost:5432/testingDB?sslmode=disable" up

down:
	migrate -path migrations -database "postgres://root:secret@localhost:5432/testingDB?sslmode=disable" down

start:
	docker-compose up -d --build

stop:
	docker-compose down -v
