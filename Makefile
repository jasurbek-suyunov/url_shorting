include .env

go:
	go run cmd/main.go
	
watch:
	google-chrome 'http://${HTTP_HOST}:${HTTP_PORT}/swagger/index.html'
	make go

login:
	docker exec -it ${DOCKER_CONTAINER_NAME} psql ${POSTGRES_DATABASE} ${POSTGRES_USER}

start:
	docker start ${DOCKER_CONTAINER_NAME}

psqlcontainer:
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -d -p ${POSTGRES_PORT}:5432 --env-file .env postgres:15-alpine3.16

rediscontainer:
	docker run --name ${DOCKER_REDIS_CONTAINER_NAME} -p ${REDIS_CLIENT_PORT}:6379 -d --env-file .env redis:7.0.5-alpine3.16


stop:
	docker stop ${DOCKER_CONTAINER_NAME}
tidy:
	go mod tidy
swag:	
	swag init -g ./cmd/main.go -o ./docs

migration-up:
	migrate -path ./migrations/ -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

migration-down:
	migrate -path ./migrations/ -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' down

createdb:
	docker exec -it ${DOCKER_POSTGRES__CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it ${DOCKER_POSTGRES__CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${POSTGRES_DB}