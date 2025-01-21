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
# This datasource reads hosts either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the filesystem 
# If id or name is provided then it reads a particular file system snapshot with that id or name
# If filesystem_id is provided then it will read all the filesystem snapshots within filesystem
# Only one of the attribute can be provided among id or name 

#Fetching filesystem snapshot using name
data "powerstore_filesystem_snapshot" "test1" {
  name = "co015nap5hot"
}

#Fetching filesystem snapshot using id
data "powerstore_filesystem_snapshot" "test2" {
  id = "6568282e-c982-62ce-5ar3-52518f324723"
}

#Fetching filesystem snapshot using filesystem id
data "powerstore_filesystem_snapshot" "test2" {
  filesystem_id = "65637292e-c982-62ce-5ar3-52518f44229"
}

#Fetching filesystem snapshot using nas server id
data "powerstore_filesystem_snapshot" "test3" {
  nas_server_id = "654b2182-f674-f39a-66fc-52518d324736"
}

#Fetching filesystem snapshot using name and nas server id
data "powerstore_filesystem_snapshot" "test4" {
  name          = "co015nap5hot"
  nas_server_id = "654b2182-f674-f39a-66fc-52518d324736"
}

#Fetching filesystem snapshot using name and file system id
data "powerstore_filesystem_snapshot" "test4" {
  name          = "co015nap5hot"
  filesystem_id = "65637292e-c982-62ce-5ar3-52518f44229"
}


# Fetching all filesystems
data "powerstore_filesystem_snapshot" "test4" {
}


output "result" {
  value = data.powerstore_filesystem_snapshot.test3.filesystems
}
