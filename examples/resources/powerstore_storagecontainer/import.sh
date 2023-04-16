# Below are the steps to import storage container :
# Step 1 - To import a storage container , we need the id of that storage container 
# Step 2 - To check the id of the storage container we can make Get request to storage container endpoint. eg. https://10.0.0.1/api/rest/storage_container which will return list of all storage container ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_storagecontainer" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_storagecontainer.resource_block_name" "id_of_the_storage_container" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file