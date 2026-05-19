/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// Test to Fetch RecycleBin items
func TestAccRecycleBinDs_FetchAll(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceAll,
			},
		},
	})
}

// Test to Fetch RecycleBin items by resource_type
func TestAccRecycleBinDs_FetchByResourceType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceByVolume,
			},
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceByVolumeGroup,
			},
		},
	})
}

// Test to Fetch RecycleBin item by invalid ID
func TestAccRecycleBinDs_FetchByInvalidID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RecycleBinDataSourceByInvalidID,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Recycle Bin Item"),
			},
		},
	})
}

// Test to Fetch RecycleBin items by filter expression
func TestAccRecycleBinDs_FetchByFilter(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinDataSourceByFilter,
			},
		},
	})
}

var RecycleBinDataSourceAll = `
data "powerstore_recycle_bin" "test" {
}
`

var RecycleBinDataSourceByVolume = `
data "powerstore_recycle_bin" "test" {
	resource_type = "volume"
}
`

var RecycleBinDataSourceByVolumeGroup = `
data "powerstore_recycle_bin" "test" {
	resource_type = "volume_group"
}
`

var RecycleBinDataSourceByInvalidID = `
data "powerstore_recycle_bin" "test" {
	id = "invalid-id-does-not-exist"
}
`

var RecycleBinDataSourceByFilter = `
data "powerstore_recycle_bin" "test" {
	filter_expression = "resource_type=eq.volume"
}
`

// Unit tests for Configure method
func TestRecycleBinDataSource_Configure_InvalidType(t *testing.T) {
	d := &recycleBinDataSource{}
	req := fwdatasource.ConfigureRequest{
		ProviderData: "invalid_type",
	}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.NotEmpty(t, resp.Diagnostics.Errors()[0].Summary)
}

func TestRecycleBinDataSource_Configure_Nil(t *testing.T) {
	d := &recycleBinDataSource{}
	req := fwdatasource.ConfigureRequest{
		ProviderData: nil,
	}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, d.client)
}

func TestRecycleBinDataSource_Configure_Success(t *testing.T) {
	d := &recycleBinDataSource{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwdatasource.ConfigureRequest{
		ProviderData: c,
	}
	resp := &fwdatasource.ConfigureResponse{}

	d.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, d.client)
}

// Test mapRecycleBinItemsToState helper function
func TestMapRecycleBinItemsToState(t *testing.T) {
	id := "test-id"
	name := "test-volume"
	resourceType := clientgen.RECYCLEBINRESOURCETYPEENUM_VOLUME
	logicalProvisioned := int64(1024)
	logicalUsed := int64(512)
	applianceID := "appliance-123"

	items := []clientgen.RecycleBinInstance{
		{
			Id:                 &id,
			Name:               &name,
			ResourceType:       &resourceType,
			LogicalProvisioned: &logicalProvisioned,
			LogicalUsed:        &logicalUsed,
			ApplianceId:        &applianceID,
		},
	}

	state := mapRecycleBinItemsToState(items)
	assert.Len(t, state, 1)
	assert.Equal(t, "test-id", state[0].ID.ValueString())
	assert.Equal(t, "test-volume", state[0].Name.ValueString())
	assert.Equal(t, "volume", state[0].ResourceType.ValueString())
}

func TestMapRecycleBinItemsToState_Empty(t *testing.T) {
	items := []clientgen.RecycleBinInstance{}
	state := mapRecycleBinItemsToState(items)
	assert.Len(t, state, 0)
}
