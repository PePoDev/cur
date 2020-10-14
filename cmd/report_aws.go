package cmd

import (
	"bytes"
	"encoding/json"

	"github.com/pepodev/cur-cli/reporter"
	"github.com/pepodev/cur-cli/resources/aws"
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	reportCmd.AddCommand(awsCmd)

	awsCmd.Flags().StringP("region", "", "", "Target region to collect usage")
}

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Collect cloud usage data and make report",
	Run: func(cmd *cobra.Command, args []string) {
		cloud := aws.Cloud{}

		region, err := cmd.Flags().GetString("region")
		if err != nil {
			xlog.Fatalln("Flag region is required !")
		}

		xlog.Debugln("Request session for region", region)
		cloud.GetSession(region)

		xlog.Infof("Fetching instance details")
		if ec2instance, err := cloud.FetchEC2Instances(); err != nil {
			xlog.Panic(err.Error())
		} else if len(ec2instance) == 0 {
			xlog.Infof("Nothing instance found\n\n")
		} else {
			xlog.Debugln("Found", len(ec2instance), "instances")
			for _, instance := range ec2instance {
				xlog.Debugf("- Id[%v] Name[%v] Type[%v] Status[%v]", instance.ID, instance.Name, instance.Type, instance.Status)

			}
		}
		xlog.Infof("Fetched")

		cloud.FetchS3Buckets()
		cloud.FetchELBs()
		cloud.FetchEBSs()

		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err = encoder.Encode(cloud)
		if err != nil {
			xlog.Fatalln("Error Encode", err)
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			xlog.Fatalln(err)
		}

		switch output {
		case "json":
			xlog.Infoln("Cloud Usage Report for AWS at region", region, "->", buffer.String())
		case "slack":
			slackReport := reporter.SlackReporter{}
			slackReport.SetTitle(cloud.ReportTitle())
			slackReport.SetContent(buffer.String())
			slackReport.Send(viper.GetString("SLACK_WEBHOOK_URL"))
			xlog.Infoln("Send successful")
		}
	},
}
