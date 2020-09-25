package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// EC2instance structure data for instance on aws
type EC2instance struct {
	ID     string
	Name   string
	Status string
	Type   string
	Region string
}

// EC2InstancesList of usage for EC2 instance
func EC2InstancesList() (ec2instances []EC2instance) {
	for _, region := range regions {
		log.Printf("Request session for region %s", region)
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(region),
		}))

		log.Printf("Fetching instance details for region %s", region)
		result, err := ec2.New(sess).DescribeInstances(&ec2.DescribeInstancesInput{})
		if err != nil {
			log.Printf("Error DescribeInstances: %v\n\n", err)
		} else {
			if len(result.Reservations) == 0 {
				log.Printf("Nothing instance found in region %s\n\n", region)
				continue
			}
			for _, reservation := range result.Reservations {
				for _, instance := range reservation.Instances {
					instanceName := ""
					for _, tag := range instance.Tags {
						if *tag.Key == "Name" {
							instanceName = *tag.Value
						}
					}
					log.Printf("Found Id[%v] Name[%v] Type[%v] Status[%v]", *instance.InstanceId, instanceName, *instance.InstanceType, *instance.State.Name)
					ec2instances = append(ec2instances, EC2instance{
						ID:     *instance.InstanceId,
						Name:   instanceName,
						Type:   *instance.InstanceType,
						Status: *instance.State.Name,
						Region: region,
					})
				}
			}

			log.Printf("Done for region %s\n", region)
		}
	}
	return ec2instances
}
