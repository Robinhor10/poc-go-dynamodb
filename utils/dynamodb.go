package utils

import (
    "errors"
    "poc-go-dynamodb/models"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "Items"

func GetItem(id string) (*models.Item, error) {
    result, err := DynamoDB.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "ID": {
                S: aws.String(id),
            },
        },
    })

    if err != nil {
        return nil, err
    }

    if result.Item == nil {
        return nil, errors.New("item não encontrado")
    }

    item := new(models.Item)
    err = dynamodbattribute.UnmarshalMap(result.Item, item)
    if err != nil {
        return nil, err
    }

    return item, nil
}

func PutItem(item models.Item) error {
    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        return err
    }

    _, err = DynamoDB.PutItem(&dynamodb.PutItemInput{
        TableName: aws.String(tableName),
        Item:      av,
    })

    return err
}
