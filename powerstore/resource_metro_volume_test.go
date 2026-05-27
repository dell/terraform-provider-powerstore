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

// Helper function to generate unique volume name
func getMetroVolumeName() string {
	if endpoint == "http://localhost:3003/api/rest" {
		return "test_acc_cvol" // Use mock server's expected name
	}
	return fmt.Sprintf("metro_test_vol-%d", time.Now().UnixNano()) // Use dynamic name for real server
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

// Acceptance test: Create metro volume
func TestAccMetroVolume_CreateOnMock(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	// Generate unique config once for this test
	volName := getMetroVolumeName()
	config := fmt.Sprintf(`
resource "powerstore_volume" "volume_create_test" {
  name = "%s"
  size = 2.5
}

resource "powerstore_metro_volume" "test" {
  volume_id        = powerstore_volume.volume_create_test.id
  remote_system_id = "%s"
}
`, volName, remoteSystemID)
	configPaused := fmt.Sprintf(`
resource "powerstore_volume" "volume_create_test" {
  name = "%s"
  size = 2.5
}

resource "powerstore_metro_volume" "test" {
  volume_id            = powerstore_volume.volume_create_test.id
  remote_system_id     = "%s"
  is_replication_paused = true
}
`, volName, remoteSystemID)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("powerstore_metro_volume.test", "id"),
					resource.TestCheckResourceAttrSet("powerstore_metro_volume.test", "metro_replication_session_id"),
					resource.TestCheckResourceAttrSet("powerstore_metro_volume.test", "remote_system_id"),
				),
			},
			// Import test
			{
				Config:            ProviderConfigForTesting + config,
				ResourceName:      "powerstore_metro_volume.test",
				ImportState:       true,
				ImportStateVerify: false,
			},
			// Update test - pause replication
			{
				Config: ProviderConfigForTesting + configPaused,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_metro_volume.test", "is_replication_paused", "true"),
				),
			},
			// Update test - resume replication
			{
				Config: ProviderConfigForTesting + config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_metro_volume.test", "is_replication_paused", "false"),
				),
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

var MetroVolumeParamsCreate = fmt.Sprintf(`
resource "powerstore_volume" "volume_create_test" {
  name = "test_acc_cvol"
  size = 2.5
}

resource "powerstore_metro_volume" "test" {
  volume_id        = powerstore_volume.volume_create_test.id
  remote_system_id = "%s"
}
`, remoteSystemID)

var MetroVolumeParamsCreatePaused = fmt.Sprintf(`
resource "powerstore_volume" "volume_create_test" {
  name = "test_acc_cvol"
  size = 2.5
}

resource "powerstore_metro_volume" "test" {
  volume_id            = powerstore_volume.volume_create_test.id
  remote_system_id     = "%s"
  is_replication_paused = true
}
`, remoteSystemID)
