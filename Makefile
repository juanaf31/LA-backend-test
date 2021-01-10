start:
	docker-compose up

start-detach:
	docker-compose up -d

stop:
	docker-compose stop

build:
	docker-compose up --build

rebuild:
	docker-compose down
	docker-compose build
	docker-compose up -d

dev:
	go test ./... -cover
	go run main.go