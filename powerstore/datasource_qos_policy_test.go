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

func TestAccQosPolicyDs_FetchAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyDataSourceParamsAll,
			},
		},
	})
}

func TestAccQosPolicyDs_FetchByName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyResourceParams + QosPolicyDataSourceParamsName,
			},
		},
	})
}

func TestAccQosPolicyDs_FetchByID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyResourceParams + QosPolicyDataSourceParamsID,
			},
		},
	})
}

func TestAccQosPolicyDs_FetchByType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyResourceParams + QosPolicyDataSourceParamsType,
			},
		},
	})
}

func TestAccQosPolicyDs_FetchByFilter(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + QosPolicyResourceParams + QosPolicyDataSourceParamsFilter,
			},
		},
	})
}

func TestAccQosPolicyDs_InvalidName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsInvalidName,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore QoS Policy"),
			},
		},
	})
}

func TestAccQosPolicyDs_InvalidID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsInvalidID,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore QoS Policy"),
			},
		},
	})
}

func TestAccQosPolicyDs_IDAndNameConflict(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsIDAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccQosPolicyDs_TypeAndNameConflict(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsTypeAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccQosPolicyDs_InvalidTypeValue(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsInvalidType,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
		},
	})
}

func TestAccQosPolicyDs_EmptyID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsEmptyID,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

func TestAccQosPolicyDs_EmptyName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + QosPolicyDataSourceParamsEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

// Unit tests

func TestQosPolicyDataSource_Metadata(t *testing.T) {
	d := &dataSourceQosPolicy{}
	req := fwdatasource.MetadataRequest{}
	resp := &fwdatasource.MetadataResponse{}

	d.Metadata(context.Background(), req, resp)

	assert.Equal(t, "_qos_policy", resp.TypeName)
}

func TestQosPolicyDataSource_Schema(t *testing.T) {
	d := &dataSourceQosPolicy{}
	req := fwdatasource.SchemaRequest{}
	resp := &fwdatasource.SchemaResponse{}

	d.Schema(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
}

func TestQosPolicyDataSource_Configure_Nil(t *testing.T) {
	d := &dataSourceQosPolicy{}
	req := fwdatasource.ConfigureRequest{ProviderData: nil}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, d.client)
}

func TestQosPolicyDataSource_Configure_InvalidType(t *testing.T) {
	d := &dataSourceQosPolicy{}
	req := fwdatasource.ConfigureRequest{ProviderData: "invalid"}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
}

func TestQosPolicyDataSource_Configure_Success(t *testing.T) {
	d := &dataSourceQosPolicy{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwdatasource.ConfigureRequest{ProviderData: c}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, d.client)
}

func TestUpdateQosPolicyState_AllFields(t *testing.T) {
	id := "policy-id-1"
	name := "test-policy"
	description := "Test Description"
	typeEnum := clientgen.POLICYTYPEENUM_QO_S
	managedBy := clientgen.POLICYMANAGEDBYENUM_USER
	managedByL10n := "User"
	managedById := "managed-id"
	typeL10n := "QoS"
	isReadOnly := false
	isReplica := false
	fileIoRuleId := "file-rule-id"
	ioRuleId := "io-rule-id"

	policies := []clientgen.PolicyInstance{
		{
			Id:                &id,
			Name:              &name,
			Description:       &description,
			Type:              &typeEnum,
			ManagedBy:         &managedBy,
			ManagedByL10n:     &managedByL10n,
			ManagedById:       &managedById,
			TypeL10n:          &typeL10n,
			IsReadOnly:        &isReadOnly,
			IsReplica:         &isReplica,
			FileIoLimitRuleId: &fileIoRuleId,
			IoLimitRule:       &clientgen.IoLimitRuleInstance{Id: &ioRuleId},
		},
	}

	state := updateQosPolicyState(policies)
	assert.Len(t, state, 1)
	assert.Equal(t, "policy-id-1", state[0].ID.ValueString())
	assert.Equal(t, "test-policy", state[0].Name.ValueString())
	assert.Equal(t, "Test Description", state[0].Description.ValueString())
	assert.Equal(t, "QoS", state[0].Type.ValueString())
	assert.Equal(t, "User", state[0].ManagedBy.ValueString())
	assert.Equal(t, "managed-id", state[0].ManagedById.ValueString())
	assert.Equal(t, false, state[0].IsReadOnly.ValueBool())
	assert.Equal(t, false, state[0].IsReplica.ValueBool())
	assert.Equal(t, "file-rule-id", state[0].FileIoLimitRuleId.ValueString())
	assert.Equal(t, "io-rule-id", state[0].IoLimitRuleId.ValueString())
}

