/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads QoS policies either by id or name where user can provide a value to any one of them
# If it is a empty datasource block , then it will read all the QoS policies
# If id or name is provided then it reads a particular QoS policy with that id or name
# Only one of the attribute can be provided among id, name, and type

data "powerstore_qos_policy" "test1" {
  name = "tf_acc_qos_policy"
}

# Output all QoS Policy Details
output "qos_policy_all_details" {
  value = data.powerstore_qos_policy.test1.qos_policies
}

# Output only QoS Policy IDs
output "qos_policy_IDs_only" {
  value = data.powerstore_qos_policy.test1.qos_policies.*.id
}

# Output QoS Policy IDs and type with name as key
output "qos_policy_id_and_type" {
  value = {
    for policy in data.powerstore_qos_policy.test1.qos_policies : policy.name => {
      id   = policy.id
      type = policy.type
    }
  }
}
