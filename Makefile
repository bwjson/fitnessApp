COMPOSE_FILE = docker-compose.yml

.PHONY: build up down restart logs clean help

build:
	docker build -t user_service:local .

run: build
	docker-compose -f $(COMPOSE_FILE) up --build

down:
	docker-compose -f $(COMPOSE_FILE) down

restart:
	docker-compose -f $(COMPOSE_FILE) down
	docker-compose -f $(COMPOSE_FILE) up --build

logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

clean: down
	rm -rf data

help:
	@echo "Доступные команды:"
	@echo "  make build    - Сборка образов"
	@echo "  make up       - Запуск контейнеров"
	@echo "  make down     - Остановка контейнеров"
	@echo "  make restart  - Перезапуск контейнеров"
	@echo "  make logs     - Просмотр логов"
	@echo "  make clean    - Очистка данных"
