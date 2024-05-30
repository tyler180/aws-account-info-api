package api

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/s3/buckets", GetS3BucketsHandler).Methods("GET")
	r.HandleFunc("/dynamodb/tables", GetDynamoDBTablesHandler).Methods("GET")
	r.HandleFunc("/ec2/instances", DescribeEC2InstancesHandler).Methods("GET")
	r.HandleFunc("/organizations/accounts", ListOrganizationAccountsHandler).Methods("GET")
	r.HandleFunc("/iam/users", ListIAMUsersHandler).Methods("GET")
}
