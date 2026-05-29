# Mundo Invest - Client & Event Management API

API desenvolvida em Go para gerenciamento de clientes e processamento de eventos integrados ao Pipefy via GraphQL.

---

## Tecnologias utilizadas

- Go 1.24.0
- Chi Router
- GORM
- SQLite (testes) / PostgreSQL (produção)
- Testify (testes)
- Pipefy GraphQL (integração simulada)

---

## Instalações necessárias
- Docker compose
- go migrate:
``
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
``

## Como executar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/gguedes14/mundo-invest.git
cd mundo-invest
```

### 2. Executar o Postgresql com docker compose

```bash
cd deployments/local
docker compose up
``` 

### 3. Instalar as dependencias do go
```bash
go mod tidy
```

### 4. Exportar as envs
```bash
source .env
```

### 5. Executar a API
```bash
go run cmd/api/main.go
```

### 6. Executar as migrations

```bash
migrate -path ./migrations \
-database "postgres://postgres:postgres@localhost:5432/mundo_invest?sslmode=disable" \
up
```

## Requisições

### Criar clientes e mapeamento de cards
```bash
curl -X POST http://localhost:3000/clientes \
  -H "Content-Type: application/json" \
  -d '{
    "cliente_nome": "João Silva",
    "cliente_email": "joao.silva@example.com",
    "tipo_solicitacao": "Atualização cadastral",
    "valor_patrimonio": 250000
  }'
```

### Atualização de card
```bash
curl -X POST http://localhost:3000/webhooks/pipefy/card-updated \
  -H "Content-Type: application/json" \
  -d '{
    "event_id": "evt_322",
    "card_id": "card_416",
    "cliente_email": "joao.silva@example.com",
    "timestamp": "2026-05-18T12:00:00Z"
  }'
```

### Executar os testes
```bash
go test ./... -v
```
