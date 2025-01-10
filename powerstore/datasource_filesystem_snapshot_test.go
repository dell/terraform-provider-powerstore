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
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsFileSystemID,
			},
			{
				Config: ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsAll,
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
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore fileSystemSnapshot Snapshots"),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore fileSystemSnapshot Snapshots"),
			},
		},
	})
}

var FileSystemSnapshotDataSourceparamsFileSystemID = `
data "powerstore_filesystem_snapshot" "test1" {
	filesystem_id = "` + fileSystemID + `"
}
`

var FileSystemSnapshotDataSourceparamsName = `
data "powerstore_filesystem_snapshot" "test1" {
	name = "` + fileSystemSnapshotName + `"
}
`

var FileSystemSnapshotDataSourceparamsNameNegative = `
data "powerstore_filesystem_snapshot" "test1" {
	name = "invalid-name"
}
`

var FileSystemSnapshotDataSourceparamsID = `
data "powerstore_filesystem_snapshot" "test1" {
	id = "` + fileSystemSnapshotID + `"
}
`

var FileSystemSnapshotDataSourceparamsAll = `
data "powerstore_filesystem_snapshot" "test1" {
}
`

var FileSystemSnapshotDataSourceparamsIDNegative = `
data "powerstore_filesystem_snapshot" "test1" {
	id = "invalid-id"
}
`

var FileSystemSnapshotDataSourceparamsIDAndNameNegative = `
data "powerstore_filesystem_snapshot" "test1" {
	id = "` + fileSystemSnapshotID + `"
	name = "` + fileSystemSnapshotName + `"
}
`

var FileSystemSnapshotDataSourceparamsEmptyIDNegative = `
data "powerstore_filesystem_snapshot" "test1" {
	id = ""
}
`

var FileSystemSnapshotDataSourceparamsEmptyNameNegative = `
data "powerstore_filesystem_snapshot" "test1" {
	name = ""
}
`
