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

# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads hosts either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the hosts
# If id or name is provided then it reads a particular host with that id or name
# Only one of the attribute can be provided among id and name

# Fetching host using filter expression
# This filter expression will fetch all the hosts where `name` is `host-ny`
data "powerstore_host" "host" {
  filter_expression = "name=eq.host-ny"
}

data "powerstore_host" "test1" {
  name = "tf_host"
}

data "powerstore_host" "all"{
}

# Output all host details
output "host_all_details" {
  value = data.powerstore_host.all.host
}

# Output only host IDs
output "host_IDs_only" {
  value = data.powerstore_host.all.host.*.id
}

# Ouptput host name and os type with host id as key
output "host_name_and_os_type" {
    value = {
    for host in data.powerstore_host.all.host : host.id => {
      name = host.name
      os_type = host.os_type
    }
  }
}