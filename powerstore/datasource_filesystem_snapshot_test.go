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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch File System Snapshots
func TestAccFileSystemSnapshotDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceParamsNasServerID,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_filesystem_snapshot.test", "filesystem_snapshots.0.nas_server_id", nasServerID),
				),
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsFileSystemID,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.powerstore_filesystem_snapshot.test1", "filesystem_snapshots.0.parent_id", "powerstore_filesystem.test_fs_create", "id"),
				),
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsName,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.powerstore_filesystem_snapshot.test2", "filesystem_snapshots.0.name", "powerstore_filesystem_snapshot.test", "name"),
				),
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsID,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.powerstore_filesystem_snapshot.test3", "filesystem_snapshots.0.id", "powerstore_filesystem_snapshot.test", "id"),
				),
			},

			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsAll,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.powerstore_filesystem_snapshot.test4", "filesystem_snapshots.#"),
				),
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceParamsFilter,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.powerstore_filesystem_snapshot.test3", "filesystem_snapshots.0.access_type", "powerstore_filesystem_snapshot.test", "access_type"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsIDAndNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsEmptyIDNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsEmptyNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore File System Snapshot by ID"),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsFilterNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore File System Snapshot"),
			},
		},
	})
}

var FileSystemSnapshotDataSourceParamsNasServerID = FileSystemSnapParamsCreate + `	
data "powerstore_filesystem_snapshot" "test" {
	nas_server_id = "` + nasServerID + `"
	depends_on = [powerstore_filesystem.test_fs_create, powerstore_filesystem_snapshot.test]
}`

var FileSystemSnapshotDataSourceParamsFilter = FileSystemSnapParamsCreate + `	
data "powerstore_filesystem_snapshot" "test" {
	filer_expression = "access_type=eq.Snapshot"
	depends_on = [powerstore_filesystem.test_fs_create, powerstore_filesystem_snapshot.test]
}`

var FileSystemSnapshotDataSourceparamsFileSystemID = FileSystemSnapParamsCreate + `
data "powerstore_filesystem_snapshot" "test1" {
	filesystem_id = powerstore_filesystem.test_fs_create.id
	depends_on = [powerstore_filesystem.test_fs_create, powerstore_filesystem_snapshot.test]
}
`

var FileSystemSnapshotDataSourceparamsName = FileSystemSnapParamsCreate + `
data "powerstore_filesystem_snapshot" "test2" {
	name = powerstore_filesystem_snapshot.test.name
}
`

var FileSystemSnapshotDataSourceparamsID = FileSystemSnapParamsCreate + `
data "powerstore_filesystem_snapshot" "test3" {
	id = powerstore_filesystem_snapshot.test.id
}
`

var FileSystemSnapshotDataSourceparamsAll = `
data "powerstore_filesystem_snapshot" "test4" {
}
`

var FileSystemSnapshotDataSourceparamsIDAndNameNegative = `
data "powerstore_filesystem_snapshot" "test5" {
	id = "unique-id"
	name = "unique-name"
}
`

var FileSystemSnapshotDataSourceparamsNameNegative = `
data "powerstore_filesystem_snapshot" "test6" {
	name = "invalid-name"
}
`

var FileSystemSnapshotDataSourceparamsIDNegative = `
data "powerstore_filesystem_snapshot" "test7" {
	id = "invalid-id"
}
`

var FileSystemSnapshotDataSourceparamsEmptyIDNegative = `
data "powerstore_filesystem_snapshot" "test8" {
	id = ""
}
`

var FileSystemSnapshotDataSourceparamsEmptyNameNegative = `
data "powerstore_filesystem_snapshot" "test9" {
	name = ""
}
`
var FileSystemSnapshotDataSourceparamsFilterNegative = `
data "powerstore_filesystem_snapshot" "test9" {
	filter_expression = "name=InvalidName"
}
`
