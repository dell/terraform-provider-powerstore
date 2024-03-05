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

type volumeGroupDataSource struct {
	client *client.Client
}

type volumeGroupDataSourceModel struct {
	VolumeGroups []models.VolumeGroupDataSource `tfsdk:"volume_groups"`
	ID           types.String                   `tfsdk:"id"`
	Name         types.String                   `tfsdk:"name"`
}

var (
	_ datasource.DataSource              = &volumeGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &volumeGroupDataSource{}
)

// newVolumeGroupDataSource returns the volume group data source object
func newVolumeGroupDataSource() datasource.DataSource {
	return &volumeGroupDataSource{}
}

func (d *volumeGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumegroup"
}

func (d *volumeGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing volumegroup from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing volumegroup from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the volume group. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the volume group. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Volume group name. Conflicts with `id`.",
				MarkdownDescription: "Volume group name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"volume_groups": schema.ListNestedAttribute{
				Description:         "List of volume groups.",
				MarkdownDescription: "List of volume groups.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Unique identifier of the volume group.",
							MarkdownDescription: "Unique identifier of the volume group.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Volume group name.",
							MarkdownDescription: "Volume group name.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "Volume group description.",
							MarkdownDescription: "Volume group description.",
							Computed:            true,
						},
						"creation_timestamp": schema.StringAttribute{
							Description:         "The time at which the volume group was created.",
							MarkdownDescription: "The time at which the volume group was created.",
							Computed:            true,
						},
						"is_protectable": schema.BoolAttribute{
							Description:         "This is a derived field that is set internally.",
							MarkdownDescription: "This is a derived field that is set internally.",
							Computed:            true,
						},
						"protection_policy_id": schema.StringAttribute{
							Description:         "Unique identifier of the protection policy assigned to the volume group.",
							MarkdownDescription: "Unique identifier of the protection policy assigned to the volume group.",
							Computed:            true,
						},
						"migration_session_id": schema.StringAttribute{
							Description:         "Unique identifier of the migration session assigned to the volume group when it is part of a migration activity.",
							MarkdownDescription: "Unique identifier of the migration session assigned to the volume group when it is part of a migration activity.",
							Computed:            true,
						},
						"is_write_order_consistent": schema.BoolAttribute{
							Description:         "For a primary or a clone volume group, this property determines whether snapshot sets of the group will be write order consistent.",
							MarkdownDescription: "For a primary or a clone volume group, this property determines whether snapshot sets of the group will be write order consistent.",
							Computed:            true,
						},
						"placement_rule": schema.StringAttribute{
							Description:         "This is set during creation, and determines resource balancer recommendations.",
							MarkdownDescription: "This is set during creation, and determines resource balancer recommendations.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							Description:         "Type of volume.",
							MarkdownDescription: "Type of volume.",
							Computed:            true,
						},
						"is_replication_destination": schema.BoolAttribute{
							Description:         "Indicates whether this volume group is a replication destination.",
							MarkdownDescription: "Indicates whether this volume group is a replication destination.",
							Computed:            true,
						},
						"protection_data": schema.ObjectAttribute{
							Description:         "Specifies the ProtectionData associated with a volume group.",
							MarkdownDescription: "Specifies the ProtectionData associated with a volume group.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"source_id":            types.StringType,
								"creator_type":         types.StringType,
								"expiration_timestamp": types.StringType,
							},
						},
						"is_importing": schema.BoolAttribute{
							Description:         "Indicates whether the volume group is being imported.",
							MarkdownDescription: "Indicates whether the volume group is being imported.",
							Computed:            true,
						},
						"location_history": schema.ListNestedAttribute{
							Description:         "Storage resource location history.",
							MarkdownDescription: "Storage resource location history.",
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
								},
							},
						},
						"type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to type",
							MarkdownDescription: "Localized message string corresponding to type",
							Computed:            true,
						},
						"protection_policy": schema.ObjectAttribute{
							Description:         "Specifies the Protection Policy associated with a volume group.",
							MarkdownDescription: "Specifies the Protection Policy associated with a volume group.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"id":          types.StringType,
								"name":        types.StringType,
								"description": types.StringType,
							},
						},
						"migration_session": schema.ObjectAttribute{
							Description:         "Properties of a migration session.",
							MarkdownDescription: "Properties of a migration session.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"id":   types.StringType,
								"name": types.StringType,
							},
						},
						"volumes": schema.ListNestedAttribute{
							Description:         "List of the volumes that are associated with this volume_group.",
							MarkdownDescription: "List of the volumes that are associated with this volume_group.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the volume instance.",
										MarkdownDescription: "Unique identifier of the volume instance.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Name of the volume.",
										MarkdownDescription: "Name of the volume.",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										Description:         "Description of the volume.",
										MarkdownDescription: "Description of the volume.",
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

func (d *volumeGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *volumeGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumeGroupDataSourceModel
	var volumeGroups []gopowerstore.VolumeGroup
	var volumeGroup gopowerstore.VolumeGroup
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the volume group based on volume group id/name and if nothing is mentioned, then it returns all the volume groups
	if state.Name.ValueString() != "" {
		volumeGroup, err = d.client.PStoreClient.GetVolumeGroupByName(context.Background(), state.Name.ValueString())
		volumeGroups = append(volumeGroups, volumeGroup)
	} else if state.ID.ValueString() != "" {
		volumeGroup, err = d.client.PStoreClient.GetVolumeGroup(context.Background(), state.ID.ValueString())
		volumeGroups = append(volumeGroups, volumeGroup)
	} else {
		volumeGroups, err = d.client.PStoreClient.GetVolumeGroups(context.Background())
	}

	//check if there is any error while getting the volume group
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Volume Group",
			err.Error(),
		)
		return
	}

	state.VolumeGroups, err = updateVolGroupState(volumeGroups, d.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update volume group state",
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

// updateVolGroupState iterates over the volume groups list and update the state
func updateVolGroupState(volumeGroups []gopowerstore.VolumeGroup, p *client.Client) (response []models.VolumeGroupDataSource, err error) {
	for _, volumeGroupValue := range volumeGroups {
		volumeGroupState := models.VolumeGroupDataSource{
			ID:                       types.StringValue(volumeGroupValue.ID),
			Name:                     types.StringValue(volumeGroupValue.Name),
			Description:              types.StringValue(volumeGroupValue.Description),
			CreationTimestamp:        types.StringValue(volumeGroupValue.CreationTimeStamp),
			IsProtectable:            types.BoolValue(volumeGroupValue.IsProtectable),
			ProtectionPolicyID:       types.StringValue(volumeGroupValue.ProtectionPolicyID),
			MigrationSessionID:       types.StringValue(volumeGroupValue.MigrationSessionID),
			IsWriteOrderConsistent:   types.BoolValue(volumeGroupValue.IsWriteOrderConsistent),
			PlacementRule:            types.StringValue(string(volumeGroupValue.PlacementRule)),
			Type:                     types.StringValue(string(volumeGroupValue.Type)),
			IsReplicationDestination: types.BoolValue(volumeGroupValue.IsReplicationDestination),
			IsImporting:              types.BoolValue(volumeGroupValue.IsImporting),
			TypeL10:                  types.StringValue(volumeGroupValue.TypeL10),

			ProtectionData: models.ProtectionData{
				SourceID:            types.StringValue(volumeGroupValue.ProtectionData.SourceID),
				CreatorType:         types.StringValue(volumeGroupValue.ProtectionData.CreatorType),
				ExpirationTimestamp: types.StringValue(volumeGroupValue.ProtectionData.ExpirationTimeStamp),
			},

			ProtectionPolicy: models.VolProtectionPolicy{
				ID:          types.StringValue(volumeGroupValue.ProtectionPolicy.ID),
				Name:        types.StringValue(volumeGroupValue.ProtectionPolicy.Name),
				Description: types.StringValue(volumeGroupValue.ProtectionPolicy.Description),
			},

			MigrationSession: models.MigrationSession{
				ID:   types.StringValue(volumeGroupValue.MigrationSession.ID),
				Name: types.StringValue(volumeGroupValue.MigrationSession.Name),
			},
		}

		for _, history := range volumeGroupValue.LocationHistory {
			volumeGroupState.LocationHistory = append(volumeGroupState.LocationHistory, models.LocationHistory{
				FromApplianceID: types.StringValue(history.FromApplianceID),
				ToApplianceID:   types.StringValue(history.ToApplianceID),
			})
		}

		for _, volume := range volumeGroupValue.Volumes {
			volumeGroupState.Volumes = append(volumeGroupState.Volumes, models.Volumes{
				ID:          types.StringValue(volume.ID),
				Name:        types.StringValue(volume.Name),
				Description: types.StringValue(volume.Description),
			})
		}
		response = append(response, volumeGroupState)
	}
	return response, nil
}
