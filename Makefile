.PHONY: gen, up, down, start, stop, psql
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

psql:
	docker exec -it monly_postgres psql -U root -d testingDB
