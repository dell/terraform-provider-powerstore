package powerstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"testing"

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
		},
	})
}

// Test to volume id of Volume group snapshot
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

// Test to volume id of Volume group snapshot
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
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test", "volume_group_id", volumeGroupID),
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

// Test to import successfully
func TestAccVolumeGroupSnapshot_ImportSuccess(t *testing.T) {

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapParamsCreate,
			},
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

var VolumeGroupSnapParamsCreate = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_id="` + volumeGroupID + `"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapParamInvalidVolumeID = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_id="5c3c103a-9373-4f50-a34a"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapParamsRename = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap_new"
  description = "Test Snapshot Resource"
  volume_group_id="` + volumeGroupID + `"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapshotParamsCreateWithoutName = `
resource "powerstore_volumegroup_snapshot" "test" {
  description = "Test Snapshot Resource"
  volume_group_id="` + volumeGroupID + `"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var VolumeGroupSnapshotParamsCreateWithoutExpiry = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_id="` + volumeGroupID + `"
}
`

var VolumeGroupSnapParamsCreateWithoutVolume = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  expiration_timestamp="2035-05-06T09:01:47Z"
}
`

var SnapParamsCreateVolumeGroupName = `
resource "powerstore_volumegroup_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_group_name="` + volumeGroupName + `"
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
