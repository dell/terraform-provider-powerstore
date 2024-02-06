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

// Test to Fetch Volume Groups
func TestAccVolumeGroupDs_FetchVolumeGroup(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupDataSourceparamsIDAndNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupDataSourceparamsEmptyIDNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupDataSourceparamsEmptyNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
		},
	})
}

var VolumeGroupDataSourceparamsName = `
data "powerstore_volumegroup" "test1" {
	name = "` + volumeGroupName + `"
}
`

var VolumeGroupDataSourceparamsNameNegative = `
data "powerstore_volumegroup" "test1" {
	name = "invalid-name"
}
`

var VolumeGroupDataSourceparamsID = `
data "powerstore_volumegroup" "test1" {
	id = "` + volumeGroupID + `"
}
`

var VolumeGroupDataSourceparamsAll = `
data "powerstore_volumegroup" "test1" {
}
`

var VolumeGroupDataSourceparamsIDNegative = `
data "powerstore_volumegroup" "test1" {
	id = "invalid-id"
}
`

var VolumeGroupDataSourceparamsIDAndNameNegative = `
data "powerstore_volumegroup" "test1" {
	id = "` + volumeGroupID + `"
	name = "` + volumeGroupName + `"
}
`

var VolumeGroupDataSourceparamsEmptyIDNegative = `
data "powerstore_volumegroup" "test1" {
	id = ""
}
`

var VolumeGroupDataSourceparamsEmptyNameNegative = `
data "powerstore_volumegroup" "test1" {
	name = ""
}
`
