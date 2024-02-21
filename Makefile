all: run

run:
	go run cmd/app/main.go

up:
	docker-compose up -d db

stop_and_delete_container:
	docker stop books
	docker rm books
	docker image rmi crud_service-db:latest

create_table:
	docker exec -it new_task_manager psql -U postgres -d booking -c "\i script.sql"

swag:
	swag init -g cmd/app/main.go
