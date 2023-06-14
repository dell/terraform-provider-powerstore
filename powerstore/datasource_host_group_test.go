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

// Test to Fetch Host Groups
func TestAccHostGroup_FetchHostGroup(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

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
		},
	})
}

var HostGroupDataSourceparamsName = `
data "powerstore_hostgroup" "test1" {
	name = "` + hostGroupName + `"
}
`

var HostGroupDataSourceparamsNameNegative = `
data "powerstore_hostgroup" "test1" {
	name = "invalid-name"
}
`

var HostGroupDataSourceparamsID = `
data "powerstore_hostgroup" "test1" {
	id = "` + hostGroupID + `"
}
`

var HostGroupDataSourceparamsAll = `
data "powerstore_hostgroup" "test1" {
}
`

var HostGroupDataSourceparamsIDNegative = `
data "powerstore_hostgroup" "test1" {
	id = "invalid-id"
}
`

var HostGroupDataSourceparamsIDAndNameNegative = `
data "powerstore_hostgroup" "test1" {
	id = "` + hostGroupID + `"
	name = "` + hostGroupName + `"
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
