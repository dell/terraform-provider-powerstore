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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create Volume
func TestAccVolume_CreateVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			// Import Succes Test
			{
				Config:            ProviderConfigForTesting + VolumeParams,
				ResourceName:      "powerstore_volume.volume_create_test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "test_acc_cvol", s[0].Attributes["name"])
					assert.Equal(t, "2.5", s[0].Attributes["size"])
					assert.Equal(t, "GB", s[0].Attributes["capacity_unit"])
					return nil
				},
			},
		},
	})
}

// Test to Rename Volume
func TestAccVolume_UpdateVolumeRename(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: ProviderConfigForTesting + VolumeParamsRename,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol_updated")),
			},
		},
	})
}

// Create Volume with size in int and capacity unit TB
func TestAccVolume_CreateVolumeWithMBInInt(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithMBInInt,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test_mb", "name", "test_acc_cvol_mb")),
			},
		},
	})
}

// Create Volume with size in float and capacity unit TB
func TestAccVolume_CreateVolumeWithTBInFloat(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithTBInFloat,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test_tb_float", "name", "test_acc_cvol_tb_float")),
			},
		},
	})
}

// Test to create volume with invalid high capacity unit, maximum volume size is 230 TB so PB is invalid
func TestAccVolume_CreateVolumeWithPB(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeParamsWithPB,
				ExpectError: regexp.MustCompile(InvalidAttributeErrorMsg),
			},
		},
	})
}

// Test to create volume with invalid capacity unit, valid are MB, TB, GB
func TestAccVolume_CreateVolumeWithInvalidCapUnit(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + VolumeParamsWithInvalidCapUnit,
				ExpectError: regexp.MustCompile(InvalidAttributeErrorMsg),
			},
		},
	})
}

// Test to Update Volume size
func TestAccVolume_UpdateVolumeGb(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeCreateForUpdateGb,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "name", "test_acc_uvol_gb"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "size", "2")),
			},
			{
				Config: ProviderConfigForTesting + VolumeUpdateGb,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "name", "test_acc_uvol_gb_updated"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "size", "2.5")),
			},
		},
	})
}

// Test to reduce volume size, Powerstore does not support decreasing volume size
func TestAccVolume_UpdateVolumeGbError1(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeCreateForUpdateGb,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "name", "test_acc_uvol_gb"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "size", "2")),
			},
			{
				Config:      ProviderConfigForTesting + VolumeUpdateGbError,
				ExpectError: regexp.MustCompile(InvalidSizeErrorMsg),
			},
		},
	})
}

// Test to update Volume size from GB to TB
func TestAccVolume_UpdateVolumeSizeGbToTb(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: ProviderConfigForTesting + VolumeUpdateGbToTb,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "TB"))),
			},
		},
	})
}

// Test to update Appliance ID in volume resource
func TestAccVolume_UpdateVolumeApplianceID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithApplianceID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "appliance_id", "A1")),
			},
		},
	})
}

// Test to update Invalid Appliance ID in volume resource
func TestAccVolume_UpdateVolumeInvalidApplianceID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithInvalidApplianceID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "appliance_id", "Z1")),
				ExpectError: regexp.MustCompile(CreateVolumeErrorMsg),
			},
		},
	})
}

// Test to update Invalid Performance Policy ID in volume resource
func TestAccVolume_UpdateVolumePerformancePolicyID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithInvalidPerformancePolicy,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
				ExpectError: regexp.MustCompile(InvalidAttributeErrorMsg),
			},
		},
	})
}

// Test to Add Volume Group ID in volume resource
func TestAccVolume_AddVolumeGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithVolumeGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
		},
	})
}

// Test to update Volume Group ID in volume resource
func TestAccVolume_UpdateVolumeGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
			},
			// Add Volume Group
			{
				Config: ProviderConfigForTesting + VolumeParamsWithVolumeGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			// Remove Volume Group
			{
				Config: ProviderConfigForTesting + VolumeParamsWithVolumeGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
		},
	})
}

// Test to Add Host ID in volume resource
func TestAccVolume_AddHostID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithHostName,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
		},
	})
}

// Test to update Host ID in volume resource
func TestAccVolume_UpdateHostID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
			},
			{
				Config: ProviderConfigForTesting + VolumeParamsWithHostID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
		},
	})
}

// Test to update Host ID in volume resource
func TestAccVolume_UpdateHostGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
			},
			{
				Config: ProviderConfigForTesting + VolumeParamsWithHostGroupName,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			// Remove host IDs from the host group before hotgroup deletion
			{
				Config: ProviderConfigForTesting + HostGroupPreReqParamsclear,
			},
		},
	})
}

