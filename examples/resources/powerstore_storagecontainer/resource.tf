# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check storage_container_import.tf for more info
# name is the required attribute to create and update
# To check which attributes of the storage container can be updated, please refer Product Guide in the documentation


resource "powerstore_storagecontainer" "test1" {
  name             = "scterraform1"
  quota            = 10737418240
  storage_protocol = "SCSI"
  high_water_mark  = 70
}