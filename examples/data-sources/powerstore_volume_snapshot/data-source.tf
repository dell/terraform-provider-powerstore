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
# This datasource reads volume snapshots either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volume snapshots
# If id or name is provided then it reads a particular volume snapshot with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_volume_snapshot" "test1" {
  name = "test_snap"
  #id = "adeeef05-aa68-4c17-b2d0-12c4a8e69176"
}

# Get volume snapshot details using filter expression
# This filter expression will fetch all the volume snapshots where `is_replication_destination` is set to `false` and `state_l10n` is `Ready`
data "powerstore_volume_snapshot" "volume_snapshot_by_filter" {
  filter_expression = "and=(is_replication_destination.eq.false, state_l10n.eq.Ready)"
}

# Output all Volume Snapshot Details
output "volume_snapshot_all_details" {
  value = data.powerstore_volume_snapshot.test1.volumes
}

# Output only Volume Snapshot IDs
output "volume_snapshot_IDs_only" {
  value = data.powerstore_volume_snapshot.test1.volumes.*.id
}

# Output Volume Snapshot IDs and sizes with Volume Snapshot name as key
output "volume_snapshot_id_and_size" {
  value = {
    for volume_snapshot in data.powerstore_volume_snapshot.test1.volumes : volume_snapshot.name => {
      id   = volume_snapshot.id
      size = volume_snapshot.size
    }
  }
}
