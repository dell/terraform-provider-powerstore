package powerstore

import (
	"context"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

var (
	_ datasource.DataSource              = &volumeDataSource{}
	_ datasource.DataSourceWithConfigure = &volumeDataSource{}
)

// newVolumeDataSource returns the volume data source object
func newVolumeDataSource() datasource.DataSource {
	return &volumeDataSource{}
}

type volumeDataSource struct {
	client *client.Client
}

type volumeDataSourceModel struct {
	Volumes []models.VolumeDataSource `tfsdk:"volumes"`
	ID      types.String              `tfsdk:"id"`
	Name    types.String              `tfsdk:"name"`
}

func (d *volumeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volume"
}

func (d *volumeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ".",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the volume instance.",
				MarkdownDescription: "Unique identifier of the volume instance.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the volume.",
				MarkdownDescription: "Name of the volume.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("id")),
				},
			},
			"volumes": schema.ListNestedAttribute{
				Description:         "List of volumes.",
				MarkdownDescription: "List of volumes.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "The ID of the volume.",
							MarkdownDescription: "The ID of the volume.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the volume.",
							MarkdownDescription: "Name of the volume.",
							Computed:            true,
						},
						"size": schema.Float64Attribute{
							Description:         "The size of the volume.",
							MarkdownDescription: "The size of the volume.",
							Computed:            true,
						},
						"capacity_unit": schema.StringAttribute{
							Description:         "The Capacity Unit corresponding to the size.",
							MarkdownDescription: "The Capacity Unit corresponding to the size.",
							Computed:            true,
						},
						"host_id": schema.StringAttribute{
							Description:         "The host id of the volume.",
							MarkdownDescription: "The host id of the volume.",
							Computed:            true,
						},
						"host_group_id": schema.StringAttribute{
							Description:         "The host group id of the volume.",
							MarkdownDescription: "The host group id of the volume.",
							Computed:            true,
						},
						"logical_unit_number": schema.Int64Attribute{
							Description:         "The current amount of data written to the volume",
							MarkdownDescription: "The current amount of data written to the volume",
							Computed:            true,
						},
						"volume_group_id": schema.StringAttribute{
							Description:         "The volume group id of the volume.",
							MarkdownDescription: "The volume group id of the volume.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "The description of the volume.",
							MarkdownDescription: "The description of the volume.",
							Computed:            true,
						},
						"appliance_id": schema.StringAttribute{
							Description:         "Unique identifier of the appliance on which the volume is provisioned.",
							MarkdownDescription: "Unique identifier of the appliance on which the volume is provisioned.",
							Computed:            true,
						},
						"protection_policy_id": schema.StringAttribute{
							Description:         "The protection policy assigned to the volume.",
							MarkdownDescription: "The protection policy assigned to the volume.",
							Computed:            true,
						},
						"performance_policy_id": schema.StringAttribute{
							Description:         "The performance policy assigned to the volume.",
							MarkdownDescription: "The performance policy assigned to the volume.",
							Computed:            true,
						},
						"creation_timestamp": schema.StringAttribute{
							Description:         "The creation timestamp of the volume.",
							MarkdownDescription: "The creation timestamp of the volume.",
							Computed:            true,
						},
						"is_replication_destination": schema.BoolAttribute{
							Description:         "Indicates whether this volume is a replication destination.",
							MarkdownDescription: "Indicates whether this volume is a replication destination.",
							Computed:            true,
						},
						"node_affinity": schema.StringAttribute{
							Description:         "The node affinity of the volume.",
							MarkdownDescription: "The node affinity of the volume.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							Description:         "The type of the volume.",
							MarkdownDescription: "The type of the volume.",
							Computed:            true,
						},
						"app_type": schema.StringAttribute{
							Description:         "The app type of the volume.",
							MarkdownDescription: "The app type of the volume.",
							Computed:            true,
						},
						"app_type_other": schema.StringAttribute{
							Description:         "The app type other of the volume.",
							MarkdownDescription: "The app type other of the volume.",
							Computed:            true,
						},
						"wwn": schema.StringAttribute{
							Description:         "The wwn of the volume.",
							MarkdownDescription: "The wwn of the volume.",
							Computed:            true,
						},
						"state": schema.StringAttribute{
							Description:         "The state of the volume.",
							MarkdownDescription: "The state of the volume.",
							Computed:            true,
						},
						"nguid": schema.StringAttribute{
							Description:         "The nguid of the volume.",
							MarkdownDescription: "The nguid of the volume.",
							Computed:            true,
						},
						"nsid": schema.Int64Attribute{
							Description:         "The nsid of the volume.",
							MarkdownDescription: "The nsid of the volume.",
							Computed:            true,
						},
						"logical_used": schema.Int64Attribute{
							Description:         "Current amount of data used by the volume.",
							MarkdownDescription: "Current amount of data used by the volume.",
							Computed:            true,
						},
						"migration_session_id": schema.StringAttribute{
							Description:         "Unique identifier of the migration session assigned to the volume",
							MarkdownDescription: "Unique identifier of the migration session assigned to the volume",
							Computed:            true,
						},
						"metro_replication_session_id": schema.StringAttribute{
							Description:         "Unique identifier of the replication session assigned to the volume",
							MarkdownDescription: "Unique identifier of the replication session assigned to the volume",
							Computed:            true,
						},
						"is_host_access_available": schema.BoolAttribute{
							Description:         "Indicates whether the volume is available to host",
							MarkdownDescription: "Indicates whether the volume is available to host",
							Computed:            true,
						},
						"type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to type",
							MarkdownDescription: "Localized message string corresponding to type",
							Computed:            true,
						},
						"state_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to state",
							MarkdownDescription: "Localized message string corresponding to state",
							Computed:            true,
						},
						"node_affinity_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to node affinity",
							MarkdownDescription: "Localized message string corresponding to node affinity",
							Computed:            true,
						},
						"app_type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to app type",
							MarkdownDescription: "Localized message string corresponding to app type",
							Computed:            true,
						},
						"protection_data": schema.ObjectAttribute{
							Description:         "Specifies the ProtectionData associated with a volume.",
							MarkdownDescription: "Specifies the ProtectionData associated with a volume.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"source_id": types.StringType,
							},
						},
						"location_history": schema.ListNestedAttribute{
							Description:         "Specifies the LocationHistory for a volume.",
							MarkdownDescription: "Specifies the LocationHistory for a volume.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"from_appliance_id": schema.StringAttribute{
										Description:         "Unique identifier of the appliance from which the volume was relocated.",
										MarkdownDescription: "Unique identifier of the appliance from which the volume was relocated.",
										Computed:            true,
									},
									"to_appliance_id": schema.StringAttribute{
										Description:         "Unique identifier of the appliance to which the volume was relocated.",
										MarkdownDescription: "Unique identifier of the appliance to which the volume was relocated.",
										Computed:            true,
									},
									"migrated_on": schema.StringAttribute{
										Description:         "Time when the storage resource location changed.",
										MarkdownDescription: "Time when the storage resource location changed.",
										Computed:            true,
									},
								},
							},
						},
						"appliance": schema.ObjectAttribute{
							Description:         "Specifies the Appliance associated for a volume.",
							MarkdownDescription: "Specifies the Appliance associated for a volume.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"id":          types.StringType,
								"name":        types.StringType,
								"service_tag": types.StringType,
							},
						},
						"migration_session": schema.ObjectAttribute{
							Description:         "Specifies the MigrationSession associated for a volume.",
							MarkdownDescription: "Specifies the MigrationSession associated for a volume.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"id":   types.StringType,
								"name": types.StringType,
							},
						},
						"mapped_volumes": schema.ListNestedAttribute{
							Description:         "Specifies the MappedVolumes associated with a volume.",
							MarkdownDescription: "Specifies the MappedVolumes associated with a volume.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of a mapping between a host and a volume.",
										MarkdownDescription: "Unique identifier of a mapping between a host and a volume.",
										Computed:            true,
									},
								},
							},
						},
						"datastores": schema.ListNestedAttribute{
							Description:         "Specifies the Datastores for a volume.",
							MarkdownDescription: "Specifies the Datastores for a volume.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the datastore instance.",
										MarkdownDescription: "Unique identifier of the datastore instance.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "User-assigned name of the datastore in vCenter.",
										MarkdownDescription: "User-assigned name of the datastore in vCenter.",
										Computed:            true,
									},
									"istance_uuid": schema.StringAttribute{
										Description:         "UUID instance of the datastore in vCenter.",
										MarkdownDescription: "UUID instance of the datastore in vCenter.",
										Computed:            true,
									},
								},
							},
						},
						"volume_groups": schema.ListNestedAttribute{
							Description:         "Specifies the VolumeGroup for a volume.",
							MarkdownDescription: "Specifies the VolumeGroup for a volume.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the volume group.",
										MarkdownDescription: "Unique identifier of the volume group.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Name of the volume group.",
										MarkdownDescription: "Name of the volume group.",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										Description:         "Description for the volume group.",
										MarkdownDescription: "Description for the volume group.",
										Computed:            true,
									},
								},
							},
						},
						"protection_policy": schema.ObjectAttribute{
							Description:         "Specifies the protection policy associated for a volume.",
							MarkdownDescription: "Specifies the protection policy associated for a volume.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"id":          types.StringType,
								"name":        types.StringType,
								"description": types.StringType,
							},
						},
					},
				},
			},
		},
	}
}

