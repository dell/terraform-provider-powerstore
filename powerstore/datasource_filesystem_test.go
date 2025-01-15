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
func TestAccFileSystemtDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + FileSystemDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + FileSystemDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + FileSystemDataSourceparamsAll,
			},
			{
				Config: ProviderConfigForTesting + FileSystemDataSourceparamsNasServerId,
			},
			{
				Config: ProviderConfigForTesting + FileSystemtDataSourceparamsNameNegative,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_filesystem.test1", "filesystems.#", "0"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FileSystemDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile(".*Unable to Read PowerStore fileSystem.*"),
			},
		},
	})
}

var FileSystemDataSourceparamsID = FsParams + `
data "powerstore_filesystem" "test1" {
	depends_on = [
		powerstore_filesystem.test_fs_create
	]
	id = resource.powerstore_filesystem.test_fs_create.id
}
`

var FileSystemDataSourceparamsName = FsParams + `
data "powerstore_filesystem" "test1" {
	depends_on = [
		powerstore_filesystem.test_fs_create
	]
	name = resource.powerstore_filesystem.test_fs_create.name
}
`

var FileSystemDataSourceparamsNasServerId = FsParams + `
data "powerstore_filesystem" "test1" {
	depends_on = [
		powerstore_filesystem.test_fs_create
	]
	nas_server_id = resource.powerstore_filesystem.test_fs_create.nas_server_id
}
`

var FileSystemDataSourceparamsAll = FsParams + `
data "powerstore_filesystem" "test1" {
	depends_on = [
		powerstore_filesystem.test_fs_create
	]
}
`

var FileSystemtDataSourceparamsNameNegative = `
data "powerstore_filesystem" "test1" {
	name = "InvalidName"
}
`

var FileSystemDataSourceparamsIDNegative = `
data "powerstore_filesystem" "test1" {
	id = "InvalidID"
}
`
