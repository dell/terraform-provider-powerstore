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

// Test to Create VolumeGroup Snapshot Resource
func TestAccVolumeGroupSnapshot_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "name", "test_snap"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "description", "Test Snapshot Resource"),
				),
			},
			// Import Test
			{
				Config:            ProviderConfigForTesting + VolumeGroupSnapParamsCreate,
				ResourceName:      "powerstore_volumegroup_snapshot.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "test_snap", s[0].Attributes["name"])
					return nil
				},
			},
		},
	})
}

// Test to add invalid volume group id to the Volume group snapshot
func TestAccVolumeGroupSnapshot_InvalidSnapshotVolumegroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapParamInvalidVolumeID,
				ExpectError: regexp.MustCompile(CreateVolumeGroupSnapshotErrorMsg),
			},
		},
	})
}

// Test to Rename Volume group snapshot
func TestAccVolumeGroupSnapshot_UpdateSnapshotRename(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapParamsCreate,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "name", "test_snap")),
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapParamsRename,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "name", "test_snap_new")),
			},
		},
	})
}

// Test to update volume group id of Volume group snapshot
func TestAccVolumeGroupSnapshot_UpdateSnapshotVolumeName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapParamsCreate,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "name", "test_snap")),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapParamInvalidVolumeID,
				ExpectError: regexp.MustCompile(VolumeGroupIDNameUpdateErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup Snapshot Resource Without Name
func TestAccVolumeGroupSnapshot_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup Snapshot Resource Without Expiration timeout
func TestAccVolumeGroupSnapshot_CreateWithoutExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotParamsCreateWithoutExpiry,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "name", "test_snap"),
				),
			},
		},
	})
}

// Test to Create VolumeGroup Snapshot Resource Without volume group name and ID
func TestAccVolumeGroupSnapshot_CreateWithoutVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapParamsCreateWithoutVolume,
				ExpectError: regexp.MustCompile(InvalidAttributeCombinationErrorMsg),
			},
		},
	})
}

// Test to Create Snapshot Resource With volume group Name
func TestAccVolumeGroupSnapshot_CreateWithVolumeGroupName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsCreateVolumeGroupName,
			},
		},
	})
}

// Test to Create Snapshot Resource With Invalid volume group Name
func TestAccVolumeGroupSnapshot_CreateWithInvalidVolumeGroupName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SnapParamsCreateInvalidVolumeGroupName,
				ExpectError: regexp.MustCompile(CreateVolumeGroupSnapshotErrorMsg),
			},
		},
	})
}

// Negative test case for import
func TestAccVolumeGroupSnapshot_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + VolumeGroupSnapParamsCreate,
				ResourceName:  "powerstore_volumegroup_snapshot.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportSnapshotDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var PreReqVolumeGroupSnap = PreReqVolumeGroup + `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  volume_group_id= powerstore_volumegroup.test.id
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`
var VolumeGroupSnapParamsCreate = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_id= powerstore_volumegroup.test.id
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapParamInvalidVolumeID = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test" {
  depends_on = [powerstore_volumegroup.test]
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_id="5c3c103a-9373-4f50-a34a"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapParamsRename = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap_new"
  description = "Test Snapshot Resource"
  volume_group_id= powerstore_volumegroup.test.id
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapshotParamsCreateWithoutName = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test" {
  description = "Test Snapshot Resource"
  volume_group_id= powerstore_volumegroup.test.id
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapshotParamsCreateWithoutExpiry = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_id= powerstore_volumegroup.test.id
}
`

var VolumeGroupSnapParamsCreateWithoutVolume = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var SnapParamsCreateVolumeGroupName = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "test" {
  depends_on = [powerstore_volumegroup.test]
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_name= powerstore_volumegroup.test.name
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var SnapParamsCreateInvalidVolumeGroupName = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_name="random_volgroup"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`
