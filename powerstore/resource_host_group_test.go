package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create HostGroup
func TestAccHostGroup_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Create Host Group"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "host_ids.0", hostID),
				),
			},
		},
	})
}

// Test to Create HostGroup with Invalid values, will result in error
func TestAccHostGroup_CreateWithInvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostGroupParamsCreateServerError,
				ExpectError: regexp.MustCompile(CreateHGDetailErrorMsg),
			},
		},
	})
}

// Test to Create HostGroup with Invalid values, will result in error
func TestAccHostGroup_CreateWithInvalidHostID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostGroupParamsCreateWithInvalidHostID,
				ExpectError: regexp.MustCompile(CreateHGInvalidHostErrorMsg),
			},
		},
	})
}

// Test to Create HostGroup with blank name, will result in error
func TestAccHostGroup_CreateWithBlankName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostGroupParamsCreateWithBlankName,
				ExpectError: regexp.MustCompile(CreateHGWithBlankName),
			},
		},
	})
}

// Test to Update existing HostGroup Params
func TestAccHostGroup_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostGroupParamsCreate1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Create Host Group"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "host_ids.0", hostID1),
				),
			},
			{
				Config: ProviderConfigForTesting + HostGroupParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Update Host Group"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "host_ids.0", hostID1),
				),
			},
		},
	})
}

// Test to Update existing HostGroup Params
func TestAccHostGroup_UpdateRemoveHost(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostGroupParamsCreate2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Create Host Group"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "host_ids.0", hostID2),
				),
			},
			{
				Config: ProviderConfigForTesting + HostGroupParamsUpdateRemoveHost,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Update Host Group"),
				),
			},
		},
	})
}

// Test to import host group successfully
func TestAccHostGroup_ImportSuccess(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostGroupParamsCreate3,
			},
			{
				Config:       ProviderConfigForTesting + HostGroupParamsCreate3,
				ResourceName: "powerstore_hostgroup.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "test_hostgroup", s[0].Attributes["name"])
					assert.Equal(t, "Test Create Host Group", s[0].Attributes["description"])
					assert.Equal(t, hostID3, s[0].Attributes["host_ids.0"])
					return nil
				},
			},
		},
	})
}

// Negative - Test to import host group
func TestAccHostGroup_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + HostGroupParamsCreate,
				ResourceName:  "powerstore_hostgroup.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportHGDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var HostGroupParamsCreate = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = ["` + hostID + `"]
}
`

var HostGroupParamsCreate1 = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = ["` + hostID1 + `"]
}
`

var HostGroupParamsCreate2 = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = ["` + hostID2 + `"]
}
`

var HostGroupParamsCreate3 = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = ["` + hostID3 + `"]
}
`

var HostGroupParamsCreateServerError = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
}
`

var HostGroupParamsCreateWithInvalidHostID = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = ["invalid-id"]
}
`

var HostGroupParamsCreateWithBlankName = `
resource "powerstore_hostgroup" "test" {
	name = ""
	description = "Test Create Host Group"
	host_ids = ["` + hostID + `"]
}
`

var HostGroupParamsUpdate = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Update Host Group"
	host_ids = ["` + hostID1 + `"]
}
`

var HostGroupParamsUpdateRemoveHost = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Update Host Group"
	host_ids = []
}
`
