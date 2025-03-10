/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"log"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/customtypes/nfshostset"
	"terraform-provider-powerstore/powerstore/helper"
)

// newNFSExportResource returns nfsExport new resource instance
func newNFSExportResource() resource.Resource {
	return &resourceNFSExport{}
}

type resourceNFSExport struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceNFSExport) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfs_export"
}

// Schema defines resource interface Schema method
func (r *resourceNFSExport) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the nfs export entity of PowerStore Array. We can Create, Update and Delete the nfs export using this resource. We can also import an existing nfs export from PowerStore array.",
		Description:         "This resource is used to manage the nfs export entity of PowerStore Array. We can Create, Update and Delete the nfs export using this resource. We can also import an existing nfs export from PowerStore array.",

		Attributes: NFSExportResourceSchema(),
	}
}

// NFSExportResourceSchema defines resource interface Schema method
func NFSExportResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Computed:            true,
			Description:         "The unique identifier of the NFS Export.",
			MarkdownDescription: "The unique identifier of the NFS Export.",
		},
		"file_system_id": schema.StringAttribute{
			MarkdownDescription: "The unique identifier of the file	system on which the NFS Export will be created.",
			Description:         "The unique identifier of the file system on which the NFS Export will be created.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the NFS Export.",
			Description:         "The name of the NFS Export.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"path": schema.StringAttribute{
			MarkdownDescription: "The local path to export relative to the nfs export root directory. With NFS, each export of a file_system or file_nfs must have a unique local path. Before you can create additional Exports within an NFS shared folder, you must create directories within it from a Linux/Unix host that is connected to the nfs export. After a directory has been created from a mounted host, you can create a corresponding Export and Set access permissions accordingly.",
			Description:         "The local path to export relative to the nfs export root directory. With NFS, each export of a file_system or file_nfs must have a unique local path. Before you can create additional Exports within an NFS shared folder, you must create directories within it from a Linux/Unix host that is connected to the nfs export. After a directory has been created from a mounted host, you can create a corresponding Export and Set access permissions accordingly.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"description": schema.StringAttribute{
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			MarkdownDescription: "A user-defined description of the NFS Export.",
			Description:         "A user-defined description of the NFS Export.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"min_security": schema.StringAttribute{
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			MarkdownDescription: "The NFS enforced security type for users accessing the NFS Export. Valid values are: 'Sys', 'Kerberos', 'Kerberos_With_Integrity', 'Kerberos_With_Encryption'.",
			Description:         "The NFS enforced security type for users accessing the NFS Export. Valid values are: 'Sys', 'Kerberos', 'Kerberos_With_Integrity', 'Kerberos_With_Encryption'.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("Sys", "Kerberos", "Kerberos_With_Integrity", "Kerberos_With_Encryption"),
			},
		},
		"anonymous_gid": schema.Int32Attribute{
			PlanModifiers: []planmodifier.Int32{
				int32planmodifier.UseStateForUnknown(),
			},
			MarkdownDescription: "The GID (Group ID) of the anonymous user. This is the group ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Description:         "The GID (Group ID) of the anonymous user. This is the group ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Optional:            true,
			Computed:            true,
		},
		"anonymous_uid": schema.Int32Attribute{
			MarkdownDescription: "The UID (User ID) of the anonymous user. This is the user ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Description:         "The UID (User ID) of the anonymous user. This is the user ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Int32{
				int32planmodifier.UseStateForUnknown(),
			},
		},
		"is_no_suid": schema.BoolAttribute{
			MarkdownDescription: "If Set, do not allow access to Set SUID. Otherwise, allow access.",
			Description:         "If Set, do not allow access to Set SUID. Otherwise, allow access.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"nfs_owner_username": schema.StringAttribute{
			MarkdownDescription: "The default owner of the NFS Export associated with the datastore. Required if secure NFS enabled. For NFSv3 or NFSv4 without Kerberos, the default owner is root. Was added in version 3.0.0.0.",
			Description:         "The default owner of the NFS Export associated with the datastore. Required if secure NFS enabled. For NFSv3 or NFSv4 without Kerberos, the default owner is root. Was added in version 3.0.0.0.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Validators: []validator.String{
				stringvalidator.LengthAtMost(32),
				stringvalidator.LengthAtLeast(1),
			},
		},
		"default_access": schema.StringAttribute{
			MarkdownDescription: "The default access level for all hosts that can access the NFS Export. The default access level is the access level that is assigned to a host that is not explicitly Seted in the 'no_access_hosts', 'read_only_hosts', 'read_only_root_hosts', 'read_write_hosts', or 'read_write_root_hosts' Sets. Valid values are: 'No_Access', 'Read_Only', 'Read_Write', 'Root', 'Read_Only_Root'.",
			Description:         "The default access level for all hosts that can access the NFS Export. The default access level is the access level that is assigned to a host that is not explicitly Seted in the 'no_access_hosts', 'read_only_hosts', 'read_only_root_hosts', 'read_write_hosts', or 'read_write_root_hosts' Sets. Valid values are: 'No_Access', 'Read_Only', 'Read_Write', 'Root', 'Read_Only_Root'.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Validators: []validator.String{
				stringvalidator.OneOf(string(gopowerstore.NoAccess), string(gopowerstore.ReadOnly), string(gopowerstore.ReadWrite), string(gopowerstore.Root), string(gopowerstore.ReadOnlyRoot)),
			},
		},
		"no_access_hosts": schema.SetAttribute{
			MarkdownDescription: "Hosts with no access to the NFS Export or its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Description:         "Hosts with no access to the NFS Export or its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Optional:            true,
			Computed:            true,
			CustomType:          nfshostset.NewHostSetType(),
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
				nfshostset.HostSetPlanModifier{},
			},
		},
		"read_only_hosts": schema.SetAttribute{
			MarkdownDescription: "Hosts with read-only access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Description:         "Hosts with read-only access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Optional:            true,
			Computed:            true,
			CustomType:          nfshostset.NewHostSetType(),
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
				nfshostset.HostSetPlanModifier{},
			},
		},
		"read_only_root_hosts": schema.SetAttribute{
			MarkdownDescription: "Hosts with read-only and read-only for root user access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Description:         "Hosts with read-only and read-only for root user access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Optional:            true,
			Computed:            true,
			CustomType:          nfshostset.NewHostSetType(),
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
				nfshostset.HostSetPlanModifier{},
			},
		},
		"read_write_hosts": schema.SetAttribute{
			MarkdownDescription: "Hosts with read and write access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Description:         "Hosts with read and write access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Optional:            true,
			Computed:            true,
			CustomType:          nfshostset.NewHostSetType(),
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
				nfshostset.HostSetPlanModifier{},
			},
		},
		"read_write_root_hosts": schema.SetAttribute{
			MarkdownDescription: "Hosts with read and write and read and write for root user access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Description:         "Hosts with read and write and read and write for root user access to the NFS Export and its nfsExports. Hosts can be entered by Hostname, IP addresses (IPv4, IPv6, IPv4/PrefixLength, IPv6/PrefixLenght, or IPv4/subnetmask), or Netgroups prefixed with @. The maximum length of an Host name is 255 bytes, and the sum of lengths of all the items in the Set is limited to 4096 bytes.",
			Optional:            true,
			Computed:            true,
			CustomType:          nfshostset.NewHostSetType(),
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
				nfshostset.HostSetPlanModifier{},
			},
		},
	}
}

