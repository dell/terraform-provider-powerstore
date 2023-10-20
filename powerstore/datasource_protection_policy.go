/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type protectionPolicyDataSource struct {
	client *client.Client
}

type protectionPolicyDataSourceModel struct {
	Policies []models.ProtectionPolicyDataSource `tfsdk:"policies"`
	ID       types.String                        `tfsdk:"id"`
	Name     types.String                        `tfsdk:"name"`
}

var (
	_ datasource.DataSource              = &protectionPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &protectionPolicyDataSource{}
)

// newProtectionPolicyDataSource returns the protecion policy data source object
func newProtectionPolicyDataSource() datasource.DataSource {
	return &protectionPolicyDataSource{}
}

func (d *protectionPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protectionpolicy"
}

func (d *protectionPolicyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the existing protection policy from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		MarkdownDescription: "This datasource is used to query the existing protection policy from PowerStore array. The information fetched from this datasource can be used for getting the details for further processing in resource block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the protection policy. Conflicts with `name`.",
				MarkdownDescription: "Unique identifier of the protection policy. Conflicts with `name`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("name")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Protection policy name. Conflicts with `id`.",
				MarkdownDescription: "Protection policy name. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"policies": schema.ListNestedAttribute{
				Description:         "List of protection policies.",
				MarkdownDescription: "List of protection policies.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Unique identifier of the protection policy.",
							MarkdownDescription: "Unique identifier of the protection policy.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "Protection policy name.",
							MarkdownDescription: "Protection policy name.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							Description:         "Protection policy description.",
							MarkdownDescription: "Protection policy description.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							Description:         "Type of the protection policy.",
							MarkdownDescription: "Type of the protection policy.",
							Computed:            true,
						},
						"managed_by": schema.StringAttribute{
							Description:         "Entity that owns and manages this instance.",
							MarkdownDescription: "Entity that owns and manages this instance.",
							Computed:            true,
						},
						"managed_by_id": schema.StringAttribute{
							Description:         "Unique identifier of the managing entity based on the value of the managed_by property",
							MarkdownDescription: "Unique identifier of the managing entity based on the value of the managed_by property",
							Computed:            true,
						},
						"is_read_only": schema.BoolAttribute{
							Description:         "Indicates whether this protection policy can be modified.",
							MarkdownDescription: "Indicates whether this protection policy can be modified.",
							Computed:            true,
						},
						"is_replica": schema.BoolAttribute{
							Description:         "Indicates if this is a replica of a protection policy on a remote system",
							MarkdownDescription: "Indicates if this is a replica of a protection policy on a remote system",
							Computed:            true,
						},
						"type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to type",
							MarkdownDescription: "Localized message string corresponding to type",
							Computed:            true,
						},
						"managed_by_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to managed_by",
							MarkdownDescription: "Localized message string corresponding to managed_by",
							Computed:            true,
						},
						"virtual_machines": schema.ListNestedAttribute{
							Description:         "Specifies the virtual machines associated with a protection policy.",
							MarkdownDescription: "Specifies the virtual machines associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "The unique identifier of the virtual machine.",
										MarkdownDescription: "The unique identifier of the virtual machine.",
										Computed:            true,
									},
									"instance_uuid": schema.StringAttribute{
										Description:         "UUID instance of the VM in vCenter.",
										MarkdownDescription: "UUID instance of the VM in vCenter.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "User-assigned name of the VM in vCenter.",
										MarkdownDescription: "User-assigned name of the VM in vCenter.",
										Computed:            true,
									},
								},
							},
						},
						"volumes": schema.ListNestedAttribute{
							Description:         "Specifies the volumes associated with a protection policy.",
							MarkdownDescription: "Specifies the volumes associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the volume instance.",
										MarkdownDescription: "Unique identifier of the volume instance.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "The name of the volume",
										MarkdownDescription: "The name of the volume",
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
						"volume_groups": schema.ListNestedAttribute{
							Description:         "Specifies the volume group associated with a protection policy.",
							MarkdownDescription: "Specifies the volume group associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "The unique identifier of the volume group.",
										MarkdownDescription: "The unique identifier of the volume group.",
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
						"file_systems": schema.ListNestedAttribute{
							Description:         "Specifies the virtual volumes associated with a protection policy.",
							MarkdownDescription: "Specifies the virtual volumes associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the file system.",
										MarkdownDescription: "Unique identifier of the file system.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Name of the file system.",
										MarkdownDescription: "Name of the file system.",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										Description:         "Description of the file system.",
										MarkdownDescription: "Description of the file system.",
										Computed:            true,
									},
								},
							},
						},
						"performance_rules": schema.ListNestedAttribute{
							Description:         "Specifies the performance rule associated with a protection policy.",
							MarkdownDescription: "Specifies the performance rule associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier representing a performance rule.",
										MarkdownDescription: "Unique identifier representing a performance rule.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Name of the performance rule.",
										MarkdownDescription: "Name of the performance rule.",
										Computed:            true,
									},
									"io_priority": schema.StringAttribute{
										Description:         "The I/O priority for quality of service rules.",
										MarkdownDescription: "The I/O priority for quality of service rules.",
										Computed:            true,
									},
								},
							},
						},
						"snapshot_rules": schema.ListNestedAttribute{
							Description:         "Specifies the snapshot rule associated with a protection policy.",
							MarkdownDescription: "Specifies the snapshot rule associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the snapshot rule.",
										MarkdownDescription: "Unique identifier of the snapshot rule.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Snapshot rule name.",
										MarkdownDescription: "Snapshot rule name.",
										Computed:            true,
									},
								},
							},
						},
						"replication_rules": schema.ListNestedAttribute{
							Description:         "Specifies the replication rule associated with a protection policy.",
							MarkdownDescription: "Specifies the replication rule associated with a protection policy.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "Unique identifier of the replication rule.",
										MarkdownDescription: "Unique identifier of the replication rule.",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										Description:         "Name of the replication rule.",
										MarkdownDescription: "Name of the replication rule.",
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

func (d *protectionPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *protectionPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state protectionPolicyDataSourceModel
	var policies []gopowerstore.ProtectionPolicy
	var policy gopowerstore.ProtectionPolicy
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Read the protection policy based on protection policy id/name and if nothing is mentioned, then it returns all the protection policies
	if state.Name.ValueString() != "" {
		policy, err = d.client.PStoreClient.GetProtectionPolicyByName(context.Background(), state.Name.ValueString())
		policies = append(policies, policy)
	} else if state.ID.ValueString() != "" {
		policy, err = d.client.PStoreClient.GetProtectionPolicy(context.Background(), state.ID.ValueString())
		policies = append(policies, policy)
	} else {
		policies, err = d.client.PStoreClient.GetProtectionPolicies(context.Background())
	}

	//check if there is any error while getting the protection policy
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Protection Policy",
			err.Error(),
		)
		return
	}

	state.Policies, err = updatePolicyState(policies, d.client)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update protection policy state",
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

