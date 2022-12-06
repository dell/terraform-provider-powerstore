---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "powerstore_snapshotrule Resource - powerstore"
subcategory: ""
description: |-
  
---

# powerstore_snapshotrule (Resource)



## Example Usage

```terraform
# powerstore_snapshotrule example

resource "powerstore_snapshotrule" "test" {
  name = "test_snapshotrule_1"
  # interval = "Four_Hours"
  time_of_day = "21:00"
  timezone = "UTC"
  days_of_week = ["Monday"]
  desired_retention = 56
  nas_access_type = "Snapshot"
  is_read_only = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `desired_retention` (Number) The Desired snapshot retention period in hours to retain snapshots for this time period.
- `name` (String) The name of the snapshot rule.

### Optional

- `days_of_week` (List of String) The days of the week when the snapshot rule should be applied.
- `interval` (String) The interval between snapshots taken by a snapshot rule.
- `is_read_only` (Boolean) Indicates whether this snapshot rule can be modified.
- `nas_access_type` (String) The NAS filesystem snapshot access method for snapshot rule.
- `time_of_day` (String) The time of the day to take a daily snapshot, with format hh:mm.
- `timezone` (String) The time zone identifier for applying the time zone to the time_of_day for a snapshot rule.

### Read-Only

- `id` (String) The ID of the snapshot rule.
- `is_replica` (Boolean) Indicates whether this is a replica of a snapshot rule on a remote system.
- `managed_by` (String) The entity that owns and manages the instance.
- `managed_by_id` (String) The unique id of the managing entity.

