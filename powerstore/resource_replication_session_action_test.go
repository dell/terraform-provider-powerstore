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

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionPause, replicationSessionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test", "action", "pause"),
					resource.TestCheckResourceAttr("powerstore_replication_session_action.test", "session_id", replicationSessionID),
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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(ReplSessionActionResume, replicationSessionID),
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

// Acceptance test: Stop failover test on mock server
// Skipped due to mock server body-parser limitation with POST responses
func TestAccReplicationSessionAction_StopFailoverTestOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	t.Skip("Mock server limitation: body-parser cannot handle POST responses with JSON body")
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

var ReplSessionActionStopFailoverTestMock = `
resource "powerstore_replication_session_action" "test_stop_ft" {
  session_id = "b5f699ec-45a2-4bac-8153-3d520bffa861"
  action     = "stop_failover_test"
}
`
