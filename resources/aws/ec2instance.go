package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

// EC2instance structure data for instance on aws
type EC2instance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

// FetchEC2Instances of usage for EC2 instance
func (cloud *Cloud) FetchEC2Instances() (ec2instances []EC2instance, err error) {
	var result *ec2.DescribeInstancesOutput
	result, err = ec2.New(cloud.session).DescribeInstances(&ec2.DescribeInstancesInput{})

	if err != nil {
		return cloud.Ec2instances, err
	}

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceName := ""
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					instanceName = *tag.Value
				}
			}
			cloud.Ec2instances = append(cloud.Ec2instances, EC2instance{
				ID:     *instance.InstanceId,
				Name:   instanceName,
				Type:   *instance.InstanceType,
				Status: *instance.State.Name,
			})
		}
	}

	return cloud.Ec2instances, err
}
