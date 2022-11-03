package powerstore

import (
	"context"
	pstore "github.com/dell/gopowerstore"
	"terraform-provider-powerstore/models"
)

func updateVolState(volState *models.Volume, volResponse pstore.Volume, hostMapping []pstore.HostVolumeMapping, volPlan *models.Volume, operation string) {
	// Update value from Volume Response to State
	volState.ID.Value = volResponse.ID
	volState.Name.Value = volResponse.Name
	volState.Size.Value = volResponse.Size
	volState.Type.Value = string(volResponse.Type)
	volState.WWN.Value = volResponse.Wwn
	volState.Description.Value = volResponse.Description
	volState.State.Value = string(volResponse.State)
	volState.WWN.Value = volResponse.Wwn
	volState.ApplianceID.Value = volResponse.ApplianceID
	volState.CreationTimeStamp.Value = volResponse.CreationTimeStamp
	volState.ProtectionPolicyID.Value = volResponse.ProtectionPolicyID
	volState.PerformancePolicyID.Value = volResponse.PerformancePolicyID

	// Only if Host is mapped to the volume then update host details
	if len(hostMapping) > 0 {
		volState.HostID.Value = hostMapping[0].HostID
		volState.HostGroupID.Value = hostMapping[0].HostGroupID
		volState.LogicalUnitNumber.Value = hostMapping[0].LogicalUnitNumber
	} else {
		volState.HostID.Value = ""
		volState.HostGroupID.Value = ""
		volState.LogicalUnitNumber.Value = 0
	}

	if operation == "Create" || operation == "Update" {
		volState.VolumeGroupID.Value = volPlan.VolumeGroupID.Value
		volState.MinimumSize.Value = volPlan.MinimumSize.Value
		volState.SectorSize.Value = volPlan.SectorSize.Value
	}

	volState.AppType.Value = volResponse.AppType
	volState.AppTypeOther.Value = volResponse.AppTypeOther
	volState.IsReplicationDestination.Value = volResponse.IsReplicationDestination
	volState.NodeAffinity.Value = volResponse.NodeAffinity
	volState.LogicalUsed.Value = volResponse.LogicalUsed
	volState.Nsid.Value = volResponse.Nsid
	volState.Nguid.Value = volResponse.Nguid
}

func creationValidation(ctx context.Context, plan models.Volume) (bool, string) {
	if plan.HostID.Value != "" && plan.HostGroupID.Value != "" {
		return false, "Either HostID or HostGroupID can be present "
	}
	if plan.PerformancePolicyID.Value == "" {
		return false, "Performance Policy if present cannot be empty. Either remove the field or set desired value"
	}
	return true, ""
}
