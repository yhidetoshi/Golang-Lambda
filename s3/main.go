package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	svc = s3.New(session.New())
)

const (
	LOCATION = "ap-northeast-1"
)

func main() {
	lambda.Start(Handler)
}

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context) {
	result := getS3Buckets()
	for v, _ := range result {
		fmt.Println(result[v])
	}
}

func getS3Buckets() []string {
	var buckets []string

	params := &s3.ListBucketsInput{}
	res, err := svc.ListBuckets(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, res := range res.Buckets {
		location := GetS3BucketLocation(res.Name)
		if location == LOCATION {
			buckets = append(buckets, *res.Name)
		}
	}
	return buckets
}

func GetS3BucketLocation(bucketname *string) string {
	var region string

	params := &s3.GetBucketLocationInput{
		Bucket: bucketname,
	}
	location, err := svc.GetBucketLocation(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if location.LocationConstraint == nil {
		region = "NULL"
	} else {
		region = *location.LocationConstraint
	}
	return region
}
