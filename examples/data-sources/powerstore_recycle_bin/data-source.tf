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
# This datasource reads Recycle Bin items from the PowerStore array
# If it is an empty datasource block, then it will read all recycle bin items
# If id is provided then it reads a particular recycle bin item
# If resource_type is provided, it filters by volume or volume_group
# filter_expression can be used for advanced filtering

# Fetch all recycle bin items
data "powerstore_recycle_bin" "all" {
}

# Fetch recycle bin items filtered by resource type (volumes only)
data "powerstore_recycle_bin" "volumes_only" {
  resource_type = "volume"
}

# Fetch recycle bin items filtered by resource type (volume groups only)
data "powerstore_recycle_bin" "volume_groups_only" {
  resource_type = "volume_group"
}

# Fetch recycle bin items using filter expression
data "powerstore_recycle_bin" "filtered" {
  filter_expression = "resource_type=eq.volume"
}

# Output all recycle bin items
output "recycle_bin_all" {
  value = data.powerstore_recycle_bin.all.recycle_bin_items
}

# Output only recycle bin item IDs
output "recycle_bin_ids" {
  value = data.powerstore_recycle_bin.all.recycle_bin_items.*.id
}

# Output recycle bin item names and types
output "recycle_bin_summary" {
  value = {
    for item in data.powerstore_recycle_bin.all.recycle_bin_items : item.id => {
      name          = item.name
      resource_type = item.resource_type
      expires       = item.expiration_timestamp
    }
  }
}
