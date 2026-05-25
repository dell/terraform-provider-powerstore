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

func TestAccIoLimitRuleDs_FetchAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleDataSourceParamsAll,
			},
		},
	})
}

func TestAccIoLimitRuleDs_FetchByName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleResourceParams + IoLimitRuleDataSourceParamsName,
			},
		},
	})
}

func TestAccIoLimitRuleDs_FetchByID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleResourceParams + IoLimitRuleDataSourceParamsID,
			},
		},
	})
}

func TestAccIoLimitRuleDs_FetchByType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleResourceParams + IoLimitRuleDataSourceParamsType,
			},
		},
	})
}

func TestAccIoLimitRuleDs_FetchByFilter(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + IoLimitRuleResourceParams + IoLimitRuleDataSourceParamsFilter,
			},
		},
	})
}

func TestAccIoLimitRuleDs_InvalidName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsInvalidName,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore I/O Limit Rule"),
			},
		},
	})
}

func TestAccIoLimitRuleDs_InvalidID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsInvalidID,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore I/O Limit Rule"),
			},
		},
	})
}

func TestAccIoLimitRuleDs_IDAndNameConflict(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsIDAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccIoLimitRuleDs_TypeAndNameConflict(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsTypeAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccIoLimitRuleDs_InvalidTypeValue(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsInvalidType,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
		},
	})
}

func TestAccIoLimitRuleDs_EmptyID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsEmptyID,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

func TestAccIoLimitRuleDs_EmptyName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + IoLimitRuleDataSourceParamsEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

// Unit tests

func TestIoLimitRuleDataSource_Metadata(t *testing.T) {
	d := &dataSourceIoLimitRule{}
	req := fwdatasource.MetadataRequest{}
	resp := &fwdatasource.MetadataResponse{}

	d.Metadata(context.Background(), req, resp)

	assert.Equal(t, "_io_limit_rule", resp.TypeName)
}

func TestIoLimitRuleDataSource_Schema(t *testing.T) {
	d := &dataSourceIoLimitRule{}
	req := fwdatasource.SchemaRequest{}
	resp := &fwdatasource.SchemaResponse{}

	d.Schema(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
}

func TestIoLimitRuleDataSource_Configure_Nil(t *testing.T) {
	d := &dataSourceIoLimitRule{}
	req := fwdatasource.ConfigureRequest{ProviderData: nil}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, d.client)
}

