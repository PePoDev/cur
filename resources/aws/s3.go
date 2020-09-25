package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Bucket structure data for instance on aws
type S3Bucket struct {
	Name         string
	CreationDate string
}

// S3BucketsList of usage for EC2 instance
func S3BucketsList(region string) (s3buckets []S3Bucket) {
	log.Printf("Request session for region %s", region)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	log.Printf("Fetching instance details for region %s", region)
	result, err := s3.New(sess).ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Printf("Error ListBucketsInput: %v\n\n", err)
	} else {
		for _, bucket := range result.Buckets {
			log.Printf("Found Name[%v] CreationDate[%v]", *bucket.Name, bucket.CreationDate.String())
			s3buckets = append(s3buckets, S3Bucket{
				Name:         *bucket.Name,
				CreationDate: bucket.CreationDate.String(),
			})
		}

		log.Printf("Done for region %s\n", region)
	}

	return s3buckets
}