func TestUpdateQosPolicyState_NilFields(t *testing.T) {
	policies := []clientgen.PolicyInstance{{}}

	state := updateQosPolicyState(policies)
	assert.Len(t, state, 1)
	assert.Equal(t, "", state[0].ID.ValueString())
	assert.Equal(t, "", state[0].Name.ValueString())
	assert.Equal(t, "", state[0].Description.ValueString())
	assert.Equal(t, "", state[0].Type.ValueString())
	assert.True(t, state[0].FileIoLimitRuleId.IsNull())
	assert.True(t, state[0].IoLimitRuleId.IsNull())
}

func TestUpdateQosPolicyState_NilIoLimitRule(t *testing.T) {
	id := "policy-id"
	policies := []clientgen.PolicyInstance{{Id: &id, IoLimitRule: nil}}

	state := updateQosPolicyState(policies)
	assert.Len(t, state, 1)
	assert.True(t, state[0].IoLimitRuleId.IsNull())
}

func TestUpdateQosPolicyState_IoLimitRuleNilId(t *testing.T) {
	id := "policy-id"
	policies := []clientgen.PolicyInstance{
		{Id: &id, IoLimitRule: &clientgen.IoLimitRuleInstance{Id: nil}},
	}

	state := updateQosPolicyState(policies)
	assert.Len(t, state, 1)
	assert.True(t, state[0].IoLimitRuleId.IsNull())
}

func TestUpdateQosPolicyState_FilePerformanceType(t *testing.T) {
	id := "policy-id-2"
	typeEnum := clientgen.POLICYTYPEENUM_FILE_PERFORMANCE

	policies := []clientgen.PolicyInstance{
		{Id: &id, Type: &typeEnum},
	}

	state := updateQosPolicyState(policies)
	assert.Len(t, state, 1)
	assert.Equal(t, "File_Performance", state[0].Type.ValueString())
}

func TestUpdateQosPolicyState_Empty(t *testing.T) {
	state := updateQosPolicyState([]clientgen.PolicyInstance{})
	assert.Len(t, state, 0)
}

func TestUpdateQosPolicyState_Multiple(t *testing.T) {
	id1, id2 := "id-1", "id-2"
	name1, name2 := "policy-1", "policy-2"
	typeQos := clientgen.POLICYTYPEENUM_QO_S
	typeFile := clientgen.POLICYTYPEENUM_FILE_PERFORMANCE

	policies := []clientgen.PolicyInstance{
		{Id: &id1, Name: &name1, Type: &typeQos},
		{Id: &id2, Name: &name2, Type: &typeFile},
	}

	state := updateQosPolicyState(policies)
	assert.Len(t, state, 2)
	assert.Equal(t, "id-1", state[0].ID.ValueString())
	assert.Equal(t, "QoS", state[0].Type.ValueString())
	assert.Equal(t, "id-2", state[1].ID.ValueString())
	assert.Equal(t, "File_Performance", state[1].Type.ValueString())
}

// Terraform configs

var QosPolicyResourceParams = `
resource "powerstore_io_limit_rule" "test" {
	name     = "tf_acc_io_limit_rule_ds"
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

var QosPolicyDataSourceParamsAll = `
data "powerstore_qos_policy" "test" {
}
`

var QosPolicyDataSourceParamsName = `
data "powerstore_qos_policy" "test" {
	depends_on = [powerstore_qos_policy.test]
	name = powerstore_qos_policy.test.name
}
`

var QosPolicyDataSourceParamsID = `
data "powerstore_qos_policy" "test" {
	depends_on = [powerstore_qos_policy.test]
	id = powerstore_qos_policy.test.id
}
`

var QosPolicyDataSourceParamsType = `
data "powerstore_qos_policy" "test" {
	depends_on = [powerstore_qos_policy.test]
	type = "QoS"
}
`

var QosPolicyDataSourceParamsFilter = `
data "powerstore_qos_policy" "test" {
	depends_on = [powerstore_qos_policy.test]
	filter_expression = "name=ilike.tf_acc_*"
}
`

var QosPolicyDataSourceParamsInvalidName = `
data "powerstore_qos_policy" "test" {
	name = "invalid-name-does-not-exist"
}
`

var QosPolicyDataSourceParamsInvalidID = `
data "powerstore_qos_policy" "test" {
	id = "invalid-id-does-not-exist"
}
`

var QosPolicyDataSourceParamsIDAndName = `
data "powerstore_qos_policy" "test" {
	id   = "some-id"
	name = "some-name"
}
`

var QosPolicyDataSourceParamsTypeAndName = `
data "powerstore_qos_policy" "test" {
	type = "QoS"
	name = "some-name"
}
`

var QosPolicyDataSourceParamsInvalidType = `
data "powerstore_qos_policy" "test" {
	type = "Invalid"
}
`

var QosPolicyDataSourceParamsEmptyID = `
data "powerstore_qos_policy" "test" {
	id = ""
}
`

var QosPolicyDataSourceParamsEmptyName = `
data "powerstore_qos_policy" "test" {
	name = ""
}
`
