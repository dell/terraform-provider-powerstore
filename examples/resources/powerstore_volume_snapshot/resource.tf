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
# To import , check volume_snapshot_import.tf for more info
# volume_id/volume_name is the required attribute to create volume snapshot.
# name, expiration_timestamp, performance_policy_id and description are the optional attributes
# if name is present in the config it cannot be blank("").
# During create operation, if expiration_timestamp is not specified or set to blank(""), snapshot will be created with infinite retention.
# During modify operation, to set infinite retention, expiration_timestamp can be set to blank("").
# Either volume_id or volume_name should be present.
# Volume DataSource can be used to fetch volume ID/Name
# To check which attributes of the volume snapshot resource can be updated, please refer Product Guide in the documentation

resource "powerstore_volume_snapshot" "test" {
  name                  = "test_snap"
  description           = "powerstore volume snapshot"
  volume_id             = "01d88dea-7d71-4a1b-abd6-be07f94aecd9"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "2023-05-06T09:01:47Z"
}
