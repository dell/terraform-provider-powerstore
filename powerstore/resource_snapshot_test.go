package powerstore

import (
	"os"
	"regexp"
	"testing"

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
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "name", "test_snap"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test", "description", "Test Snapshot Resource"),
				),
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
				Config:      ProviderConfigForTesting + SnapshotParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
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
				Config:      ProviderConfigForTesting + SnapshotParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
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

var SnapParamsCreate = `
resource "powerstore_volume_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2023-05-06T09:01:47Z"
  creator_type="User"
}
`

var SnapshotParamsCreateWithoutName = `
resource "powerstore_volume_snapshot" "test" {
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2023-05-06T09:01:47Z"
  creator_type="User"
}
`

var SnapshotParamsCreateWithoutExpiry = `
resource "powerstore_volume_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_id="` + volumeID + `"
  performance_policy_id = "default_medium"
  creator_type="User"
}
`

var SnapParamsCreateWithoutVolume = `
resource "powerstore_volume_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  performance_policy_id = "default_medium"
  expiration_timestamp="2023-05-06T09:01:47Z"
  creator_type="User"
}
`
var SnapParamsCreateVolumeName = `
resource "powerstore_volume_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_name="` + volumeName + `"
  performance_policy_id = "default_medium"
  expiration_timestamp="2023-05-06T09:01:47Z"
  creator_type="User"
}
`
var SnapParamsCreateInvalidVolumeName = `
resource "powerstore_volume_snapshot" "test" {
  name = "test_snap"
  description = "Test Snapshot Resource"
  volume_name="random_volname"
  performance_policy_id = "default_medium"
  expiration_timestamp="2023-05-06T09:01:47Z"
  creator_type="User"
}
`
