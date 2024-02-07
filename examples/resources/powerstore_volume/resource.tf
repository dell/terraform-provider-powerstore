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
# To import , check volume_import.tf for more info
# name and size are the required attributes to create and update
# To check which attributes of the volume can be updated, please refer Product Guide in the documentation


resource "powerstore_volume" "test1" {
  name                  = "test_vol1"
  size                  = 3
  capacity_unit         = "GB"
  description           = "Creating volume"
  host_id               = ""
  host_group_id         = ""
  appliance_id          = "A1"
  volume_group_id       = ""
  min_size              = 1048576
  sector_size           = 512
  protection_policy_id  = ""
  performance_policy_id = "default_medium"
  app_type              = "Relational_Databases_Other"
  app_type_other        = ""
}
