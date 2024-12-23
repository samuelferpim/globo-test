# Arquitetura do Sistema de Votação BBB

## Visão Geral

Este sistema foi projetado para lidar com um grande volume de votos em tempo real, garantindo escalabilidade, confiabilidade e performance. A arquitetura foi pensada para ser limpa e modular, facilitando manutenção e futuras expansões.

## Decisões Arquiteturais

### 1. Arquitetura Limpa

Optei por uma arquitetura mais limpa para:
- Separar as preocupações 
- Facilitar a manutenção e testabilidade do código
- Permitir a evolução independente de diferentes camadas do sistema
- Injeção de dependências para facilitar a troca de implementações

### 2. Redis para Resultados Parciais

Escolhi o Redis como banco de dados em memória para:
- Armazenar e recuperar rapidamente os resultados parciais das votações
- Lidar com alta concorrência de leitura/escrita
- Permitir operações atômicas de incremento, crucial para a contagem de votos
- Utilização de Pipeline para melhorar a performance

### 3. Banco de Dados Relacional (PostgreSQL)

Um banco de dados relacional é usado para:
- Armazenar dados completos persistentes e estruturados
- Manter o histórico completo de votações
- Realizar análises e relatórios mais complexos

### 4. Sistema de Filas (RabbitMQ)

Implementei um sistema de filas com RabbitMQ para:
- Desacoplar o processamento de votos da API de recebimento
- Garantir que nenhum voto seja perdido em momentos de pico
- Permitir o processamento assíncrono e distribuído dos votos

### 5. Microsserviços

A arquitetura de microsserviços foi adotada para:
- Escalar componentes independentemente
- Facilitar o desenvolvimento e deploy de partes específicas do sistema
- Melhorar a resiliência geral do sistema

- Gostaria de ter feito um consumer para consumir os dados da fila e salvar no banco de dados, mas não consegui implementar a tempo.
- Também implementar um serviço para realizar um SYNC entre os dados (Redis/Postgres) e para separar o write do read, para não sobrecarregar a aplicação, gerar reports e monitorar a consistência dos dados.
- Implementar um cache para os resultados parciais, para melhorar a performance.

## Fluxo de Dados

1. A API recebe os votos e os envia para a fila
2. Consumidores processam os votos da fila
3. Os resultados são atualizados no Redis para consulta rápida
4. Periodicamente, os dados são persistidos no banco de dados relacional

## Considerações de Escalabilidade

- Os componentes podem ser escalados horizontalmente conforme necessário
- O uso de Redis e filas permite lidar com picos de tráfego
- A arquitetura distribuída melhora a tolerância a falhas

## Próximos Passos

- Implementar um sistema de monitoramento e alertas
- Monitoramento de logs e métricar, Jaeger para tracing, Grafana para dashboards e Prometheus para as métricas. 
- Load Balancer/CDN para melhorar a performance e escalabilidade
- Configurar o Gateway para aplicar rate limit, aplicar estratégia para detectar e bloquear bots e DDOS.


## Considerações

- Considerei um monolito, porém escolhi a arquitetura de microsserviços para melhorar a escalabilidade e resiliência do sistema e dos dados.
- Alternativas cogitadas foram Event-Driven com CQRS, Serverless e Sharding