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

	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Acceptance tests

func TestAccFileIoLimitRuleDs_FetchAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + FileIoLimitRuleDataSourceParamsAll,
			},
		},
	})
}

func TestAccFileIoLimitRuleDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + FileIoLimitRuleResourceParams + FileIoLimitRuleDataSourceParamsName,
			},
			{
				Config: ProviderConfigForTesting + FileIoLimitRuleResourceParams + FileIoLimitRuleDataSourceParamsID,
			},
			{
				Config: ProviderConfigForTesting + FileIoLimitRuleResourceParams + FileIoLimitRuleDataSourceParamsFilter,
			},
		},
	})
}

func TestAccFileIoLimitRuleDs_InvalidName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleDataSourceParamsInvalidName,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore File I/O Limit Rule"),
			},
		},
	})
}

func TestAccFileIoLimitRuleDs_InvalidID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleDataSourceParamsInvalidID,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore File I/O Limit Rule"),
			},
		},
	})
}

func TestAccFileIoLimitRuleDs_IDAndNameConflict(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleDataSourceParamsIDAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccFileIoLimitRuleDs_EmptyID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleDataSourceParamsEmptyID,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

func TestAccFileIoLimitRuleDs_EmptyName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FileIoLimitRuleDataSourceParamsEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

// Unit tests

func TestFileIoLimitRuleDataSource_Metadata(t *testing.T) {
	d := &dataSourceFileIoLimitRule{}
	req := fwdatasource.MetadataRequest{}
	resp := &fwdatasource.MetadataResponse{}

	d.Metadata(context.Background(), req, resp)

	assert.Equal(t, "_file_io_limit_rule", resp.TypeName)
}

func TestFileIoLimitRuleDataSource_Schema(t *testing.T) {
	d := &dataSourceFileIoLimitRule{}
	req := fwdatasource.SchemaRequest{}
	resp := &fwdatasource.SchemaResponse{}

	d.Schema(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
}

func TestFileIoLimitRuleDataSource_Configure_Nil(t *testing.T) {
	d := &dataSourceFileIoLimitRule{}
	req := fwdatasource.ConfigureRequest{ProviderData: nil}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, d.client)
}

func TestFileIoLimitRuleDataSource_Configure_InvalidType(t *testing.T) {
	d := &dataSourceFileIoLimitRule{}
	req := fwdatasource.ConfigureRequest{ProviderData: "invalid"}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestFileIoLimitRuleDataSource_Configure_Success(t *testing.T) {
	d := &dataSourceFileIoLimitRule{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwdatasource.ConfigureRequest{ProviderData: c}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, d.client)
}

func TestUpdateFileIoLimitRuleState_AllFields(t *testing.T) {
	id := "rule-id-1"
	name := "test-rule"
	maxBw := int32(100)

	rules := []clientgen.FileIoLimitRuleInstance{
		{Id: &id, Name: &name, MaxBw: &maxBw},
	}

	state := updateFileIoLimitRuleState(rules)
	assert.Len(t, state, 1)
	assert.Equal(t, "rule-id-1", state[0].ID.ValueString())
	assert.Equal(t, "test-rule", state[0].Name.ValueString())
	assert.Equal(t, int32(100), state[0].MaxBw.ValueInt32())
}

func TestUpdateFileIoLimitRuleState_NilFields(t *testing.T) {
	rules := []clientgen.FileIoLimitRuleInstance{{}}

	state := updateFileIoLimitRuleState(rules)
	assert.Len(t, state, 1)
	assert.Equal(t, "", state[0].ID.ValueString())
	assert.Equal(t, "", state[0].Name.ValueString())
	assert.Equal(t, int32(0), state[0].MaxBw.ValueInt32())
}

func TestUpdateFileIoLimitRuleState_Empty(t *testing.T) {
	state := updateFileIoLimitRuleState([]clientgen.FileIoLimitRuleInstance{})
	assert.Len(t, state, 0)
}

func TestUpdateFileIoLimitRuleState_Multiple(t *testing.T) {
	id1, name1, bw1 := "id-1", "rule-1", int32(100)
	id2, name2, bw2 := "id-2", "rule-2", int32(200)

	rules := []clientgen.FileIoLimitRuleInstance{
		{Id: &id1, Name: &name1, MaxBw: &bw1},
		{Id: &id2, Name: &name2, MaxBw: &bw2},
	}

	state := updateFileIoLimitRuleState(rules)
	assert.Len(t, state, 2)
	assert.Equal(t, "id-1", state[0].ID.ValueString())
	assert.Equal(t, "id-2", state[1].ID.ValueString())
	assert.Equal(t, int32(100), state[0].MaxBw.ValueInt32())
	assert.Equal(t, int32(200), state[1].MaxBw.ValueInt32())
}

// Terraform configs

var FileIoLimitRuleResourceParams = `
resource "powerstore_file_io_limit_rule" "test" {
	name   = "tf_acc_file_io_limit_rule"
	max_bw = 100
}
`

var FileIoLimitRuleDataSourceParamsAll = `
data "powerstore_file_io_limit_rule" "test" {
}
`

var FileIoLimitRuleDataSourceParamsName = `
data "powerstore_file_io_limit_rule" "test" {
	depends_on = [powerstore_file_io_limit_rule.test]
	name = powerstore_file_io_limit_rule.test.name
}
`

var FileIoLimitRuleDataSourceParamsID = `
data "powerstore_file_io_limit_rule" "test" {
	depends_on = [powerstore_file_io_limit_rule.test]
	id = powerstore_file_io_limit_rule.test.id
}
`

var FileIoLimitRuleDataSourceParamsFilter = `
data "powerstore_file_io_limit_rule" "test" {
	depends_on = [powerstore_file_io_limit_rule.test]
	filter_expression = "name=ilike.tf_acc_*"
}
`

var FileIoLimitRuleDataSourceParamsInvalidName = `
data "powerstore_file_io_limit_rule" "test" {
	name = "invalid-name-does-not-exist"
}
`

var FileIoLimitRuleDataSourceParamsInvalidID = `
data "powerstore_file_io_limit_rule" "test" {
	id = "invalid-id-does-not-exist"
}
`

var FileIoLimitRuleDataSourceParamsIDAndName = `
data "powerstore_file_io_limit_rule" "test" {
	id   = "some-id"
	name = "some-name"
}
`

var FileIoLimitRuleDataSourceParamsEmptyID = `
data "powerstore_file_io_limit_rule" "test" {
	id = ""
}
`

var FileIoLimitRuleDataSourceParamsEmptyName = `
data "powerstore_file_io_limit_rule" "test" {
	name = ""
}
`
