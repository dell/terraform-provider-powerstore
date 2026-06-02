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
	"fmt"
	"os"
	"regexp"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/powerstore/helper"
	"testing"

	"github.com/bytedance/mockey"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Acceptance tests

func TestAccIoLimitRuleRes(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	var mockCreateApi, mockReadApi, mockDeleteApi, mockUpdateApi *mockey.Mocker

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				// Create mock error
				PreConfig: func() {
					mockCreateApi = mockey.Mock((*clientgen.IoLimitRuleApiService).PostAllIoLimitRulesExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + IoLimitRuleParamsCreate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				// read after create mock error
				PreConfig: func() {
					mockCreateApi.UnPatch()
					mockCreateApi = mockey.Mock((*clientgen.IoLimitRuleApiService).PostAllIoLimitRulesExecute).Return(&clientgen.CreateResponse{
						Id: helper.StringPtr("1"),
					}, nil, nil).Build()
					mockReadApi = mockey.Mock((*clientgen.IoLimitRuleApiService).GetIoLimitRuleByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + IoLimitRuleParamsCreate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				Config: ProviderConfigForTesting + IoLimitRuleParamsCreate,
				PreConfig: func() {
					mockReadApi.UnPatch()
					mockCreateApi.UnPatch()
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "name", "tf_acc_io_limit_rule"),
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "type", "Absolute"),
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "max_iops", "1000"),
				),
			},
			{
				// read refresh error
				PreConfig: func() {
					mockReadApi = mockey.Mock((*clientgen.IoLimitRuleApiService).GetIoLimitRuleByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + IoLimitRuleParamsCreate,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				// delete error
				PreConfig: func() {
					mockReadApi.UnPatch()
					mockDeleteApi = mockey.Mock((*clientgen.IoLimitRuleApiService).DeleteIoLimitRuleByIdExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Destroy:     true,
				Config:      ProviderConfigForTesting + IoLimitRuleParamsUpdate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				// update error
				PreConfig: func() {
					mockDeleteApi.UnPatch()
					mockUpdateApi = mockey.Mock((*clientgen.IoLimitRuleApiService).PatchIoLimitRuleByIdExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + IoLimitRuleParamsUpdate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				PreConfig: func() {
					mockUpdateApi.UnPatch()
				},
				Config: ProviderConfigForTesting + IoLimitRuleParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "max_iops", "2000"),
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "max_bw", "2000"),
				),
			},
			{
				Config:            ProviderConfigForTesting + IoLimitRuleParamsUpdate,
				ResourceName:      "powerstore_io_limit_rule.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_acc_io_limit_rule", s[0].Attributes["name"])
					assert.Equal(t, "Absolute", s[0].Attributes["type"])
					return nil
				},
			},
		},
	})
}

func TestAccIoLimitRuleRes_CreateWithDensityType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleParamsCreateDensity,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "type", "Density"),
				),
			},
		},
	})
}

func TestAccIoLimitRuleRes_CreateWithAllOptionals(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleParamsCreateAllOptionals,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "max_iops", "1000"),
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "max_bw", "2000"),
					resource.TestCheckResourceAttr("powerstore_io_limit_rule.test", "burst_percentage", "20"),
				),
			},
		},
	})
}

func TestAccIoLimitRuleRes_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(`The argument "name" is required`),
			},
		},
	})
}

func TestAccIoLimitRuleRes_CreateWithoutType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleParamsCreateWithoutType,
				ExpectError: regexp.MustCompile(`The argument "type" is required`),
			},
		},
	})
}

func TestAccIoLimitRuleRes_CreateWithInvalidType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleParamsCreateWithInvalidType,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
		},
	})
}

// Unit tests

func TestResourceIoLimitRule_Metadata(t *testing.T) {
	r := &resourceIoLimitRule{}
	req := fwresource.MetadataRequest{}
	resp := &fwresource.MetadataResponse{}

	r.Metadata(context.Background(), req, resp)

	assert.Equal(t, "_io_limit_rule", resp.TypeName)
}

