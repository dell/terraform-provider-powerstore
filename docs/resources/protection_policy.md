
# powerstore_protection_policy (Resource)

Use this resource type to manage protection policies and to view information about performance
policies.

Note: Performance policies are predefined for high, low, and medium performance. They cannot be
added to or changed.

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **description** (String) Policy description.
- **file_systems** (Block List) (see [below for nested schema](#nestedblock--file_systems)) This is the inverse of the resource type file_system
association.
- **id** (String) The ID of this resource.
- **is_replica** (Boolean) Indicates whether this is a replica policy, which is applied
to replication destination storage resources. A policy of this
type is restricted from many operations. Default : false
- **name** (String) Policy name. This property supports case-insensitive
filtering
- **policy_type** (String) Supported policy types.
Values : Protection, Performance.
Note: Terraform provider supports Protection policy only.
- **replication_rules** (Block List) (see [below for nested schema](#nestedblock--replication_rules)) List of the replication_rules that are associated with this
policy.
- **resource_id** (String) The ID of this resource.
- **snapshot_rules** (Block List) (see [below for nested schema](#nestedblock--snapshot_rules)) List of the snapshot_rules that are associated with this
policy.
- **type_l10n** (String) Localized message string corresponding to type
- **virtual_machines** (Block List) (see [below for nested schema](#nestedblock--virtual_machines)) This is the inverse of the resource type virtual_machine
association.
- **volumes** (Block List) (see [below for nested schema](#nestedblock--volumes)) This is the inverse of the resource type volume association.
- **volume_groups** (Block List) (see [below for nested schema](#nestedblock--volume_groups)) This is the inverse of the resource type volume_group
association.

<a id="nestedblock--file_systems"></a>

### Nested Schema for `file_systems`

Optional:

- **id** (String) The ID of this resource.
- **name** (String) Name of the file system


<a id="nestedblock--replication_rules"></a>

### Nested Schema for `replication_rules`

Optional:

- **id** (String) The ID of this resource.
- **name** (String) Name of the replication rules


<a id="nestedblock--snapshot_rules"></a>

### Nested Schema for `snapshot_rules`

Optional:

- **days_of_week** (List of String) Days of the week when the rule should be applied. Applies
only for rules where the time_of_day parameter is set.
Values : Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
- **id** (String) The ID of this resource.
- **interval** (String) Interval between snapshots. Either the interval parameter or the time_of_day parameter may be set. Setting one clears the other parameter.
Values : Five_Minutes, Fifteen_Minutes, Thirty_Minutes, One_Hour, Two_Hours, Three_Hours,
Four_Hours, Six_Hours, Eight_Hours, Twelve_Hours, One_Day
- **name** (String) Snapshot rule name. This property supports caseinsensitive filtering
- **time_of_day** (String) Time of the day to take a daily snapshot, with format "hh:mm" in 24 hour time format. Either the interval parameter or the time_of_day parameter will be set, but not both.


<a id="nestedblock--virtual_machines"></a>

### Nested Schema for `virtual_machines`

Optional:

- **id** (String) The ID of this resource.
- **name** (String) Name of the virtual machine


<a id="nestedblock--volumes"></a>

### Nested Schema for `volumes`

Optional:

- **id** (String) The ID of this resource.
- **name** (String) Name of the volume


<a id="nestedblock--volume_groups"></a>

### Nested Schema for `volume_groups`

Optional:

- **id** (String) The ID of this resource.
- **name** (String) Name of the volume group
