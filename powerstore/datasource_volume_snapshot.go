package powerstore

import (
	"context"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &volumeSnapshotDataSource{}
	_ datasource.DataSourceWithConfigure = &volumeSnapshotDataSource{}
)

// newVolumeSnapshotDataSource returns the volume snapshot data source object
func newVolumeSnapshotDataSource() datasource.DataSource {
	return &volumeSnapshotDataSource{}
}

type volumeSnapshotDataSource struct {
	client *client.Client
}

type volumeSnapshotDataSourceModel struct {
	Snapshots []models.VolumeSnapshotDataSource `tfsdk:"snapshots"`
	ID        types.String                      `tfsdk:"id"`
	Name      types.String                      `tfsdk:"name"`
}

func (d *volumeSnapshotDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volume_snapshot"
}

func (d *volumeSnapshotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "VolumeSnapshot DataSource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the volume snapshot instance.",
				MarkdownDescription: "Unique identifier of the volume snapshot instance.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the volume snaphshot.",
				MarkdownDescription: "Name of the volume snapshot.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"snapshots": schema.ListNestedAttribute{
				Description:         "List of snapshots.",
				MarkdownDescription: "List of snapshots.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "The ID of the volume snapshot.",
							MarkdownDescription: "The ID of the volume snapshot.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the volume snapshot.",
							MarkdownDescription: "Name of the volume snapshot.",
							Computed:            true,
						},
						"size": schema.Float64Attribute{
							Description:         "The size of the volume.",
							MarkdownDescription: "The size of the volume.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "The description of the volume snapshot.",
							MarkdownDescription: "The description of the volume snapshot.",
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
								"source_id":            types.StringType,
								"creator_type":         types.StringType,
								"expiration_timestamp": types.StringType,
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
					},
				},
			},
		},
	}
}

func (d *volumeSnapshotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *volumeSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumeSnapshotDataSourceModel
	var volumes []gopowerstore.Volume
	var volume gopowerstore.Volume
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the snapshot based on snapshot id/name and if nothing is mentioned, then it returns all the snapshots
	if state.Name.ValueString() != "" {
		volume, err = d.client.PStoreClient.GetSnapshotByName(context.Background(), state.Name.ValueString())
		volumes = append(volumes, volume)
	} else if state.ID.ValueString() != "" {
		volume, err = d.client.PStoreClient.GetSnapshot(context.Background(), state.ID.ValueString())
		volumes = append(volumes, volume)
	} else {
		volumes, err = d.client.PStoreClient.GetSnapshots(context.Background())
	}

	//check if there is any error while getting the volume snapshot
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Volume Snapshots",
			err.Error(),
		)
		return
	}

	state.Snapshots, err = updateVolumeSnapshotState(volumes, d.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update volume snapshot state",
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

// updateVolumeSnapshotState iterates over the snapshot list and update the state
func updateVolumeSnapshotState(volumes []gopowerstore.Volume, p *client.Client) (response []models.VolumeSnapshotDataSource, err error) {
	for _, volumeValue := range volumes {
		size, _ := convertFromBytes(volumeValue.Size)
		volumeSnapshotState := models.VolumeSnapshotDataSource{
			ID:                        types.StringValue(volumeValue.ID),
			Name:                      types.StringValue(volumeValue.Name),
			Size:                      types.Float64Value(size),
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
				SourceID:            types.StringValue(volumeValue.ProtectionData.SourceID),
				CreatorType:         types.StringValue(volumeValue.ProtectionData.CreatorType),
				ExpirationTimestamp: types.StringValue(volumeValue.ProtectionData.ExpirationTimeStamp),
			},
		}

		for _, history := range volumeValue.LocationHistory {
			volumeSnapshotState.LocationHistory = append(volumeSnapshotState.LocationHistory, models.LocationHistory{
				FromApplianceID: types.StringValue(history.FromApplianceId),
				ToApplianceID:   types.StringValue(history.ToApplianceId),
			})
		}

		response = append(response, volumeSnapshotState)
	}
	return response, nil
}
