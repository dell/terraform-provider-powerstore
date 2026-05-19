/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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
# This resource manages the recycle bin configuration and item recovery/deletion
# It has three modes of operation:
# 1. Config mode: Set expiration_duration (0-30 days)
# 2. Item action mode: Recover or delete items by resource_id or resource_name
# 3. Empty mode: Empty the entire recycle bin

# Example: Set recycle bin expiration duration to 7 days
resource "powerstore_recycle_bin_config" "example" {
  expiration_duration = 7
}
