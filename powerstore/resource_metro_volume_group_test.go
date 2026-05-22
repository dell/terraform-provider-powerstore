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
	"testing"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Test Schema method for metro volume group resource
func TestMetroVolumeGroupResource_Schema(t *testing.T) {
	r := newMetroVolumeGroupResource()
	resp := &fwresource.SchemaResponse{}
	r.Schema(context.Background(), fwresource.SchemaRequest{}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)

	assert.True(t, resp.Schema.Attributes["volume_group_id"].IsRequired())
	assert.True(t, resp.Schema.Attributes["remote_system_id"].IsRequired())
	assert.True(t, resp.Schema.Attributes["id"].IsComputed())
	assert.True(t, resp.Schema.Attributes["state"].IsComputed())
	assert.True(t, resp.Schema.Attributes["remote_resource_id"].IsComputed())
	assert.True(t, resp.Schema.Attributes["data_transfer_state"].IsComputed())
	assert.True(t, resp.Schema.Attributes["metro_replication_session_id"].IsComputed())
	assert.True(t, resp.Schema.Attributes["remote_appliance_id"].IsOptional())
	assert.True(t, resp.Schema.Attributes["delete_remote_volume_group"].IsOptional())
	assert.True(t, resp.Schema.Attributes["force"].IsOptional())
}

// Test Metadata method
func TestMetroVolumeGroupResource_Metadata(t *testing.T) {
	r := newMetroVolumeGroupResource()
	resp := &fwresource.MetadataResponse{}
	r.Metadata(context.Background(), fwresource.MetadataRequest{ProviderTypeName: "powerstore"}, resp)
	assert.Equal(t, "powerstore_metro_volume_group", resp.TypeName)
}

// Test Configure method with nil provider data
func TestMetroVolumeGroupResource_Configure_Nil(t *testing.T) {
	r := &resourceMetroVolumeGroup{}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: nil}, resp)
	assert.False(t, resp.Diagnostics.HasError())
}

// Test Configure method with wrong type
func TestMetroVolumeGroupResource_Configure_WrongType(t *testing.T) {
	r := &resourceMetroVolumeGroup{}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: "wrong"}, resp)
	assert.True(t, resp.Diagnostics.HasError())
}

// Test Configure method with correct type
func TestMetroVolumeGroupResource_Configure_Correct(t *testing.T) {
	r := &resourceMetroVolumeGroup{}
	resp := &fwresource.ConfigureResponse{}
	mockClient := &client.Client{GenClient: &clientgen.APIClient{}}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: mockClient}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

// Test updateState method
func TestMetroVolumeGroupResource_UpdateState(t *testing.T) {
	r := &resourceMetroVolumeGroup{}
	state := &models.MetroVolumeGroupResource{}

	sessionID := "test-session-id"
	stateEnum := clientgen.ReplicationStateEnum("OK")
	dtsEnum := clientgen.DataTransferStateEnum("Synchronized")
	localResID := "local-vg-id"
	remoteResID := "remote-vg-id"
	remoteSystemID := "remote-sys-id"

	session := &clientgen.ReplicationSessionInstance{
		Id:                &sessionID,
		State:             &stateEnum,
		DataTransferState: &dtsEnum,
		LocalResourceId:   &localResID,
		RemoteResourceId:  &remoteResID,
		RemoteSystemId:    &remoteSystemID,
	}

	r.updateState(state, session)

	assert.Equal(t, "test-session-id", state.ID.ValueString())
	assert.Equal(t, "test-session-id", state.MetroReplicationSessionID.ValueString())
	assert.Equal(t, "OK", state.State.ValueString())
	assert.Equal(t, "Synchronized", state.DataTransferState.ValueString())
	assert.Equal(t, "local-vg-id", state.VolumeGroupID.ValueString())
	assert.Equal(t, "remote-vg-id", state.RemoteResourceID.ValueString())
	assert.Equal(t, "remote-sys-id", state.RemoteSystemID.ValueString())
}

// Test updateState with nil session
func TestMetroVolumeGroupResource_UpdateState_NilSession(t *testing.T) {
	r := &resourceMetroVolumeGroup{}
	state := &models.MetroVolumeGroupResource{
		ID: types.StringValue("existing"),
	}
	r.updateState(state, nil)
	assert.Equal(t, "existing", state.ID.ValueString())
}

// Test updateState with partial nil fields
func TestMetroVolumeGroupResource_UpdateState_PartialNilFields(t *testing.T) {
	r := &resourceMetroVolumeGroup{}
	state := &models.MetroVolumeGroupResource{}

	sessionID := "test-session-id"
	session := &clientgen.ReplicationSessionInstance{
		Id:                &sessionID,
		State:             nil,
		DataTransferState: nil,
		LocalResourceId:   nil,
		RemoteResourceId:  nil,
		RemoteSystemId:    nil,
	}

	r.updateState(state, session)

	assert.Equal(t, "test-session-id", state.ID.ValueString())
	assert.True(t, state.RemoteResourceID.IsNull())
}

// Acceptance test: Create metro volume group - missing volume_group_id
func TestAccMetroVolumeGroup_MissingVolumeGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeGroupParamsMissingVGID,
				ExpectError: regexp.MustCompile(`The argument "volume_group_id" is required`),
			},
		},
	})
}

// Acceptance test: Create metro volume group - missing remote_system_id
func TestAccMetroVolumeGroup_MissingRemoteSystemID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeGroupParamsMissingRemoteSystemID,
				ExpectError: regexp.MustCompile(`The argument "remote_system_id" is required`),
			},
		},
	})
}

// Acceptance test: Create metro volume group - empty volume_group_id
func TestAccMetroVolumeGroup_EmptyVolumeGroupID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeGroupParamsEmptyVGID,
				ExpectError: regexp.MustCompile(`string length must be at least 1`),
			},
		},
	})
}

// Acceptance test: Create metro volume group
func TestAccMetroVolumeGroup_CreateOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(MetroVolumeGroupParamsCreate, remoteSystemID),
			},
		},
	})
}

// HCL config strings
var MetroVolumeGroupParamsMissingVGID = `
resource "powerstore_metro_volume_group" "test" {
  remote_system_id = "cd41130c-a751-4b39-bde1-b76f246c27b6"
}
`

var MetroVolumeGroupParamsMissingRemoteSystemID = `
resource "powerstore_metro_volume_group" "test" {
  volume_group_id = "f64ad207-06eb-4098-b907-2a204cfb5ce9"
}
`

var MetroVolumeGroupParamsEmptyVGID = `
resource "powerstore_metro_volume_group" "test" {
  volume_group_id  = ""
  remote_system_id = "cd41130c-a751-4b39-bde1-b76f246c27b6"
}
`

var MetroVolumeGroupParamsCreate = `
resource "powerstore_metro_volume_group" "test" {
  volume_group_id  = "f64ad207-06eb-4098-b907-2a204cfb5ce9"
  remote_system_id = "%s"
}
`
