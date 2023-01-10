package powerstore

// const defaultSectorSize = 512

// type resourceVolumeType struct{}

// // GetSchema returns the schema for this resource.
// func (r resourceVolumeType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
// 	return tfsdk.Schema{
// 		Attributes: map[string]tfsdk.Attribute{
// 			"id": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Description:         "The ID of the volume.",
// 				MarkdownDescription: "The ID of the volume.",
// 			},
// 			"name": {
// 				Type:                types.StringType,
// 				Required:            true,
// 				Description:         "The name of the volume.",
// 				MarkdownDescription: "The name of the volume.",
// 			},
// 			"size": {
// 				Type:                types.Float64Type,
// 				Required:            true,
// 				Description:         "The size of the volume.",
// 				MarkdownDescription: "The size of the volume.",
// 			},
// 			"capacity_unit": {
// 				Type:                types.StringType,
// 				Optional:            true,
// 				Computed:            true,
// 				Description:         "The Capacity Unit corresponding to the size.",
// 				MarkdownDescription: "The Capacity Unit corresponding to the size.",
// 				PlanModifiers: tfsdk.AttributePlanModifiers{
// 					DefaultAttribute(types.String{Value: "GB"}),
// 				},
// 			},
// 			"host_id": {
// 				Type:                types.StringType,
// 				Optional:            true,
// 				Computed:            true,
// 				Description:         "The host id of the volume.",
// 				MarkdownDescription: "The host id of the volume.",
// 			},
// 			"host_name": {
// 				Type:                types.StringType,
// 				Optional:            true,
// 				Computed:            true,
// 				Description:         "The host name of the volume.",
// 				MarkdownDescription: "The host name of the volume.",
// 			},
// 			"host_group_id": {
// 				Type:                types.StringType,
// 				Optional:            true,
// 				Computed:            true,
// 				Description:         "The host group id of the volume.",
// 				MarkdownDescription: "The host group id of the volume.",
// 			},
// 			"host_group_name": {
// 				Type:                types.StringType,
// 				Optional:            true,
// 				Computed:            true,
// 				Description:         "The host group name of the volume.",
// 				MarkdownDescription: "The host group name of the volume.",
// 			},
// 			"logical_unit_number": {
// 				Type:                types.Int64Type,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The current amount of data written to the volume.",
// 				MarkdownDescription: "The current amount of data written to the volume.",
// 			},
// 			"volume_group_id": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The volume group id of the volume.",
// 				MarkdownDescription: "The volume group id of the volume.",
// 			},
// 			"volume_group_name": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The volume group name of the volume.",
// 				MarkdownDescription: "The volume group name of the volume.",
// 			},
// 			"min_size": {
// 				Type:                types.Int64Type,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The minimum size  of the volume.",
// 				MarkdownDescription: "The minimum size of the volume.",
// 			},
// 			"sector_size": {
// 				Type:                types.Int64Type,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The sector size of the volume.",
// 				MarkdownDescription: "The sector size of the volume.",
// 				PlanModifiers: tfsdk.AttributePlanModifiers{
// 					DefaultAttribute(types.Int64{Value: defaultSectorSize}),
// 				},
// 			},
// 			"description": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The description of the volume.",
// 				MarkdownDescription: "The description of the volume.",
// 			},
// 			"appliance_id": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The appliance_id of the volume.",
// 				MarkdownDescription: "The appliance_id of the volume.",
// 			},
// 			"appliance_name": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The appliance name of the volume.",
// 				MarkdownDescription: "The appliance name of the volume.",
// 			},
// 			"protection_policy_id": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The protection_policy_id of the volume.",
// 				MarkdownDescription: "The protection_policy_id of the volume.",
// 			},
// 			"protection_policy_name": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The protection policy name of the volume.",
// 				MarkdownDescription: "The protection policy name of the volume.",
// 			},
// 			"performance_policy_id": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The performance_policy_id of the volume.",
// 				MarkdownDescription: "The performance_policy_id of the volume.",
// 				PlanModifiers: tfsdk.AttributePlanModifiers{
// 					DefaultAttribute(types.String{Value: "default_medium"}),
// 				},
// 			},
// 			"creation_timestamp": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Description:         "The creation_timestamp of the volume.",
// 				MarkdownDescription: "The creation_timestamp of the volume.",
// 			},
// 			"is_replication_destination": {
// 				Type:                types.BoolType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The is_replication_destination of the volume.",
// 				MarkdownDescription: "The is_replication_destination of the volume.",
// 			},
// 			"node_affinity": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The node_affinity of the volume.",
// 				MarkdownDescription: "The node_affinity of the volume.",
// 			},
// 			"type": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The type of the volume.",
// 				MarkdownDescription: "The type of the volume.",
// 			},
// 			"app_type": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The app type of the volume.",
// 				MarkdownDescription: "The app type of the volume.",
// 			},
// 			"app_type_other": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The app type other of the volume.",
// 				MarkdownDescription: "The app type other of the volume.",
// 			},
// 			"wwn": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The wwn of the volume.",
// 				MarkdownDescription: "The wwn of the volume.",
// 			},
// 			"state": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The state of the volume.",
// 				MarkdownDescription: "The state of the volume.",
// 			},
// 			"nguid": {
// 				Type:                types.StringType,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The nguid of the volume.",
// 				MarkdownDescription: "The nguid of the volume.",
// 			},
// 			"nsid": {
// 				Type:                types.Int64Type,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "The nsid of the volume.",
// 				MarkdownDescription: "The nsid of the volume.",
// 			},
// 			"logical_used": {
// 				Type:                types.Int64Type,
// 				Computed:            true,
// 				Optional:            true,
// 				Description:         "Current amount of data used by the volume.",
// 				MarkdownDescription: "Current amount of data used by the volume.",
// 			},
// 		},
// 	}, nil
// }

