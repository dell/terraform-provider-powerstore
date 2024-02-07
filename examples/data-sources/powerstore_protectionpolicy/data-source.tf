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
# This datasource reads protection policies either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the protection policies
# If id or name is provided then it reads a particular protection policy with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_protectionpolicy" "test1" {
  name = "terraform_protection_policy_2"
}

output "policyResult" {
  value = data.powerstore_protectionpolicy.test1.policies
}
