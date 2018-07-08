package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	svc = ec2.New(session.New())
)

func main() {
	lambda.Start(Handler)
}

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context) {
	result := getInstanceInfo()
	for v, _ := range result {
		fmt.Println(result[v])
	}

}

func getInstanceInfo() []string {
	var tagName string
	var instances []string

	params := &ec2.DescribeInstancesInput{}
	res, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, resInfo := range res.Reservations {
		for _, instanceInfo := range resInfo.Instances {
			for _, tagInfo := range instanceInfo.Tags {
				if *tagInfo.Key == "Name" {
					tagName = *tagInfo.Value
				}
			}
			instances = append(instances, tagName)
		}
	}
	return instances
}
