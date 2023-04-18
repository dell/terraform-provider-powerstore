# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads protection policies either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the protection policies
# If id or name is provided then it reads a particular protection policy with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_protectionpolicy" "test1" {
  name = "terraform_protection_policy_2"
}

output "policyResult" {
  value = data.powerstore_protectionpolicy.test1.policies
}
