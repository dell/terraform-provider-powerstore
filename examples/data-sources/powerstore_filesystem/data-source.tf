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
# This datasource reads filesystems either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the filesystems
# If id or name is provided then it reads a particular filesystem with that id or name
# Only one of the attribute can be provided among id and name

#Fetching filesystem using name
data "powerstore_filesystem" "test1" {
  name = "ephemeral-csi-0e62d1cdec3b543a0740e9a307df221187955b47af4b99b7be31ebd9cb6192ce"
}

#Fetching filesystem using id
data "powerstore_filesystem" "test2" {
  id = "6568282e-c982-62ce-5ac3-52518d324736"
}

#Fetching filesystem using nas server id
data "powerstore_filesystem" "test3" {
  nas_server_id = "654b2182-f674-f39a-66fc-52518d324736"
}

# Fetching all filesystems
data "powerstore_filesystem" "test4" {
}


output "result" {
  value = data.powerstore_filesystem.test3.filesystems
}
