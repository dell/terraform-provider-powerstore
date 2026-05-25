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

	"github.com/bytedance/mockey"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

// Helper to create metro volume config for replication session action tests
func getMetroConfigForActionTests() string {
	volName := fmt.Sprintf("repl-action-test-vol-%d", time.Now().UnixNano())
	return fmt.Sprintf(`
resource "powerstore_volume" "test_vol" {
  name = "%s"
  size = 2.5
}

resource "powerstore_metro_volume" "test" {
  volume_id        = powerstore_volume.test_vol.id
  remote_system_id = "%s"
}
`, volName, remoteSystemID)
}

// Test Schema method for replication session action resource
func TestReplicationSessionActionResource_Schema(t *testing.T) {
	r := newReplicationSessionActionResource()
	resp := &fwresource.SchemaResponse{}
	r.Schema(context.Background(), fwresource.SchemaRequest{}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)

	assert.True(t, resp.Schema.Attributes["session_id"].IsRequired())
	assert.True(t, resp.Schema.Attributes["action"].IsRequired())
	assert.True(t, resp.Schema.Attributes["id"].IsComputed())
	assert.True(t, resp.Schema.Attributes["post_state"].IsComputed())
	assert.True(t, resp.Schema.Attributes["is_planned"].IsOptional())
	assert.True(t, resp.Schema.Attributes["reverse"].IsOptional())
}

// Test Metadata method
func TestReplicationSessionActionResource_Metadata(t *testing.T) {
	r := newReplicationSessionActionResource()
	resp := &fwresource.MetadataResponse{}
	r.Metadata(context.Background(), fwresource.MetadataRequest{ProviderTypeName: "powerstore"}, resp)
	assert.Equal(t, "powerstore_replication_session_action", resp.TypeName)
}

// Test Configure method with nil provider data
func TestReplicationSessionActionResource_Configure_Nil(t *testing.T) {
	r := &resourceReplicationSessionAction{}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: nil}, resp)
	assert.False(t, resp.Diagnostics.HasError())
}

// Test Configure method with wrong type
func TestReplicationSessionActionResource_Configure_WrongType(t *testing.T) {
	r := &resourceReplicationSessionAction{}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: "wrong"}, resp)
	assert.True(t, resp.Diagnostics.HasError())
}

// Test Configure method with correct type
func TestReplicationSessionActionResource_Configure_Correct(t *testing.T) {
	r := &resourceReplicationSessionAction{}
	resp := &fwresource.ConfigureResponse{}
	mockClient := &client.Client{GenClient: &clientgen.APIClient{}}
	r.Configure(context.Background(), fwresource.ConfigureRequest{ProviderData: mockClient}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

// Test Update method returns error
func TestReplicationSessionActionResource_Update(t *testing.T) {
	r := &resourceReplicationSessionAction{}
	resp := &fwresource.UpdateResponse{}
	r.Update(context.Background(), fwresource.UpdateRequest{}, resp)
	// Update is not supported for replication session action
	assert.True(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Diagnostics[0].Summary(), "Update not supported")
}

// Acceptance test: Missing session_id
func TestAccReplicationSessionAction_MissingSessionID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + ReplSessionActionMissingSessionID,
				ExpectError: regexp.MustCompile(`The argument "session_id" is required`),
			},
		},
	})
}

// Acceptance test: Missing action
func TestAccReplicationSessionAction_MissingAction(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionMissingAction, replicationSessionID),
				ExpectError: regexp.MustCompile(`The argument "action" is required`),
			},
		},
	})
}

// Acceptance test: Invalid action value
func TestAccReplicationSessionAction_InvalidAction(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionInvalidAction, replicationSessionID),
				ExpectError: regexp.MustCompile(`value must be one of`),
			},
		},
	})
}

// Acceptance test: Empty session_id
func TestAccReplicationSessionAction_EmptySessionID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + ReplSessionActionEmptySessionID,
				ExpectError: regexp.MustCompile(`string length must be at least 1`),
			},
		},
	})
}

