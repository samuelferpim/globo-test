.PHONY: build run test clean load-test

# Variáveis
DOCKER_COMPOSE = docker-compose
K6 = k6

# Comandos
build:
	$(DOCKER_COMPOSE) build

run:
	$(DOCKER_COMPOSE) up -d

stop:
	$(DOCKER_COMPOSE) down

test:
	go test ./...

clean:
	$(DOCKER_COMPOSE) down -v
	docker system prune -f

logs:
	$(DOCKER_COMPOSE) logs -f

shell-api:
	$(DOCKER_COMPOSE) exec api sh

shell-db:
	$(DOCKER_COMPOSE) exec db psql -U user -d bbbvoting

load-test:
	$(K6) run --out json=tests/k6/results/results.json tests/k6/script.js

load-test-fast:
	$(K6) run --out json=tests/k6/results/results.json tests/k6/script-fast.js
