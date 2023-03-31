# Below are the steps to import volume group :
# Step 1 - To import a volume group , we need the id of that volume group 
# Step 2 - To check the id of the volume group we can make use of host group datasource to read required/all host group ids. Alternatively, we can make Get request to volume group endpoint. eg. https://10.0.0.1/api/rest/policy which will return list of all volume group ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_volumegroup" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_volumegroup.resource_block_name" "id_of_the_volume_group" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
