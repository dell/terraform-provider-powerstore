/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"context"
	"os"
	"regexp"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"testing"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Acceptance tests

func TestAccQosPolicy_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "name", "tf_acc_qos_policy"),
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "type", "QoS"),
				),
			},
		},
	})
}

func TestAccQosPolicy_CreateFilePerformance(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyParamsCreateFilePerformance,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "type", "File_Performance"),
				),
			},
		},
	})
}

func TestAccQosPolicy_CreateWithDescription(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyParamsCreateWithDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "description", "Test QoS policy"),
				),
			},
		},
	})
}

func TestAccQosPolicy_CreateWithIoLimitRule(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleParamsCreate + QosPolicyParamsCreateWithIoLimitRule,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "type", "QoS"),
					resource.TestCheckResourceAttrSet("powerstore_qos_policy.test", "io_limit_rule_id"),
				),
			},
		},
	})
}

func TestAccQosPolicy_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "name", "tf_acc_qos_policy"),
				),
			},
			{
				Config: ProviderConfigForTesting + QosPolicyParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_qos_policy.test", "description", "Updated description"),
				),
			},
		},
	})
}

func TestAccQosPolicy_Import(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyParamsCreate,
			},
			{
				Config:            ProviderConfigForTesting + QosPolicyParamsCreate,
				ResourceName:      "powerstore_qos_policy.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_acc_qos_policy", s[0].Attributes["name"])
					assert.Equal(t, "QoS", s[0].Attributes["type"])
					return nil
				},
			},
		},
	})
}

func TestAccQosPolicy_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(`The argument "name" is required`),
			},
		},
	})
}

func TestAccQosPolicy_CreateWithoutType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyParamsCreateWithoutType,
				ExpectError: regexp.MustCompile(`The argument "type" is required`),
			},
		},
	})
}

func TestAccQosPolicy_CreateWithInvalidType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyParamsCreateWithInvalidType,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
		},
	})
}

func TestAccQosPolicy_CreateWithoutLimitRuleID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyParamsCreateWithoutLimitRuleID,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

// Unit tests

func TestResourceQosPolicy_Metadata(t *testing.T) {
	r := &resourceQosPolicy{}
	req := fwresource.MetadataRequest{}
	resp := &fwresource.MetadataResponse{}

	r.Metadata(context.Background(), req, resp)

	assert.Equal(t, "_qos_policy", resp.TypeName)
}

func TestResourceQosPolicy_Schema(t *testing.T) {
	r := &resourceQosPolicy{}
	req := fwresource.SchemaRequest{}
	resp := &fwresource.SchemaResponse{}

	r.Schema(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
}

func TestResourceQosPolicy_Configure_Nil(t *testing.T) {
	r := &resourceQosPolicy{}
	req := fwresource.ConfigureRequest{ProviderData: nil}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, r.client)
}

