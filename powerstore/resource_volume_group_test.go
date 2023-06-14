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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
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
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
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
				Config:      ProviderConfigForTesting + VolumeGroupParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
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
				Config:      ProviderConfigForTesting + VolumeGroupParamsCreateWithInvalidPolicy,
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
				Config:      ProviderConfigForTesting + VolumeGroupParamsCreateWithInvalidVolume,
				ExpectError: regexp.MustCompile(CreateVolumeGroupDetailErrorMsg),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params
func TestAccVolumeGroup_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params, will result in error
func TestAccVolumeGroup_UpdateError(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupParamsUpdateServerError,
				ExpectError: regexp.MustCompile(UpdateVolumeGroupDetailErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup with Volume name
func TestAccVolumeGroup_CreateWithVolumeName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsWithVolumeName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "volume_names.0", volumeName),
				),
			},
		},
	})
}

// Test to Create VolumeGroup with Protection Policy name
func TestAccVolumeGroup_CreateWithPolicyName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsWithPolicyName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "protection_policy_name", policyName),
				),
			},
		},
	})
}

// Test to Create VolumeGroup with protection policy id and protection policy name, will result in error
func TestAccVolumeGroup_CreateWithInvalidPolicyName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeGroupParamsWithInvalidPolicyName,
				ExpectError: regexp.MustCompile(CreateVolumeGroupInvalidPolicyErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup with volume id and volume name, will result in error
func TestAccVolumeGroup_CreateWithVolumeIDAndName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeGroupParamsWithVolumeIDAndName,
				ExpectError: regexp.MustCompile(VolumeGroupInvalidAttributeCombinationErrorMsg),
			},
		},
	})
}

// Test to Create VolumeGroup with protection policy id and protection policy name, will result in error
func TestAccVolumeGroup_CreateWithPolicyIDAndName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeGroupParamsWithPolicyIDAndName,
				ExpectError: regexp.MustCompile(VolumeGroupInvalidAttributeCombinationErrorMsg),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params
func TestAccVolumeGroup_UpdateAddPolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdateAddPolicy,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "protection_policy_id", policyID),
				),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params
func TestAccVolumeGroup_UpdateAddVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdateAddVolume,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "volume_ids.0", volumeID),
				),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params
func TestAccVolumeGroup_UpdateAddPolicyNegative(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupParamsUpdateAddPolicyNegative,
				ExpectError: regexp.MustCompile(CreateVolumeGroupInvalidPolicyErrorMsg),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params
func TestAccVolumeGroup_UpdateRemovePolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdateAddPolicy,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "protection_policy_id", policyID),
				),
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdateRemovePolicy,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
		},
	})
}

// Test to Update existing VolumeGroup Params
func TestAccVolumeGroup_UpdateRemoveVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdateAddVolume,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "volume_ids.0", volumeID),
				),
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsUpdateRemoveVolume,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "tf_volume_group_new"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Updating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
				),
			},
		},
	})
}

// Test to import volume group successfully
func TestAccVolumeGroup_ImportSuccess(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupParamsCreate,
			},
			{
				Config:       ProviderConfigForTesting + VolumeGroupParamsCreate,
				ResourceName: "powerstore_volumegroup.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_volume_group_new", s[0].Attributes["name"])
					assert.Equal(t, "Creating Volume Group", s[0].Attributes["description"])
					assert.Equal(t, "false", s[0].Attributes["is_write_order_consistent"])
					return nil
				},
			},
		},
	})
}

// Negative - Test to import volume group
func TestAccVolumeGroup_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + VolumeGroupParamsCreate,
				ResourceName:  "powerstore_volumegroup.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportVGDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var VolumeGroupParamsCreate = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Creating Volume Group"
  is_write_order_consistent = false
}
`

var VolumeGroupParamsCreateWithoutName = `
resource "powerstore_volumegroup" "test" {
	description = "Create volume group without name"
}
`

var VolumeGroupParamsCreateWithInvalidPolicy = `
resource "powerstore_volumegroup" "test" {
	name = "tf_volume_group_new"
	description = "Creating Volume Group"
	is_write_order_consistent = false
	protection_policy_id = "invalid-id"
}
`

var VolumeGroupParamsCreateWithInvalidVolume = `
resource "powerstore_volumegroup" "test" {
	name = "tf_volume_group_new"
	description = "Creating Volume Group"
	is_write_order_consistent = false
	volume_ids = ["invalid-id"]
}
`

var VolumeGroupParamsUpdate = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
}
`

var VolumeGroupParamsUpdateServerError = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
  protection_policy_id = "invalid-id"
}
`

var VolumeGroupParamsWithVolumeName = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Creating Volume Group"
  is_write_order_consistent = false
  volume_names = ["` + volumeName + `"]
}
`

var VolumeGroupParamsWithPolicyName = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Creating Volume Group"
  is_write_order_consistent = false
  protection_policy_name = "` + policyName + `"
}
`

var VolumeGroupParamsWithInvalidPolicyName = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Creating Volume Group"
  is_write_order_consistent = false
  protection_policy_name = "invalid-name"
}
`

var VolumeGroupParamsWithVolumeIDAndName = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Creating Volume Group"
  is_write_order_consistent = false
  volume_ids = ["` + volumeID + `"]
  volume_names = ["` + volumeName + `"]
}
`

var VolumeGroupParamsWithPolicyIDAndName = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Creaing Volume Group"
  is_write_order_consistent = false
  protection_policy_id = "` + policyID + `"
  protection_policy_name = "` + policyName + `"
}
`

var VolumeGroupParamsUpdateAddPolicy = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
  protection_policy_id = "` + policyID + `"
}
`

var VolumeGroupParamsUpdateAddPolicyNegative = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
  protection_policy_name = "invalid-name"
}
`

var VolumeGroupParamsUpdateAddVolume = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
  volume_ids = ["` + volumeID + `"]
}
`

var VolumeGroupParamsUpdateRemovePolicy = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
}
`

var VolumeGroupParamsUpdateRemoveVolume = `
resource "powerstore_volumegroup" "test" {
  name = "tf_volume_group_new"
  description = "Updating Volume Group"
  is_write_order_consistent = false
}
`
