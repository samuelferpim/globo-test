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
	$(K6) run --out json=results.json tests/script.js && k6 report results.json

load-test-fast:
	$(K6) run --out json=results.json tests/script-fast.js && k6 report results.json
