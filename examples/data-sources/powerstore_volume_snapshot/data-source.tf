# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads volume snapshots either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volume snapshots
# If id or name is provided then it reads a particular volume snapshot with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_volume_snapshot" "test1" {
  name = "test_snap"
  #id = "adeeef05-aa68-4c17-b2d0-12c4a8e69176"
}

output "volumeSnapshotResult" {
  value = data.powerstore_volume_snapshot.test1.volumes
}
