# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check host_group_import.tf for more info
# name and host_ids are the required attributes to create and update
# description is the optional attribute
# Host datasource can be used to fetch host id/name.
# To check which attributes of the host resource can be updated, please refer Product Guide in the documentation

resource "powerstore_hostgroup" "test" {
  name        = "test_hostgroup"
  description = "Creating host group"
  host_ids    = ["42c60954-ea71-4b50-b172-63880cd48f99"]
}
