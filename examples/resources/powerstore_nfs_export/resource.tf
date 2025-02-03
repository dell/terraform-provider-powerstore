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

}
