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

func TestAccFileIoLimitRule(t *testing.T) {
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
					mockCreateApi = mockey.Mock((*clientgen.FileIoLimitRuleApiService).PostAllFileIoLimitRulesExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsCreate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				// read after create mock error
				PreConfig: func() {
					mockCreateApi.UnPatch()
					mockCreateApi = mockey.Mock((*clientgen.FileIoLimitRuleApiService).PostAllFileIoLimitRulesExecute).Return(&clientgen.CreateResponse{
						Id: helper.StringPtr("1"),
					}, nil, nil).Build()
					mockReadApi = mockey.Mock((*clientgen.FileIoLimitRuleApiService).GetFileIoLimitRuleByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsCreate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				Config: ProviderConfigForTesting + FileIoLimitRuleParamsCreate,
				PreConfig: func() {
					mockReadApi.UnPatch()
					mockCreateApi.UnPatch()
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_file_io_limit_rule.test", "name", "tf_acc_file_io_limit_rule"),
					resource.TestCheckResourceAttr("powerstore_file_io_limit_rule.test", "max_bw", "100"),
				),
			},
			{
				// read refresh error
				PreConfig: func() {
					mockReadApi = mockey.Mock((*clientgen.FileIoLimitRuleApiService).GetFileIoLimitRuleByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsCreate,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				// delete error
				PreConfig: func() {
					mockReadApi.UnPatch()
					mockDeleteApi = mockey.Mock((*clientgen.FileIoLimitRuleApiService).DeleteFileIoLimitRuleByIdExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Destroy:     true,
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsUpdate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				// update error
				PreConfig: func() {
					mockDeleteApi.UnPatch()
					mockUpdateApi = mockey.Mock((*clientgen.FileIoLimitRuleApiService).PatchFileIoLimitRuleByIdExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsUpdate,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				PreConfig: func() {
					mockUpdateApi.UnPatch()
				},
				Config: ProviderConfigForTesting + FileIoLimitRuleParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_file_io_limit_rule.test", "name", "tf_acc_file_io_limit_rule"),
					resource.TestCheckResourceAttr("powerstore_file_io_limit_rule.test", "max_bw", "200"),
				),
			},
			{
				Config:            ProviderConfigForTesting + FileIoLimitRuleParamsUpdate,
				ResourceName:      "powerstore_file_io_limit_rule.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_acc_file_io_limit_rule", s[0].Attributes["name"])
					return nil
				},
			},
		},
	})
}

func TestAccFileIoLimitRule_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(`The argument "name" is required`),
			},
		},
	})
}

func TestAccFileIoLimitRule_CreateWithEmptyName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsCreateWithEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

func TestAccFileIoLimitRule_CreateWithoutMaxBw(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleParamsCreateWithoutMaxBw,
				ExpectError: regexp.MustCompile(`The argument "max_bw" is required`),
			},
		},
	})
}

// Unit tests

func TestResourceFileIoLimitRule_Metadata(t *testing.T) {
	r := &resourceFileIoLimitRule{}
	req := fwresource.MetadataRequest{}
	resp := &fwresource.MetadataResponse{}

	r.Metadata(context.Background(), req, resp)

	assert.Equal(t, "_file_io_limit_rule", resp.TypeName)
}

func TestResourceFileIoLimitRule_Schema(t *testing.T) {
	r := &resourceFileIoLimitRule{}
	req := fwresource.SchemaRequest{}
	resp := &fwresource.SchemaResponse{}

	r.Schema(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
}

func TestResourceFileIoLimitRule_Configure_Nil(t *testing.T) {
	r := &resourceFileIoLimitRule{}
	req := fwresource.ConfigureRequest{ProviderData: nil}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, r.client)
}

func TestResourceFileIoLimitRule_Configure_InvalidType(t *testing.T) {
	r := &resourceFileIoLimitRule{}
	req := fwresource.ConfigureRequest{ProviderData: "invalid"}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestResourceFileIoLimitRule_Configure_Success(t *testing.T) {
	r := &resourceFileIoLimitRule{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwresource.ConfigureRequest{ProviderData: c}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

func TestResourceFileIoLimitRule_UpdateState_AllFields(t *testing.T) {
	r := resourceFileIoLimitRule{}
	id := "rule-id"
	name := "test-rule"
	maxBw := int32(100)

	response := &clientgen.FileIoLimitRuleInstance{
		Id:    &id,
		Name:  &name,
		MaxBw: &maxBw,
	}

	state := r.updateFileIoLimitRuleState(response)
	assert.Equal(t, "rule-id", state.ID.ValueString())
	assert.Equal(t, "test-rule", state.Name.ValueString())
	assert.Equal(t, int32(100), state.MaxBw.ValueInt32())
}

func TestResourceFileIoLimitRule_UpdateState_NilFields(t *testing.T) {
	r := resourceFileIoLimitRule{}
	response := &clientgen.FileIoLimitRuleInstance{}

	state := r.updateFileIoLimitRuleState(response)
	assert.Equal(t, "", state.ID.ValueString())
	assert.Equal(t, "", state.Name.ValueString())
	assert.Equal(t, int32(0), state.MaxBw.ValueInt32())
}

// Terraform configs

var FileIoLimitRuleParamsCreate = `
resource "powerstore_file_io_limit_rule" "test" {
	name   = "tf_acc_file_io_limit_rule"
	max_bw = 100
}
`

var FileIoLimitRuleParamsUpdate = `
resource "powerstore_file_io_limit_rule" "test" {
	name   = "tf_acc_file_io_limit_rule"
	max_bw = 200
}
`

var FileIoLimitRuleParamsCreateWithoutName = `
resource "powerstore_file_io_limit_rule" "test" {
	max_bw = 100
}
`

var FileIoLimitRuleParamsCreateWithEmptyName = `
resource "powerstore_file_io_limit_rule" "test" {
	name   = ""
	max_bw = 100
}
`

var FileIoLimitRuleParamsCreateWithoutMaxBw = `
resource "powerstore_file_io_limit_rule" "test" {
	name = "tf_acc_file_io_limit_rule"
}
`
