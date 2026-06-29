include .env
export

APP_NAME=ai-service-go

.PHONY: run, test, docker-build, docker-run

run:
	go run ./cmd

docker-build:
	docker build -t ${APP_NAME}:latest .

docker-run:
	docker build -t ${APP_NAME}:latest .
	docker run --env-file .env -d -p "${NODE_PORT}:${NODE_PORT}" ${APP_NAME}:latest