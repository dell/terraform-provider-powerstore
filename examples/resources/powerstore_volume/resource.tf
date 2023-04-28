# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check volume_import.tf for more info
# name and size are the required attributes to create and update
# To check which attributes of the volume can be updated, please refer Product Guide in the documentation


resource "powerstore_volume" "test1" {
  name                  = "test_vol1"
  size                  = 3
  capacity_unit         = "GB"
  description           = "Creating volume"
  host_id               = ""
  host_group_id         = ""
  appliance_id          = "A1"
  volume_group_id       = ""
  min_size              = 1048576
  sector_size           = 512
  protection_policy_id  = ""
  performance_policy_id = "default_medium"
  app_type              = "Relational_Databases_Other"
  app_type_other        = ""
}
