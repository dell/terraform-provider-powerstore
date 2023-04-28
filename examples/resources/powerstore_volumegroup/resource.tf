# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check volume_group_import.tf for more info
# name is the required attribute to create and update
# To check which attributes of the volume group can be updated, please refer Product Guide in the documentation

resource "powerstore_volumegroup" "terraform-provider-test1" {
  # (resource arguments)
  description               = "Creating Volume Group"
  name                      = "test_volume_group"
  is_write_order_consistent = "false"
  protection_policy_id      = "01b8521d-26f5-479f-ac7d-3d8666097094"
  volume_ids                = ["140bb395-1d85-49ae-bde8-35070383bd92"]
}
