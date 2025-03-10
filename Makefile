# Переменные
DOCKER_COMPOSE = docker-compose
GO_SERVICE = app
BINARY_NAME = myapp

# Команда для сборки Go-приложения
build:
	@echo "Building Go application..."
	go build -o $(BINARY_NAME) ./cmd/$(GO_SERVICE)

# Команда для запуска docker-compose
up:
	@echo "Starting docker-compose..."
	$(DOCKER_COMPOSE) up -d

# Команда для остановки docker-compose
down:
	@echo "Stopping docker-compose..."
	$(DOCKER_COMPOSE) down

# Команда для пересборки и перезапуска docker-compose
rebuild:
	@echo "Rebuilding and restarting docker-compose..."
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) up --build -d

# Команда для очистки (удаление бинарника и остановка контейнеров)
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	$(DOCKER_COMPOSE) down

# Команда для запуска тестов
test:
	@echo "Running tests..."
	go test ./...

# Команда для запуска приложения локально (без Docker)
run:
	@echo "Running Go application locally..."
	go run ./cmd/main.go

# Команда для проверки зависимостей
deps:
	@echo "Checking dependencies..."
	go mod tidy

# Команда для просмотра логов
logs:
	@echo "Showing logs..."
	$(DOCKER_COMPOSE) logs -f

# Команда для входа в контейнер с Go-приложением
ssh:
	@echo "Entering Go service container..."
	$(DOCKER_COMPOSE) exec $(GO_SERVICE) sh

migrate:
	@echo "Start automigrate..."
	go run migrations/auto.go

.PHONY: build up down rebuild clean test run deps logs ssh