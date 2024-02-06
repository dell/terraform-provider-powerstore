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

# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check protection_policy_import.tf for more info
# name and snapshot_rule_ids or replication_rule_ids or snapshot_rule_names or replication_rule_names are required attributes to create and update
# To check which attributes of the protection policy can be updated, please refer Product Guide in the documentation

resource "powerstore_protectionpolicy" "terraform-provider-test1" {
  # (resource arguments)
  description            = "Creating Protection Policy"
  name                   = "test_protection_policy1"
  snapshot_rule_names    = ["vsi_aut_snaprule", "snapshot_test_emi", "test_snapshotrule_1", "snap-use-for-nfs-test"]
  replication_rule_names = ["Emalee-SRA-7416-Rep"]
}