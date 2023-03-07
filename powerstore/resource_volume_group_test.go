package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create VolumeGroup
func TestAccVolumeGroup_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "test_volume_group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
		},
	})
}

// Test to Create VolumeGroup without required field, will result in error
func TestAccVolumeGroup_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      VolumeGroupParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateVolumeGroupMissingErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup with Invalid protection policy ID, will result in error
func TestAccVolumeGroup_CreateWithInvalidPolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      VolumeGroupParamsCreateWithInvalidPolicy,
				ExpectError: regexp.MustCompile(CreateVolumeGroupDetailErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup with Invalid volume ID, will result in error
func TestAccVolumeGroup_CreateWithInvalidVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      VolumeGroupParamsCreateWithInvalidVolume,
				ExpectError: regexp.MustCompile(CreateVolumeGroupDetailErrorMsg),
			},
		},
	})
}

var VolumeGroupParamsCreate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_volumegroup" "test" {
  name = "test_volume_group"
  description = "Creating Volume Group"
  is_write_order_consistent = false
}
`

var VolumeGroupParamsCreateWithoutName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_volumegroup" "test" {
	description = "Create volume group without name"
}
`

var VolumeGroupParamsCreateWithInvalidPolicy = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_volumegroup" "test" {
	name = "test_volume_group"
	description = "Create volume group without name"
	is_write_order_consistent = false
	protection_policy_id = "invalid-id"
}
`

var VolumeGroupParamsCreateWithInvalidVolume = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_volumegroup" "test" {
	name = "test_volume_group"
	description = "Create volume group without name"
	is_write_order_consistent = false
	volume_ids = ["invalid-id"]
}
`
