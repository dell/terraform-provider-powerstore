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
# This datasource reads Snapshot Rules either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the Snapshot Rules
# If id or name is provided then it reads a particular Snapshot Rule with that id or name
# Only one of the attribute can be provided among id and  name 

# Fetching snapshot rule using filter expression
# This filter expression will fetch the snapshot rule where name is `snapshotrule-ny`
data "powerstore_snapshotrule" "test1" {
  filter_expression = "name=eq.snapshotrule-ny"
}


data "powerstore_snapshotrule" "test1" {
  name = "test_snapshotrule_1"
}

# Output all Snapshot Rules Details
output "snapshotRules_all_details" {
  value = data.powerstore_snapshotrule.test1.snapshot_rules
}

# Output only Snapshot Rule IDs
output "snapshot_rules_IDs_only" {
  value = data.powerstore_snapshotrule.test1.snapshot_rules.*.id
}

# Output Snapshot Rule names and timezone with Snapshot Rule id as key
output "snapshot_rule_name_and_timezone" {
  value = {
    for rule in data.powerstore_snapshotrule.test1.snapshot_rules : rule.id => {
      name = rule.name
      timezone = rule.timezone
    }
  }
}