// Configure - defines configuration for nfs export resource
func (r *resourceNFSExport) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create - method to create nfs export resource
func (r *resourceNFSExport) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.NFSExport

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nfsCreate := r.planToNFSExportCreateParam(plan)

	nfsCreateResponse, err := r.client.PStoreClient.CreateNFSExport(context.Background(), nfsCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating nfs export",
			"Could not create nfs export, unexpected error: "+err.Error(),
		)
		return
	}
	// Get nfsExport Details using ID retrieved above
	nfsExportResponse, err1 := r.client.PStoreClient.GetNFSExport(context.Background(), nfsCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting nfs export after creation",
			"Could not get nfs export, unexpected error: "+err1.Error(),
		)
		return
	}

	// Update details to state
	state := r.nfsExportState(ctx, nfsExportResponse)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)

	log.Printf("Done with Create")
}

// Read - reads nfs export resource information
func (r *resourceNFSExport) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.NFSExport
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nfsExportID := state.ID.ValueString()

	// Get nfsExport details from API and then update what is in state from what the API returns
	nfsExportResponse, err := r.client.PStoreClient.GetNFSExport(context.Background(), nfsExportID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading nfs export",
			"Could not read nfs export ID with error "+nfsExportID+": "+err.Error(),
		)
		return
	}
	state = r.nfsExportState(ctx, nfsExportResponse)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	log.Printf("Done with Read")
}

// Update - updates nfs export resource
func (r *resourceNFSExport) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.NFSExport
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.NFSExport
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check not modifiable attributes
	if !plan.Name.Equal(state.Name) || !plan.FileSystemID.Equal(state.FileSystemID) || !plan.Path.Equal(state.Path) || !plan.NfsOwnerUsername.Equal(state.NfsOwnerUsername) {
		resp.Diagnostics.AddError(
			"Error updating nfs export resource",
			"nfs export attributes [name, filesystem_id, path, nfs_owner_username] are not modifiable",
		)
		return

	}

	nfsModify := r.planToNFSExportModifyParam(plan)
	_, err := r.client.PStoreClient.ModifyNFSExport(context.Background(), nfsModify, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating nfs export",
			"Could not update nfs export: "+err.Error(),
		)
		return
	}

	// Get nfsExport Details using ID retrieved above
	nfsExportResponse, err1 := r.client.PStoreClient.GetNFSExport(context.Background(), state.ID.ValueString())
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting nfs export after modify",
			"Could not get nfs export, unexpected error: "+err1.Error(),
		)
		return
	}

	// Update details to state
	state = r.nfsExportState(ctx, nfsExportResponse)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)

	log.Printf("Successfully done with Update")
}

