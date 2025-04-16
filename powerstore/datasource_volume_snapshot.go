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

func (d *volumeSnapshotDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volume_snapshot"
}

func (d *volumeSnapshotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing volume snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing volume snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the volume snapshot instance.",
				MarkdownDescription: "Unique identifier of the volume snapshot instance.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the volume snaphshot.",
				MarkdownDescription: "Name of the volume snapshot.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
				},
			},

			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Volume Snapshots by. Conflicts with `id` and `name`.",
				MarkdownDescription: "PowerStore filter expression to filter Volume Snapshots by. Conflicts with `id` and `name`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},

			"volumes": schema.ListNestedAttribute{
				Description:         "List of volumes.",
				MarkdownDescription: "List of volumes.",
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
							Description:         "The performance policy assigned to the volume snapshot.",
							MarkdownDescription: "The performance policy assigned to the volume snapshot.",
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
									"instance_uuid": schema.StringAttribute{
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

func (d *volumeSnapshotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *volumeSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumeDataSourceModel
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
	} else if state.Filters.ValueString() != "" {
		err = validateVolumeFilter(state.Filters.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("filter_expression"),
				"Invalid filter expression",
				err.Error(),
			)
			return
		}
		filterMap := convertQueriesToMap(state.Filters.ValueQueries())
		volumes, err = d.client.GetVolumesSnapshotsByFilter(ctx, filterMap)

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

	state.Volumes, err = updateVolumeState(volumes, d.client)
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
