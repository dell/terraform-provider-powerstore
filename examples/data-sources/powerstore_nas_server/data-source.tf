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
# This datasource reads NAS Servers either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the NAS Servers
# If id or name is provided then it reads a particular NAS Server with that id or name
# Only one of the attribute can be provided among id and name 

# Fetching all NAS Servers
data "powerstore_nas_server" "all" {
}

# Fetching NAS Server by id
data "powerstore_nas_server" "nas_server_by_id" {
  id = "282479293"
}

# Fetching NAS Server by name
data "powerstore_nas_server" "nas_server_by_name" {
  name = "nas_server_1"
}

# Get NAS Servers details using filter expression
# This filter expression will fetch all the NAS Servers that have `operational_status` as `Started` and 
# `is_replication_destination` as `false`
data "powerstore_nas_server" "nas_server_by_filter" {
  filter_expression = "and=(operational_status.eq.Started, is_replication_destination.eq.false)"
}

# Output all NAS Server details
output "nas_server_all_details" {
  value = data.powerstore_nas_server.all.nas_servers
}

# Output only NAS Server IDs
output "nas_server_IDs_only" {
  value = data.powerstore_nas_server.all.nas_servers.*.id
}

# Output NAS Server name and current node ID with nas server ID as key
output "nas_server_name_and_current_node_id" {
  value = {
    for nas_server in data.powerstore_nas_server.all.nas_servers : nas_server.id => {
      name = nas_server.name
      current_node_id = nas_server.current_node_id  
    }
  }
}