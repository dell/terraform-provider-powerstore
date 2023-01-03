package storagecontainer

import (
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m *model) saveSeverResponse(response gopowerstore.StorageContainer) {

	m.ID = types.StringValue(response.ID)
	m.Name = types.StringValue(response.Name)
	m.Quota = types.Int64Value(response.Quota)
	m.StorageProtocol = types.StringValue(string(response.StorageProtocol))
	m.HighWaterMark = types.Int64Value(int64(response.HighWaterMark))
}

func (m *model) serverRequest(state model) *gopowerstore.StorageContainer {

	// a workaround
	// currently PowerStore not accepting PATCH call for same values
	// so sending only updated values

	storageContainerUpdate := &gopowerstore.StorageContainer{}

	if m.Name.ValueString() != state.Name.ValueString() {
		storageContainerUpdate.Name = m.Name.ValueString()
	}

	if m.Quota.ValueInt64() != state.Quota.ValueInt64() {
		storageContainerUpdate.Quota = m.Quota.ValueInt64()
	}

	if m.StorageProtocol.ValueString() != state.StorageProtocol.ValueString() {
		storageContainerUpdate.StorageProtocol = gopowerstore.StorageContainerStorageProtocolEnum(m.StorageProtocol.ValueString())
	}

	if m.HighWaterMark.ValueInt64() != state.HighWaterMark.ValueInt64() {
		storageContainerUpdate.HighWaterMark = int16(m.HighWaterMark.ValueInt64())
	}

	return storageContainerUpdate
}
