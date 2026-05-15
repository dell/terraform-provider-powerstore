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

// Test U-16: Create secure volume group snapshot
func TestAccVolumeGroupSnapshot_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureVGSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "name", "test_vg_snap_secure"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "expiration_timestamp", "2035-05-06T09:01:47Z"),
				),
			},
		},
	})
}

// Test U-17: Create secure VG snapshot without expiration - should fail
func TestAccVolumeGroupSnapshot_CreateSecureMissingExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SecureVGSnapParamsCreateNoExpiration,
				ExpectError: regexp.MustCompile(".*[Ss]ecure snapshots require.*expiration.*"),
			},
		},
	})
}

// Test U-18/U-19: Update VG snapshot - one-way lock
func TestAccVolumeGroupSnapshot_UpdateToSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + NonSecureVGSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + VGSnapParamsUpdateToSecure,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_nonsecure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test U-19: Reject unsecure VG snapshot
func TestAccVolumeGroupSnapshot_RejectUnsecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureVGSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:      ProviderConfigForTesting + VGSnapParamsRevertSecure,
				ExpectError: regexp.MustCompile(".*[Oo]ne.way lock.*|.*cannot.*unlock.*|.*cannot.*change.*true.*false.*"),
			},
		},
	})
}

// --- HCL Test Configs ---

var SecureVGSnapParamsCreate = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test_secure" {
  name = "test_vg_snap_secure"
  description = "Secure VG Snapshot"
  volume_group_id = powerstore_volumegroup.test.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = true
}
`

var SecureVGSnapParamsCreateNoExpiration = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test_secure_no_exp" {
  name = "test_vg_snap_secure_no_exp"
  description = "Secure VG Snapshot Without Expiration"
  volume_group_id = powerstore_volumegroup.test.id
  is_secure = true
}
`

var NonSecureVGSnapParamsCreate = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test_nonsecure" {
  name = "test_vg_snap_nonsecure"
  description = "Non-Secure VG Snapshot"
  volume_group_id = powerstore_volumegroup.test.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var VGSnapParamsUpdateToSecure = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test_nonsecure" {
  name = "test_vg_snap_nonsecure"
  description = "Non-Secure VG Snapshot"
  volume_group_id = powerstore_volumegroup.test.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = true
}
`

var VGSnapParamsRevertSecure = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test_secure" {
  name = "test_vg_snap_secure"
  description = "Secure VG Snapshot"
  volume_group_id = powerstore_volumegroup.test.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = false
}
`
