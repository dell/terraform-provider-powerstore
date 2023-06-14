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

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create Host Resource
func TestAccHost_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", "tf_host_acc_new"),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
				),
			},
		},
	})
}

// Test to Create Host Resource with Single CHAP
func TestAccHost_CreateSingleCHAP(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreateWithSingleCHAP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", "tf_host_acc_new"),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
				),
			},
		},
	})
}

// Test to Create Host Resource without policy
func TestAccHost_CreateWithoutPolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostParamsCreateWithoutPolicy,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
			},
		},
	})
}

// Test to Create Host Resource Without Name
func TestAccHost_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
			},
		},
	})
}

// Test to rename Host Resource
func TestAccHost_Rename(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", "tf_host_acc_new"),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
					resource.TestCheckResourceAttr("powerstore_host.test", "os_type", "Linux"),
				),
			},
			{
				Config: ProviderConfigForTesting + HostParamsRename,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", "new_name_host_acc_test"),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
					resource.TestCheckResourceAttr("powerstore_host.test", "os_type", "Linux"),
				),
			},
		},
	})
}

// Test to Add and remove Host Resource
func TestAccHost_AddRemoveInitiators(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreateWithCHAP,
			},
			{
				Config: ProviderConfigForTesting + HostParamsAddInitiators,
			},
			{
				Config: ProviderConfigForTesting + HostParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", "tf_host_acc_new"),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
					resource.TestCheckResourceAttr("powerstore_host.test", "os_type", "Linux"),
				),
			},
		},
	})
}

// Test to Modify Host Resource
func TestAccHost_ModifyInitiators(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreateWithSingleCHAP,
			},
			{
				Config: ProviderConfigForTesting + HostParamsCreateWithCHAP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", "tf_host_acc_new"),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
					resource.TestCheckResourceAttr("powerstore_host.test", "os_type", "Linux"),
				),
			},
		},
	})
}

// Negative test case for import
func TestAccHost_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + HostParamsCreate,
				ResourceName:  "powerstore_host.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportHostDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

// Test to import successfully
func TestAccHost_ImportSuccess(t *testing.T) {

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreate,
			},
			{
				Config:            ProviderConfigForTesting + HostParamsCreate,
				ResourceName:      "powerstore_host.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_host_acc_new", s[0].Attributes["name"])
					assert.Equal(t, "Linux", s[0].Attributes["os_type"])
					return nil
				},
			},
		},
	})

}

var HostParamsCreate = `
resource "powerstore_host" "test" {
	name = "tf_host_acc_new"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606"}]
}
`

var HostParamsCreateWithCHAP = `
resource "powerstore_host" "test" {
	name = "tf_host_acc_new"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606" ,chap_single_username="chap_single_username",chap_single_password="chap_single_password", chap_mutual_username="chap_mutual_username",chap_mutual_password="chap_mutual_password"}]
}
`
var HostParamsCreateWithSingleCHAP = `
resource "powerstore_host" "test" {
	name = "tf_host_acc_new"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606",chap_single_username="chap_single_username",chap_single_password="chap_single_password"}]
}
`

var HostParamsCreateWithoutPolicy = `
resource "powerstore_host" "test" {
	name = "tf_host_acc_new"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606" }]
	description = "Test Host Resource"
}
`
var HostParamsCreateWithoutName = `
resource "powerstore_host" "test" {
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606" }]
	description = "Test Host Resource"
	os_type = "Linux"
}
`
var HostParamsRename = `
resource "powerstore_host" "test" {
	name = "new_name_host_acc_test"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606" }]
}
`
var HostParamsAddInitiators = `
resource "powerstore_host" "test" {
	name = "tf_host_acc_new"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606" ,chap_single_username="chap_single_username",chap_single_password="chap_single_password", chap_mutual_username="chap_mutual_username",chap_mutual_password="chap_mutual_password"},{port_name="iqn.1998-01.com.vmware:lgc198248-5b06fb37" ,chap_single_username="chap_single_username",chap_single_password="chap_single_password"}]
}
`
