/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
# To import , check storage_container_import.tf for more info
# name is the required attribute to create and update
# To check which attributes of the storage container can be updated, please refer Product Guide in the documentation


resource "powerstore_storagecontainer" "test1" {
  name             = "scterraform1"
  quota            = 10737418240
  storage_protocol = "SCSI"
  high_water_mark  = 70
}