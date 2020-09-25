// Package aws provide function to get resource usage from aws api
package aws

var (
	regions = []string{
		"us-east-2",
		"us-east-1",
		"us-west-1",
		"us-west-2",
		"af-south-1",
		"ap-east-1",
		"ap-south-1",
		"ap-northeast-3",
		"ap-northeast-2",
		"ap-northeast-1",
		"ap-southeast-2",
		"ap-southeast-1",
		"ca-central-1",
		"cn-north-1",
		"cn-northwest-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-south-1",
		"eu-west-3",
		"eu-north-1",
		"me-south-1",
		"sa-east-1",
		"us-gov-east-1",
		"us-gov-west-1",
	}
)

// Cloud is contain all resource of aws
type Cloud struct {
}

// Resources is contain all resource of aws
type Resources struct {
}

// GetAllResources retrun string format with all resources usage
func GetAllResources() string {
	return ""
}

// ReportTitle return title format
func (Cloud) ReportTitle() string {
	return "AWS Usage Reporting"
}

// TODO: list of resource need to catch up !! disks addresses fwRules vpc s3