func TestIoLimitRuleDataSource_Configure_InvalidType(t *testing.T) {
	d := &dataSourceIoLimitRule{}
	req := fwdatasource.ConfigureRequest{ProviderData: "invalid"}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestIoLimitRuleDataSource_Configure_Success(t *testing.T) {
	d := &dataSourceIoLimitRule{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwdatasource.ConfigureRequest{ProviderData: c}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, d.client)
}

func TestUpdateIoLimitRuleState_AllFields(t *testing.T) {
	id := "rule-id-1"
	name := "test-rule"
	typeEnum := clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE
	maxIops := int32(1000)
	maxBw := int32(500)
	burstPct := int32(20)
	typeL10n := "Absolute"

	rules := []clientgen.IoLimitRuleInstance{
		{
			Id:              &id,
			Name:            &name,
			Type:            &typeEnum,
			MaxIops:         &maxIops,
			MaxBw:           &maxBw,
			BurstPercentage: &burstPct,
			TypeL10n:        &typeL10n,
		},
	}

	state := updateIoLimitRuleState(rules)
	assert.Len(t, state, 1)
	assert.Equal(t, "rule-id-1", state[0].ID.ValueString())
	assert.Equal(t, "test-rule", state[0].Name.ValueString())
	assert.Equal(t, "Absolute", state[0].Type.ValueString())
	assert.Equal(t, int32(1000), state[0].MaxIops.ValueInt32())
	assert.Equal(t, int32(500), state[0].MaxBw.ValueInt32())
	assert.Equal(t, int32(20), state[0].BurstPercentage.ValueInt32())
	assert.Equal(t, "Absolute", state[0].TypeL10n.ValueString())
}

func TestUpdateIoLimitRuleState_NilFields(t *testing.T) {
	rules := []clientgen.IoLimitRuleInstance{{}}

	state := updateIoLimitRuleState(rules)
	assert.Len(t, state, 1)
	assert.Equal(t, "", state[0].ID.ValueString())
	assert.Equal(t, "", state[0].Name.ValueString())
	assert.Equal(t, "", state[0].Type.ValueString())
	assert.Equal(t, int32(0), state[0].MaxIops.ValueInt32())
	assert.Equal(t, int32(0), state[0].MaxBw.ValueInt32())
	assert.Equal(t, int32(0), state[0].BurstPercentage.ValueInt32())
}

func TestUpdateIoLimitRuleState_DensityType(t *testing.T) {
	id := "rule-id-2"
	name := "density-rule"
	typeEnum := clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY

	rules := []clientgen.IoLimitRuleInstance{
		{Id: &id, Name: &name, Type: &typeEnum},
	}

	state := updateIoLimitRuleState(rules)
	assert.Len(t, state, 1)
	assert.Equal(t, "Density", state[0].Type.ValueString())
}

func TestUpdateIoLimitRuleState_Empty(t *testing.T) {
	state := updateIoLimitRuleState([]clientgen.IoLimitRuleInstance{})
	assert.Len(t, state, 0)
}

func TestUpdateIoLimitRuleState_Multiple(t *testing.T) {
	id1, name1 := "id-1", "rule-1"
	id2, name2 := "id-2", "rule-2"
	typeAbsolute := clientgen.BANDWIDTHLIMITTYPEENUM_ABSOLUTE
	typeDensity := clientgen.BANDWIDTHLIMITTYPEENUM_DENSITY

	rules := []clientgen.IoLimitRuleInstance{
		{Id: &id1, Name: &name1, Type: &typeAbsolute},
		{Id: &id2, Name: &name2, Type: &typeDensity},
	}

	state := updateIoLimitRuleState(rules)
	assert.Len(t, state, 2)
	assert.Equal(t, "id-1", state[0].ID.ValueString())
	assert.Equal(t, "Absolute", state[0].Type.ValueString())
	assert.Equal(t, "id-2", state[1].ID.ValueString())
	assert.Equal(t, "Density", state[1].Type.ValueString())
}

// Terraform configs

var IoLimitRuleResourceParams = `
resource "powerstore_io_limit_rule" "test" {
	name = "tf_acc_io_limit_rule"
	type = "Absolute"
	max_iops = 1000
}
`

var IoLimitRuleDataSourceParamsAll = `
data "powerstore_io_limit_rule" "test" {
}
`

var IoLimitRuleDataSourceParamsName = `
data "powerstore_io_limit_rule" "test" {
	depends_on = [powerstore_io_limit_rule.test]
	name = powerstore_io_limit_rule.test.name
}
`

var IoLimitRuleDataSourceParamsID = `
data "powerstore_io_limit_rule" "test" {
	depends_on = [powerstore_io_limit_rule.test]
	id = powerstore_io_limit_rule.test.id
}
`

var IoLimitRuleDataSourceParamsType = `
data "powerstore_io_limit_rule" "test" {
	depends_on = [powerstore_io_limit_rule.test]
	type = "Absolute"
}
`

var IoLimitRuleDataSourceParamsFilter = `
data "powerstore_io_limit_rule" "test" {
	depends_on = [powerstore_io_limit_rule.test]
	filter_expression = "name=ilike.tf_acc_*"
}
`

var IoLimitRuleDataSourceParamsInvalidName = `
data "powerstore_io_limit_rule" "test" {
	name = "invalid-name-does-not-exist"
}
`

var IoLimitRuleDataSourceParamsInvalidID = `
data "powerstore_io_limit_rule" "test" {
	id = "invalid-id-does-not-exist"
}
`

var IoLimitRuleDataSourceParamsIDAndName = `
data "powerstore_io_limit_rule" "test" {
	id   = "some-id"
	name = "some-name"
}
`

var IoLimitRuleDataSourceParamsTypeAndName = `
data "powerstore_io_limit_rule" "test" {
	type = "Absolute"
	name = "some-name"
}
`

var IoLimitRuleDataSourceParamsInvalidType = `
data "powerstore_io_limit_rule" "test" {
	type = "Invalid"
}
`

var IoLimitRuleDataSourceParamsEmptyID = `
data "powerstore_io_limit_rule" "test" {
	id = ""
}
`

var IoLimitRuleDataSourceParamsEmptyName = `
data "powerstore_io_limit_rule" "test" {
	name = ""
}
`
