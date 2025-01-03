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
# This datasource reads replication rules either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the replication rules
# If id or name is provided then it reads a particular replication rule with that id or name
# Only one of the attribute can be provided among id and  name 

# Get replication rule details using name
resource "powerstore_replication_rule" "test" {
  name             = "terraform_replication_rule"
  rpo              = "One_Hour"
  remote_system_id = "db11abb3-789e-47f9-96b5-84b5374cbcd2"
  alert_threshold  = 1000
  is_read_only     = false
}

data "powerstore_replication_rule" "test" {
  depends_on = [powerstore_replication_rule.test]
  name       = "terraform_replication_rule"
}

# Get replication rule details using ID
resource "powerstore_replication_rule" "test" {
  name             = "terraform_replication_rule"
  rpo              = "One_Hour"
  remote_system_id = "db11abb3-789e-47f9-96b5-84b5374cbcd2"
  alert_threshold  = 1000
  is_read_only     = false
}

data "powerstore_replication_rule" "test" {
  id = powerstore_replication_rule.test.id
}

output "replicationRule" {
  value = data.powerstore_replication_rule.test.replication_rules
}
