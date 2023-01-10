package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
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
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "high_water_mark", "60"),
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
			ExpectError: regexp.MustCompile("Attribute storage_protocol value must be one of"),
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

// Test to import resource but resulting in error
func TestAccStorageContainer_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        StorageContainerParamsCreate,
				ResourceName:  "powerstore_storagecontainer.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile("Could not import storage container"),
				ImportStateId: "invalid-id",
			},
		},
	})
}

// Test to import successfully
func TestAccStorageContainer_ImportSuccess(t *testing.T) {

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
				Config:            StorageContainerParamsCreate,
				ResourceName:      "powerstore_storagecontainer.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "scterraform_acc", s[0].Attributes["name"])
					assert.Equal(t, "10737418240", s[0].Attributes["quota"])
					assert.Equal(t, "SCSI", s[0].Attributes["storage_protocol"])
					return nil
				},
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
	high_water_mark = 60
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
