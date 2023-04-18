# Below are the steps to import host group :
# Step 1 - To import a host group, we need the id of that host group
# Step 2 - To check the id of the host group we can make use of host group datasource to read required/all host group ids. Alternatively, we can make Get request to host group endpoint. eg. https://10.0.0.1/api/rest/host_group which will return list of all host group ids.
# Step 3 - Add empty resource block in tf file.
# eg.
# resource "powerstore_hostgroup" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_hostgroup.resource_block_name" "id_of_the_hostgroup" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
