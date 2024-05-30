package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IAMClient struct {
	Client *iam.Client
}

func NewIAMClient(factory *AWSClientFactory) (*IAMClient, error) {
	client, err := factory.GetClient("iam")
	if err != nil {
		return nil, err
	}
	return &IAMClient{Client: client.(*iam.Client)}, nil
}

func (i *IAMClient) ListUsers(ctx context.Context) ([]string, error) {
	result, err := i.Client.ListUsers(ctx, &iam.ListUsersInput{})
	if err != nil {
		return nil, err
	}

	var userNames []string
	for _, user := range result.Users {
		userNames = append(userNames, *user.UserName)
	}
	return userNames, nil
}
