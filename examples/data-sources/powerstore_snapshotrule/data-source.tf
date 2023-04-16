# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads Snapshot Rules either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the Snapshot Rules
# If id or name is provided then it reads a particular Snapshot Rule with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_snapshotrule" "test1" {
  name = "test_snapshotrule_1"
}

output "snapshotRule" {
  value = data.powerstore_snapshotrule.test1.snapshot_rules
}
