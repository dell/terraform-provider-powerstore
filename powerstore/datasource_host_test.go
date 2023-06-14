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

// Test to Fetch Host details
func TestAccHost_FetchHost(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostDataSourceParamsName,
			},
			{
				Config: ProviderConfigForTesting + HostDataSourceParamsID,
			},
			{
				Config: ProviderConfigForTesting + HostDataSourceParamsAll,
			},
		},
	})
}

// Test to fetch Host - Negative
func TestAccHost_FetchHostNegative(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostDataSourceParamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Host"),
			},
			{
				Config:      ProviderConfigForTesting + HostDataSourceParamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Host"),
			},
			{
				Config:      ProviderConfigForTesting + HostDataSourceParamsEmptyID,
				ExpectError: regexp.MustCompile(InvalidLengthErrorMsg),
			},
			{
				Config:      ProviderConfigForTesting + HostDataSourceParamsEmptyName,
				ExpectError: regexp.MustCompile(InvalidLengthErrorMsg),
			},
		},
	})
}

var HostDataSourceParamsName = `
data "powerstore_host" "test1" {
	name = "` + hostNameRead + `"
}
`
var HostDataSourceParamsNameNegative = `
data "powerstore_host" "test1" {
	name = "invalid-name"
}
`

var HostDataSourceParamsIDNegative = `
data "powerstore_host" "test1" {
	id = "invalid-id"
}
`

var HostDataSourceParamsID = `
data "powerstore_host" "test1" {
	id = "` + hostIDRead + `"
}
`

var HostDataSourceParamsEmptyName = `
data "powerstore_host" "test1" {
	name = ""
}
`

var HostDataSourceParamsEmptyID = `
data "powerstore_host" "test1" {
	id = ""
}
`

var HostDataSourceParamsAll = `
data "powerstore_host" "test1" {
}
`
