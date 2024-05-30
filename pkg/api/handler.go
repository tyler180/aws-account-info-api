package api

import (
	"encoding/json"
	"net/http"

	"github.com/tyler180/aws-account-info-api/pkg/aws"
)

func GetS3BucketsHandler(w http.ResponseWriter, r *http.Request) {
	factory, err := aws.NewAWSClientFactory(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s3Client, err := aws.NewS3Client(factory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buckets, err := s3Client.ListBuckets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(buckets)
}

func GetDynamoDBTablesHandler(w http.ResponseWriter, r *http.Request) {
	factory, err := aws.NewAWSClientFactory(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dynamoClient, err := aws.NewDynamoDBClient(factory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tables, err := dynamoClient.ListTables(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tables)
}

func DescribeEC2InstancesHandler(w http.ResponseWriter, r *http.Request) {
	factory, err := aws.NewAWSClientFactory(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ec2Client, err := aws.NewEC2Client(factory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	instances, err := ec2Client.DescribeInstances(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(instances)
}

func ListOrganizationAccountsHandler(w http.ResponseWriter, r *http.Request) {
	factory, err := aws.NewAWSClientFactory(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orgClient, err := aws.NewOrganizationsClient(factory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accounts, err := orgClient.ListAccounts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accounts)
}

func ListIAMUsersHandler(w http.ResponseWriter, r *http.Request) {
	factory, err := aws.NewAWSClientFactory(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	iamClient, err := aws.NewIAMClient(factory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users, err := iamClient.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
