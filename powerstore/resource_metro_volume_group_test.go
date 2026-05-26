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
	"time"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	"github.com/bytedance/mockey"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Helper function to generate unique volume group name
func getMetroVolumeGroupName() string {
	if endpoint == "http://localhost:3003/api/rest/" {
		return "tf_volume_group_new" // Use mock server's expected name
	}
	return fmt.Sprintf("metro_test_vg-%d", time.Now().UnixNano()) // Use dynamic name for real server
}

// Helper function to generate volume names for volume group tests
func getMetroVolumeNamesForVGroup() (string, string) {
	if endpoint == "http://localhost:3003/api/rest/" {
		return "test_acc_cvol", "test_acc_cvol" // Use mock server's expected names (both same)
	}
	vgName := getMetroVolumeGroupName()
	return fmt.Sprintf("%s-vol1", vgName), fmt.Sprintf("%s-vol2", vgName)
}

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
	assert.True(t, resp.Schema.Attributes["is_replication_paused"].IsOptional())
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

	// Generate unique config once for this test
	vgName := getMetroVolumeGroupName()
	vol1Name, vol2Name := getMetroVolumeNamesForVGroup()
	config := fmt.Sprintf(`
resource "powerstore_volume" "vol1" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volume" "vol2" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volumegroup" "test" {
  name        = "%s"
  description = "Creating Volume Group"
  volume_ids = [powerstore_volume.vol1.id, powerstore_volume.vol2.id]
}

resource "powerstore_metro_volume_group" "test" {
  volume_group_id  = powerstore_volumegroup.test.id
  remote_system_id = "%s"
}
`, vol1Name, vol2Name, vgName, remoteSystemID)
	configPaused := fmt.Sprintf(`
resource "powerstore_volume" "vol1" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volume" "vol2" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volumegroup" "test" {
  name        = "%s"
  description = "Creating Volume Group"
  volume_ids = [powerstore_volume.vol1.id, powerstore_volume.vol2.id]
}

resource "powerstore_metro_volume_group" "test" {
  volume_group_id      = powerstore_volumegroup.test.id
  remote_system_id     = "%s"
  is_replication_paused = true
}
`, vol1Name, vol2Name, vgName, remoteSystemID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_metro_volume_group.test", "id"),
					resource.TestCheckResourceAttrSet("powerstore_metro_volume_group.test", "metro_replication_session_id"),
					resource.TestCheckResourceAttrSet("powerstore_metro_volume_group.test", "remote_system_id"),
				),
			},
			// Import test
			{
				Config:            ProviderConfigForTesting + config,
				ResourceName:      "powerstore_metro_volume_group.test",
				ImportState:       true,
				ImportStateVerify: false,
			},
			// Update test - pause replication
			{
				Config: ProviderConfigForTesting + configPaused,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_metro_volume_group.test", "is_replication_paused", "true"),
				),
			},
			// Update test - resume replication
			{
				Config: ProviderConfigForTesting + config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_metro_volume_group.test", "is_replication_paused", "false"),
				),
			},
		},
	})
}

