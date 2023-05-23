# Commands to run this tf file : terraform init && terraform plan && terraform apply
# Create, Update, Delete is supported for this resource
# To import , check host_import.tf for more info
# name, os_type and initiators are the required attributes to create and update
# description and host_connectivity are the optional attributes
# To check which attributes of the host resource can be updated, please refer Product Guide in the documentation

resource "powerstore_host" "test" {
  name              = "new-host1"
  os_type           = "Linux"
  description       = "Creating host"
  host_connectivity = "Local_Only"
  initiators        = [{ port_name = "iqn.1994-05.com.redhat:88cb605", port_type = "iSCSI" }]
}