// Delete - method to delete nfs export resource
func (r *resourceNFSExport) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.NFSExport
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get nfsExport ID from state
	nfsExportID := state.ID.ValueString()

	// Delete nfsExport by calling API
	_, err := r.client.PStoreClient.DeleteNFSExport(context.Background(), nfsExportID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting nfsExport",
			"Could not delete nfsExportID "+nfsExportID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// ImportState - imports state for existing nfsExport
func (r *resourceNFSExport) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *resourceNFSExport) planToNFSExportCreateParam(plan models.NFSExport) *gopowerstore.NFSExportCreate {
	nfsCreate := &gopowerstore.NFSExportCreate{
		// Required
		FileSystemID: plan.FileSystemID.ValueString(),
		Name:         plan.Name.ValueString(),
		Path:         plan.Path.ValueString(),

		// Optional
		AnonymousGID:       plan.AnonymousGID.ValueInt32(),
		AnonymousUID:       plan.AnonymousUID.ValueInt32(),
		Description:        plan.Description.ValueString(),
		IsNoSUID:           plan.IsNoSUID.ValueBool(),
		NFSOwnerUsername:   plan.NfsOwnerUsername.ValueString(),
		MinSecurity:        plan.MinSecurity.ValueString(),
		DefaultAccess:      gopowerstore.NFSExportDefaultAccessEnum(plan.DefaultAccess.ValueString()),
		NoAccessHosts:      helper.TFListToSlice[string](plan.NoAccessHosts),
		ReadOnlyHosts:      helper.TFListToSlice[string](plan.ReadOnlyHosts),
		ReadOnlyRootHosts:  helper.TFListToSlice[string](plan.ReadOnlyRootHosts),
		ReadWriteHosts:     helper.TFListToSlice[string](plan.ReadWriteHosts),
		ReadWriteRootHosts: helper.TFListToSlice[string](plan.ReadWriteRootHosts),
	}
	return nfsCreate
}

func (r *resourceNFSExport) planToNFSExportModifyParam(plan models.NFSExport) *gopowerstore.NFSExportModify {
	nfsModify := &gopowerstore.NFSExportModify{
		AnonymousGID: plan.AnonymousGID.ValueInt32(),
		AnonymousUID: plan.AnonymousUID.ValueInt32(),
		Description:  plan.Description.ValueString(),
		IsNoSUID:     plan.IsNoSUID.ValueBool(),
		// NFSOwnerUsername:   plan.NfsOwnerUsername.ValueString(),
		MinSecurity:        plan.MinSecurity.ValueString(),
		DefaultAccess:      plan.DefaultAccess.ValueString(),
		NoAccessHosts:      helper.TFListToSlice[string](plan.NoAccessHosts),
		ReadOnlyHosts:      helper.TFListToSlice[string](plan.ReadOnlyHosts),
		ReadOnlyRootHosts:  helper.TFListToSlice[string](plan.ReadOnlyRootHosts),
		ReadWriteHosts:     helper.TFListToSlice[string](plan.ReadWriteHosts),
		ReadWriteRootHosts: helper.TFListToSlice[string](plan.ReadWriteRootHosts),
	}
	return nfsModify
}

// nfsExportState - method to convert gopowerstore.NFSExport to terraform state
func (r *resourceNFSExport) nfsExportState(ctx context.Context, input gopowerstore.NFSExport) models.NFSExport {
	return models.NFSExport{
		ID:                 types.StringValue(input.ID),
		AnonymousGID:       types.Int32Value(input.AnonymousGID),
		AnonymousUID:       types.Int32Value(input.AnonymousUID),
		DefaultAccess:      types.StringValue(string(input.DefaultAccess)),
		Description:        types.StringValue(input.Description),
		FileSystemID:       types.StringValue(input.FileSystemID),
		IsNoSUID:           types.BoolValue(input.IsNoSUID),
		MinSecurity:        types.StringValue(input.MinSecurity),
		Name:               types.StringValue(input.Name),
		NfsOwnerUsername:   types.StringValue(input.NFSOwnerUsername),
		Path:               types.StringValue(input.Path),
		NoAccessHosts:      nfshostset.NewHostSetValueFullyKnown(ctx, input.NoAccessHosts),
		ReadOnlyHosts:      nfshostset.NewHostSetValueFullyKnown(ctx, input.ROHosts),
		ReadOnlyRootHosts:  nfshostset.NewHostSetValueFullyKnown(ctx, input.RORootHosts),
		ReadWriteHosts:     nfshostset.NewHostSetValueFullyKnown(ctx, input.RWHosts),
		ReadWriteRootHosts: nfshostset.NewHostSetValueFullyKnown(ctx, input.RWRootHosts),
	}
}
