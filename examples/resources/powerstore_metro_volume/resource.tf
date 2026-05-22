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
# This resource configures metro (synchronous) replication on an existing volume.
# Metro replication provides RPO=0 (zero data loss) between two PowerStore clusters.
# The volume must already exist - use powerstore_volume resource to create it first.

# Example: Configure metro replication on an existing volume
resource "powerstore_metro_volume" "example" {
  volume_id        = "volume-id-here"
  remote_system_id = "remote-system-id-here"

  # Optional: specify a remote appliance
  # remote_appliance_id = "remote-appliance-id"

  # Optional: control destroy behavior
  # delete_remote_volume = false
  # force                = false
}
