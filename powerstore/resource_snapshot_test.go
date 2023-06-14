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

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create Snapshot Resource
func TestAccVolumeSnapshot_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "tf_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "description", "Test Snapshot Resource"),
				),
			},
		},
	})
}

// Test to add invalid volume id to the Volume snapshot
func TestAccVolumeSnapshot_InvalidSnapshotVolumeID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SnapParamInvalidVolumeID,
				ExpectError: regexp.MustCompile(CreateSnapshotErrorMsg),
			},
		},
	})
}

// Test to Rename Volume snapshot
func TestAccVolumeSnapshot_UpdateSnapshotRename(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsCreate,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "tf_snap_acc")),
			},
			{
				Config: ProviderConfigForTesting + SnapParamsRename,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "tf_snap_acc_new")),
			},
		},
	})
}

// Test to update volume id of Volume snapshot
func TestAccVolumeSnapshot_UpdateSnapshotVolumeName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsCreate,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "tf_snap_acc")),
			},
			{
				Config:      ProviderConfigForTesting + SnapParamInvalidVolumeID,
				ExpectError: regexp.MustCompile(VolumeIDNameUpdateErrorMsg),
			},
		},
	})
}

// Test to Create Snapshot Resource Without Name
func TestAccVolumeSnapshot_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapshotParamsCreateWithoutName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "description", "Test Snapshot Resource"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "expiration_timestamp", "2035-05-06T09:01:47Z"),
				),
			},
		},
	})
}

// Test to Create Snapshot Resource Without Expiration timeout
func TestAccVolumeSnapshot_CreateWithoutExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapshotParamsCreateWithoutExpiry,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "tf_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "volume_id", volumeID),
				),
			},
		},
	})
}

// Test to Create Snapshot Resource Without volume ID
func TestAccVolumeSnapshot_CreateWithoutVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SnapParamsCreateWithoutVolume,
				ExpectError: regexp.MustCompile(InvalidAttributeCombinationErrorMsg),
			},
		},
	})
}

// Test to Create Snapshot Resource With volume Name
func TestAccVolumeSnapshot_CreateWithVolumeName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsCreateVolumeName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "tf_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "volume_name", volumeName),
				),
			},
		},
	})
}

// Test to Create Snapshot Resource With Invalid volume Name
func TestAccVolumeSnapshot_CreateWithInvalidVolumeName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SnapParamsCreateInvalidVolumeName,
				ExpectError: regexp.MustCompile(CreateSnapshotErrorMsg),
			},
		},
	})
}

// Negative test case for import
func TestAccVolumeSnapshot_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + SnapParamsCreate,
				ResourceName:  "powerstore_volume_snapshot.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportSnapshotDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

// Test to import successfully
func TestAccVolumeSnapshot_ImportSuccess(t *testing.T) {

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsCreate,
			},
			{
				Config:       ProviderConfigForTesting + SnapParamsCreate,
				ResourceName: "powerstore_volume_snapshot.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_snap_acc", s[0].Attributes["name"])
					return nil
				},
			},
		},
	})
}

var SnapParamsCreate = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc"
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
  creator_type = "User"
}
`

var SnapParamInvalidVolumeID = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc"
  description = "Test Snapshot Resource"
  volume_id="05a959ed-6545-48fb-9887-ce4"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var SnapParamsRename = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc_new"
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`
var SnapshotParamsCreateWithoutName = `
resource "powerstore_volume_snapshot" "test" {
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var SnapshotParamsCreateWithoutExpiry = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc"
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
}
`

var SnapParamsCreateWithoutVolume = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc"
  description = "Test Snapshot Resource"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`
var SnapParamsCreateVolumeName = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc"
  description = "Test Snapshot Resource"
  volume_name="` + volumeName + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`
var SnapParamsCreateInvalidVolumeName = `
resource "powerstore_volume_snapshot" "test" {
  name = "tf_snap_acc"
  description = "Test Snapshot Resource"
  volume_name="random_volname"
  performance_policy_id = "default_medium"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`