// // NewResource is a wrapper around provider
// func (r resourceVolumeType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
// 	return resourceVolume{
// 		p: *(p.(*Pstoreprovider)),
// 	}, nil
// }

// type resourceVolume struct {
// 	p Pstoreprovider
// }

// func (r resourceVolume) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
// 	if !r.p.configured {
// 		resp.Diagnostics.AddError(
// 			"Provider not configured",
// 			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
// 		)
// 		return
// 	}
// 	var valInBytes int64

// 	log.Printf("Started Creating Volume")
// 	var plan models.Volume

// 	diags := req.Plan.Get(ctx, &plan)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	valInBytes, errmsg := convertToBytes(ctx, plan)
// 	if errmsg != "" {
// 		resp.Diagnostics.AddError(
// 			"Error creating volume",
// 			"Could not create volume "+errmsg,
// 		)
// 		return
// 	}

// 	valid, errmsg := fetchByName(r.p.client, &plan)
// 	if !valid {
// 		resp.Diagnostics.AddError(
// 			"Error creating volume",
// 			"Could not create volume, either of "+errmsg+" should be present",
// 		)
// 		return
// 	}

// 	volumeCreate := &gopowerstore.VolumeCreate{
// 		Name:                &plan.Name.Value,
// 		Description:         plan.Description.Value,
// 		Size:                &valInBytes,
// 		ApplianceID:         plan.ApplianceID.Value,
// 		VolumeGroupID:       plan.VolumeGroupID.Value,
// 		SectorSize:          &plan.SectorSize.Value,
// 		ProtectionPolicyID:  plan.ProtectionPolicyID.Value,
// 		PerformancePolicyID: plan.PerformancePolicyID.Value,
// 		AppType:             plan.AppType.Value,
// 		AppTypeOther:        plan.AppTypeOther.Value,
// 		MinimumSize:         plan.MinimumSize.Value,
// 		HostID:              plan.HostID.Value,
// 		HostGroupID:         plan.HostGroupID.Value,
// 		LogicalUnitNumber:   plan.LogicalUnitNumber.Value,
// 	}

// 	// Add validation
// 	valid, validErr := creationValidation(context.Background(), plan)
// 	if !valid {
// 		resp.Diagnostics.AddError(
// 			"Error creating volume",
// 			"Could not create volume "+validErr,
// 		)
// 		return
// 	}

// 	// Create New Volume
// 	// The function returns only ID of the newly created Volume
// 	volCreateResponse, err := r.p.client.PStoreClient.CreateVolume(context.Background(), volumeCreate)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error creating volume",
// 			"Could not create volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}

// 	// Get Volume Details using ID retrieved above
// 	volResponse, err1 := r.p.client.PStoreClient.GetVolume(context.Background(), volCreateResponse.ID)
// 	if err1 != nil {
// 		resp.Diagnostics.AddError(
// 			"Error getting volume after creation",
// 			"Could not get volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}
// 	// Get Host Mapping from volume ID
// 	hostMapping, err1 := r.p.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volCreateResponse.ID)
// 	if err1 != nil {
// 		resp.Diagnostics.AddError(
// 			"Error fetching volume host mapping",
// 			"Could not create volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}
// 	// Get Volume Group Mapping details from API
// 	volGroupMapping, err := r.p.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volCreateResponse.ID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error fetching volume host mapping",
// 			"Could not create volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}
// 	log.Printf("After Volume create call")

// 	result := models.Volume{}
// 	updateVolState(&result, volResponse, hostMapping, volGroupMapping, &plan, operationCreate)

// 	log.Printf("Added to result: %v", result)

// 	diags = resp.State.Set(ctx, result)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}
// 	log.Printf("Done with Create")
// }

// // Read resource information
// func (r resourceVolume) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {

// 	var state models.Volume
// 	diags := req.State.Get(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	// Get volume details from API and then update what is in state from what the API returns
// 	volID := state.ID.Value
// 	volResponse, err := r.p.client.PStoreClient.GetVolume(context.Background(), volID)

// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error reading volume",
// 			"Could not read volume with error "+volID+": "+err.Error(),
// 		)
// 		return
// 	}
// 	// Get Host Mapping details from API
// 	hostMapping, err := r.p.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error fetching volume host mapping",
// 			"Could not create volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}
// 	// Get Volume Group Mapping details from API
// 	volGroupMapping, err := r.p.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error fetching volume host mapping",
// 			"Could not create volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}

// 	updateVolState(&state, volResponse, hostMapping, volGroupMapping, nil, operationRead)

// 	// Set state
// 	diags = resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	log.Printf("Done with Read")
// }

// // Update resource
// func (r resourceVolume) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {

// 	log.Printf("Started Update")

// 	// Get plan values
// 	var plan models.Volume
// 	diags := req.Plan.Get(ctx, &plan)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	// Get current state
// 	var state models.Volume
// 	diags = req.State.Get(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	valid, errmsg := fetchByName(r.p.client, &plan)
// 	if !valid {
// 		resp.Diagnostics.AddError(
// 			"Error creating volume",
// 			"Could not Update volume, either of "+errmsg+" should be present",
// 		)
// 		return
// 	}

// 	// Update volume parameters. In case of validation failure, return
// 	updatedParams, updateFailedParameters, errMessages := updateVol(ctx, r.p.client, plan, state)
// 	if len(updateFailedParameters) > 0 && updateFailedParameters[0] == "Validation Failed" {
// 		resp.Diagnostics.AddError(
// 			"Validation Check Failed",
// 			errMessages[0],
// 		)
// 		return
// 	}
// 	// Get vg ID from state
// 	volID := state.ID.Value

// 	if len(errMessages) > 0 || len(updateFailedParameters) > 0 {
// 		errMessage := strings.Join(errMessages, ",\n")
// 		resp.Diagnostics.AddError(
// 			fmt.Sprintf("Failed to update all parameters of Volume, updated parameters are %v and parameters failed to update are %v", updatedParams, updateFailedParameters),
// 			errMessage)
// 	}

// 	// Get volume details from volume ID
// 	volResponse, err := r.p.client.PStoreClient.GetVolume(context.Background(), volID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error getting volume after update",
// 			"Could not get after update volID "+volID+": "+err.Error(),
// 		)
// 		return
// 	}

// 	// Get Host Mapping from volume ID
// 	hostMapping, err := r.p.client.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error fetching volume host mapping",
// 			"Could not update volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}

// 	// Get Volume Group Mapping details from API
// 	volGroupMapping, err := r.p.client.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error fetching volume host mapping",
// 			"Could not create volume, unexpected error: "+err.Error(),
// 		)
// 		return
// 	}

// 	updateVolState(&state, volResponse, hostMapping, volGroupMapping, &plan, operationUpdate)

// 	//Set State
// 	diags = resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	log.Printf("Done with Update")
// }

// // Delete resource
// func (r resourceVolume) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
// 	log.Printf("Started with Delete")

// 	var state models.Volume
// 	diags := req.State.Get(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	// Get vg ID from state
// 	volID := state.ID.Value
// 	if state.ProtectionPolicyID.Value != "" {
// 		state.ProtectionPolicyID.Value = ""
// 		modifyVolume(state, 0, volID, r.p.client)
// 	}

// 	if state.HostID.Value != "" || state.HostGroupID.Value != "" {
// 		err := detachHostFromVolume(state, models.Volume{}, r.p.client, volID)
// 		if err != nil {
// 			resp.Diagnostics.AddError(
// 				"Cannot detach volume host mapping",
// 				"Could not delete volume, unexpected error: "+err.Error(),
// 			)
// 			return
// 		}
// 	}

// 	if state.VolumeGroupID.Value != "" {
// 		err := detachVolumeGroup(ctx, state, r.p.client, volID)
// 		if err != nil {
// 			resp.Diagnostics.AddError(
// 				"Error detaching volume Group",
// 				"Could not delete volume, unexpected error: "+err.Error(),
// 			)
// 			return

// 		}
// 	}

// 	// Delete volume by calling API
// 	_, err := r.p.client.PStoreClient.DeleteVolume(context.Background(), nil, volID)
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error deleting volume",
// 			"Could not delete volID "+volID+": "+err.Error(),
// 		)
// 		return
// 	}

// 	// Remove resource from state
// 	resp.State.RemoveResource(ctx)
// 	log.Printf("Done with Delete")
// }

// // ImportState import state for existing infrastructure
// func (r resourceVolume) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {

// 	log.Printf("Started with Import")

// 	// fetching Volume information
// 	response, err := r.p.client.PStoreClient.GetVolume(context.Background(), req.ID)

// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Error importing Volume",
// 			fmt.Sprintf("Could not import volume: %s with error: %s", req.ID, err.Error()),
// 		)
// 		return
// 	}

// 	state := models.Volume{}

// 	// as state is like a plan here, a current state prior to this import operation
// 	updateVolState(&state, response, nil, gopowerstore.VolumeGroups{}, &state, operationImport)

// 	// Set state
// 	diags := resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	log.Printf("Done with Import")
// }
