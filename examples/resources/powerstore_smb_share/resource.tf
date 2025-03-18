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

# To create an SMB Share of a filesystem, we shall:
# 1. get the id of the filesystem to be shared over SMB
data "powerstore_filesystem" "sales_catalog" {
  name = "sales_catalog_fs"
  lifecycle {
    postcondition {
      condition = length(self.filesystems) == 0
      error_message = "Expected a single filesystem for sales catalog, but got ${length(self.filesystems)}"
    }
  }
}

# 2. create an SMB Share from that filesystem
resource "powerstore_smb_share" "sales_catalog_for_2024_march" {
  // Required
  file_system_id = data.powerstore_filesystem.sales_catalog.filesystems[0].id
  name           = "sales_catalog_for_2024_march"
  path           = "/sales_catalog_fs/2024/March"

  // Optional
  description                        = "SMB Share for sales catalog for 2024 March"
  aces                               = [{ "access_level" : "Full", "access_type" : "Allow", "trustee_name" : "Everyone", "trustee_type" : "WellKnown" }]
  is_abe_enabled                     = true
  is_continuous_availability_enabled = true
  is_encryption_enabled              = true
  is_branch_cache_enabled            = true
  offline_availability               = "Manual"
  umask                              = "077"
}
