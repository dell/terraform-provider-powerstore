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

# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete and Import is supported for this resource

# get NAS server id by name of the NAS server
data "powerstore_nas_server" "nas_server_us_east" {
  name = "nas_server_us_east"
  lifecycle {
    postcondition {
      condition = length(self.nas_servers) == 1
      error_message = "Expected a single NAS server, but got none"
    }
  }
}

# create filesystem on the NAS server
resource "powerstore_filesystem" "us_east_sales_catalog_fs" {
  name          = "us_east_sales_catalog_fs"
  description   = "File System for US East Sales Catalog"
  size          = 3
  nas_server_id = data.powerstore_nas_server.nas_server_us_east.nas_servers[0].id
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
