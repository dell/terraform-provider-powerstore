#Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.
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


# Below are the steps to import recycle bin config :
# Step 1 - The recycle bin config always has id "0"
# Step 2 - Add empty resource block in tf file.
# eg.
# resource "powerstore_recycle_bin_config" "resource_block_name" {
  # (resource arguments)
# }
# Step 3 - Execute the command: terraform import "powerstore_recycle_bin_config.resource_block_name" "0"
# Step 4 - After successful execution of the command, check the state file
terraform import "powerstore_recycle_bin_config.resource_block_name" "0"
