all: run

run:
	go run cmd/app/main.go

up:
	docker-compose up db

stop_and_delete_container:
	docker stop books
	docker rm books
	docker image rmi crud_service-db:latest

create_table:
	docker exec -it books psql -U postgres -d booking -c "\i script.sql"
