package utils

import (
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoDB *dynamodb.DynamoDB

func init() {
    sess := session.Must(session.NewSession(&aws.Config{
        Region:   aws.String(os.Getenv("AWS_REGION")),
        Endpoint: aws.String(os.Getenv("DYNAMODB_ENDPOINT")),
    }))
    DynamoDB = dynamodb.New(sess)
}
