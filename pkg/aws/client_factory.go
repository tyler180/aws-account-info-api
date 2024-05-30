package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type AWSClient interface{}

type AWSClientFactory struct {
	cfg aws.Config
}

func NewAWSClientFactory(ctx context.Context, roleARN string) (*AWSClientFactory, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	if roleARN != "" {
		stsClient := sts.NewFromConfig(cfg)
		cfg.Credentials = stscreds.NewAssumeRoleProvider(stsClient, roleARN)
	}

	return &AWSClientFactory{cfg: cfg}, nil
}

func (f *AWSClientFactory) GetClient(serviceType string) (AWSClient, error) {
	switch serviceType {
	case "s3":
		return s3.NewFromConfig(f.cfg), nil
	case "dynamodb":
		return dynamodb.NewFromConfig(f.cfg), nil
	case "ec2":
		return ec2.NewFromConfig(f.cfg), nil
	case "organizations":
		return organizations.NewFromConfig(f.cfg), nil
	case "iam":
		return iam.NewFromConfig(f.cfg), nil
	default:
		return nil, fmt.Errorf("unsupported service type: %s", serviceType)
	}
}
