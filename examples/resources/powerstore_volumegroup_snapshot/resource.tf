# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check volumegroup_snapshot_import.tf for more info
# name and volume_group_id/volume_group_name are the required attributes to create and update.
# description and expiration_timestamp are the optional attributes.
# Either volume_group_id or volume_group_name should be present.
# VolumeGroup DataSource can be used to fetch volume group ID/Name
# expiration_timestamp if present in config cannot be ""
# To check which attributes of the volume group snapshot resource can be updated, please refer Product Guide in the documentation

resource "powerstore_volumegroup_snapshot" "test" {
  name                 = "test_snap"
  volume_group_id      = "075aeb23-c782-4cce-9372-5a2e31dc5138"
  expiration_timestamp = "2023-05-06T09:01:47Z"
}
