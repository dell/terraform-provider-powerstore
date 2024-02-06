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
# To import , check volume_group_import.tf for more info
# name is the required attribute to create and update
# Volume datasource can be used to fetch volume id/name.
# Protection policy datasource can be used to fetch protection policy id/name.
# To check which attributes of the volume group can be updated, please refer Product Guide in the documentation

resource "powerstore_volumegroup" "terraform-provider-test1" {
  # (resource arguments)
  description               = "Creating Volume Group"
  name                      = "test_volume_group"
  is_write_order_consistent = "false"
  protection_policy_id      = "01b8521d-26f5-479f-ac7d-3d8666097094"
  volume_ids                = ["140bb395-1d85-49ae-bde8-35070383bd92"]
}