func (d *volumeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *volumeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumeDataSourceModel
	var volumes []gopowerstore.Volume
	var volume gopowerstore.Volume
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	//Read the volumes based on volume id/name and if nothing is mentioned, then it returns all the volumes
	if state.Name.ValueString() != "" {
		volume, err = d.client.PStoreClient.GetVolumeByName(context.Background(), state.Name.ValueString())
	} else if state.ID.ValueString() != "" {
		volume, err = d.client.PStoreClient.GetVolume(context.Background(), state.ID.ValueString())
	} else {
		volumes, err = d.client.PStoreClient.GetVolumes(context.Background())
	}
	//check if there is any error while getting the volume
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Volumes",
			err.Error(),
		)
		return
	}
	volumes = append(volumes, volume)
	if len(volumes) == 0 {
		resp.Diagnostics.AddError(
			"No Volume found corresponding to the given name/ID",
			err.Error(),
		)
		return
	}
	state.Volumes, err = updateVolumeState(volumes, d.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update volume state",
			err.Error(),
		)
		return
	}
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// updateVolumeState iterates over the volume list and update the state
func updateVolumeState(volumes []gopowerstore.Volume, p *client.Client) (response []models.VolumeDataSource, err error) {
	for _, volumeValue := range volumes {
		var volGroup types.String
		var hostID string
		var hostGroupID string
		var logicalUnit int64
		size, unit := convertFromBytes(volumeValue.Size)
		hostMapping, err := p.PStoreClient.GetHostVolumeMappingByVolumeID(context.Background(), volumeValue.ID)
		if err != nil {
			return nil, err
		}
		if len(hostMapping) > 0 {
			hostID, hostGroupID, logicalUnit = hostMapping[0].HostID, hostMapping[0].HostGroupID, hostMapping[0].LogicalUnitNumber
		}
		volGroupMapping, err := p.PStoreClient.GetVolumeGroupsByVolumeID(context.Background(), volumeValue.ID)
		if err != nil {
			return nil, err
		}
		if len(volGroupMapping.VolumeGroup) > 0 {
			volGroup = types.StringValue(volGroupMapping.VolumeGroup[0].ID)
		}
		volumeState := models.VolumeDataSource{
			ID:                        types.StringValue(volumeValue.ID),
			Name:                      types.StringValue(volumeValue.Name),
			Size:                      types.Float64Value(size),
			CapacityUnit:              types.StringValue(unit),
			HostID:                    types.StringValue(hostID),
			HostGroupID:               types.StringValue(hostGroupID),
			LogicalUnitNumber:         types.Int64Value(logicalUnit),
			VolumeGroupID:             volGroup,
			Description:               types.StringValue(volumeValue.Description),
			ApplianceID:               types.StringValue(volumeValue.ApplianceID),
			ProtectionPolicyID:        types.StringValue(volumeValue.ProtectionPolicyID),
			PerformancePolicyID:       types.StringValue(volumeValue.PerformancePolicyID),
			CreationTimeStamp:         types.StringValue(volumeValue.CreationTimeStamp),
			IsReplicationDestination:  types.BoolValue(volumeValue.IsReplicationDestination),
			NodeAffinity:              types.StringValue(string(volumeValue.NodeAffinity)),
			Type:                      types.StringValue(string(volumeValue.Type)),
			WWN:                       types.StringValue(volumeValue.Wwn),
			State:                     types.StringValue(string(volumeValue.State)),
			LogicalUsed:               types.Int64Value(volumeValue.LogicalUsed),
			AppType:                   types.StringValue(string(volumeValue.AppType)),
			AppTypeOther:              types.StringValue(volumeValue.AppTypeOther),
			Nsid:                      types.Int64Value(volumeValue.Nsid),
			Nguid:                     types.StringValue(volumeValue.Nguid),
			MigrationSessionID:        types.StringValue(volumeValue.MigrationSessionID),
			MetroReplicationSessionID: types.StringValue(volumeValue.MetroReplicationSessionID),
			TypeL10n:                  types.StringValue(volumeValue.TypeL10n),
			StateL10n:                 types.StringValue(volumeValue.StateL10n),
			NodeAffinityL10n:          types.StringValue(volumeValue.NodeAffinityL10n),
			AppTypeL10n:               types.StringValue(volumeValue.AppTypeL10n),
			ProtectionData: models.ProtectionData{
				SourceID: types.StringValue(volumeValue.ProtectionData.SourceID),
			},

			Appliance: models.Appliance{
				ID:   types.StringValue(volumeValue.Appliance.ID),
				Name: types.StringValue(volumeValue.Appliance.Name),
			},

			ProtectionPolicy: models.VolProtectionPolicy{
				ID:          types.StringValue(volumeValue.ProtectionPolicy.ID),
				Name:        types.StringValue(volumeValue.ProtectionPolicy.Name),
				Description: types.StringValue(volumeValue.ProtectionPolicy.Description),
			},
			MigrationSession: models.MigrationSession{
				ID:   types.StringValue(volumeValue.MigrationSession.ID),
				Name: types.StringValue(volumeValue.MigrationSession.Name),
			},
		}

		for _, history := range volumeValue.LocationHistory {
			volumeState.LocationHistory = append(volumeState.LocationHistory, models.LocationHistory{
				FromApplianceID: types.StringValue(history.FromApplianceId),
				ToApplianceID:   types.StringValue(history.ToApplianceId),
			})
		}
		for _, volume := range volumeValue.MappedVolumes {
			volumeState.MappedVolumes = append(volumeState.MappedVolumes, models.MappedVolumes{
				ID: types.StringValue(volume.ID),
			})
		}
		for _, volumeGroup := range volumeValue.VolumeGroup {
			volumeState.VolumeGroup = append(volumeState.VolumeGroup, models.VolumeGroup{
				ID:   types.StringValue(volumeGroup.ID),
				Name: types.StringValue(volumeGroup.Name),
			})
		}
		for _, datastore := range volumeValue.Datastores {
			volumeState.Datastores = append(volumeState.Datastores, models.Datastores{
				ID:   types.StringValue(datastore.ID),
				Name: types.StringValue(datastore.Name),
			})
		}

		response = append(response, volumeState)
	}
	return response, nil
}
