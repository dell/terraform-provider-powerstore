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

# Example: PowerStore-to-PowerStore remote system (TCP)
resource "powerstore_remote_system" "ps_to_ps" {
  management_address   = "100.1.1.1"
  description          = "Remote PowerStore for replication"
  data_network_latency = "Low"
}

# Example: Universal FC remote system
resource "powerstore_remote_system" "fc_universal" {
  management_address   = "100.2.2.2"
  description          = "Universal FC remote system for block replication"
  type                 = "Universal"
  data_connection_type = "FC"
  data_network_latency = "Low"
  universal_details = {
    fc_targets = [
      {
        wwnn = "58:cc:f0:98:49:21:07:00"
        wwpn = "58:cc:f0:98:49:21:07:01"
      },
      {
        wwnn = "58:cc:f0:98:49:21:07:00"
        wwpn = "58:cc:f0:98:49:21:07:02"
      }
    ]
  }
}
