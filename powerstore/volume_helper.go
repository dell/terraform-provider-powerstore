/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package powerstore

import (
	"context"
	"fmt"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	pstore "github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
		volState.ProtectionPolicyName = volPlan.ProtectionPolicyName
	} else if operation == operationImport || operation == operationRead {
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
		updateFailedParameters = append(updateFailedParameters, "name,size,protection policy,performance policy, description")
		errorMessages = append(errorMessages, fmt.Sprintf("Failed to Update : %s", err.Error()))
	} else {
		updatedParameters = append(updatedParameters, "name, size, protection policy, performance policy, description")
	}

	// If there's any mismatch between planned and state value of Host and HostGroup ID then either Mapping or UnMapping of host is performed
	if planVol.HostGroupID.ValueString() != stateVol.HostGroupID.ValueString() || planVol.HostID.ValueString() != stateVol.HostID.ValueString() {
		// Detach host from volume
		err := detachHostFromVolume(stateVol, planVol, client, volID)
		if err != nil {
			updateFailedParameters = append(updateFailedParameters, "unmap volume from host")
			errorMessages = append(errorMessages, fmt.Sprintf("Failed to unmap volume from host: %s", err.Error()))
		} else {
			updatedParameters = append(updatedParameters, "unmapped volume from host")
		}
		// Attach host to volume
		err = attachHostFromVolume(stateVol, planVol, client, volID)
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
			err := detachVolumeGroup(ctx, stateVol, client, volID)
			if err != nil {
				updateFailedParameters = append(updateFailedParameters, "Unmap volume group ID")
				errorMessages = append(errorMessages, fmt.Sprintf("Failed to unmap volume group ID: %s", err.Error()))
			} else {
				updatedParameters = append(updatedParameters, "unmapped volume group ID")
			}
		}
		if planVol.VolumeGroupID.ValueString() != "" {
			err = attachVolumeGroup(ctx, planVol, client, volID)
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
	if !planVol.LogicalUnitNumber.IsUnknown() && planVol.LogicalUnitNumber != stateVol.LogicalUnitNumber {
		return false, "Logical Unit Number cannot be modified."
	}
	if !planVol.ApplianceID.IsUnknown() && planVol.ApplianceID.ValueString() != stateVol.ApplianceID.ValueString() {
		return false, "Appliance ID cannot be modified."
	}
	if planVol.ApplianceName.ValueString() != stateVol.ApplianceName.ValueString() {
		return false, "Appliance Name cannot be modified."
	}
	if planVol.SectorSize != stateVol.SectorSize {
		return false, "Sector Size cannot be modified."
	}
	if planVol.MinimumSize != stateVol.MinimumSize {
		return false, "Minimum Size cannot be modified."
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
	if unit > 0 {
		return newSize, units[unit-1]
	}
	return newSize, units[unit]
}

// fetchByName updates IDs of the corresponding name present in plan
func fetchByName(client client.Client, plan *models.Volume) string {
	if plan.HostName.ValueString() != "" {
		hostMap, err := client.PStoreClient.GetHostByName(context.Background(), plan.HostName.ValueString())
		if err != nil {
			return "Invalid host name"
		}
		plan.HostID = types.StringValue(hostMap.ID)
	}
	if plan.HostGroupName.ValueString() != "" {
		hostGroupMap, err := client.PStoreClient.GetHostGroupByName(context.Background(), plan.HostGroupName.ValueString())
		if err != nil {
			return "Invalid host group name"
		}
		plan.HostGroupID = types.StringValue(hostGroupMap.ID)
	}
	if plan.VolumeGroupName.ValueString() != "" {
		volGroupMap, err := client.PStoreClient.GetVolumeGroupByName(context.Background(), plan.VolumeGroupName.ValueString())
		if err != nil {
			return "Invalid volume group name"
		}
		plan.VolumeGroupID = types.StringValue(volGroupMap.ID)
	}
	if plan.ApplianceName.ValueString() != "" {
		applianceMap, err := client.PStoreClient.GetApplianceByName(context.Background(), plan.ApplianceName.ValueString())
		if err != nil {
			return "Invalid Appliance name"
		}
		plan.ApplianceID = types.StringValue(applianceMap.ID)
	}
	if plan.ProtectionPolicyName.ValueString() != "" {
		policyMap, err := client.PStoreClient.GetProtectionPolicyByName(context.Background(), plan.ProtectionPolicyName.ValueString())
		if err != nil {
			return "Invalid Protection policy name"
		}
		plan.ProtectionPolicyID = types.StringValue(policyMap.ID)
	}
	return ""
}

func detachHostFromVolume(stateVol, planVol models.Volume, client client.Client, volID string) error {
	var err error
	if stateVol.HostID.ValueString() != "" || stateVol.HostGroupID.ValueString() != "" {
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
	if planVol.HostID.ValueString() != "" || planVol.HostGroupID.ValueString() != "" {
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
	_, err := client.PStoreClient.AddMembersToVolumeGroup(ctx, &pstore.VolumeGroupMembers{VolumeIDs: []string{volID}}, planVol.VolumeGroupID.ValueString())
	return err
}

func detachVolumeGroup(ctx context.Context, stateVol models.Volume, client client.Client, volID string) error {
	_, err := client.PStoreClient.RemoveMembersFromVolumeGroup(ctx, &pstore.VolumeGroupMembers{VolumeIDs: []string{volID}}, stateVol.VolumeGroupID.ValueString())
	return err
}

func modifyVolume(planVol models.Volume, valInBytes int64, volID string, client client.Client) error {
	protectionPolicy := planVol.ProtectionPolicyID.ValueString()
	vgModify := &pstore.VolumeModify{
		Name:                planVol.Name.ValueString(),
		Size:                valInBytes,
		ProtectionPolicyID:  protectionPolicy,
		PerformancePolicyID: planVol.PerformancePolicyID.ValueString(),
		Description:         planVol.Description.ValueString(),
		AppType:             planVol.AppType.ValueString(),
		AppTypeOther:        planVol.AppTypeOther.ValueString(),
	}

	_, err := client.PStoreClient.ModifyVolume(context.Background(), vgModify, volID)
	return err
}
