# Below are the steps to import host :
# Step 1 - To import a host , we need the id of that host
# Step 2 - To check the id of the host we can make Get request to host endpoint. eg. https://10.0.0.1/api/rest/host which will return list of all host ids.
# Step 3 - Add empty resource block in tf file.
# eg.
# resource "powerstore_host" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_host.resource_block_name" "id_of_the_host" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
