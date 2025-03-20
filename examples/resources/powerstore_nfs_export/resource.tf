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
