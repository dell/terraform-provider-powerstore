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

title: "powerstore_nfs_export resource"
linkTitle: "powerstore_nfs_export"
page_title: "powerstore_nfs_export Resource - powerstore"
subcategory: "File Storage Management"
description: |-
  This resource is used to manage the nfs export entity of PowerStore Array. We can Create, Update and Delete the nfs export using this resource. We can also import an existing nfs export from PowerStore array.
---

# powerstore_nfs_export (Resource)

This resource is used to manage the nfs export entity of PowerStore Array. We can Create, Update and Delete the nfs export using this resource. We can also import an existing nfs export from PowerStore array.

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

# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete and Import is supported for this resource

# To create an NFS export of a filesystem, we shall:
# 1. get the id of the filesystem to be shared over NFS
data "powerstore_filesystem" "sales_catalog" {
  name = "sales_catalog_fs"
  lifecycle {
    postcondition {
      condition     = length(self.filesystems) == 0
      error_message = "Expected a single filesystem for sales catalog, but got ${length(self.filesystems)}"
    }
  }
}

# 2. create an NFS export from that filesystem
resource "powerstore_nfs_export" "sales_catalog_for_2024_march" {
  // Required
  file_system_id = data.powerstore_filesystem.sales_catalog.filesystems[0].id
  name           = "sales_catalog_for_2024_march"
  path           = "/sales_catalog_fs/2024/March"

  // Optional
  anonymous_gid  = -24
  anonymous_uid  = -24
  description    = "NFS export of Sales Catalog for 2024 March"
  is_no_suid     = false
  min_security   = "Sys"        # Options: "Sys", "Kerberos", "Kerberos_With_Integrity", "Kerberos_With_Encryption"
  default_access = "Read_Write" # Options: "No_Access", "Read_Only", "Read_Write", "Root", "Read_Only_Root"

  # host access related fields (optional)
  no_access_hosts = [
    "192.168.1.0/24", # subnet - ipv4/prefixlength form
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",      # subnet - ipv4/subnet mask form  
    "2001:db8:85a3::8a2e:370:7334/64", # subnet - ipv6/prefixLength form
    "2001:db8:85a3::8a2e:370:7334",    # ipv6 address
    "2001:db8:85a3::/64",              # subnet - ipv6/prefixLength normalized form
  ]

  read_only_hosts = [
    "10.168.1.0/24",
    "11.28.1.0", # ipv4 address
  ]

  read_only_root_hosts = [
    "11.168.1.0/24",
    "hostname1", # hostname
    "hostname2",
    "@netgroup1", # netgroup (must be prefixed with @)
  ]

  read_write_hosts = [
    "12.168.1.0/24",
    "dell.com" # dns domain
  ]

  read_write_root_hosts = [
    "13.168.1.0/24",
  ]

  # set the above lists to empty list to remove all hosts from the NFS export, as below
  # no_access_hosts = []
  # read_only_hosts = []
  # read_only_root_hosts = []
  # read_write_hosts = []
  # read_write_root_hosts = []

}

# To expose a snapshot of a filesystem via NFS, we shall:
# 1. create a snapshot of type "Protocol" of the given filesystem
resource "powerstore_filesystem_snapshot" "sales_catalog_snap" {
  name          = "sales_catalog_snap"
  description   = "Snapshot of Sales Catalog Filesystem"
  filesystem_id = data.powerstore_filesystem.sales_catalog.filesystems[0].id
  access_type   = "Protocol"
}

# 2. Expose the snapshot over NFS (here, we are sharing the /2024/March directory from the snapshot)
resource "powerstore_nfs_export" "sales_catalog_for_2024_march_snap" {
  file_system_id = powerstore_filesystem_snapshot.sales_catalog_snap.id
  name           = "sales_catalog_for_2024_march_snap"
  path           = "/${powerstore_filesystem_snapshot.sales_catalog_snap.name}/2024/March"
  description    = "NFS export of Sales Catalog for 2024 March from snapshot"
}
```

After the execution of above resource block, NFS Export would have been created on the PowerStore array. For more information, Please check the terraform state file.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `file_system_id` (String) The unique identifier of the file	system on which the NFS Export will be created.
- `name` (String) The name of the NFS Export.
- `path` (String) The local path to export relative to the nfs export root directory. With NFS, each export of a file_system or file_nfs must have a unique local path. Before you can create additional Exports within an NFS shared folder, you must create directories within it from a Linux/Unix host that is connected to the nfs export. After a directory has been created from a mounted host, you can create a corresponding Export and Set access permissions accordingly.

### Optional

- `anonymous_gid` (Number) The GID (Group ID) of the anonymous user. This is the group ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.
- `anonymous_uid` (Number) The UID (User ID) of the anonymous user. This is the user ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.
- `default_access` (String) The default access level for all hosts that can access the NFS Export. The default access level is the access level that is assigned to a host that is not explicitly Seted in the 'no_access_hosts', 'read_only_hosts', 'read_only_root_hosts', 'read_write_hosts', or 'read_write_root_hosts' Sets. Valid values are: 'No_Access', 'Read_Only', 'Read_Write', 'Root', 'Read_Only_Root'.
- `description` (String) A user-defined description of the NFS Export.
- `is_no_suid` (Boolean) If Set, do not allow access to Set SUID. Otherwise, allow access.
- `min_security` (String) The NFS enforced security type for users accessing the NFS Export. Valid values are: 'Sys', 'Kerberos', 'Kerberos_With_Integrity', 'Kerberos_With_Encryption'.
- `nfs_owner_username` (String) The default owner of the NFS Export associated with the datastore. Required if secure NFS enabled. For NFSv3 or NFSv4 without Kerberos, the default owner is root. Was added in version 3.0.0.0.
- `no_access_hosts` (Set of String) Hosts with no access to the NFS export or its snapshots. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLength, or IPv4/subnetmask), or Netgroups prefixed with @.
- `read_only_hosts` (Set of String) Hosts with read-only access to the NFS export and its snapshots. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLength, or IPv4/subnetmask), or Netgroups prefixed with @.
- `read_only_root_hosts` (Set of String) Hosts with read-only and read-only for root user access to the NFS Export and its snapshots. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLength, or IPv4/subnetmask), or Netgroups prefixed with @.
- `read_write_hosts` (Set of String) Hosts with read and write access to the NFS Export and its snapshots. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLength, or IPv4/subnetmask), or Netgroups prefixed with @.
- `read_write_root_hosts` (Set of String) Hosts with read and write and read and write for root user access to the NFS Export and its snapshots. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLength, or IPv4/subnetmask), or Netgroups prefixed with @.

### Read-Only

- `id` (String) The unique identifier of the NFS Export.

## Import

Import is supported using the following syntax:

```shell
#Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.
#
#Licensed under the Mozilla Public License Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://mozilla.org/MPL/2.0/
#
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.


# Below are the steps to import nfs export :
# Step 1 - To import a nfs export , we need the id of that nfs export 
# Step 2 - To check the id of the nfs export we can make use of nfs export datasource to read required/all nfs export ids. Alternatively, we can make GET request to nfs export endpoint. eg. https://10.0.0.1/api/rest/nfs_export which will return list of all nfs export ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_nfs_export" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_nfs_export.resource_block_name" "id_of_the_nfs_export" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
```
