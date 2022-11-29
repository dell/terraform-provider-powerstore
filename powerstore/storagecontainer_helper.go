package powerstore

import (
	pstore "github.com/dell/gopowerstore"
	"terraform-provider-powerstore/models"
)

func updateStorageContainerState(scState *models.StorageContainer, scResponse pstore.StorageContainer, plan *models.StorageContainer, operation string) {
	scState.ID.Value = scResponse.ID
	scState.Name.Value = scResponse.Name
	scState.Quota.Value = scResponse.Quota
	scState.StorageProtocol.Value = string(scResponse.StorageProtocol)
	if operation == "Create" {
		scState.HighWaterMark.Value = plan.HighWaterMark.Value
	}
}
