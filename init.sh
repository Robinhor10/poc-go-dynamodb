#!/bin/sh

# Esperar o LocalStack iniciar
sleep 10

# Criar a tabela DynamoDB
aws dynamodb create-table \
    --endpoint-url=http://localstack:4566 \
    --region us-east-1 \
    --table-name Items \
    --attribute-definitions AttributeName=ID,AttributeType=S \
    --key-schema AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

echo "Tabela DynamoDB criada."

# Verificar se a tabela foi criada com sucesso
aws dynamodb describe-table --endpoint-url=http://localstack:4566 --region us-east-1 --table-name Items > /dev/null 2>&1

if [ $? -eq 0 ]; then
    echo "Tabela Items criada com sucesso."
else
    echo "Erro: Tabela Items não foi criada."
    exit 1
fi
