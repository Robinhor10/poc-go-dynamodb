# Go DynamoDB API

Esta é uma API em Go que interage com o DynamoDB usando LocalStack. A API permite consultar e gravar registros no DynamoDB.

## Estrutura do Projeto

go-dynamodb-api/
├── Dockerfile
├── docker-compose.yml
├── main.go
├── go.mod
├── handlers/
│ ├── get_item.go
│ ├── put_item.go
├── models/
│ └── item.go
├── utils/
│ ├── config.go
│ └── dynamodb.go
├── README.md


## Configuração e Execução

### Pré-requisitos

- Docker
- Docker Compose

### Passos

1. Clone o repositório:

```bash
git clone <repo-url>
cd poc-go-dynamodb

2. Construa e inicie os contêineres:

```bash
docker-compose up --build

3. Acesse a aplicação em http://localhost:8080

Endpoints
GET /item/{id}: Consulta um item pelo ID.
POST /item: Cria um novo item.

Estrutura dos Dados

{
    "id": "string",
    "name": "string",
    "value": "string"
}

### Exemplos de Uso

Consultar Item

curl -X GET http://localhost:8080/item/{id}

Criar Item

curl -X POST http://localhost:8080/item -H "Content-Type: application/json" -d '{
    "id": "1",
    "name": "Item 1",
    "value": "Valor 1"
}'

### Explicação dos Módulos

main.go
O ponto de entrada da aplicação, onde os roteadores e o servidor HTTP são configurados.

handlers/get_item.go
Lida com a consulta de itens no DynamoDB.

handlers/put_item.go
Lida com a criação de novos itens no DynamoDB.

models/item.go
Define a estrutura do item armazenado no DynamoDB.

utils/config.go
Configura a sessão do DynamoDB.

utils/dynamodb.go
Contém funções para interagir com o DynamoDB.

Aprendizado sobre Go
Go é uma linguagem de programação compilada, estaticamente tipada e conhecida por sua simplicidade e eficiência.
Os pacotes são organizados de forma modular, como visto nas pastas handlers, models e utils.
A biblioteca net/http é utilizada para criar servidores HTTP.
gorilla/mux é um roteador HTTP poderoso para criar endpoints de API.
AWS SDK para Go (aws-sdk-go) é usado para interagir com o DynamoDB.