// Acceptance test: Pause action
func TestAccReplicationSessionAction_PauseOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	metroConfig := getMetroConfigForActionTests()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create metro volume (creates replication session)
			{
				Config: ProviderConfigForTesting + metroConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_metro_volume.test", "metro_replication_session_id"),
				),
			},
			// Step 2: Pause the session
			{
				Config: ProviderConfigForTesting + metroConfig + `
resource "powerstore_replication_session_action" "test_pause" {
  session_id = powerstore_metro_volume.test.metro_replication_session_id
  action = "pause"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_pause", "action", "pause"),
					resource.TestCheckResourceAttrSet("powerstore_replication_session_action.test_pause", "session_id"),
				),
			},
		},
	})
}

// Acceptance test: Sync action
func TestAccReplicationSessionAction_SyncOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if endpoint != "http://localhost:3003/api/rest/" {
		t.Skip("Skipping on real server - sync is not supported for synchronous (metro) replication sessions")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionSync, replicationSessionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_sync", "action", "sync"),
				),
			},
		},
	})
}

// Acceptance test: Failover action
func TestAccReplicationSessionAction_FailoverOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if endpoint != "http://localhost:3003/api/rest/" {
		t.Skip("Skipping on real server - operation not supported for metro sessions")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionFailover, replicationSessionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_failover", "action", "failover"),
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_failover", "is_planned", "true"),
				),
			},
		},
	})
}

// Acceptance test: Resume action
func TestAccReplicationSessionAction_ResumeOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	// Generate unique name for this test
	volName := fmt.Sprintf("resume-action-test-%d", time.Now().UnixNano())
	metroConfig := fmt.Sprintf(`
resource "powerstore_volume" "test_vol" {
  name = "%s"
  size = 2.5
}
resource "powerstore_metro_volume" "test" {
  volume_id        = powerstore_volume.test_vol.id
  remote_system_id = "%s"
}
`, volName, remoteSystemID)
	metroConfigPaused := fmt.Sprintf(`
resource "powerstore_volume" "test_vol" {
  name = "%s"
  size = 2.5
}
resource "powerstore_metro_volume" "test" {
  volume_id             = powerstore_volume.test_vol.id
  remote_system_id      = "%s"
  is_replication_paused = true
}
`, volName, remoteSystemID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create metro volume
			{
				Config: ProviderConfigForTesting + metroConfig,
				Check:  resource.TestCheckResourceAttrSet("powerstore_metro_volume.test", "metro_replication_session_id"),
			},
			// Step 2: Pause the session
			{
				Config: ProviderConfigForTesting + metroConfigPaused,
				Check:  resource.TestCheckResourceAttr("powerstore_metro_volume.test", "is_replication_paused", "true"),
			},
			// Step 3: Resume via action resource
			{
				Config: ProviderConfigForTesting + metroConfigPaused + `
resource "powerstore_replication_session_action" "test_resume" {
  session_id = powerstore_metro_volume.test.metro_replication_session_id
  action = "resume"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_resume", "action", "resume"),
				),
			},
		},
	})
}

// Acceptance test: Reprotect action
func TestAccReplicationSessionAction_ReprotectOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if endpoint != "http://localhost:3003/api/rest/" {
		t.Skip("Skipping on real server - requires session in FAILED OVER state")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionReprotect, replicationSessionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_reprotect", "action", "reprotect"),
				),
			},
		},
	})
}

// Acceptance test: Start failover test
func TestAccReplicationSessionAction_StartFailoverTestOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if endpoint != "http://localhost:3003/api/rest/" {
		t.Skip("Skipping on real server - must be run from destination system")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionStartFailoverTest, replicationSessionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test_start_ft", "action", "start_failover_test"),
				),
			},
		},
	})
}

// Test mocked error paths for replication session action create
func TestAccReplicationSessionAction_CreateErrors(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if endpoint != "http://localhost:3003/api/rest/" {
		t.Skip("Skipping on real server - mocked error test")
	}

	var mocker1 *mockey.Mocker
	defer func() {
		if mocker1 != nil {
			mocker1.UnPatch()
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Create error - sync action fails
			{
				PreConfig: func() {
					mocker1 = mockey.Mock((*clientgen.ReplicationSessionApiService).ReplicationSessionSyncExecute).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionSync, replicationSessionID),
				ExpectError: regexp.MustCompile(`.*Error performing replication session action.*`),
			},
			// Create error - post-action read fails
			{
				PreConfig: func() {
					mocker1.UnPatch()
					mocker1 = mockey.Mock((*clientgen.ReplicationSessionApiService).GetReplicationSessionByIdExecute).Return(nil, nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionPause, replicationSessionID),
				ExpectError: regexp.MustCompile(`.*Error reading replication session after action.*`),
			},
		},
	})
}

// HCL config strings
var ReplSessionActionMissingSessionID = `
resource "powerstore_replication_session_action" "test" {
  action = "pause"
}
`

var ReplSessionActionMissingAction = `
resource "powerstore_replication_session_action" "test" {
  session_id = "%s"
}
`

var ReplSessionActionInvalidAction = `
resource "powerstore_replication_session_action" "test" {
  session_id = "%s"
  action     = "invalid_action"
}
`

var ReplSessionActionEmptySessionID = `
resource "powerstore_replication_session_action" "test" {
  session_id = ""
  action     = "pause"
}
`

var ReplSessionActionPause = `
resource "powerstore_replication_session_action" "test" {
  session_id = "%s"
  action     = "pause"
}
`

var ReplSessionActionSync = `
resource "powerstore_replication_session_action" "test_sync" {
  session_id = "%s"
  action     = "sync"
}
`

var ReplSessionActionFailover = `
resource "powerstore_replication_session_action" "test_failover" {
  session_id = "%s"
  action     = "failover"
  is_planned = true
  reverse    = false
}
`

var ReplSessionActionResume = `
resource "powerstore_replication_session_action" "test_resume" {
  session_id = "%s"
  action     = "resume"
}
`

var ReplSessionActionReprotect = `
resource "powerstore_replication_session_action" "test_reprotect" {
  session_id = "%s"
  action     = "reprotect"
}
`

var ReplSessionActionStartFailoverTest = `
resource "powerstore_replication_session_action" "test_start_ft" {
  session_id = "%s"
  action     = "start_failover_test"
}
`
