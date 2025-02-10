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

# fetching all NFS exports on the array
data powerstore_nfs_export all_nfs_exports {
}

# fetching NFS export using id
data powerstore_nfs_export nfs_export_by_id {
  id = "67974f74-6688-b677-9d08-5692f12c6aa4"
}

# fetching NFS exports using name
data powerstore_nfs_export nfs_export_by_name {
  name = "nfs-export-1"
}

# fetching all NFS exports from a filesystem
data powerstore_filesystem us_east_sales_catalog {
  name = "us-east-sales-catalog"
}

data powerstore_nfs_export nfs_export_by_filesystem {
  file_system_id = data.powerstore_filesystem.us_east_sales_catalog.filesystems[0].id
}

# name and filesystem_id filters can be used together
data powerstore_nfs_export nfs_export_by_filesystem_and_name {
  file_system_id = data.powerstore_filesystem.us_east_sales_catalog.filesystems[0].id
  name          = "nfs-export-1"
}

# fetching NFS exports using filter expression
data powerstore_nfs_export nfs_export_by_name_regex {
  filter_expression = "name=ilike.us-east*"
}

output "nfs_exports_with_name_regex" {
  value = data.powerstore_nfs_export.nfs_export_by_name_regex.nfs_exports
}
