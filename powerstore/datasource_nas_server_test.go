/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// Test to Fetch NAS Servers
func TestAccNasServerDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				//Get all NAS Servers
				Config: ProviderConfigForTesting + NasServerDsConfig,
			},

			{
				//Get NAS Server by ID
				Config: ProviderConfigForTesting + NasServerIDDsConfig,
			},

			{
				//Get NAS Server by name
				Config: ProviderConfigForTesting + NasServerNameDsConfig,
			},

			{
				//Get NAS Server by Invalid Name
				Config:      ProviderConfigForTesting + NasServerNameDsConfigNeg,
				ExpectError: regexp.MustCompile(".*Unable to Read PowerStore NAS Servers.*"),
			},

			{
				//Get NAS Server by Invalid ID
				Config:      ProviderConfigForTesting + NasServerIDDsConfigNeg,
				ExpectError: regexp.MustCompile(".*Unable to Read PowerStore NAS Servers.*"),
			},
		},
	})
}

var NasServerDsConfig = `
data "powerstore_nas_server" "test" {
}
`
var NasServerIDDsConfig = `
data "powerstore_nas_server" "test" {
	id = "` + nasServerID + `"
}
`
var NasServerNameDsConfig = `
data "powerstore_nas_server" "test" {
	name = "` + nasServerName + `"
}
`

var NasServerNameDsConfigNeg = `
data "powerstore_nas_server" "test" {
	name = "invalid_nas00"
}
`
var NasServerIDDsConfigNeg = `
data "powerstore_nas_server" "test" {
	id = "0"
}
`
