package powerstore

// const (
// 	// MiB to convert size in megabytes
// 	MiB = 1024 * 1024
// 	// GiB to convert size in gigabytes
// 	GiB = 1024 * MiB
// 	// TiB to convert size in terabytes
// 	TiB = 1024 * GiB
// )

// func updateVolState(volState *models.Volume, volResponse pstore.Volume, hostMapping []pstore.HostVolumeMapping, volGroupMapping pstore.VolumeGroups, volPlan *models.Volume, operation operation) {
// 	// Update value from Volume Response to State
// 	volState.ID.Value = volResponse.ID
// 	volState.Name.Value = volResponse.Name
// 	size, unit := convertFromBytes(volResponse.Size)
// 	volState.Size.Value = size
// 	volState.CapacityUnit.Value = unit
// 	volState.Type.Value = string(volResponse.Type)
// 	volState.WWN.Value = volResponse.Wwn
// 	volState.Description.Value = volResponse.Description
// 	volState.State.Value = string(volResponse.State)
// 	volState.WWN.Value = volResponse.Wwn
// 	volState.ApplianceID.Value = volResponse.ApplianceID
// 	volState.CreationTimeStamp.Value = volResponse.CreationTimeStamp
// 	volState.ProtectionPolicyID.Value = volResponse.ProtectionPolicyID
// 	volState.PerformancePolicyID.Value = volResponse.PerformancePolicyID

// 	if len(volGroupMapping.VolumeGroup) > 0 {
// 		volState.VolumeGroupID.Value = volGroupMapping.VolumeGroup[0].ID
// 	} else {
// 		volState.VolumeGroupID.Value = ""
// 	}

// 	// Only if Host is mapped to the volume then update host details
// 	if len(hostMapping) > 0 {
// 		volState.HostID.Value = hostMapping[0].HostID
// 		volState.HostGroupID.Value = hostMapping[0].HostGroupID
// 		volState.LogicalUnitNumber.Value = hostMapping[0].LogicalUnitNumber
// 	} else {
// 		volState.HostID.Value = ""
// 		volState.HostGroupID.Value = ""
// 		volState.LogicalUnitNumber.Value = 0
// 	}

// 	if operation == operationCreate || operation == operationUpdate {
// 		volState.MinimumSize.Value = volPlan.MinimumSize.Value
// 		volState.SectorSize.Value = volPlan.SectorSize.Value
// 		volState.ApplianceName.Value = volPlan.ApplianceName.Value
// 		volState.HostName.Value = volPlan.HostName.Value
// 		volState.HostGroupName.Value = volPlan.HostGroupName.Value
// 		volState.VolumeGroupName.Value = volPlan.VolumeGroupName.Value
// 	} else if operation == operationImport {
// 		volState.SectorSize.Value = defaultSectorSize
// 	}

// 	volState.AppType.Value = volResponse.AppType
// 	volState.AppTypeOther.Value = volResponse.AppTypeOther
// 	volState.IsReplicationDestination.Value = volResponse.IsReplicationDestination
// 	volState.NodeAffinity.Value = volResponse.NodeAffinity
// 	volState.LogicalUsed.Value = volResponse.LogicalUsed
// 	volState.Nsid.Value = volResponse.Nsid
// 	volState.Nguid.Value = volResponse.Nguid
// }

// func updateVol(ctx context.Context, client client.Client, planVol, stateVol models.Volume) ([]string, []string, []string) {
// 	updatedParameters := []string{}
// 	updateFailedParameters := []string{}
// 	errorMessages := []string{}
// 	volID := stateVol.ID.Value

// 	valid, errmsg := validateUpdate(ctx, planVol, stateVol)
// 	if !valid {
// 		updateFailedParameters = append(updateFailedParameters, "Validation Failed")
// 		errorMessages = append(errorMessages, fmt.Sprintf("Validation Error: %s", errmsg))
// 		return updatedParameters, updateFailedParameters, errorMessages
// 	}

// 	valInBytes, errmsg := convertToBytes(ctx, planVol)
// 	if len(errmsg) > 0 {
// 		updateFailedParameters = append(updateFailedParameters, "Validation Failed")
// 		errorMessages = append(errorMessages, fmt.Sprintf("Validation Failed: %s", errmsg))
// 		return updatedParameters, updateFailedParameters, errorMessages
// 	}

// 	err := modifyVolume(planVol, valInBytes, volID, client)

// 	if err != nil {
// 		updateFailedParameters = append(updateFailedParameters, "name,size,protection policy,performance policy")
// 		errorMessages = append(errorMessages, fmt.Sprintf("Failed to Update : %s", err.Error()))
// 	} else {
// 		updatedParameters = append(updatedParameters, "name, size, protection policy, performance policy")
// 	}