func TestResourceIoLimitRule_Schema(t *testing.T) {
	r := &resourceIoLimitRule{}
	req := fwresource.SchemaRequest{}
	resp := &fwresource.SchemaResponse{}

	r.Schema(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
}

func TestResourceIoLimitRule_Configure_Nil(t *testing.T) {
	r := &resourceIoLimitRule{}
	req := fwresource.ConfigureRequest{ProviderData: nil}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, r.client)
}

func TestResourceIoLimitRule_Configure_InvalidType(t *testing.T) {
	r := &resourceIoLimitRule{}
	req := fwresource.ConfigureRequest{ProviderData: "invalid"}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestResourceIoLimitRule_Configure_Success(t *testing.T) {
	r := &resourceIoLimitRule{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwresource.ConfigureRequest{ProviderData: c}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

func TestResourceIoLimitRule_UpdateState_AllFields(t *testing.T) {
	r := resourceIoLimitRule{}
	id := "rule-id"
	name := "test-rule"
	typeEnum := clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE
	maxIops := int32(1000)
	maxBw := int32(500)
	burstPct := int32(20)

	response := &clientgen.IoLimitRuleInstance{
		Id:              &id,
		Name:            &name,
		Type:            &typeEnum,
		MaxIops:         &maxIops,
		MaxBw:           &maxBw,
		BurstPercentage: &burstPct,
	}

	state := r.updateIoLimitRuleState(response)
	assert.Equal(t, "rule-id", state.ID.ValueString())
	assert.Equal(t, "test-rule", state.Name.ValueString())
	assert.Equal(t, "Absolute", state.Type.ValueString())
	assert.Equal(t, int32(1000), state.MaxIops.ValueInt32())
	assert.Equal(t, int32(500), state.MaxBw.ValueInt32())
	assert.Equal(t, int32(20), state.BurstPercentage.ValueInt32())
}

func TestResourceIoLimitRule_UpdateState_NilOptionals(t *testing.T) {
	r := resourceIoLimitRule{}
	id := "rule-id"
	name := "test-rule"
	typeEnum := clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY

	response := &clientgen.IoLimitRuleInstance{
		Id:   &id,
		Name: &name,
		Type: &typeEnum,
	}

	state := r.updateIoLimitRuleState(response)
	assert.Equal(t, "rule-id", state.ID.ValueString())
	assert.Equal(t, "Density", state.Type.ValueString())
	assert.Equal(t, int32(0), state.MaxIops.ValueInt32())
	assert.Equal(t, int32(0), state.MaxBw.ValueInt32())
	assert.Equal(t, int32(0), state.BurstPercentage.ValueInt32())
}

func TestResourceIoLimitRule_UpdateState_NilFields(t *testing.T) {
	r := resourceIoLimitRule{}
	response := &clientgen.IoLimitRuleInstance{}

	state := r.updateIoLimitRuleState(response)
	assert.Equal(t, "", state.ID.ValueString())
	assert.Equal(t, "", state.Name.ValueString())
	assert.Equal(t, "", state.Type.ValueString())
	assert.Equal(t, int32(0), state.MaxIops.ValueInt32())
}

// Terraform configs

var IoLimitRuleParamsCreate = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule"
	type     = "Absolute"
	max_iops = 1000
}
`

var IoLimitRuleParamsCreateDensity = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule_density"
	type     = "Density"
	max_iops = 10
}
`

var IoLimitRuleParamsCreateAllOptionals = `
resource "powerstore_io_limit_rule" "test" {
	name             = "tf_acc_io_limit_rule_full"
	type             = "Absolute"
	max_iops         = 1000
	max_bw           = 2000
	burst_percentage = 20
}
`

var IoLimitRuleParamsUpdate = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule"
	type     = "Absolute"
	max_iops = 2000
	max_bw   = 2000
}
`

var IoLimitRuleParamsCreateWithoutName = `
resource "powerstore_io_limit_rule" "test" {
	type     = "Absolute"
	max_iops = 1000
}
`

var IoLimitRuleParamsCreateWithoutType = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule"
	max_iops = 1000
}
`

var IoLimitRuleParamsCreateWithInvalidType = `
resource "powerstore_io_limit_rule" "test" {
	name = "tf_acc_io_limit_rule"
	type = "Invalid"
}
`
