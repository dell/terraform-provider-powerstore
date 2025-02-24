#Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.
#
#Licensed under the Mozilla Public License Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://mozilla.org/MPL/2.0/
#
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.


# Below are the steps to import file system :
# Step 1 - To import a file system , we need the id of that file system 
# Step 2 - To check the id of the file system we can make use of file system datasource to read required/all file system ids. Alternatively, we can make GET request to file system endpoint. eg. https://10.0.0.1/api/rest/file_system which will return list of all file system ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_filesystem" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_filesystem.resource_block_name" "id_of_the_file_system" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
