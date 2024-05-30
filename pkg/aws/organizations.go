package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
)

type OrganizationsClient struct {
	Client *organizations.Client
}

func NewOrganizationsClient(factory *AWSClientFactory) (*OrganizationsClient, error) {
	client, err := factory.GetClient("organizations")
	if err != nil {
		return nil, err
	}
	return &OrganizationsClient{Client: client.(*organizations.Client)}, nil
}

func (o *OrganizationsClient) ListAccounts(ctx context.Context) ([]string, error) {
	result, err := o.Client.ListAccounts(ctx, &organizations.ListAccountsInput{})
	if err != nil {
		return nil, err
	}

	var accountIds []string
	for _, account := range result.Accounts {
		accountIds = append(accountIds, *account.Id)
	}
	return accountIds, nil
}
