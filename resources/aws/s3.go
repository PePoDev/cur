package aws

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pepodev/xlog"
)

// S3Bucket structure data for instance on aws
type S3Bucket struct {
	Name         string
	CreationDate string
}

// FetchS3Buckets of usage for EC2 instance
func (cloud *Cloud) FetchS3Buckets() {
	xlog.Infof("Fetching s3 bucket details for region %s", cloud.Region)
	result, err := s3.New(cloud.session).ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		xlog.Infof("Error ListBucketsInput: %v\n\n", err)
	}

	for _, bucket := range result.Buckets {
		xlog.Infof("Found Name[%v] CreationDate[%v]", *bucket.Name, bucket.CreationDate.String())
		cloud.S3buckets = append(cloud.S3buckets, S3Bucket{
			Name:         *bucket.Name,
			CreationDate: bucket.CreationDate.String(),
		})
	}

	xlog.Infof("Done for region %s\n", cloud.Region)
}
