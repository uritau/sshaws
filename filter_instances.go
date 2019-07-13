package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func filter_instances(filters []instanceFilter, region string) *ec2.DescribeInstancesOutput {

	fmt.Printf("Application: %s   Environment: %s   Name: %s   Region: %s\n", app, env, name, region)
	fmt.Printf("---------------------------------------------------------\n\n")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	ec2svc := ec2.New(sess)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Environment"),
				Values: []*string{aws.String(env)},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String("*" + name + "*")},
			},
			{
				Name:   aws.String("tag:Application"),
				Values: []*string{aws.String(app + "*")},
			}, {
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}
	resp, err := ec2svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}
	return resp
}
