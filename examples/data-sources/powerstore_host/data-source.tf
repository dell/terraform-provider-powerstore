# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads hosts either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the hosts
# If id or name is provided then it reads a particular host with that id or name
# Only one of the attribute can be provided among id and name 

data "powerstore_host" "test1" {
  name = "tf_host"
}

output "hostResult" {
  value = data.powerstore_host.test1.hosts
}
