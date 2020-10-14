package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ebs"
)

// EBS structure data on aws
type EBS struct {
	ID     string
	Name   string
	Status string
	Type   string
	Region string
}

// FetchEBSs get all EBS resource on aws
func (cloud *Cloud) FetchEBSs() {
	svc := ebs.New(cloud.session)
	x := ""
	result, err := svc.ListSnapshotBlocks(&ebs.ListSnapshotBlocksInput{
		SnapshotId: &x,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ebs.ErrCodeAccessDeniedException:
				fmt.Println(ebs.ErrCodeAccessDeniedException, aerr.Error())
			case ebs.ErrCodeValidationException:
				fmt.Println(ebs.ErrCodeValidationException, aerr.Error())
			case ebs.ErrCodeResourceNotFoundException:
				fmt.Println(ebs.ErrCodeResourceNotFoundException, aerr.Error())
			case ebs.ErrCodeRequestThrottledException:
				fmt.Println(ebs.ErrCodeRequestThrottledException, aerr.Error())
			case ebs.ErrCodeServiceQuotaExceededException:
				fmt.Println(ebs.ErrCodeServiceQuotaExceededException, aerr.Error())
			case ebs.ErrCodeInternalServerException:
				fmt.Println(ebs.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
