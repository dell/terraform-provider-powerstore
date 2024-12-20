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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
		"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
		"github.com/stretchr/testify/assert"
		"regexp"
)

func TestAccFileSystem_CreateFS(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + FsParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "name", "test_fs"),
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "size", "5"),
				),
			},

			// Import Testing
			{
				Config:       ProviderConfigForTesting + FsParams,
				ResourceName: "powerstore_filesystem.test_fs_create",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "test_fs", s[0].Attributes["name"])
					return nil
				},
			},
		},
	})
}

func TestAccFileSystem_InvalidUpdate(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FsParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "name", "test_fs"),
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "size", "5"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsUpdate,
				ExpectError: regexp.MustCompile(".*Name of the file system can't be updated.*"),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsUpdate2,
				ExpectError: regexp.MustCompile(".*NAS server ID can't be updated.*"),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsUpdate3,
				ExpectError: regexp.MustCompile(".*Mode of the flr attributes can't be updated.*"),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsUpdate4,
				ExpectError: regexp.MustCompile(".*Config type can't be updated.*"),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsUpdate5,
				ExpectError: regexp.MustCompile(".*Host IO size can't be updated.*"),
			},
		},
	})
}

func TestAccFileSystem_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FsParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "name", "test_fs"),
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "size", "5"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsValidUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "description", "testing file system Update"),
				),
			},
		},
	})
}

func TestAccFileSystem_CreateErr(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + FsParamsCreateErr,
				ExpectError: regexp.MustCompile(".*Error creating file system.*"),
			},
			{
				Config:      ProviderConfigForTesting + FsParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "name", "test_fs"),
					resource.TestCheckResourceAttr("powerstore_filesystem.test_fs_create", "size", "5"),
				),
			},
			{
				Config:      ProviderConfigForTesting + FsParamsModifyErr,
				ExpectError: regexp.MustCompile(".*Error updating file system.*"),
			},
		},
	})
}


var FsParams = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system"
	size      = 5
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = true
	is_smb_no_notify_enabled = true
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = true
	file_events_publishing_mode = "All"
}
`

var FsParamsUpdate = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs_rename"
	description = "testing file system"
	size      = 5
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = false
	is_smb_no_notify_enabled = false
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = false
	file_events_publishing_mode = "All"
}
`

var FsParamsUpdate2 = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system update"
	size      = 5
	nas_server_id = "Id Updated"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = false
	is_smb_no_notify_enabled = false
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = false
	file_events_publishing_mode = "All"
}
`

var FsParamsUpdate3 = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system update"
	size      = 5
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "None"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = false
	is_smb_no_notify_enabled = false
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = false
	file_events_publishing_mode = "All"
}
`

var FsParamsUpdate4 = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system update"
	size      = 5
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "VMware"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = false
	is_smb_no_notify_enabled = false
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = false
	file_events_publishing_mode = "All"
}
`

var FsParamsUpdate5 = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system update"
	size      = 5
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = false
	is_smb_no_notify_enabled = false
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = false
	file_events_publishing_mode = "All"
	host_io_size = "VMware_8K"
}
`

var FsParamsValidUpdate = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system Update"
	size      = 4
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = false
	is_smb_no_notify_enabled = false
	is_smb_op_locks_enabled = true
	is_smb_notify_on_access_enabled = false
	is_smb_notify_on_write_enabled = false
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = false
	file_events_publishing_mode = "All"
}
`

var FsParamsCreateErr = `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system"
	size      = 5
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "VMware"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = true
	is_smb_no_notify_enabled = true
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = true
	file_events_publishing_mode = "All"
}
`

var FsParamsModifyErr= `
resource "powerstore_filesystem" "test_fs_create" {
	name                 = "test_fs"
	description = "testing file system"
	size      = 1
	nas_server_id = "` + nasServerID + `"
	flr_attributes = {
	  mode = "Enterprise"
	}
	config_type = "General"
	access_policy = "UNIX"
	locking_policy = "Advisory"
	folder_rename_policy = "All_Allowed"
	is_smb_sync_writes_enabled = true
	is_smb_no_notify_enabled = true
	is_smb_op_locks_enabled = false
	is_smb_notify_on_access_enabled = true
	is_smb_notify_on_write_enabled = true
	smb_notify_on_change_dir_depth = 12
	is_async_mtime_enabled = true
	file_events_publishing_mode = "All"
}
`