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

# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads hosts either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the filesystem 
# If id or name is provided then it reads a particular file system snapshot with that id or name
# If filesystem_id is provided then it will read all the filesystem snapshots within filesystem
# Only one of the attribute can be provided among id or name or filesystem_id


// create filesystem
resource "powerstore_filesystem" "test" {
  name          = "test_fs"
  description   = "testing file system"
  size          = 3
  nas_server_id = "<nas_server_id>"
  flr_attributes = {
    mode              = "Enterprise"
    minimum_retention = "1D"
    default_retention = "1D"
    maximum_retention = "infinite"
  }
  config_type                     = "General"
  access_policy                   = "UNIX"
  locking_policy                  = "Advisory"
  folder_rename_policy            = "All_Allowed"
  is_smb_sync_writes_enabled      = true
  is_smb_no_notify_enabled        = true
  is_smb_op_locks_enabled         = false
  is_smb_notify_on_access_enabled = true
  is_smb_notify_on_write_enabled  = false
  smb_notify_on_change_dir_depth  = 12
  is_async_mtime_enabled          = true
  file_events_publishing_mode     = "All"
}

// create snapshot from filesystem
resource "powerstore_filesystem_snapshot" "test" {
  name                 = "tf_fs_snap"
  description          = "Test File System Snapshot Resource"
  filesystem_id        = resource.powerstore_filesystem.test.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type          = "Snapshot"
}

data "powerstore_filesystem_snapshot" "test1" {
  name = resource.powerstore_filesystem_snapshot.test.name
  # id = resource.powerstore_filesystem_snapshot.test.id
  # filesystem_id= resource.powerstore_filesystem.test.id
}


output "fileSystemSnapshotResult" {
  value = data.powerstore_filesystem_snapshot.test1.filesystem_snapshots
}
