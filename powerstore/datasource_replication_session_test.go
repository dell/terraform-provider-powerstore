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
	"strings"
	"testing"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Test Schema method for replication session datasource
func TestReplicationSessionDataSource_Schema(t *testing.T) {
	d := newReplicationSessionDataSource()
	resp := &fwdatasource.SchemaResponse{}
	d.Schema(context.Background(), fwdatasource.SchemaRequest{}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)

	assert.True(t, resp.Schema.Attributes["id"].IsOptional())
	assert.True(t, resp.Schema.Attributes["replication_sessions"].IsComputed())
}

// Test Metadata method
func TestReplicationSessionDataSource_Metadata(t *testing.T) {
	d := newReplicationSessionDataSource()
	resp := &fwdatasource.MetadataResponse{}
	d.Metadata(context.Background(), fwdatasource.MetadataRequest{ProviderTypeName: "powerstore"}, resp)
	assert.Equal(t, "powerstore_replication_session", resp.TypeName)
}

// Test Configure method with nil provider data
func TestReplicationSessionDataSource_Configure_Nil(t *testing.T) {
	d := &datasourceReplicationSession{}
	resp := &fwdatasource.ConfigureResponse{}
	d.Configure(context.Background(), fwdatasource.ConfigureRequest{ProviderData: nil}, resp)
	assert.False(t, resp.Diagnostics.HasError())
}

// Test Configure method with wrong type
func TestReplicationSessionDataSource_Configure_WrongType(t *testing.T) {
	d := &datasourceReplicationSession{}
	resp := &fwdatasource.ConfigureResponse{}
	d.Configure(context.Background(), fwdatasource.ConfigureRequest{ProviderData: "wrong"}, resp)
	assert.True(t, resp.Diagnostics.HasError())
}

// Test Configure method with correct type
func TestReplicationSessionDataSource_Configure_Correct(t *testing.T) {
	d := &datasourceReplicationSession{}
	resp := &fwdatasource.ConfigureResponse{}
	mockClient := &client.Client{GenClient: &clientgen.APIClient{}}
	d.Configure(context.Background(), fwdatasource.ConfigureRequest{ProviderData: mockClient}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, d.client)
}

// Test mapReplicationSessionsToState with empty slice
func TestMapReplicationSessionsToState_Empty(t *testing.T) {
	result := mapReplicationSessionsToState([]clientgen.ReplicationSessionInstance{})
	assert.Equal(t, 0, len(result))
}

// Test mapReplicationSessionsToState with nil fields
func TestMapReplicationSessionsToState_NilFields(t *testing.T) {
	sessions := []clientgen.ReplicationSessionInstance{
		{
			Id:                     nil,
			State:                  nil,
			Role:                   nil,
			ResourceType:           nil,
			DataTransferState:      nil,
			Type:                   nil,
			ProgressPercentage:     nil,
			LastSyncDuration:       nil,
			FailoverTestInProgress: nil,
		},
	}
	result := mapReplicationSessionsToState(sessions)
	assert.Equal(t, 1, len(result))
	assert.True(t, result[0].ID.IsNull())
	assert.True(t, result[0].State.IsNull())
	assert.True(t, result[0].Role.IsNull())
	assert.True(t, result[0].ResourceType.IsNull())
	assert.True(t, result[0].DataTransferState.IsNull())
	assert.True(t, result[0].Type.IsNull())
	assert.True(t, result[0].ProgressPercentage.IsNull())
	assert.True(t, result[0].LastSyncDuration.IsNull())
	assert.True(t, result[0].FailoverTestInProgress.IsNull())
}

