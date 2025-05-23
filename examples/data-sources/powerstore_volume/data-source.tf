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
# This datasource reads volumes either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volumes
# If id or name is provided then it reads a particular volume with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_volume" "test1" {
  name = "tf_vol"
}

# Get volume details using filter expression
# This filter expression will fetch all the volumes where `is_replication_destination` is set to `false` and `name` contains `vol`
data "powerstore_volume" "volume_by_filter" {
  filter_expression = "and=(is_replication_destination.eq.false, name.ilike.*vol*)"
}

# Output all Volume Details
output "volume_all_details" {
  value = data.powerstore_volume.test1.volumes
}

# Output only Volume IDs
output "volumes_IDs_only" {
  value = data.powerstore_volume.test1.volumes.*.id
}

# Output Volume IDs and sizes with Volume name as key
output "volume_id_and_size" {
  value = {
    for volume in data.powerstore_volume.test1.volumes : volume.name => {
      id = volume.id
      size = volume.size
    }
  }
}