// updatePolicyState iterates over the protection policy list and update the state
func updatePolicyState(policies []gopowerstore.ProtectionPolicy, p *client.Client) (response []models.ProtectionPolicyDataSource, err error) {
	for _, policyValue := range policies {
		policyState := models.ProtectionPolicyDataSource{
			ID:           types.StringValue(policyValue.ID),
			Name:         types.StringValue(policyValue.Name),
			Description:  types.StringValue(policyValue.Description),
			Type:         types.StringValue(policyValue.Type),
			ManagedBy:    types.StringValue(policyValue.ManagedBy),
			ManagedByID:  types.StringValue(policyValue.ManagedByID),
			IsReadOnly:   types.BoolValue(policyValue.IsReadOnly),
			IsReplica:    types.BoolValue(policyValue.IsReplica),
			TypeL10:      types.StringValue(policyValue.TypeL10),
			ManagedByL10: types.StringValue(policyValue.ManagedByL10),
		}

		for _, volume := range policyValue.VirtualMachines {
			policyState.VirtualMachines = append(policyState.VirtualMachines, models.VirtualMachines{
				ID:           types.StringValue(volume.ID),
				InstanceUUID: types.StringValue(volume.InstanceUUID),
				Name:         types.StringValue(volume.Name),
			})
		}

		for _, volume := range policyValue.Volumes {
			policyState.Volumes = append(policyState.Volumes, models.Volumes{
				ID:          types.StringValue(volume.ID),
				Name:        types.StringValue(volume.Name),
				Description: types.StringValue(volume.Description),
			})
		}

		for _, volume := range policyValue.VolumeGroups {
			policyState.VolumeGroups = append(policyState.VolumeGroups, models.VolumeGroups{
				ID:          types.StringValue(volume.ID),
				Name:        types.StringValue(volume.Name),
				Description: types.StringValue(volume.Description),
			})
		}

		for _, fileSystem := range policyValue.FileSystems {
			policyState.FileSystems = append(policyState.FileSystems, models.FileSystems{
				ID:          types.StringValue(fileSystem.ID),
				Name:        types.StringValue(fileSystem.Name),
				Description: types.StringValue(fileSystem.Description),
			})
		}

		for _, performanceRule := range policyValue.PerformanceRules {
			policyState.PerformanceRules = append(policyState.PerformanceRules, models.PerformanceRules{
				ID:         types.StringValue(performanceRule.ID),
				Name:       types.StringValue(performanceRule.Name),
				IoPriority: types.StringValue(performanceRule.IoPriority),
			})
		}

		for _, snapshotRule := range policyValue.SnapshotRules {
			policyState.SnapshotRules = append(policyState.SnapshotRules, models.SnapshotRules{
				ID:   types.StringValue(snapshotRule.ID),
				Name: types.StringValue(snapshotRule.Name),
			})
		}

		for _, replicationRule := range policyValue.ReplicationRules {
			policyState.ReplicationRules = append(policyState.ReplicationRules, models.ReplicationRules{
				ID:   types.StringValue(replicationRule.ID),
				Name: types.StringValue(replicationRule.Name),
			})
		}

		response = append(response, policyState)
	}
	return response, nil
}
