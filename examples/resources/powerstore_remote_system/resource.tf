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

# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import, check import.sh for more info
# management_address is the required attribute to create
# description and data_network_latency are the optional attributes
# For PowerStore-to-PowerStore connections, only management_address and data_network_latency are required
# For non-PowerStore remote systems, additional parameters like type, remote_username, remote_password are required

# Example: PowerStore-to-PowerStore remote system
resource "powerstore_remote_system" "ps_to_ps" {
  management_address   = "100.1.1.1"
  description          = "Remote PowerStore for replication"
  data_network_latency = "Low"
}