// 	// If there's any mismatch between planned and state value of Host and HostGroup ID then either Mapping or UnMapping of host is performed
// 	if planVol.HostGroupID.Value != stateVol.HostGroupID.Value || planVol.HostID.Value != stateVol.HostID.Value {
// 		// Detach host from volume
// 		detachHostFromVolume(stateVol, planVol, client, volID)
// 		if err != nil {
// 			updateFailedParameters = append(updateFailedParameters, "unmap volume from host")
// 			errorMessages = append(errorMessages, fmt.Sprintf("Failed to unmap volume from host: %s", err.Error()))
// 		} else {
// 			updatedParameters = append(updatedParameters, "unmapped volume from host")
// 		}
// 		// Attach host to volume
// 		attachHostFromVolume(stateVol, planVol, client, volID)
// 		if err != nil {
// 			updateFailedParameters = append(updateFailedParameters, "map volume to host")
// 			errorMessages = append(errorMessages, fmt.Sprintf("Failed to map volume to host: %s", err.Error()))
// 		} else {
// 			updatedParameters = append(updatedParameters, "mapped volume to host")
// 		}
// 	}

// 	// If there's any mismatch between planned and state value of VolumeGroup ID then either Mapping or UnMapping of Volume Group is performed
// 	if planVol.VolumeGroupID.Value != stateVol.VolumeGroupID.Value {
// 		if stateVol.VolumeGroupID.Value != "" {
// 			detachVolumeGroup(ctx, stateVol, client, volID)
// 			if err != nil {
// 				updateFailedParameters = append(updateFailedParameters, "Unmap volume group ID")
// 				errorMessages = append(errorMessages, fmt.Sprintf("Failed to unmap volume group ID: %s", err.Error()))
// 			} else {
// 				updatedParameters = append(updatedParameters, "unmapped volume group ID")
// 			}
// 		}
// 		if planVol.VolumeGroupID.Value != "" {
// 			attachVolumeGroup(ctx, planVol, client, volID)
// 			if err != nil {
// 				updateFailedParameters = append(updateFailedParameters, "Map volume group ID")
// 				errorMessages = append(errorMessages, fmt.Sprintf("Failed to Map volume group ID : %s", err.Error()))
// 			} else {
// 				updatedParameters = append(updatedParameters, "Mapped volume group ID")
// 			}
// 		}
// 	}
// 	return updatedParameters, updateFailedParameters, errorMessages
// }

// // validations while updating volume
// func validateUpdate(ctx context.Context, planVol, stateVol models.Volume) (bool, string) {
// 	if planVol.HostID.Value != "" && planVol.HostGroupID.Value != "" {
// 		return false, "Either of HostID and Host GroupID should be present."
// 	}
// 	if (planVol.HostGroupID.Value != "" && stateVol.HostID.Value != "" && planVol.HostID.Value != "") ||
// 		(planVol.HostID.Value != "" && stateVol.HostGroupID.Value != "" && planVol.HostGroupID.Value != "") {
// 		return false, "Either of HostID and Host GroupID is already present. Both cannot be present."
// 	}

// 	return true, ""
// }

// // validations while creating volume
// func creationValidation(ctx context.Context, plan models.Volume) (bool, string) {
// 	if plan.HostID.Value != "" && plan.HostGroupID.Value != "" {
// 		return false, "Either HostID or HostGroupID can be present "
// 	}
// 	if plan.PerformancePolicyID.Value == "" {
// 		return false, "Performance Policy if present cannot be empty. Either remove the field or set desired value"
// 	}
// 	return true, ""
// }

// func convertToBytes(ctx context.Context, plan models.Volume) (int64, string) {
// 	var valInBytes float64
// 	switch plan.CapacityUnit.Value {
// 	case "MB":
// 		valInBytes = plan.Size.Value * MiB
// 	case "TB":
// 		valInBytes = plan.Size.Value * TiB
// 	case "GB":
// 		valInBytes = plan.Size.Value * GiB
// 	default:
// 		return 0, "Invalid Capacity unit"
// 	}
// 	return int64(valInBytes), ""
// }

// func convertFromBytes(bytes int64) (float64, string) {
// 	var newSize float64
// 	var unit int
// 	var units = []string{"KB", "MB", "GB", "TB"}
// 	for newSize = float64(bytes); newSize >= 1024 && unit < len(units); unit++ {
// 		newSize = newSize / 1024
// 	}
// 	return newSize, units[unit-1]
// }

