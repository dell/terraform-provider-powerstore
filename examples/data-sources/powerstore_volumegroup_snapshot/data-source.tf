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
# This datasource reads volume group snapshots either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volume group snapshots
# If id or name is provided then it reads a particular volume group snapshot with that id or name
# Only one of the attribute can be provided among id and name

# Fetch all volume group snapshots
data "powerstore_volumegroup_snapshot" "all" {
  depends_on = [powerstore_volumegroup_snapshot.test]
}

# Fetch volume group snapshots by name
data "powerstore_volumegroup_snapshot" "name" {
  name = data.powerstore_volumegroup_snapshot.all.volume_groups[0].name
}

# Fetch volume group snapshots by id
data "powerstore_volumegroup_snapshot" "id" {
  id = data.powerstore_volumegroup_snapshot.all.volume_groups[0].id
}

# Get volume group snapshots details using filter expression
data "powerstore_volumegroup_snapshot" "filter" {
  filter_expression = "name=ilike.snap"
}

output "volumeGroupSnapshotResult" {
  value = data.powerstore_volumegroup_snapshot.all.volume_groups
}