func TestResourceQosPolicy_Configure_InvalidType(t *testing.T) {
	r := &resourceQosPolicy{}
	req := fwresource.ConfigureRequest{ProviderData: "invalid"}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestResourceQosPolicy_Configure_Success(t *testing.T) {
	r := &resourceQosPolicy{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwresource.ConfigureRequest{ProviderData: c}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

func TestResourceQosPolicy_UpdateState_AllFields(t *testing.T) {
	r := resourceQosPolicy{}
	id := "policy-id"
	name := "test-policy"
	description := "Test description"
	typeEnum := clientgen.POLICYTYPEENUM_QO_S
	managedBy := clientgen.POLICYMANAGEDBYENUM_USER
	managedById := "managed-id"
	isReadOnly := false
	isReplica := false
	fileIoRuleId := "file-rule-id"
	ioRuleId := "io-rule-id"

	response := &clientgen.PolicyInstance{
		Id:                &id,
		Name:              &name,
		Description:       &description,
		Type:              &typeEnum,
		ManagedBy:         &managedBy,
		ManagedById:       &managedById,
		IsReadOnly:        &isReadOnly,
		IsReplica:         &isReplica,
		FileIoLimitRuleId: &fileIoRuleId,
		IoLimitRule:       &clientgen.IoLimitRuleInstance{Id: &ioRuleId},
	}

	state := r.updateQosPolicyState(response)
	assert.Equal(t, "policy-id", state.ID.ValueString())
	assert.Equal(t, "test-policy", state.Name.ValueString())
	assert.Equal(t, "Test description", state.Description.ValueString())
	assert.Equal(t, "QoS", state.Type.ValueString())
	assert.Equal(t, "file-rule-id", state.FileIoLimitRuleId.ValueString())
	assert.Equal(t, "io-rule-id", state.IoLimitRuleId.ValueString())
}

func TestResourceQosPolicy_UpdateState_NilFields(t *testing.T) {
	r := resourceQosPolicy{}
	response := &clientgen.PolicyInstance{}

	state := r.updateQosPolicyState(response)
	assert.Equal(t, "", state.ID.ValueString())
	assert.Equal(t, "", state.Name.ValueString())
	assert.Equal(t, "", state.Description.ValueString())
	assert.Equal(t, "", state.Type.ValueString())
	assert.True(t, state.FileIoLimitRuleId.IsNull())
	assert.True(t, state.IoLimitRuleId.IsNull())
}

func TestResourceQosPolicy_UpdateState_NilIoLimitRule(t *testing.T) {
	r := resourceQosPolicy{}
	id := "policy-id"
	response := &clientgen.PolicyInstance{Id: &id, IoLimitRule: nil}

	state := r.updateQosPolicyState(response)
	assert.True(t, state.IoLimitRuleId.IsNull())
}

func TestResourceQosPolicy_UpdateState_IoLimitRuleNilId(t *testing.T) {
	r := resourceQosPolicy{}
	id := "policy-id"
	response := &clientgen.PolicyInstance{
		Id:          &id,
		IoLimitRule: &clientgen.IoLimitRuleInstance{Id: nil},
	}

	state := r.updateQosPolicyState(response)
	assert.True(t, state.IoLimitRuleId.IsNull())
}

func TestResourceQosPolicy_UpdateState_FilePerformanceType(t *testing.T) {
	r := resourceQosPolicy{}
	id := "policy-id"
	typeEnum := clientgen.POLICYTYPEENUM_FILE_PERFORMANCE
	fileIoRuleId := "file-rule-id"

	response := &clientgen.PolicyInstance{
		Id:                &id,
		Type:              &typeEnum,
		FileIoLimitRuleId: &fileIoRuleId,
	}

	state := r.updateQosPolicyState(response)
	assert.Equal(t, "File_Performance", state.Type.ValueString())
	assert.Equal(t, "file-rule-id", state.FileIoLimitRuleId.ValueString())
}

// Terraform configs

var QosPolicyParamsCreate = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule_qos"
	type     = "Absolute"
	max_iops = 1000
	max_bw   = 2000
}

resource "powerstore_qos_policy" "test" {
	name             = "tf_acc_qos_policy"
	type             = "QoS"
	io_limit_rule_id = powerstore_io_limit_rule.test.id
}
`

var QosPolicyParamsCreateFilePerformance = `
resource "powerstore_file_io_limit_rule" "test" {
	name   = "tf_acc_file_io_limit_rule_qos"
	max_bw = 100
}

resource "powerstore_qos_policy" "test" {
	name                  = "tf_acc_qos_policy_file"
	type                  = "File_Performance"
	file_io_limit_rule_id = powerstore_file_io_limit_rule.test.id
}
`

var QosPolicyParamsCreateWithDescription = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule_desc"
	type     = "Absolute"
	max_iops = 1000
	max_bw   = 2000
}

resource "powerstore_qos_policy" "test" {
	name             = "tf_acc_qos_policy_desc"
	type             = "QoS"
	description      = "Test QoS policy"
	io_limit_rule_id = powerstore_io_limit_rule.test.id
}
`

var QosPolicyParamsCreateWithIoLimitRule = `
resource "powerstore_qos_policy" "test" {
	name             = "tf_acc_qos_policy_with_rule"
	type             = "QoS"
	io_limit_rule_id = powerstore_io_limit_rule.test.id
}
`

var QosPolicyParamsUpdate = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule_qos"
	type     = "Absolute"
	max_iops = 1000
	max_bw   = 2000
}

resource "powerstore_qos_policy" "test" {
	name             = "tf_acc_qos_policy"
	type             = "QoS"
	description      = "Updated description"
	io_limit_rule_id = powerstore_io_limit_rule.test.id
}
`

var QosPolicyParamsCreateWithoutName = `
resource "powerstore_qos_policy" "test" {
	type             = "QoS"
	io_limit_rule_id = "some-rule-id"
}
`

var QosPolicyParamsCreateWithoutType = `
resource "powerstore_qos_policy" "test" {
	name             = "tf_acc_qos_policy"
	io_limit_rule_id = "some-rule-id"
}
`

var QosPolicyParamsCreateWithInvalidType = `
resource "powerstore_qos_policy" "test" {
	name             = "tf_acc_qos_policy"
	type             = "Invalid"
	io_limit_rule_id = "some-rule-id"
}
`

var QosPolicyParamsCreateWithoutLimitRuleID = `
resource "powerstore_qos_policy" "test" {
	name = "tf_acc_qos_policy"
	type = "QoS"
}
`
