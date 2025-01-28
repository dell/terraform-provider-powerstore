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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create NFSExport
func TestAccNFSExport_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Create Error
			{
				Config:      ProviderConfigForTesting + nfsCreateError,
				ExpectError: regexp.MustCompile(".*Error creating nfs export.*"),
			},
			// Create Testing
			{
				Config: ProviderConfigForTesting + nfsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "name", "terraform_nfs"),
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "description", "terraform nfs export"),
				),
			},
			// Import Testing
			{
				Config:       ProviderConfigForTesting + nfsCreate,
				ResourceName: "powerstore_nfs_export.test1",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "terraform_nfs", s[0].Attributes["name"])
					assert.Equal(t, "terraform nfs export", s[0].Attributes["description"])
					return nil
				},
			},
			// Update Testing
			{
				Config: ProviderConfigForTesting + nfsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "name", "terraform_nfs"),
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "description", "terraform nfs export update"),
				),
			},
			// Update Error
			{
				Config:      ProviderConfigForTesting + nfsInvalidUpate,
				ExpectError: regexp.MustCompile(".*Error updating nfs export resource.*"),
			},
			// Import Error
			{
				Config:        ProviderConfigForTesting + nfsCreate,
				ResourceName:  "powerstore_nfs_export.test1",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(".*Error reading nfs export.*"),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var nfsCreate = FsParams + `
resource "powerstore_nfs_export" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_nfs"
  path = "/test_fs"

  anonymous_gid = -23
  anonymous_uid = -23 
  description = "terraform nfs export"
  is_no_suid = true 

  min_security = "Sys"
  default_access = "Read_Only"
}
`

var nfsCreateError = `
resource "powerstore_nfs_export" "test2" {
  file_system_id = "invalid"
  name = "terraform_nfs"
  path = "/test_fs"
}
`

var nfsUpdate = FsParams + `
resource "powerstore_nfs_export" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_nfs"
  path = "/test_fs"

  anonymous_gid = 23
  anonymous_uid = 23 
  description = "terraform nfs export update"
  is_no_suid = false 

  min_security = "Kerberos"
  default_access = "Read_Write"
}
`

var nfsInvalidUpate = FsParams + `
resource "powerstore_nfs_export" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform-nfs1"
  path = "/test_fs1"
}
`
