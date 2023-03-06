package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch SnapshotRule
func TestAccSnapshotRule_FetchSnapshotRule(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: SnapshotRuleDataSourceparamsName,
			},
			{
				Config: SnapshotRuleDataSourceparamsID,
			},
			{
				Config: SnapshotRuleDataSourceparamsAll,
			},
			{
				Config:      SnapshotRuleDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Snapshot Rules"),
			},
		},
	})
}

var SnapshotRuleDataSourceparamsName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_snapshotrule" "test1" {
	name = "` + snapshotRuleName + `"
}
`
var SnapshotRuleDataSourceparamsNameNegative = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_snapshotrule" "test1" {
	name = "invalid-name"
}
`

var SnapshotRuleDataSourceparamsID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_snapshotrule" "test1" {
	id = "` + snapshotRuleID + `"
}
`

var SnapshotRuleDataSourceparamsAll = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_snapshotrule" "test1" {
}
`
