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

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

type hostGroupDataSource struct {
	client *client.Client
}

type hostGroupDataSourceModel struct {
	HostGroups []models.HostGroupDataSource `tfsdk:"host_groups"`
	ID         types.String                 `tfsdk:"id"`
	Name       types.String                 `tfsdk:"name"`
	Filters    models.FilterExpressionValue `tfsdk:"filter_expression"`
}

var (
	_ datasource.DataSource              = &hostGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &hostGroupDataSource{}
)

// newHostGroupDataSource returns the host group data source object
func newHostGroupDataSource() datasource.DataSource {
	return &hostGroupDataSource{}
}

func (d *hostGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hostgroup"
}

func (d *hostGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing host group from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing host group from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the host group. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the host group. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Host group name. Conflicts with `id`.",
				MarkdownDescription: "Host group name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter HostGroup by. Conflicts with `id` and `name`.",
				MarkdownDescription: "PowerStore filter expression to filter HostGroup by. Conflicts with `id` and `name`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"host_groups": schema.ListNestedAttribute{
				Description:         "List of host groups.",
				MarkdownDescription: "List of host groups.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Unique identifier of the host group.",
							MarkdownDescription: "Unique identifier of the host group.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Host group name.",
							MarkdownDescription: "Host group name.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "Host group description.",
							MarkdownDescription: "Host group description.",
							Computed:            true,
						},
						"host_connectivity": schema.StringAttribute{
							Description:         "Connectivity type for hosts and host groups.",
							MarkdownDescription: "Connectivity type for hosts and host groups.",
							Computed:            true,
						},
						"host_connectivity_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to host_connectivity",
							MarkdownDescription: "Localized message string corresponding to host_connectivity",
							Computed:            true,
						},
						"hosts": schema.ListNestedAttribute{
							Description:         "Properties of a host.",
							MarkdownDescription: "Properties of a host.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the host.",
										MarkdownDescription: "Unique identifier of the host.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "The host name.",
										MarkdownDescription: "The host name.",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										Description:         "A description for the host.",
										MarkdownDescription: "A description for the host.",
										Computed:            true,
									},
								},
							},
						},
						"mapped_host_groups": schema.ListNestedAttribute{
							Description:         "Details about a configured host or host group attached to a volume.",
							MarkdownDescription: "Details about a configured host or host group attached to a volume.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of a mapping between a host and a volume.",
										MarkdownDescription: "Unique identifier of a mapping between a host and a volume.",
										Computed:            true,
									},
									"host_id": schema.StringAttribute{
										Description:         "Unique identifier of a host attached to a volume.",
										MarkdownDescription: "Unique identifier of a host attached to a volume.",
										Computed:            true,
									},
									"volume_id": schema.StringAttribute{
										Description:         "Unique identifier of the volume to which the host is attached.",
										MarkdownDescription: "Unique identifier of the volume to which the host is attached.",
										Computed:            true,
									},
									"volume_name": schema.StringAttribute{
										Description:         "Name of the volume to which the host is attached.",
										MarkdownDescription: "Name of the volume to which the host is attached.",
										Computed:            true,
									},
								},
							},
						},
						"host_virtual_volume_mappings": schema.ListNestedAttribute{
							Description:         "Virtual volume mapping details.",
							MarkdownDescription: "Virtual volume mapping details.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of a mapping between a host and a virtual volume.",
										MarkdownDescription: "Unique identifier of a mapping between a host and a virtual volume.",
										Computed:            true,
									},
									"host_id": schema.StringAttribute{
										Description:         "Unique identifier of a host attached to a volume.",
										MarkdownDescription: "Unique identifier of a host attached to a volume.",
										Computed:            true,
									},
									"virtual_volume_id": schema.StringAttribute{
										Description:         "Unique identifier of the virtual volume to which the host is attached.",
										MarkdownDescription: "Unique identifier of the virtual volume to which the host is attached.",
										Computed:            true,
									},
									"virtual_volume_name": schema.StringAttribute{
										Description:         "Name of the virtual volume to which the host is attached.",
										MarkdownDescription: "Name of the virtual volume to which the host is attached.",
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

func (d *hostGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *hostGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state hostGroupDataSourceModel
	var hostGroups []gopowerstore.HostGroup
	var hostGroup gopowerstore.HostGroup
	var err error
	filterMap := make(map[string]string)

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the host group based on host group id/name and if nothing is mentioned, then it returns all the host groups
	if state.Name.ValueString() != "" {
		hostGroup, err = d.client.PStoreClient.GetHostGroupByName(context.Background(), state.Name.ValueString())
		hostGroups = append(hostGroups, hostGroup)
	} else if state.ID.ValueString() != "" {
		hostGroup, err = d.client.PStoreClient.GetHostGroup(context.Background(), state.ID.ValueString())
		hostGroups = append(hostGroups, hostGroup)
	} else if state.Filters.ValueString() != "" {
		err = validateFileSystemFilter(state.Filters.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid filter expression",
				err.Error(),
			)
			return
		}
		filterMap = convertQueriesToMap(state.Filters.ValueQueries())
		hostGroups, err = d.client.GetHostGroupByFilter(ctx, filterMap)
	} else {
		hostGroups, err = d.client.PStoreClient.GetHostGroups(context.Background())
	}

	//check if there is any error while getting the host group
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Host Group",
			err.Error(),
		)
		return
	}

	state.HostGroups, err = updateHostGroupState(hostGroups, d.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update host group state",
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

// updateHostGroupState iterates over the host groups list and update the state
func updateHostGroupState(hostGroups []gopowerstore.HostGroup, p *client.Client) (response []models.HostGroupDataSource, err error) {
	for _, hostGroupValue := range hostGroups {
		hostGroupState := models.HostGroupDataSource{
			ID:                   types.StringValue(hostGroupValue.ID),
			Name:                 types.StringValue(hostGroupValue.Name),
			Description:          types.StringValue(hostGroupValue.Description),
			HostConnectivity:     types.StringValue(string(hostGroupValue.HostConnectivity)),
			HostConnectivityL10n: types.StringValue(hostGroupValue.HostConnectivityL10n),
		}

		for _, host := range hostGroupValue.Hosts {
			hostGroupState.Hosts = append(hostGroupState.Hosts, models.Hosts{
				ID:          types.StringValue(host.ID),
				Name:        types.StringValue(host.Name),
				Description: types.StringValue(host.Description),
			})
		}

		for _, mappedHostGroup := range hostGroupValue.MappedHostGroups {
			hostGroupState.MappedHostGroups = append(hostGroupState.MappedHostGroups, models.MappedHostGroup{
				ID:         types.StringValue(mappedHostGroup.ID),
				HostID:     types.StringValue(mappedHostGroup.HostID),
				VolumeID:   types.StringValue(mappedHostGroup.VolumeID),
				VolumeName: types.StringValue(mappedHostGroup.Volume.Name),
			})
		}

		for _, hostVirtualVolumeMapping := range hostGroupValue.HostVirtualVolumeMappings {
			hostGroupState.HostVirtualVolumeMappings = append(hostGroupState.HostVirtualVolumeMappings, models.HostVirtualVolumeMappings{
				ID:                types.StringValue(hostVirtualVolumeMapping.ID),
				HostID:            types.StringValue(hostVirtualVolumeMapping.HostID),
				VirtualVolumeID:   types.StringValue(hostVirtualVolumeMapping.VirtualVolumeID),
				VirtualVolumeName: types.StringValue(hostVirtualVolumeMapping.VirtualVolume.Name),
			})
		}

		response = append(response, hostGroupState)
	}
	return response, nil
}
