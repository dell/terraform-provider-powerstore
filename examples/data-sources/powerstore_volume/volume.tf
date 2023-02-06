# commands to run this tf file : terraform init && terraform apply --auto-approve
# This datasource reads volumes either by id or name where user can provide a value to any one of them
# If it is a empty datsource block , then it will read all the volumes
# If id or name is provided then it reads a particular volume with that id or name
# Only one of the attribute can be provided among id and  name 

data "powerstore_volume" "test1" {
        name = "tf_vol"
}

output "volumeResult" {
  value = data.powerstore_volume.test1.volumes
}
