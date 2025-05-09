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
# To import, check import.tf for more info
# name, rpo and remote_system_id are required attributes to create and update
# To check which attributes of the replication rule can be updated, please refer Product Guide in the documentation

# To create a replication rule, we shall:
# 1. fetching the Remote System for which the replication rule is to be created
# Here, we are fetching by name of the Remote System
data "powerstore_remote_system" "backup_1hr" {
  name = "RT-D4538"
  lifecycle {
    postcondition {
      condition     = length(self.remote_systems) == 1
      error_message = "Expected one Remote System with name RT-D4538, but got ${length(self.remote_systems)}."
    }
  }
}

# 2. create the replication rule
resource "powerstore_replication_rule" "backup_1hr" {
  name             = "RT_D4538_1hr"
  rpo              = "One_Hour"
  remote_system_id = data.powerstore_remote_system.backup_1hr.remote_systems[0].id
  alert_threshold  = 1000
  is_read_only     = false
}

//Below example is for import operation
/*resource "powerstore_replication_rule" "terraform-provider-test-import" {
}*/