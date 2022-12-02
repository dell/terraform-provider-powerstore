package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create StorageContainer
func TestAccStorageContainer_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: StorageContainerParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "scterraform_acc"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "quota", "10737418240"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "storage_protocol", "SCSI"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "high_water_mark", "70"),
				),
			},
		},
	})
}

// Test to update existing StorageContainer params
func TestAccStorageContainer_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: StorageContainerParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "scterraform_acc"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "quota", "10737418240"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "storage_protocol", "SCSI"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "high_water_mark", "70"),
				),
			},
			{
				Config: StorageContainerParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "scterraform_acc_new"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "quota", "10737418242"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "storage_protocol", "NVMe"),
				),
			},
		},
	})
}

// Test to Create StorageContainer with invalid values, will result in error
func TestAccStorageContainer_CreateWithInvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	tests := []resource.TestStep{
		{
			Config:      StorageContainerParamsInvalidStorageProtocol,
			ExpectError: regexp.MustCompile("Attribute storage_protocol must be one of these"),
		},
		{
			Config:      StorageContainerParamsCreateServerError,
			ExpectError: regexp.MustCompile("Could not create Storage Container"),
		},
	}

	for i := range tests {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testProviderFactory,
			Steps:                    []resource.TestStep{tests[i]},
		})
	}
}

// Test to update existing StorageContainer params but will result in error
func TestAccStorageContainer_UpdateError(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: StorageContainerParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "scterraform_acc"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "quota", "10737418240"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "storage_protocol", "SCSI"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "high_water_mark", "70"),
				),
			},
			{
				Config:      StorageContainerParamsCreateServerError,
				ExpectError: regexp.MustCompile("Could not update"),
			},
		},
	})
}

var StorageContainerParamsCreate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_storagecontainer" "test" {
	name = "scterraform_acc"
	quota = 10737418240
	storage_protocol = "SCSI"
	high_water_mark = 70
}
`

var StorageContainerParamsUpdate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_storagecontainer" "test" {
	name = "scterraform_acc_new"
	quota = 10737418242
	storage_protocol = "NVMe"	
}
`

var StorageContainerParamsInvalidStorageProtocol = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_storagecontainer" "test" {
	name = "scterraform_acc_new"
	quota = 10737418242
	storage_protocol = "invalid"	
}
`

var StorageContainerParamsCreateServerError = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

resource "powerstore_storagecontainer" "test" {
	name = "scterraform_acc_new12312313212313121213212312312312312313131321212121"
	quota = 10737418240
	storage_protocol = "SCSI"	
}
`