// Test to update Host and HostGroup ID in volume resource
func TestAccVolume_UpdateHostAndHostGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParams,
			},
			{
				Config: ProviderConfigForTesting + VolumeParamsWithHostAndHostGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
				ExpectError: regexp.MustCompile(HostIDHostGroupErrorMsg),
			},
			// Remove host IDs from the host group before hotgroup deletion
			{
				Config: ProviderConfigForTesting + HostGroupPreReqParamsclear,
			},
		},
	})
}

// Test to update Host Group ID in volume resource
func TestAccVolume_AddHostGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithHostGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			// Remove host IDs from the host group before hotgroup deletion
			{
				Config: ProviderConfigForTesting + HostGroupPreReqParamsclear,
			},
		},
	})
}

// Test to update Host as well as Host Group ID in volume resource. Only 1 can be present at a time.
func TestAccVolume_AddHostAndHostGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeParamsWithHostAndHostGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
				ExpectError: regexp.MustCompile(CreateVolumeHostIDErrorMsg),
			},
			// Remove host IDs from the host group before hotgroup deletion
			{
				Config: ProviderConfigForTesting + HostGroupPreReqParamsclear,
			},
		},
	})
}

// Negative test case for import
func TestAccVolume_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + VolumeParams,
				ResourceName:  "powerstore_volume.volume_create_test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportVolumeDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var PreReqVolume = `
resource "powerstore_volume" "pre_req_volume" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
}
`

var VolumeParams = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	volume_group_id = ""
	host_id = ""
	host_group_id = ""
	sector_size = 512
}
`

var VolumeParamsRename = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol_updated"
	size = 2.5
	capacity_unit = "GB"
}
`

var VolumeParamsWithPB = `
resource "powerstore_volume" "volume_create_test_mb" {
	name = "test_acc_cvol_pb"
	size = 1
	capacity_unit = "PB"
}
`

var VolumeParamsWithTBInFloat = `
resource "powerstore_volume" "volume_create_test_tb_float" {
	name = "test_acc_cvol_tb_float"
	size = 2.5
	capacity_unit = "TB"
}
`

var VolumeParamsWithMBInInt = `
resource "powerstore_volume" "volume_create_test_mb" {
	name = "test_acc_cvol_mb"
	size = 200
	capacity_unit = "MB"
}
`

var VolumeParamsWithInvalidCapUnit = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "LB"
}
`

var VolumeUpdateGbToTb = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol_tb_float"
	size = 2.5
	capacity_unit = "TB"
}
`

var VolumeCreateForUpdateGb = `
resource "powerstore_volume" "volume_update_test_gb" {
	name = "test_acc_uvol_gb"
	size = 2
	capacity_unit = "GB"
}
`

var VolumeUpdateGb = `
resource "powerstore_volume" "volume_update_test_gb" {
	name = "test_acc_uvol_gb_updated"
	size = 2.5
	capacity_unit = "GB"
}
`

var VolumeUpdateGbError = `
resource "powerstore_volume" "volume_update_test_gb" {
	name = "test_acc_uvol_gb"
	size = 1
	capacity_unit = "GB"
}
`
var VolumeParamsWithApplianceID = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	appliance_id = "A1"
}
`
var VolumeParamsWithInvalidApplianceID = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	appliance_id = "Z2"
}
`

var VolumeParamsWithVolumeGroupID = VolumeGroupParamsCreate + `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	volume_group_id = powerstore_volumegroup.test.id
}
`

var VolumeParamsWithHostID = HostPreReqForVolume + `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_id = powerstore_host.test.id
}
`
var VolumeParamsWithHostGroupID = HostGroupPreReqParams + `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_group_id = powerstore_hostgroup.test.id
}
`
var VolumeParamsWithHostAndHostGroupID = HostGroupPreReqParams + `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_group_id = powerstore_hostgroup.test.id
	host_id =  powerstore_host.test.name
}
`
var VolumeParamsWithInvalidPerformancePolicy = `
resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	performance_policy_id = ""
}
`
var VolumeParamsWithHostGroupName = HostGroupPreReqParams + `
resource "powerstore_volume" "volume_create_test" {
	depends_on = [powerstore_hostgroup.test]
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_group_name = powerstore_hostgroup.test.name
}
`
var VolumeParamsWithHostName = HostPreReqForVolume + `
resource "powerstore_volume" "volume_create_test" {
	depends_on = [powerstore_host.test]
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_name = powerstore_host.test.name
}
`