// Test mocked error paths for metro volume group create
func TestAccMetroVolumeGroup_CreateErrors(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	var mocker1, mocker2 *mockey.Mocker
	defer func() {
		if mocker1 != nil {
			mocker1.UnPatch()
		}
		if mocker2 != nil {
			mocker2.UnPatch()
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Create error - configure metro API fails
			{
				PreConfig: func() {
					mocker1 = mockey.Mock((*clientgen.VolumeGroupApiService).VolumeGroupConfigureMetroExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + MetroVolumeGroupParamsCreate,
				ExpectError: regexp.MustCompile(`.*Error configuring metro volume group.*`),
			},
			// Create error - empty session ID returned
			{
				PreConfig: func() {
					mocker1.UnPatch()
					mocker1 = mockey.Mock((*clientgen.VolumeGroupApiService).VolumeGroupConfigureMetroExecute).Return(&clientgen.VolumeGroupConfigureMetroResponse{}, nil, nil).Build()
				},
				Config:      ProviderConfigForTesting + MetroVolumeGroupParamsCreate,
				ExpectError: regexp.MustCompile(`.*no session ID was returned.*`),
			},
			// Create error - read session after create fails
			{
				PreConfig: func() {
					mocker1.UnPatch()
					sid := "mock-session-id"
					mocker1 = mockey.Mock((*clientgen.VolumeGroupApiService).VolumeGroupConfigureMetroExecute).Return(&clientgen.VolumeGroupConfigureMetroResponse{MetroReplicationSessionId: &sid}, nil, nil).Build()
					mocker2 = mockey.Mock((*clientgen.ReplicationSessionApiService).GetReplicationSessionByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + MetroVolumeGroupParamsCreate,
				ExpectError: regexp.MustCompile(`.*Error reading replication session after metro configuration.*`),
			},
		},
	})
}

// Test mocked error paths for metro volume group read, update, and delete
func TestAccMetroVolumeGroup_ReadUpdateDeleteErrors(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	var mocker1 *mockey.Mocker
	defer func() {
		if mocker1 != nil {
			mocker1.UnPatch()
		}
	}()

	// Generate unique config once for this test
	vgName := getMetroVolumeGroupName()
	vol1Name, vol2Name := getMetroVolumeNamesForVGroup()
	config := fmt.Sprintf(`
resource "powerstore_volume" "vol1" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volume" "vol2" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volumegroup" "test" {
  name        = "%s"
  description = "Creating Volume Group"
  volume_ids = [powerstore_volume.vol1.id, powerstore_volume.vol2.id]
}

resource "powerstore_metro_volume_group" "test" {
  volume_group_id  = powerstore_volumegroup.test.id
  remote_system_id = "%s"
}
`, vol1Name, vol2Name, vgName, remoteSystemID)
	configPaused := fmt.Sprintf(`
resource "powerstore_volume" "vol1" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volume" "vol2" {
  name = "%s"
  size = 2.5
}

resource "powerstore_volumegroup" "test" {
  name        = "%s"
  description = "Creating Volume Group"
  volume_ids = [powerstore_volume.vol1.id, powerstore_volume.vol2.id]
}

resource "powerstore_metro_volume_group" "test" {
  volume_group_id      = powerstore_volumegroup.test.id
  remote_system_id     = "%s"
  is_replication_paused = true
}
`, vol1Name, vol2Name, vgName, remoteSystemID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create successfully
			{
				Config: ProviderConfigForTesting + config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_metro_volume_group.test", "id"),
				),
			},
			// Step 2: Read error on refresh
			{
				PreConfig: func() {
					mocker1 = mockey.Mock((*clientgen.ReplicationSessionApiService).GetReplicationSessionByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + config,
				ExpectError: regexp.MustCompile(`.*Error reading metro volume group replication session.*`),
			},
			// Step 3: Update (pause) error
			{
				PreConfig: func() {
					mocker1.UnPatch()
					mocker1 = mockey.Mock((*clientgen.ReplicationSessionApiService).ReplicationSessionPauseExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + configPaused,
				ExpectError: regexp.MustCompile(`.*Error pausing metro volume group replication.*`),
			},
			// Step 4: Delete error
			{
				PreConfig: func() {
					mocker1.UnPatch()
					mocker1 = mockey.Mock((*clientgen.VolumeGroupApiService).VolumeGroupEndMetroExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + config,
				Destroy:     true,
				ExpectError: regexp.MustCompile(`.*Error ending metro volume group configuration.*`),
			},
			// Step 5: Cleanup - unpatch so destroy succeeds
			{
				PreConfig: func() {
					mocker1.UnPatch()
					mocker1 = nil
				},
				Config: ProviderConfigForTesting + config,
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

var MetroVolumeGroupParamsCreate = fmt.Sprintf(`
resource "powerstore_volumegroup" "test" {
  name        = "tf_volume_group_new"
  description = "Creating Volume Group"
}

resource "powerstore_metro_volume_group" "test" {
  volume_group_id  = powerstore_volumegroup.test.id
  remote_system_id = "%s"
}
`, remoteSystemID)

var MetroVolumeGroupParamsCreatePaused = fmt.Sprintf(`
resource "powerstore_volumegroup" "test" {
  name        = "tf_volume_group_new"
  description = "Creating Volume Group"
}

resource "powerstore_metro_volume_group" "test" {
  volume_group_id      = powerstore_volumegroup.test.id
  remote_system_id     = "%s"
  is_replication_paused = true
}
`, remoteSystemID)
