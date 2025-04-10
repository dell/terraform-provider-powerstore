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
	_ datasource.DataSource              = &hostDataSource{}
	_ datasource.DataSourceWithConfigure = &hostDataSource{}
)

// newHostDataSource returns the host data source object
func newHostDataSource() datasource.DataSource {
	return &hostDataSource{}
}

type hostDataSource struct {
	client *client.Client
}

type hostDataSourceModel struct {
	Hosts   []models.HostDataSource      `tfsdk:"host"`
	ID      types.String                 `tfsdk:"id"`
	Name    types.String                 `tfsdk:"name"`
	Filters models.FilterExpressionValue `tfsdk:"filter_expression"`
}

func (d *hostDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_host"
}

func (d *hostDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing host from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing host from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the host instance. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the host instance. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the host. Conflicts with `id`.",
				MarkdownDescription: "Name of the host. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter Host by. Conflicts with `id` and `name`.",
				MarkdownDescription: "PowerStore filter expression to filter Host by. Conflicts with `id` and `name`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"host": schema.ListNestedAttribute{
				Description:         "List of hosts.",
				MarkdownDescription: "List of hosts.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "The ID of the host.",
							MarkdownDescription: "The ID of the host.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Name of the host.",
							MarkdownDescription: "Name of the host.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Computed:            true,
							Description:         "Description of the host.",
							MarkdownDescription: "Description of the host.",
						},
						"host_group_id": schema.StringAttribute{
							Computed:            true,
							Description:         "Associated host group, if host is part of host group.",
							MarkdownDescription: "Associated host group, if host is part of host group.",
						},
						"os_type": schema.StringAttribute{
							Computed:            true,
							Description:         "Operating system of the host.",
							MarkdownDescription: "Operating system of the host.",
						},
						"host_connectivity": schema.StringAttribute{
							Computed:            true,
							Description:         "Connectivity type for hosts.",
							MarkdownDescription: "Connectivity type for hosts.",
						},
						"type": schema.StringAttribute{
							Computed:            true,
							Description:         "Type of hosts.",
							MarkdownDescription: "Type of hosts.",
						},
						"type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to type",
							MarkdownDescription: "Localized message string corresponding to type",
							Computed:            true,
						},
						"os_type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to OS type.",
							MarkdownDescription: "Localized message string corresponding to OS type.",
							Computed:            true,
						},
						"host_connectivity_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to host connectivity.",
							MarkdownDescription: "Localized message string corresponding to host connectivity.",
							Computed:            true,
						},
						"initiators": schema.ListNestedAttribute{
							Description:         "Initiator instance.",
							MarkdownDescription: "Initiator instance.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"port_type": schema.StringAttribute{
										Description:         "Protocol type of the host initiator.",
										MarkdownDescription: "Protocol type of the host initiator.",
										Computed:            true,
									},
									"port_name": schema.StringAttribute{
										Description:         "The port name, one of: IQN, WWN, or NQN..",
										MarkdownDescription: "The port name, one of: IQN, WWN, or NQN..",
										Computed:            true,
									},
									"chap_mutual_username": schema.StringAttribute{
										Description:         "Username for CHAP authentication.",
										MarkdownDescription: "Username for CHAP authentication.",
										Computed:            true,
									},
									"chap_single_username": schema.StringAttribute{
										Description:         "Username for CHAP authentication.",
										MarkdownDescription: "Username for CHAP authentication.",
										Computed:            true,
									},
								},
							},
						},
						"import_host_system": schema.ObjectAttribute{
							Description:         "Details about an import host system.",
							MarkdownDescription: "Details about an import host system.",
							Computed:            true,
							AttributeTypes: map[string]attr.Type{
								"id":            types.StringType,
								"agent_address": types.StringType,
							},
						},
						"mapped_hosts": schema.ListNestedAttribute{
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
										Description:         "Unique identifier of a host attached to a virtual volume.",
										MarkdownDescription: "Unique identifier of a host attached to a virtual volume.",
										Computed:            true,
									},
									"virtual_volume_id": schema.StringAttribute{
										Description:         "Unique identifier of the virtual volume to which the host is attached.",
										MarkdownDescription: "Unique identifier of the virtual volume to which the host is attached.",
										Computed:            true,
									},
								},
							},
						},
						"vsphere_hosts": schema.ListNestedAttribute{
							Description:         "List of the vsphere hosts that are associated with this host.",
							MarkdownDescription: "List of the vsphere hosts that are associated with this host.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the vsphere_host instance.",
										MarkdownDescription: "Unique identifier of the vsphere_host instance.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "User-assigned name of the ESXi host in vCenter.",
										MarkdownDescription: "User-assigned name of the ESXi host in vCenter.",
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

func (d *hostDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *hostDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state hostDataSourceModel
	var hosts []gopowerstore.Host
	var host gopowerstore.Host
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the hosts based on host id/name and if nothing is mentioned, then it returns all the hosts
	if state.Name.ValueString() != "" {
		host, err = d.client.PStoreClient.GetHostByName(context.Background(), state.Name.ValueString())
		hosts = append(hosts, host)
	} else if state.ID.ValueString() != "" {
		host, err = d.client.PStoreClient.GetHost(context.Background(), state.ID.ValueString())
		hosts = append(hosts, host)
	} else if state.Filters.ValueString() != "" {
		err = validateFileSystemFilter(state.Filters.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid filter expression",
				err.Error(),
			)
			return
		}
		filterMap := convertQueriesToMap(state.Filters.ValueQueries())
		hosts, err = d.client.GetHostByFilter(ctx, filterMap)
	} else {
		hosts, err = d.client.PStoreClient.GetHosts(context.Background())
	}
	//check if there is any error while getting the host
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Hosts",
			err.Error(),
		)
		return
	}

	// Update the state
	state.Hosts, err = updateHostState(hosts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update host state",
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

// updateHostState iterates over the host list and update the state
func updateHostState(Hosts []gopowerstore.Host) (response []models.HostDataSource, err error) {
	for _, HostValue := range Hosts {
		var hostState = models.HostDataSource{
			ID:                   types.StringValue(HostValue.ID),
			Name:                 types.StringValue(HostValue.Name),
			Description:          types.StringValue(HostValue.Description),
			HostGroupID:          types.StringValue(HostValue.HostGroupID),
			OsType:               types.StringValue(string(HostValue.OsType)),
			HostConnectivity:     types.StringValue(string(HostValue.HostConnectivity)),
			Type:                 types.StringValue(string(HostValue.Type)),
			TypeL10n:             types.StringValue(HostValue.TypeL10n),
			OsTypeL10n:           types.StringValue(HostValue.OSTypeL10n),
			HostConnectivityL10n: types.StringValue(HostValue.HostConnectivityL10n),
			ImportHostSystem: models.ImportHostSystem{
				ID:           types.StringValue(HostValue.ImportHostSystem.ID),
				AgentAddress: types.StringValue(HostValue.ImportHostSystem.AgentAddress),
			},
		}

		for _, mappedHost := range HostValue.MappedHosts {
			hostState.MappedHosts = append(hostState.MappedHosts, models.MappedHosts{
				ID:       types.StringValue(mappedHost.ID),
				HostID:   types.StringValue(mappedHost.HostID),
				VolumeID: types.StringValue(mappedHost.VolumeID),
			})
		}

		for _, hostVirtualVolumeMappings := range HostValue.HostVirtualVolumeMappings {
			hostState.HostVirtualVolumeMappings = append(hostState.HostVirtualVolumeMappings, models.HostVirtualVolumeMappings{
				ID:              types.StringValue(hostVirtualVolumeMappings.ID),
				HostID:          types.StringValue(hostVirtualVolumeMappings.HostID),
				VirtualVolumeID: types.StringValue(hostVirtualVolumeMappings.VirtualVolumeID),
			})
		}

		for _, vsphereHosts := range HostValue.VsphereHosts {
			hostState.VsphereHosts = append(hostState.VsphereHosts, models.VsphereHosts{
				ID:   types.StringValue(vsphereHosts.ID),
				Name: types.StringValue(vsphereHosts.Name),
			})
		}

		for _, initiators := range HostValue.Initiators {
			hostState.Initiators = append(hostState.Initiators, models.InitiatorInstance{
				PortName:           types.StringValue(initiators.PortName),
				PortType:           types.StringValue(string(initiators.PortType)),
				ChapMutualUsername: types.StringValue(initiators.ChapMutualUsername),
				ChapSingleUsername: types.StringValue(initiators.ChapSingleUsername),
			})
		}
		response = append(response, hostState)
	}
	return response, nil
}
