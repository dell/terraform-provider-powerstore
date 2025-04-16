/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
				Config: ProviderConfigForTesting + SnapshotRuleDataSourceparamsName,
			},
			{
				Config:      ProviderConfigForTesting + SnapshotRuleDataSourceparamsNameEmpty,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Snapshot Rules"),
			},
			{
				Config: ProviderConfigForTesting + SnapshotRuleDataSourceparamsID,
			},
			{
				Config:      ProviderConfigForTesting + SnapshotRuleDataSourceparamsIDEmpty,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Snapshot Rules"),
			},
			{
				Config: ProviderConfigForTesting + SnapshotRuleDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + SnapshotRuleDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Snapshot Rules"),
			},
			{
				//Get Snapshot Rule by Filter Expression
				Config: ProviderConfigForTesting + SnapshotRuleFilterConfig,
			},
		},
	})
}

var SnapshotRuleDataSourceparamsName = SnapshotRuleParamsWithTimeOfDay + `
data "powerstore_snapshotrule" "test1" {
	depends_on = [powerstore_snapshotrule.test]
	name = powerstore_snapshotrule.test.name
}
`

var SnapshotRuleFilterConfig = SnapshotRuleParamsWithTimeOfDay + `
data "powerstore_snapshotrule" "test1" {
	depends_on = [powerstore_snapshotrule.test]
	filter_expression = format("name=eq.%s",powerstore_snapshotrule.test.name)
}
`
var SnapshotRuleFilterConfigNeg = `
data "powerstore_snapshotrule" "test1" {
  filter_expression = "name=invalidName"
}
`

var SnapshotRuleDataSourceparamsNameEmpty = `
data "powerstore_snapshotrule" "test1" {
	name = " "
}
`

var SnapshotRuleDataSourceparamsIDEmpty = `
data "powerstore_snapshotrule" "test1" {
	id = " "
}
`

var SnapshotRuleDataSourceparamsNameNegative = `
data "powerstore_snapshotrule" "test1" {
	name = "invalid-name"
}
`

var SnapshotRuleDataSourceparamsID = SnapshotRuleParamsWithTimeOfDay + `
data "powerstore_snapshotrule" "test1" {
	id = powerstore_snapshotrule.test.id
}
`

var SnapshotRuleDataSourceparamsAll = SnapshotRuleParamsWithTimeOfDay + `
data "powerstore_snapshotrule" "test1" {
	depends_on = [powerstore_snapshotrule.test]
}
`
