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
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"

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

// remoteExchangeUsername is used for certificate exchange
var remoteExchangeUsername = setDefault(os.Getenv("POWERSTORE_REMOTE_USERNAME"), "admin")

// remoteExchangePassword is used for certificate exchange
var remoteExchangePassword = setDefault(os.Getenv("POWERSTORE_REMOTE_PASSWORD"), "Password123!")

// deleteExistingRemoteSystem deletes any existing remote system with the given management address
func deleteExistingRemoteSystem(t *testing.T) {
	if isMockServer() {
		return // No cleanup needed for mock server
	}
	c, err := getClientForTest()
	if err != nil {
		t.Logf("Warning: Could not create client for cleanup: %v", err)
		return
	}
	// Find remote system by management address using GenClient
	queryParams := url.Values{}
	queryParams.Set("select", "id,management_address")
	remoteSystems, _, err := c.GenClient.RemoteSystemApi.GetAllRemoteSystems(context.Background()).Queries(queryParams).Execute()
	if err != nil {
		t.Logf("Warning: Could not list remote systems for cleanup: %v", err)
		return
	}
	for _, rs := range remoteSystems {
		if rs.ManagementAddress != nil && *rs.ManagementAddress == remoteManagementAddress {
			deleteBody := map[string]interface{}{}
			_, err := c.GenClient.RemoteSystemApi.DeleteRemoteSystemById(context.Background(), *rs.Id).Body(deleteBody).Execute()
			if err != nil {
				t.Logf("Warning: Could not delete existing remote system %s: %v", *rs.Id, err)
			} else {
				t.Logf("Deleted existing remote system %s (%s) for clean test state", *rs.Id, *rs.ManagementAddress)
			}
		}
	}
}

// restoreRemoteSystem recreates the remote system after tests complete
func restoreRemoteSystem(t *testing.T) {
	if isMockServer() {
		return // No restore needed for mock server
	}
	c, err := getClientForTest()
	if err != nil {
		t.Logf("Warning: Could not create client for restore: %v", err)
		return
	}
	// Check if remote system already exists using GenClient
	queryParams := url.Values{}
	queryParams.Set("select", "id,management_address")
	remoteSystems, _, err := c.GenClient.RemoteSystemApi.GetAllRemoteSystems(context.Background()).Queries(queryParams).Execute()
	if err != nil {
		t.Logf("Warning: Could not list remote systems for restore check: %v", err)
		return
	}
	for _, rs := range remoteSystems {
		if rs.ManagementAddress != nil && *rs.ManagementAddress == remoteManagementAddress {
			t.Logf("Remote system %s already exists, no restore needed", *rs.ManagementAddress)
			return
		}
	}
	// Perform certificate exchange first using GenClient
	exchangeBody := clientgen.X509CertificateExchange{
		Service:  clientgen.X509CertificateServiceEnum("Replication_HTTP"),
		Address:  remoteManagementAddress,
		Port:     443,
		Username: remoteExchangeUsername,
		Password: remoteExchangePassword,
	}
	_, err = c.GenClient.X509CertificateApi.PostX509CertificateById(context.Background()).Body(exchangeBody).Execute()
	if err != nil {
		t.Logf("Warning: Certificate exchange failed during restore: %v", err)
	}
	// Create remote system using GenClient
	latency := clientgen.RemoteSystemLatencyEnum("Low")
	createBody := clientgen.RemoteSystemCreate{
		ManagementAddress:  &remoteManagementAddress,
		DataNetworkLatency: &latency,
	}
	_, _, err = c.GenClient.RemoteSystemApi.PostAllRemoteSystems(context.Background()).Body(createBody).Execute()
	if err != nil {
		t.Logf("Warning: Could not restore remote system: %v", err)
	} else {
		t.Logf("Restored remote system %s after tests", remoteManagementAddress)
	}
}

// getClientForTest creates a PowerStore client for test cleanup/restore operations
func getClientForTest() (*client.Client, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint not set")
	}
	return client.NewClient(endpoint, username, password, true, 120)
}

// --- Acceptance Tests ---

// TestAccRemoteSystemResource_CRUD covers Create, Read, Update (description + latency), ImportState, and Delete.
func TestAccRemoteSystemResource_CRUD(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })

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

// TestAccRemoteSystemResource_UpdateName covers the name field update path in Update.
// Note: PowerStore-to-PowerStore remote systems auto-generate name from remote system - custom name not supported.
func TestAccRemoteSystemResource_UpdateName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - PowerStore remote systems auto-generate name, custom name not supported")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })

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

// TestAccRemoteSystemResource_CreateWithCredentials covers remote_username and remote_password in Create.
func TestAccRemoteSystemResource_CreateWithCredentials(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })

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

