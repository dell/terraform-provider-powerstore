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
	"strings"
	"testing"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Test Schema method for metro volume resource
func TestMetroVolumeResource_Schema(t *testing.T) {
	r := newMetroVolumeResource()
	resp := &fwresource.SchemaResponse{}
	r.Schema(context.Background(), fwresource.SchemaRequest{}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)

	// Check required attributes
	assert.True(t, resp.Schema.Attributes["volume_id"].IsRequired())
	assert.True(t, resp.Schema.Attributes["remote_system_id"].IsRequired())

	// Check computed attributes
	assert.True(t, resp.Schema.Attributes["id"].IsComputed())
	assert.True(t, resp.Schema.Attributes["state"].IsComputed())
	assert.True(t, resp.Schema.Attributes["remote_resource_id"].IsComputed())
	assert.True(t, resp.Schema.Attributes["data_transfer_state"].IsComputed())
	assert.True(t, resp.Schema.Attributes["metro_replication_session_id"].IsComputed())

	// Check optional attributes
	assert.True(t, resp.Schema.Attributes["remote_appliance_id"].IsOptional())
	assert.True(t, resp.Schema.Attributes["delete_remote_volume"].IsOptional())
	assert.True(t, resp.Schema.Attributes["force"].IsOptional())
}

// Test Metadata method
func TestMetroVolumeResource_Metadata(t *testing.T) {
	r := newMetroVolumeResource()
	resp := &fwresource.MetadataResponse{}
	r.Metadata(context.Background(), fwresource.MetadataRequest{ProviderTypeName: "powerstore"}, resp)
	assert.Equal(t, "powerstore_metro_volume", resp.TypeName)
}

// Test Configure method with nil provider data
func TestMetroVolumeResource_Configure_Nil(t *testing.T) {
	r := &resourceMetroVolume{}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: nil}, resp)
	assert.False(t, resp.Diagnostics.HasError())
}

// Test Configure method with wrong type
func TestMetroVolumeResource_Configure_WrongType(t *testing.T) {
	r := &resourceMetroVolume{}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: "wrong"}, resp)
	assert.True(t, resp.Diagnostics.HasError())
}

// Test Configure method with correct type
func TestMetroVolumeResource_Configure_Correct(t *testing.T) {
	r := &resourceMetroVolume{}
	resp := &fwresource.ConfigureResponse{}
	mockClient := &client.Client{GenClient: &clientgen.APIClient{}}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: mockClient}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

// Test updateState method
func TestMetroVolumeResource_UpdateState(t *testing.T) {
	r := &resourceMetroVolume{}
	state := &models.MetroVolumeResource{}

	sessionID := "test-session-id"
	stateEnum := clientgen.ReplicationStateEnum("OK")
	dtsEnum := clientgen.DataTransferStateEnum("Synchronized")
	localResID := "local-vol-id"
	remoteResID := "remote-vol-id"
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
	assert.Equal(t, "local-vol-id", state.VolumeID.ValueString())
	assert.Equal(t, "remote-vol-id", state.RemoteResourceID.ValueString())
	assert.Equal(t, "remote-sys-id", state.RemoteSystemID.ValueString())
}

// Test updateState with nil session
func TestMetroVolumeResource_UpdateState_NilSession(t *testing.T) {
	r := &resourceMetroVolume{}
	state := &models.MetroVolumeResource{
		ID: types.StringValue("existing"),
	}
	r.updateState(state, nil)
	assert.Equal(t, "existing", state.ID.ValueString())
}

// Test updateState with partial nil fields
func TestMetroVolumeResource_UpdateState_PartialNilFields(t *testing.T) {
	r := &resourceMetroVolume{}
	state := &models.MetroVolumeResource{}

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

// Acceptance test: Create metro volume - missing volume_id
func TestAccMetroVolume_MissingVolumeID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeParamsMissingVolumeID,
				ExpectError: regexp.MustCompile(`The argument "volume_id" is required`),
			},
		},
	})
}

// Acceptance test: Create metro volume - missing remote_system_id
func TestAccMetroVolume_MissingRemoteSystemID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeParamsMissingRemoteSystemID,
				ExpectError: regexp.MustCompile(`The argument "remote_system_id" is required`),
			},
		},
	})
}

// Acceptance test: Create metro volume - empty volume_id
func TestAccMetroVolume_EmptyVolumeID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeParamsEmptyVolumeID,
				ExpectError: regexp.MustCompile(`string length must be at least 1`),
			},
		},
	})
}

// Acceptance test: Create metro volume - empty remote_system_id
func TestAccMetroVolume_EmptyRemoteSystemID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + MetroVolumeParamsEmptyRemoteSystemID,
				ExpectError: regexp.MustCompile(`string length must be at least 1`),
			},
		},
	})
}

// Acceptance test: Create metro volume (mock only - requires primary volume not in replication)
func TestAccMetroVolume_CreateOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	// Only run on mock server - requires primary volume not in existing replication
	if !strings.HasPrefix(endpoint, "http://localhost:3003") {
		t.Skip("This test only runs on the mock server - requires primary volume not in existing replication")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(MetroVolumeParamsCreate, remoteSystemID),
			},
		},
	})
}

// HCL config strings
var MetroVolumeParamsMissingVolumeID = `
resource "powerstore_metro_volume" "test" {
  remote_system_id = "cd41130c-a751-4b39-bde1-b76f246c27b6"
}
`

var MetroVolumeParamsMissingRemoteSystemID = `
resource "powerstore_metro_volume" "test" {
  volume_id = "volume_post_id"
}
`

var MetroVolumeParamsEmptyVolumeID = `
resource "powerstore_metro_volume" "test" {
  volume_id        = ""
  remote_system_id = "cd41130c-a751-4b39-bde1-b76f246c27b6"
}
`

var MetroVolumeParamsEmptyRemoteSystemID = `
resource "powerstore_metro_volume" "test" {
  volume_id        = "volume_post_id"
  remote_system_id = ""
}
`

var MetroVolumeParamsCreate = `
resource "powerstore_metro_volume" "test" {
  volume_id        = "volume_post_id"
  remote_system_id = "%s"
}
`
