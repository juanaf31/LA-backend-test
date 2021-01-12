start: ## for starting docker container in foreground
	docker-compose up

start-detach: ## for starting docker container in background
	docker-compose up -d

stop: ## for stopping docker container services
	docker-compose stop

build: ## for build application
	docker-compose up --build

rebuild: ## for rebuild application after any changes
	docker-compose down
	docker-compose build
	docker-compose up -d

logs: ## Show logs
	docker-compose logs

dev: ## for run unit testing and application
	go test ./... -cover
	go run main.go