// TestAccRemoteSystemResource_UpdateCredentials covers remote_username and remote_password in Update.
func TestAccRemoteSystemResource_UpdateCredentials(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })

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

// TestAccRemoteSystemResource_MissingAddress validates that management_address is required.
func TestAccRemoteSystemResource_MissingAddress(t *testing.T) {
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

// TestAccRemoteSystemResource_EmptyAddress validates that management_address cannot be empty.
func TestAccRemoteSystemResource_EmptyAddress(t *testing.T) {
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

// TestAccRemoteSystemResource_InvalidLatency validates that invalid latency values are rejected.
func TestAccRemoteSystemResource_InvalidLatency(t *testing.T) {
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

// TestAccRemoteSystemResource_CreateWithType covers Type field in Create.
func TestAccRemoteSystemResource_CreateWithType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateWithTypeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "type", "PowerStore"),
				),
			},
		},
	})
}

// TestAccRemoteSystemResource_UpdateType covers Type field update path in Update.
// Note: Mock-only test - "NonPowerStore" type is not a valid enum value on real server.
func TestAccRemoteSystemResource_UpdateType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - NonPowerStore type is not a valid enum value")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateWithTypeConfig,
			},
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateTypeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "type", "NonPowerStore"),
				),
			},
		},
	})
}

// TestAccRemoteSystemResource_UpdateAddress covers management_address update path in Update.
func TestAccRemoteSystemResource_UpdateAddress(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - address update requires different remote array")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateConfig,
			},
			{
				Config: ProviderConfigForTesting + RemoteSystemUpdateAddressConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "management_address", "10.230.24.10"),
				),
			},
		},
	})
}

// TestAccRemoteSystemResource_CreateWithDataConnectionType covers DataConnectionType field in Create.
// Note: data_connection_type is not valid for PowerStore-to-PowerStore remote systems.
func TestAccRemoteSystemResource_CreateWithDataConnectionType(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - data_connection_type not valid for PowerStore remote systems")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateWithDataConnectionTypeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "data_connection_type", "iSCSI"),
				),
			},
		},
	})
}

// TestAccRemoteSystemResource_CreateWithIscsiAddresses covers IscsiAddresses field in Create.
// Note: iscsi_addresses is not valid for PowerStore-to-PowerStore remote systems.
func TestAccRemoteSystemResource_CreateWithIscsiAddresses(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	if !isMockServer() {
		t.Skip("Skipping on real server - iscsi_addresses not valid for PowerStore remote systems")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateWithIscsiAddressesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "iscsi_addresses.0", "192.168.1.1"),
					resource.TestCheckResourceAttr("powerstore_remote_system.test", "iscsi_addresses.1", "192.168.1.2"),
				),
			},
		},
	})
}

// TestAccRemoteSystemResource_NoChangeUpdate covers no-op update (no fields changed).
func TestAccRemoteSystemResource_NoChangeUpdate(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	deleteExistingRemoteSystem(t)
	t.Cleanup(func() { restoreRemoteSystem(t) })
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateConfig,
			},
			{
				Config: ProviderConfigForTesting + RemoteSystemCreateConfig, // Same config - no changes
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_remote_system.test", "id"),
				),
			},
		},
	})
}

// --- Test Config Strings ---

var RemoteSystemCreateConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	description          = "Terraform acceptance test remote system"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateDescConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	description          = "Updated description"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateLatencyConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	description          = "Updated description"
	data_network_latency = "High"
}
`

var RemoteSystemUpdateNameConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	name                 = "Updated Name"
	description          = "Updated with name change"
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
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	remote_username      = "admin"
	remote_password      = "Password123!"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateCredentialsConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	remote_username      = "newuser"
	remote_password      = "NewPassword123!"
	description          = "Updated with credentials change"
	data_network_latency = "Low"
}
`

var RemoteSystemCreateWithTypeConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	type                 = "PowerStore"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateTypeConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	type                 = "NonPowerStore"
	data_network_latency = "Low"
}
`

var RemoteSystemUpdateAddressConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "10.230.24.10"
	description          = "Terraform acceptance test remote system"
	data_network_latency = "Low"
}
`

var RemoteSystemCreateWithDataConnectionTypeConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	data_connection_type = "iSCSI"
	data_network_latency = "Low"
}
`

var RemoteSystemCreateWithIscsiAddressesConfig = `
resource "powerstore_remote_system" "test" {
	management_address   = "` + remoteManagementAddress + `"
	exchange_username    = "` + remoteExchangeUsername + `"
	exchange_password    = "` + remoteExchangePassword + `"
	iscsi_addresses      = ["192.168.1.1", "192.168.1.2"]
	data_network_latency = "Low"
}
`
