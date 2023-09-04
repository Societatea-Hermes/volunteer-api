include .env

stop_containters:
	@echo "Stopping other docker containers"
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped conainters"; \
		docker stop $$(docker ps -q); \
	else \
		echo "no containers running..."; \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:15-alpine

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_containers:
	docker start ${DB_DOCKER_CONTAINER}

migrate_up:
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

migrate_down:
	sqlx migrate revert --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary..."
	go build -o ${BINARY} cmd/api/*.go

run: stop build
	@env DSN=${DSN} PORT=${PORT} ./${BINARY} &
	@echo "API started..."

stop:
	@echo "stopping server..."
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "server stopped..."