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

# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads volume groups either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volume groups
# If id or name is provided then it reads a particular volume group with that id or name
# Only one of the attribute can be provided among id and  name 

# Fetch all volume groups
data "powerstore_volumegroup" "all" {
}

# Fetch volume group by name
data "powerstore_volumegroup" "name" {
  name = "volume_group"
}

# Fetch volume group by id
data "powerstore_volumegroup" "id" {
  id = "2d0780e3-2ce7-4d8b-b2ec-349c5e9e65"
}

# Get volume groups details using filter expression
# This filter expression will fetch all the volume groups where name starts with `vg`
data "powerstore_volumegroup" "filter" {
  filter_expression = "name=ilike.vg*"
}

# Output all Volume Group Details
output "volumegroup_all_details" {
  value = data.powerstore_volumegroup.all.volume_groups
}

# Output only Volume Group names
output "volume_group_names_only" {
  value = data.powerstore_volumegroup.all.volume_groups.*.name
}

# Output Volume Group creation_timestamps and location_history with Volume Group ID as key
output "volume_group_id_and_size" {
  value = {
    for volume_group in data.powerstore_volumegroup.all.volume_groups : volume_group.id => {
      creation_timestamp  = volume_group.creation_timestamp 
      location_history  = volume_group.location_history 
    }
  }
}
