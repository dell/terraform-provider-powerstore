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
	"net/url"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type volumeGroupDataSource struct {
	client *clientgen.APIClient
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
	c := req.ProviderData.(*client.Client)
	d.client = c.GenClient
}

func (d *volumeGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumeGroupDataSourceModel
	var volumeGroups []clientgen.VolumeGroupInstance
	var volumeGroup *clientgen.VolumeGroupInstance
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sel := "id,name,description,creation_timestamp,is_protectable,protection_policy_id,migration_session_id,is_write_order_consistent,placement_rule,is_importing,is_replication_destination,type,type_l10n,volumes(id,name,description),protection_policy(id,name,description),protection_data(source_id,creator_type,expiration_timestamp),location_history(from_appliance_id,to_appliance_id),migration_session(id,name)"
	//Read the volume group based on volume group id/name and if nothing is mentioned, then it returns all the volume groups
	if state.Name.ValueString() != "" {
		hreq := d.client.VolumeGroupApi.GetAllVolumeGroups(ctx)
		hreq.Queries = url.Values{
			"name":   []string{state.Name.ValueString()},
			"select": []string{sel},
		}
		volumeGroups, _, err = hreq.Execute()
	} else if state.ID.ValueString() != "" {
		hreq := d.client.VolumeGroupApi.GetVolumeGroupById(ctx, state.ID.ValueString())
		hreq.Queries = url.Values{
			"select": []string{sel},
		}
		volumeGroup, _, err = hreq.Execute()
		volumeGroups = append(volumeGroups, *volumeGroup)
	} else {
		hreq := d.client.VolumeGroupApi.GetAllVolumeGroups(ctx)
		hreq.Queries = url.Values{
			"select": []string{sel},
		}
		volumeGroups, _, err = hreq.Execute()
	}

	//check if there is any error while getting the volume group
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Volume Group",
			err.Error(),
		)
		return
	}

	state.VolumeGroups, err = updateVolGroupState(volumeGroups)
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

func pointerToStringType(in *string) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue(*in)
}

func pTimeToStringType(in *time.Time) types.String {
	if in == nil {
		return types.StringNull()
	}
	return types.StringValue((*in).String())
}

func pointerToBoolType(in *bool) types.Bool {
	if in == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*in)
}

// updateVolGroupState iterates over the volume groups list and update the state
func updateVolGroupState(volumeGroups []clientgen.VolumeGroupInstance) (response []models.VolumeGroupDataSource, err error) {
	for _, volumeGroupValue := range volumeGroups {
		volumeGroupState := models.VolumeGroupDataSource{
			ID:                       pointerToStringType(volumeGroupValue.Id),
			Name:                     pointerToStringType(volumeGroupValue.Name),
			Description:              pointerToStringType(volumeGroupValue.Description),
			CreationTimestamp:        pTimeToStringType(volumeGroupValue.CreationTimestamp),
			IsProtectable:            pointerToBoolType(volumeGroupValue.IsProtectable),
			ProtectionPolicyID:       pointerToStringType(volumeGroupValue.ProtectionPolicyId),
			MigrationSessionID:       pointerToStringType(volumeGroupValue.MigrationSessionId),
			IsWriteOrderConsistent:   pointerToBoolType(volumeGroupValue.IsWriteOrderConsistent),
			PlacementRule:            types.StringValue(volumeGroupValue.PlacementRule.Value()),
			Type:                     types.StringValue(volumeGroupValue.Type.Value()),
			IsReplicationDestination: pointerToBoolType(volumeGroupValue.IsReplicationDestination),
			IsImporting:              pointerToBoolType(volumeGroupValue.IsImporting),
			TypeL10:                  pointerToStringType(volumeGroupValue.TypeL10n),

			ProtectionData: models.ProtectionData{
				SourceID:            pointerToStringType(volumeGroupValue.ProtectionData.SourceId),
				CreatorType:         types.StringValue(volumeGroupValue.ProtectionData.CreatorType.Value()),
				ExpirationTimestamp: pTimeToStringType(volumeGroupValue.ProtectionData.ExpirationTimestamp),
			},

			ProtectionPolicy: models.VolProtectionPolicy{
				ID:          pointerToStringType(volumeGroupValue.ProtectionPolicy.Id),
				Name:        pointerToStringType(volumeGroupValue.ProtectionPolicy.Name),
				Description: pointerToStringType(volumeGroupValue.ProtectionPolicy.Description),
			},

			MigrationSession: models.MigrationSession{
				ID:   pointerToStringType(volumeGroupValue.MigrationSession.Id),
				Name: pointerToStringType(volumeGroupValue.MigrationSession.Name),
			},
		}

		for _, history := range volumeGroupValue.LocationHistory {
			volumeGroupState.LocationHistory = append(volumeGroupState.LocationHistory, models.LocationHistory{
				FromApplianceID: pointerToStringType(history.FromApplianceId),
				ToApplianceID:   pointerToStringType(history.ToApplianceId),
			})
		}

		for _, volume := range volumeGroupValue.Volumes {
			volumeGroupState.Volumes = append(volumeGroupState.Volumes, models.Volumes{
				ID:          pointerToStringType(volume.Id),
				Name:        pointerToStringType(volume.Name),
				Description: pointerToStringType(volume.Description),
			})
		}
		response = append(response, volumeGroupState)
	}
	return response, nil
}
