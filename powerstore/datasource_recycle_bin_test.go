/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// Test to Fetch RecycleBin items
func TestAccRecycleBinDs_FetchAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceAll,
			},
		},
	})
}

// Test to Fetch RecycleBin items by resource_type
func TestAccRecycleBinDs_FetchByResourceType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceByVolume,
			},
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceByVolumeGroup,
			},
		},
	})
}

// Test to Fetch RecycleBin item by invalid ID
func TestAccRecycleBinDs_FetchByInvalidID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RecycleBinDataSourceByInvalidID,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Recycle Bin Item"),
			},
		},
	})
}

// Test to Fetch RecycleBin items by filter expression
func TestAccRecycleBinDs_FetchByFilter(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceByFilter,
			},
		},
	})
}

var RecycleBinDataSourceAll = `
data "powerstore_recycle_bin" "test" {
}
`

var RecycleBinDataSourceByVolume = `
data "powerstore_recycle_bin" "test" {
	resource_type = "volume"
}
`

var RecycleBinDataSourceByVolumeGroup = `
data "powerstore_recycle_bin" "test" {
	resource_type = "volume_group"
}
`

var RecycleBinDataSourceByInvalidID = `
data "powerstore_recycle_bin" "test" {
	id = "invalid-id-does-not-exist"
}
`

var RecycleBinDataSourceByFilter = `
data "powerstore_recycle_bin" "test" {
	filter_expression = "resource_type=eq.volume"
}
`
