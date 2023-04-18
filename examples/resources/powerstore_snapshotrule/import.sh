# Below are the steps to import snapshot rule :
# Step 1 - To import a snapshot rule , we need the id of that snapshot rule 
# Step 2 - To check the id of the snapshot rule we can make Get request to snapshot rule endpoint. eg. https://10.0.0.1/api/rest/snapshot_rule which will return list of all snapshot rule ids.
# Step 3 - Add empty resource block in tf file. 
# eg. 
# resource "powerstore_snapshotrule" "resource_block_name" {
  # (resource arguments)
# }
# Step 4 - Execute the command: terraform import "powerstore_snapshotrule.resource_block_name" "id_of_the_snapshot_rule" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file