package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
}

func NewS3Client(factory *AWSClientFactory) (*S3Client, error) {
	client, err := factory.GetClient("s3")
	if err != nil {
		return nil, err
	}
	return &S3Client{Client: client.(*s3.Client)}, nil
}

func (s *S3Client) ListBuckets(ctx context.Context) ([]string, error) {
	result, err := s.Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	var buckets []string
	for _, b := range result.Buckets {
		buckets = append(buckets, *b.Name)
	}
	return buckets, nil
}
