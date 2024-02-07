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
# To import , check host_group_import.tf for more info
# name and host_ids are the required attributes to create and update
# description is the optional attribute
# Host datasource can be used to fetch host id/name.
# To check which attributes of the host resource can be updated, please refer Product Guide in the documentation

resource "powerstore_hostgroup" "test" {
  name        = "test_hostgroup"
  description = "Creating host group"
  host_ids    = ["42c60954-ea71-4b50-b172-63880cd48f99"]
}
