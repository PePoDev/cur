// Package aws provide function to get resource usage from aws api
package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// GetSession return struct of aws session
func (cloud *Cloud) GetSession(region string) error {
	var err error
	cloud.session, err = session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	cloud.Region = region
	return err
}

// Cloud is contain all resource of aws
type Cloud struct {
	session *session.Session

	Region        string
	Ec2instances  []EC2instance
	S3buckets     []S3Bucket
	ELBs          []ELB
	EBSs          []EBS
	VPCs          []EBS
	Addresss      []EBS
	Disks         []EBS
	FirewallRules []EBS
}

// ReportTitle return title format
func (Cloud) ReportTitle() string {
	return "AWS Usage Report"
}

// String return string data of struct
func (cloud Cloud) String() string {
	return ""
}
