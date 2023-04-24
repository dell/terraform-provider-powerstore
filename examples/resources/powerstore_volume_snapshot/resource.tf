# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check volume_snapshot_import.tf for more info
# volume_id/volume_name is the required attribute to create and update
# name, expiration_timestamp, performance_policy_id and creator_type are the optional attributes
# name and expiration_timestamp if present in config cannot be ""
# Either volume_id or volume_name should be present.
# To check which attributes of the volume snapshot resource can be updated, please refer Product Guide in the documentation

resource "powerstore_volume_snapshot" "test" {
  name = "test_snap"
  volume_id="01d88dea-7d71-4a1b-abd6-be07f94aecd9"
  performance_policy_id = "default_medium"
  expiration_timestamp="2023-05-06T09:01:47Z"
  creator_type="User"
}