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
	"strings"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"testing"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// isMockServer returns true when tests run against the local mock server
func isMockServer() bool {
	return strings.Contains(endpoint, "localhost") || strings.Contains(endpoint, "127.0.0.1")
}

// remoteManagementAddress extracts the host from POWERSTORE_REMOTE_ENDPOINT or uses the default
var remoteManagementAddress = func() string {
	addr := setDefault(os.Getenv("POWERSTORE_REMOTE_ENDPOINT"), "10.230.45.71")
	for _, prefix := range []string{"https://", "http://"} {
		if len(addr) > len(prefix) && addr[:len(prefix)] == prefix {
			addr = addr[len(prefix):]
		}
	}
	for i, c := range addr {
		if c == '/' {
			addr = addr[:i]
			break
		}
	}
	return addr
}()

// --- Acceptance Tests ---

// TestAccRemoteSystem_CRUD covers Create, Read, Update (description + latency), ImportState, and Delete.
func TestAccRemoteSystem_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping CRUD on real server - remote system to 10.230.45.71 already exists on 10.230.24.184 with active replication sessions. Use mock server or set POWERSTORE_REMOTE_ENDPOINT to a different array without active sessions.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "management_address", remoteManagementAddress),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "description", "Terraform acceptance test remote system"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "data_network_latency", "Low"),
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "state"),
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "serial_number"),
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "version"),
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "type"),
				),
			},
			// Step 2: Import and verify state
			{
				Config:       ProviderConfigForTesting + RemoteSystemCreateConfig,
				ResourceName: "powerstore_remote_system.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, remoteManagementAddress, s[0].Attributes["management_address"])
					assert.NotEmpty(t, s[0].Attributes["id"])
					assert.NotEmpty(t, s[0].Attributes["serial_number"])
					assert.NotEmpty(t, s[0].Attributes["state"])
					return nil
				},
			},
			// Step 3: Update description
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateDescConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "description", "Updated description"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "data_network_latency", "Low"),
				),
			},
			// Step 4: Update data_network_latency
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateLatencyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "data_network_latency", "High"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccRemoteSystem_UpdateName covers the name field update path in Update.
func TestAccRemoteSystem_UpdateName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - requires creating a remote system. Use mock server or set POWERSTORE_REMOTE_ENDPOINT to a different array without active sessions.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateConfig,
			},
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateNameConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "name", "Updated Name"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "data_network_latency", "Low"),
				),
			},
		},
	})
}

// TestAccRemoteSystem_CreateWithCredentials covers remote_username and remote_password in Create.
func TestAccRemoteSystem_CreateWithCredentials(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - requires creating a remote system with credentials. Use mock server or set POWERSTORE_REMOTE_ENDPOINT to a different array without active sessions.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateWithCredentialsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "management_address", remoteManagementAddress),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "remote_username", "admin"),
				),
			},
		},
	})
}

// TestAccRemoteSystem_UpdateCredentials covers remote_username and remote_password in Update.
func TestAccRemoteSystem_UpdateCredentials(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - requires creating a remote system with credentials. Use mock server or set POWERSTORE_REMOTE_ENDPOINT to a different array without active sessions.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateWithCredentialsConfig,
			},
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateCredentialsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "remote_username", "newuser"),
				),
			},
		},
	})
}

// TestAccRemoteSystem_MissingAddress validates that management_address is required.
func TestAccRemoteSystem_MissingAddress(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RemoteSystemMissingAddress,
				ExpectError: regexp.MustCompile(`.*management_address.*`),
			},
		},
	})
}

// TestAccRemoteSystem_EmptyAddress validates that management_address cannot be empty.
func TestAccRemoteSystem_EmptyAddress(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RemoteSystemEmptyAddress,
				ExpectError: regexp.MustCompile(`.*Attribute management_address string length must be at least 1.*`),
			},
		},
	})
}

// TestAccRemoteSystem_InvalidLatency validates that invalid latency values are rejected.
func TestAccRemoteSystem_InvalidLatency(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RemoteSystemInvalidLatency,
				ExpectError: regexp.MustCompile(`.*value must be one of.*`),
			},
		},
	})
}

// --- Unit Tests ---

func TestResourceRemoteSystem_Configure_InvalidType(t *testing.T) {
	r := &resourceRemoteSystem{}
	req := fwresource.ConfigureRequest{ProviderData: "invalid_type"}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), req, resp)
	assert.True(t, resp.Diagnostics.HasError())
	assert.NotEmpty(t, resp.Diagnostics.Errors()[0].Summary)
}

func TestResourceRemoteSystem_Configure_Nil(t *testing.T) {
	r := &resourceRemoteSystem{}
	req := fwresource.ConfigureRequest{ProviderData: nil}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, r.client)
}

