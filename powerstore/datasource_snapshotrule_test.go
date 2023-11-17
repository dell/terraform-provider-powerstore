/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch SnapshotRule
func TestAccSnapshotRuleDs_FetchSnapshotRule(t *testing.T) {
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
				Config:      SnapshotRuleDataSourceparamsNameEmpty,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Snapshot Rules"),
			},
			{
				Config: SnapshotRuleDataSourceparamsID,
			},
			{
				Config:      SnapshotRuleDataSourceparamsIDEmpty,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Snapshot Rules"),
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
var SnapshotRuleDataSourceparamsNameEmpty = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_snapshotrule" "test1" {
	name = " "
}
`

var SnapshotRuleDataSourceparamsIDEmpty = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_snapshotrule" "test1" {
	id = " "
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
