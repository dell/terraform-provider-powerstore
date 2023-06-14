/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
# To import , check snapshot_rule_import.tf for more info
# name and interval OR name, time_of_day, days_of_week and timezone are required attributes to create and update
# To check which attributes of the snapshot rule can be updated, please refer Product Guide in the documentation


resource "powerstore_snapshotrule" "test1" {
  name = "test_snapshotrule_2"
  # interval = "Four_Hours"
  time_of_day       = "21:00"
  timezone          = "UTC"
  days_of_week      = ["Monday"]
  desired_retention = 56
  nas_access_type   = "Snapshot"
  is_read_only      = false
  delete_snaps      = true
}

//Below example is for import operation
/*resource "powerstore_snapshotrule" "terraform-provider-test-import" {
}*/