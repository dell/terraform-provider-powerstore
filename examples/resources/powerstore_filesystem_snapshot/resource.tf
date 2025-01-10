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

resource "powerstore_filesystem_snapshot" "test1" {
  name                 = "tf_fs_snap"
  description          = "Test File System Snapshot Resource"
  filesystem_id        = "67608dc7-b69c-b762-0522-42848bc63a0b"
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type          = "Snapshot"
}