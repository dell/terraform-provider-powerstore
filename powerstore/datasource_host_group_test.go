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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Host Groups
func TestAccHostGroupDs_FetchHostGroup(t *testing.T) {
	// if os.Getenv("TF_ACC") == "" {
	// 	t.Skip("Dont run with units tests because it will try to create the context")
	// }

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostGroupDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + HostGroupDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + HostGroupDataSourceparamsAll,
			},
			{
				//Get Host Group by Filter Expression
				Config: ProviderConfigForTesting + HostGroupDataSourceFilterConfig,
			},
			{
				Config:      ProviderConfigForTesting + HostGroupDataSourceparamsIDAndNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + HostGroupDataSourceparamsEmptyIDNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + HostGroupDataSourceparamsEmptyNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + HostGroupDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Host Group"),
			},
			{
				Config:      ProviderConfigForTesting + HostGroupDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Host Group"),
			},
			{
				Config: ProviderConfigForTesting + HostGroupParamsUpdateRemoveHost,
			},
			{
				//Get Host Group by Invalid Filter Expression
				Config:      ProviderConfigForTesting + HostGroupDataSourceFilterConfigNeg,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Host Group"),
			},
		},
	})
}

var HostGroupDataSourceparamsName = HostGroupParamsCreate + `
data "powerstore_hostgroup" "test1" {
	depends_on = [powerstore_hostgroup.test]
	name = powerstore_hostgroup.test.name
}

`

var HostGroupDataSourceFilterConfig = HostGroupParamsCreate + `
data "powerstore_hostgroup" "test1" {
	depends_on = [powerstore_hostgroup.test]
	filter_expression = format("name=eq.%s",powerstore_hostgroup.test.name)
}
`
var HostGroupDataSourceFilterConfigNeg = `
data "powerstore_hostgroup" "test" {
  filter_expression = "name=invalidName"
}
`

var HostGroupDataSourceparamsNameNegative = `
data "powerstore_hostgroup" "test1" {
	name = "invalid-name"
}
`

var HostGroupDataSourceparamsID = HostGroupParamsCreate + `
data "powerstore_hostgroup" "test1" {
	id = powerstore_hostgroup.test.id
}
`

var HostGroupDataSourceparamsAll = HostGroupParamsCreate + `
data "powerstore_hostgroup" "test1" {
	depends_on = [powerstore_hostgroup.test]
}
`

var HostGroupDataSourceparamsIDNegative = `
data "powerstore_hostgroup" "test1" {
	id = "invalid-id"
}
`

var HostGroupDataSourceparamsIDAndNameNegative = HostGroupParamsCreate + `
data "powerstore_hostgroup" "test1" {
	id = powerstore_hostgroup.test.id
	name = powerstore_hostgroup.test.name
}
`

var HostGroupDataSourceparamsEmptyIDNegative = `
data "powerstore_hostgroup" "test1" {
	id = ""
}
`

var HostGroupDataSourceparamsEmptyNameNegative = `
data "powerstore_hostgroup" "test1" {
	name = ""
}
`