func TestResourceRemoteSystem_Configure_Success(t *testing.T) {
	r := &resourceRemoteSystem{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwresource.ConfigureRequest{ProviderData: c}
	resp := &fwresource.ConfigureResponse{}
	r.Configure(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

func TestResourceRemoteSystem_Metadata(t *testing.T) {
	r := &resourceRemoteSystem{}
	req := fwresource.MetadataRequest{}
	resp := &fwresource.MetadataResponse{}
	r.Metadata(context.Background(), req, resp)
	assert.Contains(t, resp.TypeName, "remote_system")
}

func TestResourceRemoteSystem_Schema(t *testing.T) {
	r := &resourceRemoteSystem{}
	req := fwresource.SchemaRequest{}
	resp := &fwresource.SchemaResponse{}
	r.Schema(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, resp.Schema)
	_, ok := resp.Schema.Attributes["management_address"]
	assert.True(t, ok, "management_address attribute should exist")
	_, ok = resp.Schema.Attributes["id"]
	assert.True(t, ok, "id attribute should exist")
	_, ok = resp.Schema.Attributes["state"]
	assert.True(t, ok, "state attribute should exist")
	_, ok = resp.Schema.Attributes["capabilities"]
	assert.True(t, ok, "capabilities attribute should exist")
	_, ok = resp.Schema.Attributes["serial_number"]
	assert.True(t, ok, "serial_number attribute should exist")
	_, ok = resp.Schema.Attributes["data_network_latency"]
	assert.True(t, ok, "data_network_latency attribute should exist")
}

func TestRemoteSystem_UpdateState_NilInstance(t *testing.T) {
	r := &resourceRemoteSystem{}
	state := &models.RemoteSystemResource{}
	r.updateState(context.Background(), state, nil)
	assert.Empty(t, state.ID.ValueString())
}

func TestRemoteSystem_UpdateState_WithAllFields(t *testing.T) {
	r := &resourceRemoteSystem{}
	state := &models.RemoteSystemResource{}

	id := "test-id"
	name := "test-name"
	desc := "test-desc"
	addr := "10.0.0.1"
	serial := "PS1234"
	version := "4.1.0.0"
	rsType := clientgen.RemoteSystemTypeEnum("PowerStore")
	rsState := clientgen.RemoteSystemStateEnum("OK")
	dataConnState := clientgen.DataConnectionStateEnum("OK")
	latency := clientgen.RemoteSystemLatencyEnum("Low")
	connType := clientgen.DataConnectionTypeEnum("iSCSI")

	rs := &clientgen.RemoteSystemInstance{
		Id:                  &id,
		Name:                &name,
		Description:         &desc,
		ManagementAddress:   &addr,
		SerialNumber:        &serial,
		Version:             &version,
		Type:                &rsType,
		State:               &rsState,
		DataConnectionState: &dataConnState,
		DataNetworkLatency:  &latency,
		DataConnectionType:  &connType,
		Capabilities:        []clientgen.RemoteProtectionCapabilityEnum{"Asynchronous_Block_Replication"},
	}

	r.updateState(context.Background(), state, rs)

	assert.Equal(t, "test-id", state.ID.ValueString())
	assert.Equal(t, "test-name", state.Name.ValueString())
	assert.Equal(t, "test-desc", state.Description.ValueString())
	assert.Equal(t, "10.0.0.1", state.ManagementAddress.ValueString())
	assert.Equal(t, "PS1234", state.SerialNumber.ValueString())
	assert.Equal(t, "4.1.0.0", state.Version.ValueString())
	assert.Equal(t, "PowerStore", state.Type.ValueString())
	assert.Equal(t, "OK", state.State.ValueString())
	assert.Equal(t, "OK", state.DataConnectionState.ValueString())
	assert.Equal(t, "Low", state.DataNetworkLatency.ValueString())
	assert.Equal(t, "iSCSI", state.DataConnectionType.ValueString())
	assert.Equal(t, 1, len(state.Capabilities.Elements()))
}

func TestRemoteSystem_UpdateState_WithNilOptionalFields(t *testing.T) {
	r := &resourceRemoteSystem{}
	state := &models.RemoteSystemResource{}

	id := "test-id"
	addr := "10.0.0.1"
	rs := &clientgen.RemoteSystemInstance{
		Id:                &id,
		ManagementAddress: &addr,
	}

	r.updateState(context.Background(), state, rs)

	assert.Equal(t, "test-id", state.ID.ValueString())
	assert.Equal(t, "10.0.0.1", state.ManagementAddress.ValueString())
	assert.Empty(t, state.Type.ValueString())
	assert.Empty(t, state.State.ValueString())
	assert.Empty(t, state.DataNetworkLatency.ValueString())
	assert.True(t, state.Capabilities.IsNull())
}

func TestRemoteSystem_UpdateState_EmptyCapabilities(t *testing.T) {
	r := &resourceRemoteSystem{}
	state := &models.RemoteSystemResource{}

	id := "test-id"
	rs := &clientgen.RemoteSystemInstance{
		Id:           &id,
		Capabilities: []clientgen.RemoteProtectionCapabilityEnum{},
	}

	r.updateState(context.Background(), state, rs)
	assert.False(t, state.Capabilities.IsNull())
	assert.Equal(t, 0, len(state.Capabilities.Elements()))
}

// --- Test Config Strings ---

var RemoteSystemCreateConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = "Terraform acceptance test remote system"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateDescConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = "Updated description"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateLatencyConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = "Updated description"
	data_network_latency = "High"
}
`

var RemoteSystemUpdateNameConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	name                 = "Updated Name"
	description          = "Terraform acceptance test remote system"
	data_network_latency = "Low"
}
`

var RemoteSystemMissingAddress = `
resource "powerstore_remote_system" "test" {
	description = "missing address"
}
`

var RemoteSystemEmptyAddress = `
resource "powerstore_remote_system" "test" {
	management_address = ""
}
`

var RemoteSystemInvalidLatency = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	data_network_latency = "Invalid"
}
`

var RemoteSystemCreateWithCredentialsConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	remote_username      = "admin"
	remote_password      = "Password123!"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateCredentialsConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	remote_username      = "newuser"
	remote_password      = "NewPassword123!"
	data_network_latency = "Low"
}
`
