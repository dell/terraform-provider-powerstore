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

data "powerstore_filesystem" "filesystems" {
}


resource "powerstore_smb_share" "smbShare" {
  // Required
  file_system_id = data.powerstore_filesystem.filesystems.filesystems[0].id
  name           = "terraform-smb"
  path           = "/terraform-fs"

  // Optional
  description                        = "smb share"
  aces                               = [{ "access_level" : "Full", "access_type" : "Allow", "trustee_name" : "Everyone", "trustee_type" : "WellKnown" }]
  is_abe_enabled                     = true
  is_continuous_availability_enabled = true
  is_encryption_enabled              = true
  is_branch_cache_enabled            = true
  offline_availability               = "Manual"
  umask                              = "077"
}