// Test mapReplicationSessionsToState with full data
func TestMapReplicationSessionsToState_FullData(t *testing.T) {
	id := "test-id"
	stateEnum := clientgen.ReplicationStateEnum("OK")
	roleEnum := clientgen.ReplicationRoleEnum("Metro_Preferred")
	resTypeEnum := clientgen.ReplicatedResourceTypeEnum("Volume")
	dtsEnum := clientgen.DataTransferStateEnum("Synchronized")
	typeEnum := clientgen.ReplicationSessionTypeEnum("Metro_Active_Active")
	localResID := "local-vol"
	remoteResID := "remote-vol"
	remoteSystemID := "remote-sys"
	progress := int32(100)
	duration := int32(500)
	ftInProgress := false
	stateL10n := "OK"
	roleL10n := "Metro Preferred"
	resTypeL10n := "Volume"
	typeL10n := "Metro Active Active"

	sessions := []clientgen.ReplicationSessionInstance{
		{
			Id:                     &id,
			State:                  &stateEnum,
			Role:                   &roleEnum,
			ResourceType:           &resTypeEnum,
			DataTransferState:      &dtsEnum,
			Type:                   &typeEnum,
			LocalResourceId:        &localResID,
			RemoteResourceId:       &remoteResID,
			RemoteSystemId:         &remoteSystemID,
			ProgressPercentage:     &progress,
			LastSyncDuration:       &duration,
			FailoverTestInProgress: &ftInProgress,
			StateL10n:              &stateL10n,
			RoleL10n:               &roleL10n,
			ResourceTypeL10n:       &resTypeL10n,
			TypeL10n:               &typeL10n,
		},
	}
	result := mapReplicationSessionsToState(sessions)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "test-id", result[0].ID.ValueString())
	assert.Equal(t, "OK", result[0].State.ValueString())
	assert.Equal(t, "Metro_Preferred", result[0].Role.ValueString())
	assert.Equal(t, "Volume", result[0].ResourceType.ValueString())
	assert.Equal(t, "Synchronized", result[0].DataTransferState.ValueString())
	assert.Equal(t, "Metro_Active_Active", result[0].Type.ValueString())
	assert.Equal(t, "local-vol", result[0].LocalResourceID.ValueString())
	assert.Equal(t, "remote-vol", result[0].RemoteResourceID.ValueString())
	assert.Equal(t, "remote-sys", result[0].RemoteSystemID.ValueString())
	assert.Equal(t, int64(100), result[0].ProgressPercentage.ValueInt64())
	assert.Equal(t, int64(500), result[0].LastSyncDuration.ValueInt64())
	assert.Equal(t, false, result[0].FailoverTestInProgress.ValueBool())
}

// Test mapReplicationSessionsToState with multiple sessions
func TestMapReplicationSessionsToState_MultipleSessions(t *testing.T) {
	id1 := "session-1"
	id2 := "session-2"
	sessions := []clientgen.ReplicationSessionInstance{
		{Id: &id1},
		{Id: &id2},
	}
	result := mapReplicationSessionsToState(sessions)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "session-1", result[0].ID.ValueString())
	assert.Equal(t, "session-2", result[1].ID.ValueString())
}

// Acceptance test: Read all replication sessions
func TestAccReplicationSessionDataSource_ReadAllMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ReplSessionDSReadAll,
			},
		},
	})
}

// Acceptance test: Read replication session by ID
func TestAccReplicationSessionDataSource_ReadByIDMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	// Skip on real array if session doesn't exist
	if !strings.HasPrefix(endpoint, "http://localhost:3003") {
		t.Skip("This test requires an existing replication session - skipping on real array")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionDSReadByID, replicationSessionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_replication_session.by_id", "id", replicationSessionID),
					resource.TestCheckResourceAttr("data.powerstore_replication_session.by_id", "replication_sessions.#", "1"),
				),
			},
		},
	})
}

// Test ReplicationSessionDataSource model
func TestReplicationSessionDataSource_ModelDefaults(t *testing.T) {
	m := models.ReplicationSessionDataSource{}
	assert.True(t, m.ID.IsNull())
	assert.Nil(t, m.ReplicationSessions)
}

// HCL config strings
var ReplSessionDSReadAll = `
data "powerstore_replication_session" "all" {
}
`

var ReplSessionDSReadByID = `
data "powerstore_replication_session" "by_id" {
  id = "%s"
}
`
