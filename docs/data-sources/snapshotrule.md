---
# Copyright (c) 2023-2025 Dell Inc., or its subsidiaries. All Rights Reserved.
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

title: "powerstore_snapshotrule data source"
linkTitle: "powerstore_snapshotrule"
page_title: "powerstore_snapshotrule Data Source - powerstore"
subcategory: "Data Protection Management"
description: |-
  This datasource is used to query the existing snapshot rule from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.
---

# powerstore_snapshotrule (Data Source)

This datasource is used to query the existing snapshot rule from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.

> **Note:** Only one of `name` or `id` can be provided at a time.

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

# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads Snapshot Rules either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the Snapshot Rules
# If id or name is provided then it reads a particular Snapshot Rule with that id or name
# Only one of the attribute can be provided among id and  name 

# Fetching snapshot rule using filter expression
# This filter expression will fetch the snapshot rule where name is `snapshotrule-ny`
data "powerstore_snapshotrule" "test1" {
  filter_expression = "name=eq.snapshotrule-ny"
}


data "powerstore_snapshotrule" "test1" {
  name = "test_snapshotrule_1"
}

# Output all Snapshot Rules Details
output "snapshotRules_all_details" {
  value = data.powerstore_snapshotrule.test1.snapshot_rules
}

# Output only Snapshot Rule IDs
output "snapshot_rules_IDs_only" {
  value = data.powerstore_snapshotrule.test1.snapshot_rules.*.id
}

# Output Snapshot Rule names and timezone with Snapshot Rule id as key
output "snapshot_rule_name_and_timezone" {
  value = {
    for rule in data.powerstore_snapshotrule.test1.snapshot_rules : rule.id => {
      name = rule.name
      timezone = rule.timezone
    }
  }
}
```

After the successful execution of above said block, We can see the output by executing `terraform output` command. Also, we can fetch information via the variable: `data.powerstore_replication_rule.test1.attribute_name` where attribute_name is the attribute which user wants to fetch.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter_expression` (String) PowerStore filter expression to filter Host by. Conflicts with `id` and `name`.
- `id` (String) Unique identifier of the snapshot rule instance. Conflicts with `name`.
- `name` (String) Name of the snapshot rule. Conflicts with `id`.

### Read-Only

- `snapshot_rules` (Attributes List) List of snapshot rules. (see [below for nested schema](#nestedatt--snapshot_rules))

<a id="nestedatt--snapshot_rules"></a>
### Nested Schema for `snapshot_rules`

Optional:

- `days_of_week` (List of String) The days of the week when the snapshot rule should be applied.

Read-Only:

- `days_of_week_l10n` (List of String) Localized message array corresponding to days_of_week
- `desired_retention` (Number) The Desired snapshot retention period in hours to retain snapshots for this time period.
- `id` (String) The ID of the snapshot rule.
- `interval` (String) The interval of the snapshot rule.
- `interval_l10n` (String) Localized message string corresponding to interval
- `is_read_only` (Boolean) Indicates whether this snapshot rule can be modified.
- `is_replica` (Boolean) Indicates whether this is a replica of a snapshot rule on a remote system.
- `managed_by` (String) The entity that owns and manages the instance.
- `managed_by_id` (String) The unique id of the managing entity.
- `managed_by_l10n` (String) Localized message string corresponding to managed_by.
- `name` (String) Name of the snapshot rule.
- `nas_access_type` (String) The NAS filesystem snapshot access method for snapshot rule.
- `nas_access_type_l10n` (String) Localized message string corresponding to nas_access_type.
- `policies` (Attributes List) List of the protection policies that are associated with the snapshot_rule.. (see [below for nested schema](#nestedatt--snapshot_rules--policies))
- `time_of_day` (String) The time of the day to take a daily snapshot, with format hh:mm.
- `timezone` (String) The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.
- `timezone_l10n` (String) Localized message string corresponding to timezone.

<a id="nestedatt--snapshot_rules--policies"></a>
### Nested Schema for `snapshot_rules.policies`

Read-Only:

- `description` (String) Description of the protection policy.
- `id` (String) Unique identifier of the protection policy.
- `name` (String) Name of the protection policy.
