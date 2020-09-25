package main

import (
	"fmt"

	"github.com/pepodev/cur/reporter/slack"
	"github.com/pepodev/cur/resources/aws"
)

func main() {
	slackReport := slack.Reporter{}
	slackReport.SetTitle(aws.Cloud{}.ReportTitle())
	slackReport.SetImageURL("https://media.makeameme.org/created/being-healthy-feeling.jpg")

	content := ""
	content += "##########################################\n"
	ec2instances := aws.EC2InstancesList()
	content += "_EC2 Instances Detail_\n"
	for _, instance := range ec2instances {
		content += fmt.Sprintf("* ID[%v] Name[%v] Type[%v] Status[%v] Region[%v]\n", instance.ID, instance.Name, instance.Type, instance.Status, instance.Region)
	}
	content += "##########################################\n"
	ec2buckets := aws.S3BucketsList("ap-southeast-1")
	content += "_S3 Buckets Detail_\n"
	for _, bucket := range ec2buckets {
		content += fmt.Sprintf("* ID[%v] CreationDate[%v]\n", bucket.Name, bucket.CreationDate)
	}
	content += "##########################################\n"

	slackReport.SetContent(content)
	slackReport.Send()
}
