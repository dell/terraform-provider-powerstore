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

// Test to Fetch NFS Exports
func TestAccNFSExportDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// validate nfs export with invalid url query expression
			{
				Config: ProviderConfigForTesting + `
				data "powerstore_nfs_export" "test" {
					filter_expression = "invalid;"
				}
				`,
				ExpectError: regexp.MustCompile(".*Invalid PowerStore filter expression.*"),
			},
			// validate nfs export with empty filter
			{
				Config: ProviderConfigForTesting + `
				data "powerstore_nfs_export" "test" {
					filter_expression = ""
				}
				`,
				ExpectError: regexp.MustCompile(".*got empty string value.*"),
			},
			{
				// Get all NFS Exports
				Config: ProviderConfigForTesting + `
				data "powerstore_nfs_export" "test" {
				}
				`,
			},
			{
				// Get NFS Export by Invalid ID
				Config: ProviderConfigForTesting + `
				data "powerstore_nfs_export" "test" {
					id = "invalid"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read nfs export with id.*"),
			},
			{
				// Get NFS Export by invalid filter expression
				Config: ProviderConfigForTesting + `
				data "powerstore_nfs_export" "test" {
					filter_expression = "name=inv.terraform"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read nfs exports with error.*"),
			},
			{
				//Get NFS Export by ID
				Config: ProviderConfigForTesting + nfsCreate + `
				data "powerstore_nfs_export" "test" {
					id = powerstore_nfs_export.test1.id
				}
				`,
			},
			{
				// Get NFS Export by name and filesystem id
				Config: ProviderConfigForTesting + nfsCreate + `
				data "powerstore_nfs_export" "test" {
					name = powerstore_nfs_export.test1.name
					file_system_id = powerstore_nfs_export.test1.file_system_id
				}
				`,
			},
		},
	})
}
