# Below are the steps to import volume group snapshot :
# Step 1 - To import a volume group snapshot , we need the id of that volume group snapshot
# Step 2 - To check the id of the volume group snapshot we can make Get request to volume group snapshot endpoint. eg. https://10.0.0.1/api/rest/volume_group --header 'type: Snapshot' which will return list of all volume group snapshots ids.
# Step 3 - Add empty resource block in tf file.
# eg.
#resource "powerstore_volumeGroupSnapshot" "resource_block_name" {
  # (resource arguments)
#}
# Step 4 - Execute the command: terraform import "powerstore_volumeGroupSnapshot.resource_block_name" "id_of_the_snapshot" (resource_block_name must be taken from step 3 and id must be taken from step 2)
# Step 5 - After successful execution of the command , check the state file
