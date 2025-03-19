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
# Create, Update, Delete is supported for this resource
# To import , check powerstore_filesystem_snapshot/import.tf for more info
# filesystem_id is the required attribute to create file system snapshot.
# name, expiration_timestamp, access_type and description are the optional attributes
# if name is present in the config it cannot be blank("").
# During create operation, if expiration_timestamp is not specified or set to blank(""), snapshot will be created with infinite retention.
# During modify operation, to set infinite retention, expiration_timestamp can be set to blank("").
# To check which attributes of the file system snapshot resource can be updated, please refer Product Guide in the documentation

# To create a file system snapshot, we shall:
# 1. get the id of the filesystem to be snapshotted
data "powerstore_filesystem" "us_east_sales_catalog_fs" {
  name = "us_east_sales_catalog_fs"
  lifecycle {
    postcondition {
      condition = length(self.filesystems) == 0
      error_message = "Expected a single filesystem for US East sales catalog, but got ${length(self.filesystems)}"
    }
  }
}

# 2. create an expiration timestamp in the RFC3339 format
resource "time_offset" "us_east_sales_catalog_snapshot_expiration_timestamp" {
  // this will set expiration timestamp to 2 years 1 month from the time of creation of the snapshot
  offset_years  = 2
  offset_months = 1
}

# 3. take the snapshot 
resource "powerstore_filesystem_snapshot" "us_east_sales_catalog_snapshot" {
  name                 = "us_east_sales_catalog_snapshot"
  description          = "Snapshot of US East Sales Catalog"
  filesystem_id        = data.powerstore_filesystem.us_east_sales_catalog_fs.filesystems[0].id
  expiration_timestamp = time_offset.us_east_sales_catalog_snapshot_expiration_timestamp.rfc3339
  access_type          = "Snapshot"
}
