# Below are the steps to import protection policy :
# Step 1 - To import a protection policy , we need the id of that protection policy 
# Step 2 - To check the id of the protection policy we can make Get request to protection policy endpoint. eg. https://10.0.0.1/api/rest/policy which will return list of all protection policy ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_protectionpolicy" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_protectionpolicy.resource_block_name" "id_of_the_protection_policy" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file