package powerstore

import (
	"context"
	"fmt"
	pstore "github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/types"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

const (
	// MiB to convert size in megabytes
	MiB = 1024 * 1024
	// GiB to convert size in gigabytes
	GiB = 1024 * MiB
	// TiB to convert size in terabytes
	TiB = 1024 * GiB
)

func updateVolState(volState *models.Volume, volResponse pstore.Volume, hostMapping []pstore.HostVolumeMapping, volGroupMapping pstore.VolumeGroups, volPlan *models.Volume, operation operation) {
	// Update value from Volume Response to State
	volState.ID = types.StringValue(volResponse.ID)
	volState.Name = types.StringValue(volResponse.Name)
	size, unit := convertFromBytes(volResponse.Size)
	volState.Size = types.Float64Value(size)
	volState.CapacityUnit = types.StringValue(unit)
	volState.Type = types.StringValue(string(volResponse.Type))
	volState.WWN = types.StringValue(volResponse.Wwn)
	volState.Description = types.StringValue(volResponse.Description)
	volState.State = types.StringValue(string(volResponse.State))
	volState.WWN = types.StringValue(volResponse.Wwn)
	volState.ApplianceID = types.StringValue(volResponse.ApplianceID)
	volState.CreationTimeStamp = types.StringValue(volResponse.CreationTimeStamp)
	volState.ProtectionPolicyID = types.StringValue(volResponse.ProtectionPolicyID)
	volState.PerformancePolicyID = types.StringValue(volResponse.PerformancePolicyID)

	if len(volGroupMapping.VolumeGroup) > 0 {
		volState.VolumeGroupID = types.StringValue(volGroupMapping.VolumeGroup[0].ID)
	} else {
		volState.VolumeGroupID = types.StringValue("")
	}

	// Only if Host is mapped to the volume then update host details
	if len(hostMapping) > 0 {
		volState.HostID = types.StringValue(hostMapping[0].HostID)
		volState.HostGroupID = types.StringValue(hostMapping[0].HostGroupID)
		volState.LogicalUnitNumber = types.Int64Value(hostMapping[0].LogicalUnitNumber)
	} else {
		volState.HostID = types.StringValue("")
		volState.HostGroupID = types.StringValue("")
		volState.LogicalUnitNumber = types.Int64Value(0)
	}

	if operation == operationCreate || operation == operationUpdate {
		volState.MinimumSize = volPlan.MinimumSize
		volState.SectorSize = volPlan.SectorSize
		volState.ApplianceName = volPlan.ApplianceName
		volState.HostName = volPlan.HostName
		volState.HostGroupName = volPlan.HostGroupName
		volState.VolumeGroupName = volPlan.VolumeGroupName
	} else if operation == operationImport {
		volState.SectorSize = types.Int64Value(defaultSectorSize)
	}

	volState.AppType = types.StringValue(string(volResponse.AppType))
	volState.AppTypeOther = types.StringValue(volResponse.AppTypeOther)
	volState.IsReplicationDestination = types.BoolValue(volResponse.IsReplicationDestination)
	volState.NodeAffinity = types.StringValue(string(volResponse.NodeAffinity))
	volState.LogicalUsed = types.Int64Value(volResponse.LogicalUsed)
	volState.Nsid = types.Int64Value(volResponse.Nsid)
	volState.Nguid = types.StringValue(volResponse.Nguid)
}

func updateVol(ctx context.Context, client client.Client, planVol, stateVol models.Volume) ([]string, []string, []string) {
	updatedParameters := []string{}
	updateFailedParameters := []string{}
	errorMessages := []string{}
	volID := stateVol.ID.ValueString()

	valid, errmsg := validateUpdate(ctx, planVol, stateVol)
	if !valid {
		updateFailedParameters = append(updateFailedParameters, "Validation Failed")
		errorMessages = append(errorMessages, fmt.Sprintf("Validation Error: %s", errmsg))
		return updatedParameters, updateFailedParameters, errorMessages
	}

	valInBytes, errmsg := convertToBytes(ctx, planVol)
	if len(errmsg) > 0 {
		updateFailedParameters = append(updateFailedParameters, "Validation Failed")
		errorMessages = append(errorMessages, fmt.Sprintf("Validation Failed: %s", errmsg))
		return updatedParameters, updateFailedParameters, errorMessages
	}

	err := modifyVolume(planVol, valInBytes, volID, client)

	if err != nil {
		updateFailedParameters = append(updateFailedParameters, "name,size,protection policy,performance policy")
		errorMessages = append(errorMessages, fmt.Sprintf("Failed to Update : %s", err.Error()))
	} else {
		updatedParameters = append(updatedParameters, "name, size, protection policy, performance policy")
	}

	// If there's any mismatch between planned and state value of Host and HostGroup ID then either Mapping or UnMapping of host is performed
	if planVol.HostGroupID.ValueString() != stateVol.HostGroupID.ValueString() || planVol.HostID.ValueString() != stateVol.HostID.ValueString() {
		// Detach host from volume
		detachHostFromVolume(stateVol, planVol, client, volID)
		if err != nil {
			updateFailedParameters = append(updateFailedParameters, "unmap volume from host")
			errorMessages = append(errorMessages, fmt.Sprintf("Failed to unmap volume from host: %s", err.Error()))
		} else {
			updatedParameters = append(updatedParameters, "unmapped volume from host")
		}
		// Attach host to volume
		attachHostFromVolume(stateVol, planVol, client, volID)
		if err != nil {
			updateFailedParameters = append(updateFailedParameters, "map volume to host")
			errorMessages = append(errorMessages, fmt.Sprintf("Failed to map volume to host: %s", err.Error()))
		} else {
			updatedParameters = append(updatedParameters, "mapped volume to host")
		}
	}

	// If there's any mismatch between planned and state value of VolumeGroup ID then either Mapping or UnMapping of Volume Group is performed
	if planVol.VolumeGroupID.ValueString() != stateVol.VolumeGroupID.ValueString() {
		if stateVol.VolumeGroupID.ValueString() != "" {
			detachVolumeGroup(ctx, stateVol, client, volID)
			if err != nil {
				updateFailedParameters = append(updateFailedParameters, "Unmap volume group ID")
				errorMessages = append(errorMessages, fmt.Sprintf("Failed to unmap volume group ID: %s", err.Error()))
			} else {
				updatedParameters = append(updatedParameters, "unmapped volume group ID")
			}
		}
		if planVol.VolumeGroupID.ValueString() != "" {
			attachVolumeGroup(ctx, planVol, client, volID)
			if err != nil {
				updateFailedParameters = append(updateFailedParameters, "Map volume group ID")
				errorMessages = append(errorMessages, fmt.Sprintf("Failed to Map volume group ID : %s", err.Error()))
			} else {
				updatedParameters = append(updatedParameters, "Mapped volume group ID")
			}
		}
	}
	return updatedParameters, updateFailedParameters, errorMessages
}

// validations while updating volume
func validateUpdate(ctx context.Context, planVol, stateVol models.Volume) (bool, string) {
	if planVol.HostID.ValueString() != "" && planVol.HostGroupID.ValueString() != "" {
		return false, "Either of HostID and Host GroupID should be present."
	}
	if (planVol.HostGroupID.ValueString() != "" && stateVol.HostID.ValueString() != "" && planVol.HostID.ValueString() != "") ||
		(planVol.HostID.ValueString() != "" && stateVol.HostGroupID.ValueString() != "" && planVol.HostGroupID.ValueString() != "") {
		return false, "Either of HostID and Host GroupID is already present. Both cannot be present."
	}
	if !planVol.PerformancePolicyID.IsUnknown() && planVol.PerformancePolicyID.ValueString() == "" {
		return false, "Performance Policy if present cannot be empty. Either remove the field or set desired value"
	}

	return true, ""
}

// validations while creating volume
func creationValidation(ctx context.Context, plan models.Volume) (bool, string) {
	if plan.HostID.ValueString() != "" && plan.HostGroupID.ValueString() != "" {
		return false, "Either HostID or HostGroupID can be present "
	}
	if !plan.PerformancePolicyID.IsUnknown() && plan.PerformancePolicyID.ValueString() == "" {
		return false, "Performance Policy if present cannot be empty. Either remove the field or set desired value"
	}
	return true, ""
}

func convertToBytes(ctx context.Context, plan models.Volume) (int64, string) {
	var valInBytes float64
	switch plan.CapacityUnit.ValueString() {
	case "MB":
		valInBytes = plan.Size.ValueFloat64() * MiB
	case "TB":
		valInBytes = plan.Size.ValueFloat64() * TiB
	case "GB":
		valInBytes = plan.Size.ValueFloat64() * GiB
	default:
		return 0, "Invalid Capacity unit"
	}
	return int64(valInBytes), ""
}

func convertFromBytes(bytes int64) (float64, string) {
	var newSize float64
	var unit int
	var units = []string{"KB", "MB", "GB", "TB"}
	for newSize = float64(bytes); newSize >= 1024 && unit < len(units); unit++ {
		newSize = newSize / 1024
	}
	return newSize, units[unit-1]
}

// fetchByName updates IDs of the corresponding name present in plan
func fetchByName(client client.Client, plan *models.Volume) (bool, string) {
	if plan.HostID.ValueString() != "" && plan.HostName.ValueString() != "" {
		return false, "Host ID or Host Name"
	} else if plan.HostName.ValueString() != "" {
		hostMap, _ := client.PStoreClient.GetHostByName(context.Background(), plan.HostName.ValueString())
		plan.HostID = types.StringValue(hostMap.ID)
	}
	if plan.HostGroupID.ValueString() != "" && plan.HostGroupName.ValueString() != "" {
		return false, "Host Group ID or Host Group Name"
	} else if plan.HostGroupName.ValueString() != "" {
		hostGroupMap, _ := client.PStoreClient.GetHostGroupByName(context.Background(), plan.HostGroupName.ValueString())
		plan.HostGroupID = types.StringValue(hostGroupMap.ID)
	}
	if plan.VolumeGroupID.ValueString() != "" && plan.VolumeGroupName.ValueString() != "" {
		return false, "Volume Group ID or Volume Group Name"
	} else if plan.VolumeGroupName.ValueString() != "" {
		volGroupMap, _ := client.PStoreClient.GetVolumeGroupByName(context.Background(), plan.VolumeGroupName.ValueString())
		plan.VolumeGroupID = types.StringValue(volGroupMap.ID)
	}
	if plan.ApplianceID.ValueString() != "" && plan.ApplianceName.ValueString() != "" {
		return false, "Appliance ID or Appliance Name"
	} else if plan.ApplianceName.ValueString() != "" {
		applianceMap, _ := client.PStoreClient.GetApplianceByName(context.Background(), plan.ApplianceName.ValueString())
		plan.ApplianceID = types.StringValue(applianceMap.ID)
	}
	if plan.ProtectionPolicyID.ValueString() != "" && plan.ProtectionPolicyName.ValueString() != "" {
		return false, "Protection Policy ID or Protection Policy Name"
	} else if plan.ProtectionPolicyName.ValueString() != "" {
		policyMap, _ := client.PStoreClient.GetApplianceByName(context.Background(), plan.ProtectionPolicyName.ValueString())
		plan.ProtectionPolicyID = types.StringValue(policyMap.ID)
	}
	return true, ""
}

func detachHostFromVolume(stateVol, planVol models.Volume, client client.Client, volID string) error {
	var err error
	if (stateVol.HostID.ValueString() != "" && planVol.HostID.ValueString() == "") || (stateVol.HostGroupID.ValueString() != "" && planVol.HostGroupID.ValueString() == "") {
		volumeHostMapping := &pstore.HostVolumeDetach{
			VolumeID: &volID,
		}
		if stateVol.HostID.ValueString() != "" {
			_, err = client.PStoreClient.DetachVolumeFromHost(context.Background(), stateVol.HostID.ValueString(), volumeHostMapping)
		} else {
			_, err = client.PStoreClient.DetachVolumeFromHostGroup(context.Background(), stateVol.HostGroupID.ValueString(), volumeHostMapping)
		}
	}
	return err
}

func attachHostFromVolume(stateVol, planVol models.Volume, client client.Client, volID string) error {
	var err error
	if (stateVol.HostID.ValueString() == "" && planVol.HostID.ValueString() != "") || (stateVol.HostGroupID.ValueString() == "" && planVol.HostGroupID.ValueString() != "") {
		volumeHostMapping := &pstore.HostVolumeAttach{
			VolumeID: &volID,
		}

		if planVol.HostID.ValueString() != "" {
			_, err = client.PStoreClient.AttachVolumeToHost(context.Background(), planVol.HostID.ValueString(), volumeHostMapping)
		} else {
			_, err = client.PStoreClient.AttachVolumeToHostGroup(context.Background(), planVol.HostGroupID.ValueString(), volumeHostMapping)
		}
	}
	return err
}

func attachVolumeGroup(ctx context.Context, planVol models.Volume, client client.Client, volID string) error {
	_, err := client.PStoreClient.AddMembersToVolumeGroup(ctx, &pstore.VolumeGroupMembers{VolumeIds: []string{volID}}, planVol.VolumeGroupID.ValueString())
	return err
}

func detachVolumeGroup(ctx context.Context, stateVol models.Volume, client client.Client, volID string) error {
	_, err := client.PStoreClient.RemoveMembersFromVolumeGroup(ctx, &pstore.VolumeGroupMembers{VolumeIds: []string{volID}}, stateVol.VolumeGroupID.ValueString())
	return err
}

func modifyVolume(planVol models.Volume, valInBytes int64, volID string, client client.Client) error {

	vgModify := &pstore.VolumeModify{
		Name:                planVol.Name.ValueString(),
		Size:                valInBytes,
		ProtectionPolicyID:  planVol.ProtectionPolicyID.ValueString(),
		PerformancePolicyID: planVol.PerformancePolicyID.ValueString(),
	}

	_, err := client.PStoreClient.ModifyVolume(context.Background(), vgModify, volID)
	return err
}
