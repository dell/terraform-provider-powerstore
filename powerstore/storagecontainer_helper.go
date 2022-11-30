package powerstore

import (
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	pstore "github.com/dell/gopowerstore"
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

func (r resourceStorageContainer) updateRequestPayload(plan, state models.StorageContainer) *gopowerstore.StorageContainer {

	// a workaround
	// currently PowerStore not accepting PATCH call for same values
	// so sending only updated values

	storageContainerUpdate := &gopowerstore.StorageContainer{}

	if plan.Name.Value != state.Name.Value {
		storageContainerUpdate.Name = plan.Name.Value
	}

	if plan.Quota.Value != state.Quota.Value {
		storageContainerUpdate.Quota = plan.Quota.Value
	}

	if plan.StorageProtocol.Value != state.StorageProtocol.Value {
		storageContainerUpdate.StorageProtocol = gopowerstore.StorageContainerStorageProtocolEnum(plan.StorageProtocol.Value)
	}

	return storageContainerUpdate
}
