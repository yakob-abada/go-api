CONTAINER_NAME ?= go-app

container_login:
	docker exec -it ${CONTAINER_NAME} bash

tests: 
	docker exec -it ${CONTAINER_NAME} go test ./...

start:
	docker compose up --build -d