# BBB Voting System

## Prerequisites
- Docker
- Docker Compose
- Make

## Setup
1. Clone the repository
2. Copy `.env.example` to `.env` and adjust if necessary

## Usage
- Build the project: `make build`
- Run the project: `make run`
- Stop the project: `make stop`
- Run tests: `make test`
- Clean up: `make clean`
- View logs: `make logs`
- Access API shell: `make shell-api`
- Access DB shell: `make shell-db`
- load tests: `load-test || load-test-fast`
- Generate swager: `swag init -g cmd/main.go`

## Libs

- **gin-gonic/gin**: Framework web para API REST. Escolhido pela praticidade e performance em relação a outros frameworks. Nunca usei, então resolvi testar.
- **go-redis/redis**: Cliente Go para Redis, utilizado para armazenamento de resultados parciais.
- **google/uuid**: Geração de UUIDs, para identificadores únicos.
- **onsi/ginkgo e onsi/gomega**: Frameworks para testes BDD em Go. Gosto de BDD e acho que facilita a escrita de testes.
- **spf13/viper**: Gerenciamento de configuração para aplicações Go. 
- **streadway/amqp**: Cliente AMQP para Go, usado para integração com RabbitMQ.
- **swaggo/files, swaggo/gin-swagger, swaggo/swag**: Geração automática de documentação Swagger para APIs Go.
- **uber-go/dig**: Injeção de dependência para Go, facilitando a estruturação do aplicativo.
- **uber-go/mock**: Geração de mocks para testes unitários.