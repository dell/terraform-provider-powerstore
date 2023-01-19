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
  delete_snaps = true
}