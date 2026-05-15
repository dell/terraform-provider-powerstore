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

// Test U-21: Create secure filesystem snapshot
func TestAccFileSystemSnapshot_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureFSSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "name", "tf_fs_snap_secure_acc"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "expiration_timestamp", "2035-05-06T09:01:47Z"),
				),
			},
		},
	})
}

// Test U-22: Create secure FS snapshot missing expiration - should fail
func TestAccFileSystemSnapshot_CreateSecureMissingExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SecureFSSnapParamsCreateNoExpiration,
				ExpectError: regexp.MustCompile(".*[Ss]ecure snapshots require.*expiration.*"),
			},
		},
	})
}

// Test U-24/U-25: Update FS snapshot - one-way lock
func TestAccFileSystemSnapshot_UpdateToSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + NonSecureFSSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + FSSnapParamsUpdateToSecure,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_nonsecure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test U-25: Reject unsecure FS snapshot
func TestAccFileSystemSnapshot_RejectUnsecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureFSSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FSSnapParamsRevertSecure,
				ExpectError: regexp.MustCompile(".*[Oo]ne.way lock.*|.*cannot.*unlock.*|.*cannot.*change.*true.*false.*"),
			},
		},
	})
}

// --- HCL Test Configs ---

var SecureFSSnapParamsCreate = FsParams + `
resource "powerstore_filesystem_snapshot" "test_secure" {
  name = "tf_fs_snap_secure_acc"
  description = "Test Secure FS Snapshot"
  filesystem_id = powerstore_filesystem.test_fs_create.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type = "Snapshot"
  is_secure = true
  depends_on = [powerstore_filesystem.test_fs_create]
}
`

var SecureFSSnapParamsCreateNoExpiration = FsParams + `
resource "powerstore_filesystem_snapshot" "test_secure_no_exp" {
  name = "tf_fs_snap_secure_no_exp"
  description = "Secure FS Snapshot Without Expiration"
  filesystem_id = powerstore_filesystem.test_fs_create.id
  access_type = "Snapshot"
  is_secure = true
  depends_on = [powerstore_filesystem.test_fs_create]
}
`

var NonSecureFSSnapParamsCreate = FsParams + `
resource "powerstore_filesystem_snapshot" "test_nonsecure" {
  name = "tf_fs_snap_nonsecure_acc"
  description = "Non-Secure FS Snapshot"
  filesystem_id = powerstore_filesystem.test_fs_create.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type = "Snapshot"
  depends_on = [powerstore_filesystem.test_fs_create]
}
`

var FSSnapParamsUpdateToSecure = FsParams + `
resource "powerstore_filesystem_snapshot" "test_nonsecure" {
  name = "tf_fs_snap_nonsecure_acc"
  description = "Non-Secure FS Snapshot"
  filesystem_id = powerstore_filesystem.test_fs_create.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type = "Snapshot"
  is_secure = true
  depends_on = [powerstore_filesystem.test_fs_create]
}
`

var FSSnapParamsRevertSecure = FsParams + `
resource "powerstore_filesystem_snapshot" "test_secure" {
  name = "tf_fs_snap_secure_acc"
  description = "Test Secure FS Snapshot"
  filesystem_id = powerstore_filesystem.test_fs_create.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type = "Snapshot"
  is_secure = false
  depends_on = [powerstore_filesystem.test_fs_create]
}
`
