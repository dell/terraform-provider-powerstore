---
# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.
# 
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://mozilla.org/MPL/2.0/
# 
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "powerstore_filesystem_snapshot data source"
linkTitle: "powerstore_filesystem_snapshot"
page_title: "powerstore_filesystem_snapshot Data Source - powerstore"
subcategory: "Data Protection Management"
description: |-
  This datasource is used to query the existing File System Snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.
---

# powerstore_filesystem_snapshot (Data Source)

This datasource is used to query the existing File System Snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.

> **Note:** Only one of `name` or `id` can be provided at a time.

## Example Usage

```terraform
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
# Only one of the attribute can be provided among id or name 

# Fetching filesystem snapshot(s) using name
data "powerstore_filesystem_snapshot" "sales_catalog_snapshots_q4" {
  name = "sales_catalog_snapshot_q4"
}

#Fetching filesystem snapshot using id
data "powerstore_filesystem_snapshot" "fs_snap_by_id" {
  id = "6568282e-c982-62ce-5ar3-52518f324723"
}


# Fetching all snapshots of a particular filesystem
# Step 1: Fetching the filesystem whose snapshots are to be fetched.
data "powerstore_filesystem" "us_east_sales_catalog_fs" {
  name = "us_east_sales_catalog_fs"
  lifecycle {
    postcondition {
      condition     = length(self.filesystems) == 1
      error_message = "Expected one filesystem with name us_east_sales_catalog_fs, but got ${length(self.filesystems)}"
    }
  }
}

# Step 2: Fetching the filesystem snapshots using the file system id from step 1
data "powerstore_filesystem_snapshot" "us_east_sales_catalog_snapshots" {
  filesystem_id = data.powerstore_filesystem.us_east_sales_catalog_fs.filesystems[0].id
}

# Fetching all filesystem snapshots under a particular nas server
# Step 1: Fetching the NAS server whose children snapshots are to be fetched.
data "powerstore_nas_server" "nas_server_us_east" {
  name = "nas_server_us_east"
  lifecycle {
    postcondition {
      condition     = length(self.nas_servers) == 1
      error_message = "Expected a single NAS server for US East region, but got none"
    }
  }
}
# Step 2: Fetching the filesystem snapshots using the nas server id from step 1
data "powerstore_filesystem_snapshot" "us_east_nas_server_fs_snapshots" {
  nas_server_id = data.powerstore_nas_server.nas_server_us_east.nas_servers[0].id
}

# Fetching the filesystem snapshot with a particular name under a given nas server
data "powerstore_filesystem_snapshot" "sales_catalog_snapshot_q4_under_nas_server" {
  name          = "sales_catalog_snapshot_q4"
  nas_server_id = data.powerstore_nas_server.nas_server_us_east.nas_servers[0].id
}

# Fetching snapshot with given name for a particular file system
data "powerstore_filesystem_snapshot" "sales_catalog_snapshot_q4_under_filesystem" {
  name          = "sales_catalog_snapshot_q4"
  filesystem_id = data.powerstore_filesystem.us_east_sales_catalog_fs.filesystems[0].id
}

# Fetching filesystem snapshots using filter expression
# This filter expression will fetch all the filesystems where creator type is User and size used is greater than 6788440
data "powerstore_filesystem_snapshot" "user_or_root_created_snapshots" {
  filter_expression = "and=(creator_type.eq.User, size_used.gt.6788440)"
}

# Fetching all filesystem snapshots
data "powerstore_filesystem_snapshot" "all" {
}

// Output all filesystem snapshot details
output "filesystem_snapshot_all_details" {
  value = data.powerstore_filesystem_snapshot.all.filesystem_snapshots
}

// Output only filesystem snapshot IDs
output "filesystem_snapshot_IDs_only" {
  value = data.powerstore_filesystem_snapshot.all.filesystem_snapshots.*.id
}

// Output filesystem snapshot name and size used with filesystem snapshot id as key
output "filesystem_snapshot_name_and_size_used" {
  value = {
    for filesystem_snapshot in data.powerstore_filesystem_snapshot.all.filesystem_snapshots : filesystem_snapshot.id => {
      name = filesystem_snapshot.name
      size_used = filesystem_snapshot.size_used
    }
  }
}
```

