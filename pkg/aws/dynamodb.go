package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	Client *dynamodb.Client
}

func NewDynamoDBClient(factory *AWSClientFactory) (*DynamoDBClient, error) {
	client, err := factory.GetClient("dynamodb")
	if err != nil {
		return nil, err
	}
	return &DynamoDBClient{Client: client.(*dynamodb.Client)}, nil
}

func (d *DynamoDBClient) ListTables(ctx context.Context) ([]string, error) {
	result, err := d.Client.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	return result.TableNames, nil
}
