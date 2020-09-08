package main

import (
	"sshaws/helpers"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func getInstancesInfo(describeInstancesOutput *ec2.DescribeInstancesOutput) []helpers.Instance {
	var instanceList []helpers.Instance
	for idx := range describeInstancesOutput.Reservations {
		for _, inst := range describeInstancesOutput.Reservations[idx].Instances {
			name := ""
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			newInstance := helpers.NewInstance(name, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType)
			instanceList = append(instanceList, *newInstance)
		}
	}
	return instanceList
}