// // fetchByName updates IDs of the corresponding name present in plan
// func fetchByName(client client.Client, plan *models.Volume) (bool, string) {
// 	if plan.HostID.Value != "" && plan.HostName.Value != "" {
// 		return false, "Host ID or Host Name"
// 	} else if plan.HostName.Value != "" {
// 		hostMap, _ := client.PStoreClient.GetHostByName(context.Background(), plan.HostName.Value)
// 		plan.HostID.Value = hostMap.ID
// 	}
// 	if plan.HostGroupID.Value != "" && plan.HostGroupName.Value != "" {
// 		return false, "Host Group ID or Host Group Name"
// 	} else if plan.HostGroupName.Value != "" {
// 		hostGroupMap, _ := client.PStoreClient.GetHostGroupByName(context.Background(), plan.HostGroupName.Value)
// 		plan.HostGroupID.Value = hostGroupMap.ID
// 	}
// 	if plan.VolumeGroupID.Value != "" && plan.VolumeGroupName.Value != "" {
// 		return false, "Volume Group ID or Volume Group Name"
// 	} else if plan.VolumeGroupName.Value != "" {
// 		volGroupMap, _ := client.PStoreClient.GetVolumeGroupByName(context.Background(), plan.VolumeGroupName.Value)
// 		plan.VolumeGroupID.Value = volGroupMap.ID
// 	}
// 	if plan.ApplianceID.Value != "" && plan.ApplianceName.Value != "" {
// 		return false, "Appliance ID or Appliance Name"
// 	} else if plan.ApplianceName.Value != "" {
// 		applianceMap, _ := client.PStoreClient.GetApplianceByName(context.Background(), plan.ApplianceName.Value)
// 		plan.ApplianceID.Value = applianceMap.ID
// 	}
// 	if plan.ProtectionPolicyID.Value != "" && plan.ProtectionPolicyName.Value != "" {
// 		return false, "Protection Policy ID or Protection Policy Name"
// 	} else if plan.ProtectionPolicyName.Value != "" {
// 		policyMap, _ := client.PStoreClient.GetApplianceByName(context.Background(), plan.ProtectionPolicyName.Value)
// 		plan.ProtectionPolicyID.Value = policyMap.ID
// 	}
// 	return true, ""
// }

// func detachHostFromVolume(stateVol, planVol models.Volume, client client.Client, volID string) error {
// 	var err error
// 	if (stateVol.HostID.Value != "" && planVol.HostID.Value == "") || (stateVol.HostGroupID.Value != "" && planVol.HostGroupID.Value == "") {
// 		volumeHostMapping := &pstore.HostVolumeDetach{
// 			VolumeID: &volID,
// 		}
// 		if stateVol.HostID.Value != "" {
// 			_, err = client.PStoreClient.DetachVolumeFromHost(context.Background(), stateVol.HostID.Value, volumeHostMapping)
// 		} else {
// 			_, err = client.PStoreClient.DetachVolumeFromHostGroup(context.Background(), stateVol.HostGroupID.Value, volumeHostMapping)
// 		}
// 	}
// 	return err
// }

// func attachHostFromVolume(stateVol, planVol models.Volume, client client.Client, volID string) error {
// 	var err error
// 	if (stateVol.HostID.Value == "" && planVol.HostID.Value != "") || (stateVol.HostGroupID.Value == "" && planVol.HostGroupID.Value != "") {
// 		volumeHostMapping := &pstore.HostVolumeAttach{
// 			VolumeID: &volID,
// 		}

// 		if planVol.HostID.Value != "" {
// 			_, err = client.PStoreClient.AttachVolumeToHost(context.Background(), planVol.HostID.Value, volumeHostMapping)
// 		} else {
// 			_, err = client.PStoreClient.AttachVolumeToHostGroup(context.Background(), planVol.HostGroupID.Value, volumeHostMapping)
// 		}
// 	}
// 	return err
// }

// func attachVolumeGroup(ctx context.Context, planVol models.Volume, client client.Client, volID string) error {
// 	_, err := client.PStoreClient.AddMembersToVolumeGroup(ctx, &pstore.VolumeGroupMembers{VolumeIds: []string{volID}}, planVol.VolumeGroupID.Value)
// 	return err
// }

// func detachVolumeGroup(ctx context.Context, stateVol models.Volume, client client.Client, volID string) error {
// 	_, err := client.PStoreClient.RemoveMembersFromVolumeGroup(ctx, &pstore.VolumeGroupMembers{VolumeIds: []string{volID}}, stateVol.VolumeGroupID.Value)
// 	return err
// }

// func modifyVolume(planVol models.Volume, valInBytes int64, volID string, client client.Client) error {

// 	vgModify := &pstore.VolumeModify{
// 		Name:                planVol.Name.Value,
// 		Size:                valInBytes,
// 		ProtectionPolicyID:  planVol.ProtectionPolicyID.Value,
// 		PerformancePolicyID: planVol.PerformancePolicyID.Value,
// 	}

// 	_, err := client.PStoreClient.ModifyVolume(context.Background(), vgModify, volID)
// 	return err
// }
