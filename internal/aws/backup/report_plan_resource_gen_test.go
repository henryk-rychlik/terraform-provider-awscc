// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSBackupReportPlan_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Backup::ReportPlan", "awscc_backup_report_plan", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