After the successful execution of above said block, We can see the output by executing `terraform output` command. Also, we can fetch information via the variable: `data.powerstore_filesystem_snapshot.test1.attribute_name` where attribute_name is the attribute which user wants to fetch.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filesystem_id` (String) File System ID of the Snapshot. Conflicts with `id` and `nas_server_id`.
- `filter_expression` (String) PowerStore filter expression to filter Filesystem Snapshots by. Conflicts with `id`, `name`, `nas_server_id` and `file_system_id`.
- `id` (String) Unique identifier of the File System Snapshot. Conflicts with `name` and `filesystem_id`.
- `name` (String) File System Snapshot name. Conflicts with `id`.
- `nas_server_id` (String) Nas Server ID of the Snapshot. Conflicts with `id` and `filesystem_id`.

### Read-Only

- `filesystem_snapshots` (Attributes List) List of File System Snapshots. (see [below for nested schema](#nestedatt--filesystem_snapshots))

<a id="nestedatt--filesystem_snapshots"></a>
### Nested Schema for `filesystem_snapshots`

Read-Only:

- `access_policy` (String) Access Policy of the File System
- `access_type` (String) Access Type of the File System
- `config_type` (String) Config Type of the File System
- `creation_timestamp` (String) Creation Timestamp of the File System
- `creator_type` (String) Creator Type of the File System
- `default_hard_limit` (Number) Default Hard Limit of the File System
- `default_soft_limit` (Number) Default Soft Limit of the File System
- `description` (String) Description of the File System
- `expiration_timestamp` (String) Expiration Timestamp of the File System
- `file_events_publishing_mode` (String) State of the event notification services for all file systems
- `filesystem_type` (String) Filesystem Type of the File System
- `flr_attributes` (Attributes) Flr Attributes of the File System (see [below for nested schema](#nestedatt--filesystem_snapshots--flr_attributes))
- `folder_rename_policy` (String) Folder Rename Policy of the File System
- `grace_period` (Number) Grace Period of the File System
- `host_io_size` (String) Typical size of writes
- `id` (String) ID of the File System
- `is_async_m_time_enabled` (Boolean) Is Async MTime Enabled of the File System
- `is_modified` (Boolean) Is Modified of the File System
- `is_quota_enabled` (Boolean) Is Quota Enabled of the File System
- `is_smb_no_notify_enabled` (Boolean) Is Smb No Notify Enabled of the File System
- `is_smb_notify_on_access_enabled` (Boolean) Is Smb Notify On Access Enabled of the File System
- `is_smb_notify_on_write_enabled` (Boolean) Is Smb Notify On Write Enabled of the File System
- `is_smb_op_locks_enabled` (Boolean) Is Smb Op Locks Enabled of the File System
- `is_smb_sync_writes_enabled` (Boolean) Is Smb Sync Writes Enabled of the File System
- `last_refresh_timestamp` (String) Last Refresh Timestamp of the File System
- `last_writable_timestamp` (String) Last Writable Timestamp of the File System
- `locking_policy` (String) Locking Policy of the File System
- `name` (String) Name of the File System
- `nas_server_id` (String) Nas Server ID of the File System
- `parent_id` (String) Parent ID of the File System
- `protection_policy_id` (String) Protection Policy ID of the File System
- `size_total` (Number) Size Total of the File System
- `size_used` (Number) Size Used of the File System
- `smb_notify_on_change_dir_depth` (Number) Smb Notify On Change Dir Depth of the File System

<a id="nestedatt--filesystem_snapshots--flr_attributes"></a>
### Nested Schema for `filesystem_snapshots.flr_attributes`

Read-Only:

- `auto_delete` (Boolean) Auto Delete of the File System
- `auto_lock` (Boolean) Auto Lock of the File System
- `clock_time` (String) Clock Time of the File System
- `default_retention` (String) Default Retention of the File System
- `has_protected_files` (Boolean) Has Protected Files of the File System
- `maximum_retention` (String) Maximum Retention of the File System
- `maximum_retention_date` (String) Maximum Retention Date of the File System
- `minimum_retention` (String) Minimum Retention of the File System
- `mode` (String) Mode of the File System
- `policy_interval` (Number) Policy Interval of the File System
