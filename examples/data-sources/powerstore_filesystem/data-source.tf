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

# Fetching all filesystems with a specific name
data "powerstore_filesystem" "us_east_sales_catalog_fs" {
  name = "us_east_sales_catalog_fs"
  lifecycle {
    postcondition {
      condition     = length(self.filesystems) > 0
      error_message = "Expected atleast one filesystem, but got none"
    }
  }
}

# Fetching filesystems using filter expression
# This filter expression will fetch all the filesystems where name contains _east_sales_catalog_fs and size used is greater than 594849
data "powerstore_filesystem" "east_sales_catalog_fs" {
  filter_expression = "and=(name.ilike.*_east_sales_catalog_fs*, size_used.gt.594849)"
}

# Fetching a filesystem using id
data "powerstore_filesystem" "filesystem_by_id" {
  id = "6568282e-c982-62ce-5ac3-52518d324736"
}

# Fetching all filesystems under a NAS server
data "powerstore_nas_server" "nas_server_us_east" {
  name = "nas_server_us_east"
  lifecycle {
    postcondition {
      condition     = length(self.nas_servers) == 1
      error_message = "Expected a single NAS server for US East region, but got none"
    }
  }
}

data "powerstore_filesystem" "file_systems_us_east" {
  nas_server_id = data.powerstore_nas_server.nas_server_us_east.nas_servers[0].id
}

# Fetching all filesystems
data "powerstore_filesystem" "all_file_systems" {
}

// Output all filesystem details
output "filesystem_all_details" {
  value = data.powerstore_filesystem.all_file_systems.filesystems
}

// Output only filesystem names
output "filesystem_names_only" {
  value = data.powerstore_filesystem.all_file_systems.filesystems.*.name
}

// Output filesystem name and access policy with filesystem id as key
output "name_and_access_policy" {
  value = {
    for filesystem in data.powerstore_filesystem.all_file_systems.filesystems : filesystem.id => {
      name = filesystem.name
      access_policy = filesystem.access_policy
    }
  }
}