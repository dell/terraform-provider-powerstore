# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads volume groups either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volume group snapshots
# If id or name is provided then it reads a particular volume group snapshot with that id or name
# Only one of the attribute can be provided among id and name

data "powerstore_volumegroup_snapshot" "test1" {
  name = "test_volumegroup_snap"
}

output "volumeGroupSnapshotResult" {
  value = data.powerstore_volumegroup_snapshot.test1.volume_groups
}
