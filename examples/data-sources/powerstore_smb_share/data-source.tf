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

# fetching all SMB Shares on the array
data "powerstore_smb_share" "all_smb_shares" {
}

# fetching SMB Share using id
data "powerstore_smb_share" "smb_share_by_id" {
  id = "6732e829-29c9-7fed-686a-ee23cab1d298"
}

# fetching SMB Shares using name
data "powerstore_smb_share" "smb_share_by_name" {
  name = "smb-share-1"
}

# fetching all SMB Shares of a filesystem
# Step 1: Fetching the filesystem whose SMB Shares are to be fetched.
data "powerstore_filesystem" "us_east_sales_catalog" {
  name = "us-east-sales-catalog"
  lifecycle {
    postcondition {
      condition     = length(self.filesystems) == 1
      error_message = "error: US East sales catalog filesystem list length should be 1, received: ${length(self.filesystems)}"
    }
  }
}

#  Step 2: Fetching the SMB Shares using the filesystem id from step 1
data "powerstore_smb_share" "smb_share_by_filesystem" {
  file_system_id = data.powerstore_filesystem.us_east_sales_catalog.filesystems[0].id
}

# Fetching SMB Shares using filter expression
# This filter expression will fetch all SMB Shares of subdirectories of /us-east-revenue/sports_cars
# with encryption enabled and offline availability as either Documents or None.
data "powerstore_smb_share" "smb_share_by_filters" {
  filter_expression = "path=ilike./us-east-revenue/sports_cars/*&is_encryption_enabled=is.true&offline_availability=in.(Documents,None)"
}

# Output all SMB Shares Details
output "smb_shares_all_details" {
  value = data.powerstore_smb_share.all_smb_shares.smb_shares
}

# Output only Replication share IDs
output "smb_shares_IDs_only" {
  value = data.powerstore_smb_share.all_smb_shares.smb_shares.*.id
}

# Output Replication share IDs and policies with Replication share ID as key
output "smb_shares_id_and_policies" {
  value = {
    for share in data.powerstore_smb_share.all_smb_shares.smb_shares : share.id => {
      file_system_id = share.file_system_id
      path = share.path
    }
  }
}