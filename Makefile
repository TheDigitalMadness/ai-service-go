APP_NAME=ai-service-go

.PHONY: run, test, docker-build, docker-run

run:
	go run ./cmd

docker-build:
	docker build -t ${APP_NAME}:latest .

docker-run:
	docker build -t ${APP_NAME}:latest .
	docker run ${APP_NAME}:latest --env-file .env