/*
Copyright (c) 2024-2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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

func TestAccHost_Validations(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// validate with known port name and valid unknown chap creds
			{
				Config: ProviderConfigForTesting + `
					resource terraform_data username {
						input = null
					}
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "nqn.1994-05.com.redhat:88cb606"
							chap_single_username = terraform_data.username.output
							chap_single_password = terraform_data.username.output
							chap_mutual_username = terraform_data.username.output
							chap_mutual_password = terraform_data.username.output
						}]
					}
					`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// validate with unknown iSCSI port name and known chap creds
			{
				Config: ProviderConfigForTesting + `
					resource terraform_data portname {
						input = "iqn.port"
					}
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = terraform_data.portname.output
							chap_single_username = "whatever"
							chap_single_password = "whatever"
							chap_mutual_username = "whatever"
							chap_mutual_password = "whatever"
						}]
					}
					`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// validate with unknown iSCSI port name and unknown chap creds
			{
				Config: ProviderConfigForTesting + `
					resource terraform_data portname {
						input = "iqn.port"
					}
					resource terraform_data username {
						input = null
					}
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = terraform_data.portname.output
							chap_single_username = terraform_data.username.output
							chap_single_password = terraform_data.username.output
							chap_mutual_username = terraform_data.username.output
							chap_mutual_password = terraform_data.username.output
						}]
					}
					`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// validate with known nvme port name and known chap creds - neg
			{
				Config: ProviderConfigForTesting + `
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "nqn.1994-05.com.redhat:88cb606"
							chap_single_username = "whatever"
							chap_single_password = "whatever"
							chap_mutual_username = "whatever"
							chap_mutual_password = "whatever"
						}]
					}
					`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("chap credentials are supported only with iSCSI protocol"),
			},
			// validate that chap creds work in loops
			{
				Config: ProviderConfigForTesting + `
					resource "powerstore_host" "test" {
						for_each = tomap({
							"192.168.10.156" = "iqn.1993-08.org.debian.iscsi:01:107dc7e4254a"
							"192.168.10.157" = "iqn.1993-08.org.debian.iscsi:01:107dc7e4254b"
						})
						name              = each.key
						os_type           = "ESXi"
						host_connectivity = "Local_Only"
						description       = each.value
						initiators =  [{port_name = "${each.value}", chap_single_username = "testuser", chap_single_password = "testuser123"}]
					}
					`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// validate chap_mutual_username cannot be present without chap_single_username - neg
			{
				Config: ProviderConfigForTesting + `
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "iqn.1994-05.com.redhat:88cb606"
							chap_mutual_username = "whatever"
						}]
					}
					`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("chap_mutual_username.*cannot be present without.*chap_single_username"),
			},
			// validate chap_mutual_password cannot be present without chap_mutual_username - neg
			{
				Config: ProviderConfigForTesting + `
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "iqn.1994-05.com.redhat:88cb606"
							chap_mutual_password = "whatever"
						}]
					}
					`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("chap_mutual_password.*cannot be present without.*chap_mutual_username"),
			},
			// validate chap_single_password cannot be present without chap_single_username - neg
			{
				Config: ProviderConfigForTesting + `
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "iqn.1994-05.com.redhat:88cb606"
							chap_single_password = "whatever"
						}]
					}
					`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("chap_single_password.*cannot be present without.*chap_single_username"),
			},
			// validate with unknown null chap passwords
			{
				Config: ProviderConfigForTesting + `
					resource terraform_data password {
						input = null
					}
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "iqn.1994-05.com.redhat:88cb606"
							chap_single_username = "whatever"
							chap_single_password = terraform_data.password.output
							chap_mutual_username = "whatever"
							chap_mutual_password = terraform_data.password.output
						}]
					}
					`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// validate with unknown non-null chap single username
			{
				Config: ProviderConfigForTesting + `
					resource terraform_data username {
						input = "whatever"
					}
					resource "powerstore_host" "test" {
						name = "tf_host_acc_new"
						description = "Test Host Resource"
						os_type = "Linux"
						initiators = [{
							port_name = "iqn.1994-05.com.redhat:88cb606"
							chap_single_username = terraform_data.username.output
							chap_mutual_username = "whatever"
						}]
					}
					`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

var HostPreReqForVolume = `
resource "powerstore_host" "test" {
	name = "tf_host_acc_new"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb606"}]
}
`
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
