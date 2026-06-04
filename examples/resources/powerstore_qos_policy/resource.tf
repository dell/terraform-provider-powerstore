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
# To import, check import.sh for more info
# name and type are required attributes to create
# type must be "QoS" for block storage or "File_Performance" for file storage
# io_limit_rule_id is optional for QoS type
# file_io_limit_rule_id is optional for File_Performance type

resource "powerstore_qos_policy" "terraform-provider-test1" {
  name        = "test_qos_policy1"
  description = "Test QoS policy for block storage"
  type        = "QoS"
  io_limit_rule_id = "643d8f3c-7b8c-4d1e-9f2a-1b3c4d5e6f7g"
}
