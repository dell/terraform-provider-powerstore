package powerstore

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"regexp"
	"testing"
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
				Config: VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
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
				Config: VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeParamsRename,
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
				Config: VolumeParamsWithMBInInt,
				Check:  resource.ComposeTestCheckFunc(checkCreateVolume(t, testProvider, "test_acc_cvol_mb")),
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
				Config: VolumeParamsWithTBInFloat,
				Check:  resource.ComposeTestCheckFunc(checkCreateVolume(t, testProvider, "test_acc_cvol_tb_float")),
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
				Config:      VolumeParamsWithPB,
				ExpectError: regexp.MustCompile("Invalid Capacity unit"),
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
				Config:      VolumeParamsWithInvalidCapUnit,
				ExpectError: regexp.MustCompile("Invalid Capacity unit"),
			},
		},
	})
}

// Test to create volume with invalid capacity unit, valid are MB, TB, GB
func TestAccVolume_CreateVolumeWithInvalidCapUnit2(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config:      VolumeParamsWithInvalidCapUnit,
				ExpectError: regexp.MustCompile("Invalid Capacity unit"),
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
				Config: VolumeCreateForUpdateGb,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "name", "test_acc_uvol_gb"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "size", "2")),
			},
			{
				Config: VolumeUpdateGb,
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
				Config: VolumeCreateForUpdateGb,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "name", "test_acc_uvol_gb"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_update_test_gb", "size", "2")),
			},
			{
				Config:      VolumeUpdateGbError,
				ExpectError: regexp.MustCompile("Error: Failed to update all parameters of Volume"),
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
				Config: VolumeParams,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeUpdateGbToTb,
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
				Config: VolumeParamsWithApplianceID,
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
				Config: VolumeParamsWithInvalidApplianceID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "appliance_id", "Z1")),
				ExpectError: regexp.MustCompile("Unable to find an appliance"),
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
				Config: VolumeParamsWithVolumeGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeParams,
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
				Config: VolumeParams,
			},
			{
				Config: VolumeParamsWithVolumeGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeParams,
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
				Config: VolumeParamsWithHostID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeParams,
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
				Config: VolumeParams,
			},
			{
				Config: VolumeParamsWithHostID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeParams,
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
				Config: VolumeParams,
			},
			{
				Config: VolumeParamsWithHostAndHostGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
				ExpectError: regexp.MustCompile("Either of HostID and Host GroupID should be present."),
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
				Config: VolumeParamsWithHostGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
			},
			{
				Config: VolumeParams,
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
				Config: VolumeParamsWithHostAndHostGroupID,
				Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "name", "test_acc_cvol"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "size", "2.5"),
					resource.TestCheckResourceAttr("powerstore_volume.volume_create_test", "capacity_unit", "GB")),
				ExpectError: regexp.MustCompile("Could not create volume Either HostID or HostGroupID can be present"),
			},
		},
	})
}

func checkCreateVolume(t *testing.T, p tfsdk.Provider, volName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providers := p.(*Pstoreprovider)
		_, err := providers.client.PStoreClient.GetVolumeByName(context.Background(), volName)
		if err != nil {
			return fmt.Errorf("failed to fetch volume")
		}
		if !providers.configured {
			return fmt.Errorf("provider not configured")
		}

		if providers.client.PStoreClient == nil {
			return fmt.Errorf("provider not configured")
		}
		return nil
	}
}

var VolumeParams = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	volume_group_id = ""
	host_id = ""
	host_group_id = ""
}
`

var VolumeParamsRename = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol_updated"
	size = 2.5
	capacity_unit = "GB"
}
`

var VolumeParamsWithPB = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test_mb" {
	name = "test_acc_cvol_pb"
	size = 1
	capacity_unit = "PB"
}
`

var VolumeParamsWithTBInFloat = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test_tb_float" {
	name = "test_acc_cvol_tb_float"
	size = 2.5
	capacity_unit = "TB"
}
`

var VolumeParamsWithMBInInt = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test_mb" {
	name = "test_acc_cvol_mb"
	size = 200
	capacity_unit = "MB"
}
`

var VolumeParamsWithInvalidCapUnit = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "LB"
}
`

var VolumeUpdateGbToTb = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "TB"
}
`

var VolumeCreateForUpdateGb = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_update_test_gb" {
	name = "test_acc_uvol_gb"
	size = 2
	capacity_unit = "GB"
}
`

var VolumeUpdateGb = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_update_test_gb" {
	name = "test_acc_uvol_gb_updated"
	size = 2.5
	capacity_unit = "GB"
}
`

var VolumeUpdateGbError = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_update_test_gb" {
	name = "test_acc_uvol_gb"
	size = 1
	capacity_unit = "GB"
}
`
var VolumeParamsWithApplianceID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	appliance_id = "A1"
}
`
var VolumeParamsWithInvalidApplianceID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	appliance_id = "Z2"
}
`

var VolumeParamsWithVolumeGroupID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	volume_group_id = "` + volumeGroupID + `"
}
`

var VolumeParamsWithHostID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_id = "` + hostID + `"
}
`
var VolumeParamsWithHostGroupID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_group_id = "` + hostGroupID + `"
}
`
var VolumeParamsWithHostAndHostGroupID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_volume" "volume_create_test" {
	name = "test_acc_cvol"
	size = 2.5
	capacity_unit = "GB"
	host_group_id = "` + hostGroupID + `"
	host_id =  "` + hostID + `"
}
`
