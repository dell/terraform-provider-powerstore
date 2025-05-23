/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// Test to Fetch Volume Group snapshots
func TestAccVolumeGroupSnapshotDs_FetchVolumeGroupSnapshot(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsAll,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsFilter,
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsIDAndNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsEmptyIDNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsEmptyNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
		},
	})
}

var VolumeGroupSnapshotDataSourceparamsName = PreReqVolumeGroupSnap + `
data "powerstore_volumegroup_snapshot" "test1" {
	depends_on = [powerstore_volumegroup_snapshot.test]
	name = powerstore_volumegroup_snapshot.test.name
}
`

var VolumeGroupSnapshotDataSourceparamsNameNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	name = "invalid-name"
}
`

var VolumeGroupSnapshotDataSourceparamsID = PreReqVolumeGroupSnap + `
data "powerstore_volumegroup_snapshot" "test1" {
	id = powerstore_volumegroup_snapshot.test.id
}
`

var VolumeGroupSnapshotDataSourceparamsAll = PreReqVolumeGroupSnap + `
data "powerstore_volumegroup_snapshot" "test1" {
	depends_on = [powerstore_volumegroup_snapshot.test]
}
`

var VolumeGroupSnapshotDataSourceparamsIDNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	id = "invalid-id"
}
`

var VolumeGroupSnapshotDataSourceparamsIDAndNameNegative = PreReqVolumeGroupSnap + `
data "powerstore_volumegroup_snapshot" "test1" {
	id = powerstore_volumegroup_snapshot.test.id
	name = powerstore_volumegroup_snapshot.test.name
}
`

var VolumeGroupSnapshotDataSourceparamsEmptyIDNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	id = ""
}
`

var VolumeGroupSnapshotDataSourceparamsEmptyNameNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	name = ""
}
`

var VolumeGroupSnapshotDataSourceparamsFilter = PreReqVolumeGroupSnap + `
data "powerstore_volumegroup_snapshot" "test1" {
	depends_on = [powerstore_volumegroup_snapshot.test]
	filter_expression = "name=ilike.test_snap*"
}
`
