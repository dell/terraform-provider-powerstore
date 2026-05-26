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
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Helper to create metro volume config for replication session datasource tests
func getMetroConfigForDSTests() string {
	var volName string
	if endpoint == "http://localhost:3003/api/rest/" {
		volName = "test_acc_cvol" // Use mock server's expected name
	} else {
		volName = fmt.Sprintf("repl-ds-test-vol-%d", time.Now().UnixNano()) // Use dynamic name for real server
	}
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

	metroConfig := getMetroConfigForDSTests()

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
			// Step 2: Read the session by ID
			{
				Config: ProviderConfigForTesting + metroConfig + `
data "powerstore_replication_session" "by_id" {
  id = powerstore_metro_volume.test.metro_replication_session_id
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.powerstore_replication_session.by_id", "id"),
					resource.TestCheckResourceAttr("data.powerstore_replication_session.by_id", "replication_sessions.#", "1"),
				),
			},
		},
	})
}

// Acceptance test: Read replication session by invalid ID
func TestAccReplicationSessionDataSource_ReadByInvalidID(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + ReplSessionDSReadByInvalidID,
				ExpectError: regexp.MustCompile(`Error reading replication session`),
			},
		},
	})
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

var ReplSessionDSReadByInvalidID = `
data "powerstore_replication_session" "invalid" {
  id = "invalid-id-does-not-exist"
}
`
