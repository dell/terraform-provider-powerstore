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

# fetching all Remote Systems on the array
data "powerstore_remote_system" "all_remote_systems" {
}

# fetching Remote System using id
data "powerstore_remote_system" "remote_system_by_id" {
  id = "6732e829-29c9-7fed-686a-ee23cab1d298"
}

# fetching Remote System using name
data "powerstore_remote_system" "remote_system_by_name" {
  name = "RT-D4538"
}

# fetching Remote Systems using filter expression
# Please refer to the guides section for filter expression syntax
# here, we are fetching a Remote System with a particular management IP
data "powerstore_remote_system" "remote_system_by_filters" {
  filter_expression = "management_address=eq.10.225.225.10"
}

output "all_remote_systems" {
  value = data.powerstore_remote_system.all_remote_systems.remote_systems
}
