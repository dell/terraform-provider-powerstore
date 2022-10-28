package powerstore

import (
	pstoreTypes "github.com/dell/gopowerstore"
	"terraform-provider-powerstore/models"
)

func updateVolState(volState *models.Volume, volResponse pstoreTypes.Volume, hostMapping []pstoreTypes.HostVolumeMapping, volPlan *models.Volume, operation string) {
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
	if len(hostMapping) > 0 {
		volState.HostID.Value = hostMapping[0].HostID
		volState.HostGroupID.Value = hostMapping[0].HostGroupID
		volState.LogicalUnitNumber.Value = hostMapping[0].LogicalUnitNumber
	}

	if operation == "Create" {
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

