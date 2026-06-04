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
# This datasource reads file I/O limit rules either by id or name where user can provide a value to any one of them
# If it is a empty datasource block , then it will read all the file I/O limit rules
# If id or name is provided then it reads a particular file I/O limit rule with that id or name
# Only one of the attribute can be provided among id and name

data "powerstore_file_io_limit_rule" "test1" {
  name = "tf_acc_file_io_limit_rule"
}

# Output all File I/O Limit Rule Details
output "file_io_limit_rule_all_details" {
  value = data.powerstore_file_io_limit_rule.test1.file_io_limit_rules
}

# Output only File I/O Limit Rule IDs
output "file_io_limit_rule_IDs_only" {
  value = data.powerstore_file_io_limit_rule.test1.file_io_limit_rules.*.id
}

# Output File I/O Limit Rule IDs and max_bw with name as key
output "file_io_limit_rule_id_and_max_bw" {
  value = {
    for rule in data.powerstore_file_io_limit_rule.test1.file_io_limit_rules : rule.name => {
      id     = rule.id
      max_bw = rule.max_bw
    }
  }
}
