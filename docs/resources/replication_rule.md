---
# Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.
# 
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://mozilla.org/MPL/2.0/
# 
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "powerstore_replication_rule resource"
linkTitle: "powerstore_replication_rule"
page_title: "powerstore_replication_rule Resource - powerstore"
subcategory: ""
description: |-
  This resource is used to manage the replication rule entity of PowerStore Array. We can Create, Update and Delete the replication rule using this resource. We can also import an existing replication rule from PowerStore array.
---

# powerstore_replication_rule (Resource)

This resource is used to manage the replication rule entity of PowerStore Array. We can Create, Update and Delete the replication rule using this resource. We can also import an existing replication rule from PowerStore array.

## Example Usage

```terraform
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
# To import, check import.tf for more info
# name, rpo and remote_system_id are required attributes to create and update
# To check which attributes of the replication rule can be updated, please refer Product Guide in the documentation

resource "powerstore_replication_rule" "test" {
  name             = "terraform_replication_rule"
  rpo              = "One_Hour"
  remote_system_id = "db11abb3-789e-47f9-96b5-84b5374cbcd2"
  alert_threshold  = 1000
  is_read_only     = false
}

//Below example is for import operation
/*resource "powerstore_snapshotrule" "terraform-provider-test-import" {
}*/
```

After the execution of above resource block replication rule would have been created on the PowerStore array. For more information, Please check the terraform state file.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the replication rule.
- `remote_system_id` (String) Unique identifier of the remote system associated with the replication rule.
- `rpo` (String) Recovery Point Objective (RPO) of the replication rule.

### Optional

- `alert_threshold` (Number) Alert threshold for the replication rule.
- `is_read_only` (Boolean) Indicates whether the replication rule is read-only.

### Read-Only

- `id` (String) The ID of the replication rule.

## Import

Import is supported using the following syntax:

```shell
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


# Below are the steps to import replication rule :
# Step 1 - To import a replication rule, we need the id of that replication rule 
# Step 2 - To check the id of the replication rule, we can make Get request to replication rule endpoint. eg. https://10.0.0.1/api/rest/replication_rule which will return list of all replication rule ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_replication_rule" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_replication_rule.resource_block_name" "id_of_the_replication_rule" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
``` 