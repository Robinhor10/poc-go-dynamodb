#!/bin/sh

# Esperar o LocalStack iniciar
sleep 10

# Criar a tabela DynamoDB
sh /app/init.sh

# Iniciar a aplicação
/poc-go-dynamodb
