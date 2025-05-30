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
	"net/url"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type volumeGroupSnapshotDataSource struct {
	client *clientgen.APIClient
}

var (
	_ datasource.DataSource              = &volumeGroupSnapshotDataSource{}
	_ datasource.DataSourceWithConfigure = &volumeGroupSnapshotDataSource{}
)

// newVolumeGroupSnapshotDataSource returns the volume group snapshot data source object
func newVolumeGroupSnapshotDataSource() datasource.DataSource {
	return &volumeGroupSnapshotDataSource{}
}

func (d *volumeGroupSnapshotDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_volumegroup_snapshot"
}

func (d *volumeGroupSnapshotDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing volumegroup snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing volumegroup snapshot from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the volume group snapshot. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the volume group snapshot. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Volume group snapshot name. Conflicts with `id`.",
				MarkdownDescription: "Volume group snapshot name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Replication rules by. Conflicts with `id` and `name`.",
				MarkdownDescription: "PowerStore filter expression to filter Replication rules by. Conflicts with `id` and `name`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"volume_groups": schema.ListNestedAttribute{
				Description:         "List of volume group snapshots.",
				MarkdownDescription: "List of volume group snapshots.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Unique identifier of the volume group snapshot.",
							MarkdownDescription: "Unique identifier of the volume group snapshot.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Volume group snapshot name.",
							MarkdownDescription: "Volume group snapshot name.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "Volume group snapshot description.",
							MarkdownDescription: "Volume group snapshot description.",
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

func (d *volumeGroupSnapshotDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(*client.Client)
	d.client = client.GenClient
}

func (d *volumeGroupSnapshotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state volumeGroupDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the volume group snapshot based on volume group snapshot id/name and if nothing is mentioned, then it returns all the volume group snapshots
	sel := "*,volumes(*),protection_policy(*),protection_data,location_history,migration_session(*)"
	queries := make(url.Values)
	queries.Set("select", sel)
	queries.Set("type", fmt.Sprintf("eq.%s", clientgen.VOLUMETYPEENUM_SNAPSHOT))
	dsreq := helper.DsReq[clientgen.VolumeGroupInstance, clientgen.ApiGetVolumeGroupByIdRequest, clientgen.ApiGetAllVolumeGroupsRequest]{
		Instance:   d.client.VolumeGroupApi.GetVolumeGroupById,
		Collection: d.client.VolumeGroupApi.GetAllVolumeGroups,
	}

	id := state.ID.ValueString()
	if !state.Name.IsNull() {
		queries.Set("name", "eq."+state.Name.ValueString())
	} else if !state.Filters.IsNull() {
		err := validateVolumeFilter(state.Filters.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("filter_expression"),
				"Invalid filter expression",
				err.Error(),
			)
			return
		}
		queries = helper.MergeValues(queries, state.Filters.ValueQueries())
	}
	volumeGroups, err := dsreq.Execute(ctx, queries, id)

	//check if there is any error while getting the volume group
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Volume Group Snapshot",
			err.Error(),
		)
		return
	}

	// check that there is atleast one volume group if name is provided
	if state.Name.ValueString() != "" && len(volumeGroups) == 0 {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Volume Group Snapshot",
			"There is no volume group with name "+state.Name.ValueString(),
		)
		return
	}

	state.VolumeGroups = updateVolGroupState(volumeGroups)
	state.ID = types.StringValue("placeholder")
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
