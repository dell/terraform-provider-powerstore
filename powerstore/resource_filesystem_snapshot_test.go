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

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create Snapshot Resource
func TestAccFileSystemSnapshot(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileSystemSnapParamsCreateInvalidFsID,
				ExpectError: regexp.MustCompile("Error creating filesystem snapshot"),
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test", "name", "tf_fs_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test", "description", "Test File System Snapshot Resource"),
				),
			},
			// Import Testing
			{
				Config:       ProviderConfigForTesting + FileSystemSnapParamsCreate,
				ResourceName: "powerstore_filesystem_snapshot.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_fs_snap_acc", s[0].Attributes["name"])
					return nil
				},
			},
			// Import Negative Testing
			{
				Config:        ProviderConfigForTesting + FileSystemSnapParamsCreate,
				ResourceName:  "powerstore_filesystem_snapshot.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportSnapshotDetailErrorMsg),
				ImportStateId: "invalid-id",
			},

			{
				Config: ProviderConfigForTesting + FileSystemSnapParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test", "expiration_timestamp", "2035-10-06T09:01:47Z"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test", "description", "Test File System Snapshot Resource Updated"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapParamsUpdateError,
				ExpectError: regexp.MustCompile("Error updating filesystem snapshot resource"),
			},
		},
	})
}

var FileSystemSnapParamsCreateInvalidFsID = `
resource "powerstore_filesystem_snapshot" "test" {
  filesystem_id="invalid"
}
`

var FileSystemSnapParamsCreate = `
resource "powerstore_filesystem_snapshot" "test" {
  name = "tf_fs_snap_acc"
  description = "Test File System Snapshot Resource"
  filesystem_id="`+ fileSystemID +`" 
  expiration_timestamp="2035-05-06T09:01:47Z"
  access_type = "Snapshot"
}
`
var FileSystemSnapParamsUpdate = `
resource "powerstore_filesystem_snapshot" "test" {
  name = "tf_fs_snap_acc"
  description = "Test File System Snapshot Resource Updated"
  filesystem_id="`+ fileSystemID +`" 
  expiration_timestamp="2035-10-06T09:01:47Z"
  access_type = "Snapshot"
}
`
var FileSystemSnapParamsUpdateError = `
resource "powerstore_filesystem_snapshot" "test" {
  name = "invalid"
  description = "Test File System Snapshot Resource Updated"
  filesystem_id="invalid" 
  expiration_timestamp="2035-10-06T09:01:47Z"
}
`
