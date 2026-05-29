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
	"net/url"
	"os"
	"regexp"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

var remoteEndpoint = setDefault(os.Getenv("POWERSTORE_REMOTE_ENDPOINT"), "https://10.230.45.71/api/rest/")
var remoteUsername = setDefault(os.Getenv("POWERSTORE_REMOTE_USERNAME"), "admin")
var remotePassword = setDefault(os.Getenv("POWERSTORE_REMOTE_PASSWORD"), "Password123!")

// remoteManagementAddress extracts the host from the remote endpoint for creating remote systems
var remoteManagementAddress = func() string {
	addr := setDefault(os.Getenv("POWERSTORE_REMOTE_ENDPOINT"), "10.230.45.71")
	// Strip protocol and path if present
	for _, prefix := range []string{"https://", "http://"} {
		if len(addr) > len(prefix) && addr[:len(prefix)] == prefix {
			addr = addr[len(prefix):]
		}
	}
	// Strip trailing path
	for i, c := range addr {
		if c == '/' {
			addr = addr[:i]
			break
		}
	}
	return addr
}()

// TestAccRemoteSystem_CreateOnMock - Create PowerStore remote system
func TestAccRemoteSystem_CreateOnMock(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "management_address", remoteManagementAddress),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "description", "Terraform acceptance test remote system"),
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "state"),
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "serial_number"),
				),
			},
			// Import
			{
				Config:       ProviderConfigForTesting + RemoteSystemCreateConfig,
				ResourceName: "powerstore_remote_system.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, remoteManagementAddress, s[0].Attributes["management_address"])
					assert.NotEmpty(t, s[0].Attributes["id"])
					return nil
				},
			},
			// Update description
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccRemoteSystem_MissingAddress - Create without management_address (negative)
func TestAccRemoteSystem_MissingAddress(t *testing.T) {
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

// TestAccRemoteSystem_EmptyAddress - Create with empty management_address (negative)
func TestAccRemoteSystem_EmptyAddress(t *testing.T) {
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

// TestAccRemoteSystem_InvalidAddress - Create with invalid management_address (negative)
func TestAccRemoteSystem_InvalidAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RemoteSystemInvalidAddress,
				ExpectError: regexp.MustCompile(`.*Error creating remote system.*`),
			},
		},
	})
}

// TestAccRemoteSystem_InvalidLatency - Create with invalid data_network_latency (negative)
func TestAccRemoteSystem_InvalidLatency(t *testing.T) {
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

// TestAccRemoteSystem_ImportOnReal - Import existing remote system and verify attributes on real array
func TestAccRemoteSystem_ImportOnReal(t *testing.T) {
	if os.Getenv("POWERSTORE_ENDPOINT") == "" || os.Getenv("POWERSTORE_ENDPOINT") == "http://localhost:3003/api/rest" {
		t.Skip("Skipping real array test - set POWERSTORE_ENDPOINT to a real PowerStore array")
	}
	existingID := os.Getenv("TF_VAR_remote_system_id")
	if existingID == "" {
		t.Skip("Skipping real array test - TF_VAR_remote_system_id not set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:            ProviderConfigForTesting + RemoteSystemImportConfig,
				ResourceName:      "powerstore_remote_system.real_test",
				ImportState:       true,
				ImportStateId:     existingID,
				ImportStateVerify: false,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, existingID, s[0].Attributes["id"])
					assert.Equal(t, remoteManagementAddress, s[0].Attributes["management_address"])
					assert.NotEmpty(t, s[0].Attributes["state"])
					assert.NotEmpty(t, s[0].Attributes["serial_number"])
					assert.Equal(t, "PowerStore", s[0].Attributes["type"])
					return nil
				},
			},
		},
	})
}

// TestAccRemoteSystem_UpdateOnReal - Update description of existing remote system on real array
func TestAccRemoteSystem_UpdateOnReal(t *testing.T) {
	if os.Getenv("POWERSTORE_ENDPOINT") == "" || os.Getenv("POWERSTORE_ENDPOINT") == "http://localhost:3003/api/rest" {
		t.Skip("Skipping real array test - set POWERSTORE_ENDPOINT to a real PowerStore array")
	}
	existingID := os.Getenv("TF_VAR_remote_system_id")
	if existingID == "" {
		t.Skip("Skipping real array test - TF_VAR_remote_system_id not set")
	}

	// Use the API directly to verify update works
	t.Run("update_description_via_api", func(t *testing.T) {
		ctx := context.Background()
		pstoreClient, err := client.NewClient(
			endpoint, username, password, true, 120,
		)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Update description
		desc := "Updated by Terraform AT"
		modifyBody := clientgen.RemoteSystemModify{
			Description: &desc,
		}
		_, err = pstoreClient.GenClient.RemoteSystemApi.PatchRemoteSystemById(ctx, existingID).Body(modifyBody).Execute()
		assert.NoError(t, err, "Failed to update remote system description")

		// Read back and verify
		sel := "*"
		queries := make(url.Values)
		queries.Set("select", sel)
		rs, _, err := pstoreClient.GenClient.RemoteSystemApi.GetRemoteSystemById(ctx, existingID).Queries(queries).Execute()
		assert.NoError(t, err, "Failed to read remote system")
		assert.Equal(t, desc, *rs.Description)

		// Restore description
		restore := ""
		restoreBody := clientgen.RemoteSystemModify{
			Description: &restore,
		}
		_, err = pstoreClient.GenClient.RemoteSystemApi.PatchRemoteSystemById(ctx, existingID).Body(restoreBody).Execute()
		assert.NoError(t, err, "Failed to restore remote system description")
	})
}

// Test configurations

var RemoteSystemCreateConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = "Terraform acceptance test remote system"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = "Updated description"
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

var RemoteSystemInvalidAddress = `
resource "powerstore_remote_system" "test" {
	management_address   = "invalid-address-that-does-not-exist"
	data_network_latency = "Low"
}
`

var RemoteSystemInvalidLatency = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	data_network_latency = "Invalid"
}
`

var RemoteSystemImportConfig = `
resource "powerstore_remote_system" "real_test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = ""
	data_network_latency = "Low"
}
`

var RemoteSystemRealUpdateConfig = `
resource "powerstore_remote_system" "real_test" {
	management_address   = "` + remoteManagementAddress + `"
	description          = "Updated by Terraform AT"
	data_network_latency = "Low"
}
`
