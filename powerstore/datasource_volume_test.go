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

// Test to Fetch Volume
func TestAccVolumeDs_FetchVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + VolumeDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volumes"),
			},
		},
	})
}

var VolumeDataSourceparamsName = VolumeParams + `
data "powerstore_volume" "test1" {
	name = powerstore_volume.volume_create_test.name
}
`
var VolumeDataSourceparamsNameNegative = `
data "powerstore_volume" "test1" {
	name = "invalid-name"
}
`

var VolumeDataSourceparamsID = VolumeParams + `
data "powerstore_volume" "test1" {
	id = powerstore_volume.volume_create_test.id
}
`

var VolumeDataSourceparamsAll = VolumeParams + `
data "powerstore_volume" "test1" {
	depends_on = [powerstore_volume.volume_create_test]
}
`
