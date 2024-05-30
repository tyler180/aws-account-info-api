package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2Client struct {
	Client *ec2.Client
}

func NewEC2Client(factory *AWSClientFactory) (*EC2Client, error) {
	client, err := factory.GetClient("ec2")
	if err != nil {
		return nil, err
	}
	return &EC2Client{Client: client.(*ec2.Client)}, nil
}

func (e *EC2Client) DescribeInstances(ctx context.Context) ([]string, error) {
	result, err := e.Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
	if err != nil {
		return nil, err
	}

	var instanceIds []string
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceIds = append(instanceIds, *instance.InstanceId)
		}
	}
	return instanceIds, nil
}
