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
# Read and Update is supported for this resource
# The PowerStore system maintains a single recycle bin configuration instance (id="0")
# Creation and deletion of the recycle bin configuration is not supported — only update
# To check which attributes of the recycle bin config can be updated, please refer Product Guide in the documentation

# Set recycle bin expiration to 7 days
resource "powerstore_recycle_bin_config" "example" {
  expiration_duration = 7
}

# Disable recycle bin (items expire immediately)
# resource "powerstore_recycle_bin_config" "disabled" {
#   expiration_duration = 0
# }

# Set maximum retention (30 days)
# resource "powerstore_recycle_bin_config" "max_retention" {
#   expiration_duration = 30
# }

# Below example is for import operation
# resource "powerstore_recycle_bin_config" "import_example" {
# }
# terraform import powerstore_recycle_bin_config.import_example 0
