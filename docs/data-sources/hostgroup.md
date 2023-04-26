---
# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
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

title: "powerstore_hostgroup data source"
linkTitle: "powerstore_hostgroup"
page_title: "powerstore_hostgroup Data Source - powerstore"
subcategory: ""
description: |-
  HostGroup DataSource.
---

# powerstore_hostgroup (Data Source)

HostGroup DataSource.

## Example Usage

```terraform
# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads host groups either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the host groups
# If id or name is provided then it reads a particular host group with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_hostgroup" "test1" {
  name = "test_hostgroup1"
}

output "hostGroupResult" {
  value = data.powerstore_hostgroup.test1.host_groups
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) Unique identifier of the host group.
- `name` (String) Host group name.

### Read-Only

- `host_groups` (Attributes List) List of host groups. (see [below for nested schema](#nestedatt--host_groups))

<a id="nestedatt--host_groups"></a>
### Nested Schema for `host_groups`

Read-Only:

- `description` (String) Host group description.
- `host_connectivity` (String) Connectivity type for hosts and host groups.
- `host_connectivity_l10n` (String) Localized message string corresponding to host_connectivity
- `host_virtual_volume_mappings` (Attributes List) Virtual volume mapping details. (see [below for nested schema](#nestedatt--host_groups--host_virtual_volume_mappings))
- `hosts` (Attributes List) Properties of a host. (see [below for nested schema](#nestedatt--host_groups--hosts))
- `id` (String) Unique identifier of the host group.
- `mapped_host_groups` (Attributes List) Details about a configured host or host group attached to a volume. (see [below for nested schema](#nestedatt--host_groups--mapped_host_groups))
- `name` (String) Host group name.

<a id="nestedatt--host_groups--host_virtual_volume_mappings"></a>
### Nested Schema for `host_groups.host_virtual_volume_mappings`

Read-Only:

- `host_id` (String) Unique identifier of a host attached to a volume.
- `id` (String) Unique identifier of a mapping between a host and a virtual volume.
- `virtual_volume_id` (String) Unique identifier of the virtual volume to which the host is attached.


<a id="nestedatt--host_groups--hosts"></a>
### Nested Schema for `host_groups.hosts`

Read-Only:

- `host_group_id` (String) Associated host group, if host is part of host group.
- `id` (String) Unique identifier of the host.
- `name` (String) The host name.


<a id="nestedatt--host_groups--mapped_host_groups"></a>
### Nested Schema for `host_groups.mapped_host_groups`

Read-Only:

- `host_group_id` (String) Unique identifier of a host group attached to a volume.
- `host_id` (String) Unique identifier of a host attached to a volume.
- `id` (String) Unique identifier of a mapping between a host and a volume.