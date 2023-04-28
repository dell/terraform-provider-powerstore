# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check snapshot_rule_import.tf for more info
# name and interval OR name, time_of_day, days_of_week and timezone are required attributes to create and update
# To check which attributes of the snapshot rule can be updated, please refer Product Guide in the documentation


resource "powerstore_snapshotrule" "test1" {
  name = "test_snapshotrule_2"
  # interval = "Four_Hours"
  time_of_day       = "21:00"
  timezone          = "UTC"
  days_of_week      = ["Monday"]
  desired_retention = 56
  nas_access_type   = "Snapshot"
  is_read_only      = false
  delete_snaps      = true
}

//Below example is for import operation
/*resource "powerstore_snapshotrule" "terraform-provider-test-import" {
}*/