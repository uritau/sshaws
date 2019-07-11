package main

import "github.com/aws/aws-sdk-go/service/ec2"

func get_instances_info(describeInstancesOutput *ec2.DescribeInstancesOutput) []instance {
	var instance_list []instance
	for idx, _ := range describeInstancesOutput.Reservations {
		for _, inst := range describeInstancesOutput.Reservations[idx].Instances {
			name := ""
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			new_instance := NewInstance(name, *inst.PrivateIpAddress, *inst.InstanceId, *inst.InstanceType)
			instance_list = append(instance_list, *new_instance)
		}
	}
	return instance_list
}
