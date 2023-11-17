/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Volume Snapshot
func TestAccVolumeDs_FetchVolumeSnapshot(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeSnapshotDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeSnapshotDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeSnapshotDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsIDAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsEmptyID,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Snapshots"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Snapshots"),
			},
		},
	})
}

var VolumeSnapshotDataSourceparamsID = `
data "powerstore_volume_snapshot" "test1" {
	id = "` + volumeSnapshotID + `"
}
`

var VolumeSnapshotDataSourceparamsIDNegative = `
data "powerstore_volume_snapshot" "test1" {
	id = "invalid-id"
}
`

var VolumeSnapshotDataSourceparamsEmptyID = `
data "powerstore_volume_snapshot" "test1" {
	id = ""
}
`

var VolumeSnapshotDataSourceparamsName = `
data "powerstore_volume_snapshot" "test1" {
	name = "` + volumeSnapshotName + `"
}
`

var VolumeSnapshotDataSourceparamsNameNegative = `
data "powerstore_volume_snapshot" "test1" {
	name = "invalid-name"
}
`

var VolumeSnapshotDataSourceparamsEmptyName = `
data "powerstore_volume_snapshot" "test1" {
	name = ""
}
`

var VolumeSnapshotDataSourceparamsAll = `
data "powerstore_volume_snapshot" "test1" {
}
`

var VolumeSnapshotDataSourceparamsIDAndName = `
data "powerstore_volume_snapshot" "test1" {
	id = "` + volumeSnapshotID + `"
	name = "` + volumeSnapshotName + `"
}
`
