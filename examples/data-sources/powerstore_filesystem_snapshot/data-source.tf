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
      condition = length(self.filesystems) == 1
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
      condition = length(self.nas_servers) == 1
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


# Fetching all filesystems
data "powerstore_filesystem_snapshot" "all" {
}


output "result" {
  value = data.powerstore_filesystem_snapshot.all.filesystems
}
