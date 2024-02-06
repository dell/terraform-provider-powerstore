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
			// Import Testing
			{
				Config:       ProviderConfigForTesting + HostGroupParamsCreate,
				ResourceName: "powerstore_hostgroup.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "test_hostgroup", s[0].Attributes["name"])
					assert.Equal(t, "Test Create Host Group", s[0].Attributes["description"])
					assert.Equal(t, hostID, s[0].Attributes["host_ids.0"])
					return nil
				},
			},
			// Update Testing
			{
				Config: ProviderConfigForTesting + HostGroupParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Update Host Group"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "host_ids.0", hostID),
				),
			},
			// Remove host before cleanup
			{
				Config: ProviderConfigForTesting + HostGroupParamsUpdateRemoveHost,
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
				ExpectError: regexp.MustCompile(CreateHGInvalidHostErrorMsg),
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

// Test to Create HostGroup with Host name
func TestAccHostGroup_CreateWithHostName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostGroupParamsCreateWithHostName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "name", "test_hostgroup"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "description", "Test Create Host Group"),
					resource.TestCheckResourceAttr("powerstore_hostgroup.test", "host_names.0", hostName),
				),
			},
			{
				Config: ProviderConfigForTesting + HostGroupParamsUpdateRemoveHostWithName,
			},
		},
	})
}

// Test to Create HostGroup with invalid host name, will result in error
func TestAccHostGroup_CreateWithInvalidHostName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostGroupParamsCreateWithInvalidHostName,
				ExpectError: regexp.MustCompile(CreateHostGroupInvalidHostErrorMsg),
			},
		},
	})
}

// Test to Create HostGroup with host id and host name, will result in error
func TestAccHostGroup_CreateWithHostIDAndName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostGroupParamsWithHostIDAndHostName,
				ExpectError: regexp.MustCompile(InvalidAttributeCombinationErrorMsg),
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

var HostGroupParamsCreateWithHostName = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_names = ["` + hostName + `"]
}
`

var HostGroupParamsCreateWithInvalidHostName = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_names = ["invalid-name"]
}
`

var HostGroupParamsWithHostIDAndHostName = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = ["` + hostID + `"]
	host_names = ["` + hostName + `"]
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
	host_ids = ["` + hostID + `"]
}
`

var HostGroupParamsUpdateRemoveHost = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_ids = []
}
`

var HostGroupParamsUpdateRemoveHostWithName = `
resource "powerstore_hostgroup" "test" {
	name = "test_hostgroup"
	description = "Test Create Host Group"
	host_names = []
}
`
