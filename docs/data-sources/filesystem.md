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

title: "powerstore_filesystem data source"
linkTitle: "powerstore_filesystem"
page_title: "powerstore_filesystem Data Source - powerstore"
subcategory: ""
description: |-
  This datasource is used to query the existing File System from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.
---

# powerstore_filesystem (Data Source)

This datasource is used to query the existing File System from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.

> **Note:** Only one of `name` or `id` or `nas_server_id` can be provided at a time.

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
# This datasource reads filesystems either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the filesystems
# If id or name is provided then it reads a particular filesystem with that id or name
# Only one of the attribute can be provided among id and name

#Fetching filesystem using name
data "powerstore_filesystem" "test1" {
  name = "ephemeral-csi-0e62d1cdec3b543a0740e9a307df221187955b47af4b99b7be31ebd9cb6192ce"
}

#Fetching filesystem using id
data "powerstore_filesystem" "test2" {
  id = "6568282e-c982-62ce-5ac3-52518d324736"
}

#Fetching filesystem using nas server id
data "powerstore_filesystem" "test3" {
  nas_server_id = "654b2182-f674-f39a-66fc-52518d324736"
}

# Fetching all filesystems
data "powerstore_filesystem" "test4" {
}


output "result" {
  value = data.powerstore_filesystem.test3.filesystems
}
```
After the successful execution of above said block, We can see the output by executing `terraform output` command. Also, we can fetch information via the variable: `data.powerstore_filesystem_snapshot.test1` where name is the attribute which user wants to fetch.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) Unique identifier of the File System. Conflicts with `name` and `nas_server_id`.
- `name` (String) File System name. Conflicts with `id` and `nas_server_id`.
- `nas_server_id` (String) Nas server ID. Conflicts with `id` and `name`.

### Read-Only

- `filesystems` (Attributes List) List of File System. (see [below for nested schema](#nestedatt--filesystems))

<a id="nestedatt--filesystems"></a>
### Nested Schema for `filesystems`

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
- `flr_attributes` (Attributes) Flr Attributes of the File System (see [below for nested schema](#nestedatt--filesystems--flr_attributes))
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

<a id="nestedatt--filesystems--flr_attributes"></a>
### Nested Schema for `filesystems.flr_attributes`

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