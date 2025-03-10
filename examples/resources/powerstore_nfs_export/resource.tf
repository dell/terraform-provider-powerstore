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

data "powerstore_filesystem" "test4" {
}


resource "powerstore_nfs_export" "test1" {
  // Required
  file_system_id = data.powerstore_filesystem.test3.filesystems[0].id
  name           = "terraform-nfs"
  path           = "/terraform-fs"

  // Optional
  anonymous_gid  = -24
  anonymous_uid  = -24
  description    = "nfs export"
  is_no_suid     = false
  min_security   = "Sys"        # Options: "Sys", "Kerberos", "Kerberos_With_Integrity", "Kerberos_With_Encryption"
  default_access = "Read_Write" # Options: "No_Access", "Read_Only", "Read_Write", "Root", "Read_Only_Root"

  no_access_hosts = [
    "192.168.1.0/24",
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",
    "192.168.1.54/255.1009.255.0",    
    "2001:db8:85a3::8a2e:370:7334/255.255.255.0",
    "2001:db8:85a3::8a2e:370:7334",
    "2001:db8:85a3::/64",
  ]

  read_only_hosts = [
    "10.168.1.0/24",
    "11.28.1.0",
  ]

  read_only_root_hosts = [
    "11.168.1.0/24",
    "hostname1",
    "hostname2",
    "@netgroup1",
  ]

  read_write_hosts = [
    "12.168.1.0/24",
  ]

  read_write_root_hosts = [
    "13.168.1.0/24",
  ]

}
