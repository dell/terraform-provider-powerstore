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

// Helper function to generate unique volume group name
func getMetroVolumeGroupName() string {
	if endpoint == "http://localhost:3003/api/rest" {
		return "tf_volume_group_new" // Use mock server's expected name
	}
	return fmt.Sprintf("metro_test_vg-%d", time.Now().UnixNano()) // Use dynamic name for real server
}

// Helper function to generate volume names for volume group tests
func getMetroVolumeNamesForVGroup() (string, string) {
	if endpoint == "http://localhost:3003/api/rest" {
		return "test_acc_cvol", "test_acc_cvol" // Use mock server's expected names (both same)
	}
	vgName := getMetroVolumeGroupName()
	return fmt.Sprintf("%s-vol1", vgName), fmt.Sprintf("%s-vol2", vgName)
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
