# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads host groups either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the host groups
# If id or name is provided then it reads a particular host group with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_hostgroup" "test1" {
  name = "test_hostgroup1"
}

output "hostGroupResult" {
  value = data.powerstore_hostgroup.test1.host_groups
}
