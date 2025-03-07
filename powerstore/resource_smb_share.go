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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"
)

// newSMBShareResource returns SMBShare new resource instance
func newSMBShareResource() resource.Resource {
	return &resourceSMBShare{}
}

type resourceSMBShare struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceSMBShare) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smb_share"
}

// Schema defines resource interface Schema method
func (r *resourceSMBShare) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the smb share entity of PowerStore Array. We can Create, Update and Delete the smb share using this resource. We can also import an existing smb share from PowerStore array.",
		Description:         "This resource is used to manage the smb share entity of PowerStore Array. We can Create, Update and Delete the smb share using this resource. We can also import an existing smb share from PowerStore array.",

		Attributes: SMBShareSchema(),
	}
}

// SMBShareSchema defines resource interface Schema method
func SMBShareSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			Description:         "The unique identifier of the SMB Share.",
			MarkdownDescription: "The unique identifier of the SMB Share.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"file_system_id": schema.StringAttribute{
			MarkdownDescription: "The unique identifier of the file system on which the SMB Share is created.",
			Description:         "The unique identifier of the file system on which the SMB Share is created.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the SMB Share.",
			Description:         "The name of the SMB Share.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthBetween(1, 80),
			},
		},
		"path": schema.StringAttribute{
			MarkdownDescription: "The local path to export relative to the NAS Server.",
			Description:         "The local path to export relative to the NAS Server.",
			Required:            true,
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(1),
			},
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "User defined SMB share description.",
			Description:         "User defined SMB share description.",
			Optional:            true,
			Computed:            true,
			Validators: []validator.String{
				stringvalidator.LengthBetween(1, 255),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"is_continuous_availability_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether continuous availability for Server Message Block (SMB) 3.0 is enabled for the SMB Share.",
			Description:         "Whether continuous availability for Server Message Block (SMB) 3.0 is enabled for the SMB Share.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"is_encryption_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether encryption for Server Message Block (SMB) 3.0 is enabled at the shared folder level.",
			Description:         "Whether encryption for Server Message Block (SMB) 3.0 is enabled at the shared folder level.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"is_abe_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether Access-based Enumeration (ABE) is enabled.",
			Description:         "Whether Access-based Enumeration (ABE) is enabled.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"is_branch_cache_enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether BranchCache optimization is enabled.",
			Description:         "Whether BranchCache optimization is enabled.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Bool{
				boolplanmodifier.UseStateForUnknown(),
			},
		},
		"offline_availability": schema.StringAttribute{
			MarkdownDescription: "Defines valid states of Offline Availability, where the states are: `Manual` - Only specified files will be available offline. `Documents` - All files that users open will be available offline. `Programs` - Program will preferably run from the offline cache even when connected to the network. All files that users open will be available offline. `None` - Prevents clients from storing documents and programs in offline cache.",
			Description:         "Defines valid states of Offline Availability, where the states are: `Manual` - Only specified files will be available offline. `Documents` - All files that users open will be available offline. `Programs` - Program will preferably run from the offline cache even when connected to the network. All files that users open will be available offline. `None` - Prevents clients from storing documents and programs in offline cache.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Validators: []validator.String{
				stringvalidator.OneOf([]string{
					"Manual",
					"Documents",
					"Programs",
					"None",
				}...),
			},
		},
		"umask": schema.StringAttribute{
			MarkdownDescription: "The default UNIX umask for new files created on the Share.",
			Description:         "The default UNIX umask for new files created on the Share.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Validators: []validator.String{
				stringvalidator.LengthAtLeast(3),
				stringvalidator.LengthAtMost(3),
			},
		},
		"aces": schema.SetNestedAttribute{
			MarkdownDescription: "To specify the ACL access options.",
			Description:         "To specify the ACL access options.",
			Optional:            true,
			Computed:            true,
			PlanModifiers: []planmodifier.Set{
				setplanmodifier.UseStateForUnknown(),
			},
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"trustee_name": schema.StringAttribute{
						MarkdownDescription: "The name of the trustee.",
						Description:         "The name of the trustee.",
						Required:            true,
					},
					"trustee_type": schema.StringAttribute{
						MarkdownDescription: "The type of the trustee.",
						Description:         "The type of the trustee.",
						Required:            true,
						Validators: []validator.String{
							stringvalidator.OneOf([]string{
								"Manual",
								"SID",
								"User",
								"Group",
								"WellKnown",
							}...),
						},
					},
					"access_level": schema.StringAttribute{
						MarkdownDescription: "The access level.",
						Description:         "The access level.",
						Required:            true,
						Validators: []validator.String{
							stringvalidator.OneOf([]string{
								"Read",
								"Full",
								"Change",
							}...),
						},
					},

					"access_type": schema.StringAttribute{
						MarkdownDescription: "The access type.",
						Description:         "The access type.",
						Required:            true,
						Validators: []validator.String{
							stringvalidator.OneOf([]string{
								"Allow",
								"Deny",
							}...),
						},
					},
				},
			},
		},
	}
}

// Configure - defines configuration for smb share resource
func (r *resourceSMBShare) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create - method to create smb share resource
func (r *resourceSMBShare) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.SMBShare

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Create new smb share
	smbCreate := r.planToSMBShareCreateParam(plan)
	var aclParams *gopowerstore.ModifySMBShareACL

	smbCreateResponse, err := r.client.PStoreClient.CreateSMBShare(ctx, smbCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating smb share",
			"Could not create smb share, unexpected error: "+err.Error(),
		)
		return
	}
	ID := smbCreateResponse.ID
	// Get SMBShare Details using ID retrieved above
	SMBShareDetails, err := r.client.PStoreClient.GetSMBShare(context.Background(), ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading smb share details",
			"Could not get smb share, unexpected error: "+err.Error(),
		)
		return
	}

	// update state with current details
	state, diags := r.SMBShareStateACL(SMBShareDetails, gopowerstore.SMBShareACL{})
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// update state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if ACL needs to be created
	if !plan.SMBShareACL.IsUnknown() {
		aclParams, diags = r.planToSMBShareACLUpdate(ctx, plan)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		// Update SMBShare ACL
		diags = r.updateSMBShareACL(ctx, resp, ID, aclParams)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	smbShareState, diags := r.getSMBShareDetails(ctx, ID)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags = resp.State.Set(ctx, smbShareState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads smb share resource information
func (r *resourceSMBShare) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state models.SMBShare
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	SMBShareID := state.ID.ValueString()
	smbShareState, diags := r.getSMBShareDetails(ctx, SMBShareID)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Set state
	diags = resp.State.Set(ctx, smbShareState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Read")
}

// Update - updates smb share resource
func (r *resourceSMBShare) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.SMBShare
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.SMBShare
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	SMBShareID := state.ID.ValueString()

	// Check not modifiable attributes
	if !plan.Name.Equal(state.Name) || !plan.FileSystemID.Equal(state.FileSystemID) || !plan.Path.Equal(state.Path) {
		resp.Diagnostics.AddError(
			"Error updating smb share resource",
			"smb share attributes [name, filesystem_id, path] are not modifiable",
		)
		return
	}

	smbModify := r.planToSMBShareModifyParam(plan)

	// Modify SMBShare
	_, err := r.client.PStoreClient.ModifySMBShare(context.Background(), SMBShareID, smbModify)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating smb share",
			"Could not update smb share: "+err.Error(),
		)
		return
	}

	SMBShareDetails, err := r.client.PStoreClient.GetSMBShare(context.Background(), SMBShareID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading smb share details",
			"Could not get smb share, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = r.SMBShareStateACL(SMBShareDetails, gopowerstore.SMBShareACL{})
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// update state
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	aclParams, diags := r.planToSMBShareACLUpdate(ctx, plan)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Update SMBShare ACL
	_, err = r.client.PStoreClient.SetSMBShareACL(context.Background(), SMBShareID, aclParams)
	if err != nil {
		diags.AddError(
			"Error updating smb share ACL deatils",
			"Could not update smb share ACL, unexpected error: "+err.Error(),
		)
	}

	smbShareState, diags := r.getSMBShareDetails(ctx, SMBShareID)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	diags = resp.State.Set(ctx, smbShareState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Successfully done with Update")
}

// Delete - method to delete smb share resource
func (r *resourceSMBShare) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.SMBShare
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Get SMBShare ID from state
	SMBShareID := state.ID.ValueString()

	// Delete SMBShare by calling API
	_, err := r.client.PStoreClient.DeleteSMBShare(context.Background(), SMBShareID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting SMBShare",
			"Could not delete SMBShareID "+SMBShareID+": "+err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
	log.Printf("Done with Delete")
}

// ImportState - imports state for existing SMBShare
func (r *resourceSMBShare) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// planToSMBShareCreateParam - method to read SMBShare from plan for create operation
func (r *resourceSMBShare) planToSMBShareCreateParam(plan models.SMBShare) *gopowerstore.SMBShareCreate {
	smbShare := &gopowerstore.SMBShareCreate{
		FileSystemID:                    plan.FileSystemID.ValueString(),
		Name:                            plan.Name.ValueString(),
		Path:                            plan.Path.ValueString(),
		Description:                     plan.Description.ValueString(),
		IsContinuousAvailabilityEnabled: plan.IsContinuousAvailabilityEnabled.ValueBool(),
		IsEncryptionEnabled:             plan.IsEncryptionEnabled.ValueBool(),
		IsABEEnabled:                    plan.IsABEEnabled.ValueBool(),
		IsBranchCacheEnabled:            plan.IsBranchCacheEnabled.ValueBool(),
		Umask:                           plan.Umask.ValueString(),
		OfflineAvailability:             plan.OfflineAvailability.ValueString(),
	}
	return smbShare
}

// planToSMBShareACLUpdate - method to read SMBShare ACL from plan
func (r *resourceSMBShare) planToSMBShareModifyParam(plan models.SMBShare) *gopowerstore.SMBShareModify {
	description := plan.Description.ValueString()
	isContinuousAvailabilityEnabled := plan.IsContinuousAvailabilityEnabled.ValueBool()
	isEncryptionEnabled := plan.IsEncryptionEnabled.ValueBool()
	isABEEnabled := plan.IsABEEnabled.ValueBool()
	isBranchCacheEnabled := plan.IsBranchCacheEnabled.ValueBool()

	smbShareModify := &gopowerstore.SMBShareModify{
		Description:                     &description,
		IsContinuousAvailabilityEnabled: &isContinuousAvailabilityEnabled,
		IsEncryptionEnabled:             &isEncryptionEnabled,
		IsABEEnabled:                    &isABEEnabled,
		IsBranchCacheEnabled:            &isBranchCacheEnabled,
		Umask:                           plan.Umask.ValueString(),
		OfflineAvailability:             plan.OfflineAvailability.ValueString(),
	}
	return smbShareModify
}

// planToSMBShareACLUpdate - method to read SMBShare ACL from plan
func (r *resourceSMBShare) planToSMBShareACLUpdate(ctx context.Context, plan models.SMBShare) (*gopowerstore.ModifySMBShareACL, diag.Diagnostics) {
	var diags diag.Diagnostics
	var aclParams *gopowerstore.ModifySMBShareACL
	var aces []gopowerstore.SMBShareAce
	var acesModel []models.SMBShareAce

	diags = plan.SMBShareACL.ElementsAs(ctx, &acesModel, false)
	if diags.HasError() {
		diags.AddError(
			"Error creating smb share",
			"unable to parse SMB share ACL elements from plan",
		)
		return nil, diags
	}

	for _, v := range acesModel {
		acl := gopowerstore.SMBShareAce{
			TrusteeType: v.TrusteeType.ValueString(),
			TrusteeName: v.TrusteeName.ValueString(),
			AccessLevel: v.AccessLevel.ValueString(),
			AccessType:  v.AccessType.ValueString(),
		}
		aces = append(aces, acl)
	}
	aclParams = &gopowerstore.ModifySMBShareACL{
		Aces: aces,
	}
	return aclParams, diags
}

// SMBShareStateACL - method to update terraform state
func (r *resourceSMBShare) SMBShareStateACL(shareDetails gopowerstore.SMBShare, aclDetails gopowerstore.SMBShareACL) (models.SMBShare, diag.Diagnostics) {
	// Update the SMB Share ACL details in the state
	smbShare := models.SMBShare{
		ID:                              types.StringValue(shareDetails.ID),
		Description:                     types.StringValue(shareDetails.Description),
		FileSystemID:                    types.StringValue(shareDetails.FileSystemID),
		Name:                            types.StringValue(shareDetails.Name),
		Path:                            types.StringValue(shareDetails.Path),
		IsABEEnabled:                    types.BoolValue(shareDetails.IsABEEnabled),
		IsBranchCacheEnabled:            types.BoolValue(shareDetails.IsBranchCacheEnabled),
		IsContinuousAvailabilityEnabled: types.BoolValue(shareDetails.IsContinuousAvailabilityEnabled),
		IsEncryptionEnabled:             types.BoolValue(shareDetails.IsEncryptionEnabled),
		Umask:                           types.StringValue(shareDetails.Umask),
		OfflineAvailability:             types.StringValue(shareDetails.OfflineAvailability),
	}
	smbShareACLAttrTypes := map[string]attr.Type{
		"trustee_type": types.StringType,
		"trustee_name": types.StringType,
		"access_level": types.StringType,
		"access_type":  types.StringType,
	}

	SmbShareElemTypes := types.ObjectType{
		AttrTypes: smbShareACLAttrTypes,
	}

	objectIPs := []attr.Value{}
	var diags diag.Diagnostics
	for _, ip := range aclDetails.Aces {
		obj := map[string]attr.Value{
			"trustee_type": types.StringValue(ip.TrusteeType),
			"trustee_name": types.StringValue(ip.TrusteeName),
			"access_level": types.StringValue(ip.AccessLevel),
			"access_type":  types.StringValue(ip.AccessType),
		}
		objVal, dgs := types.ObjectValue(smbShareACLAttrTypes, obj)
		diags = append(diags, dgs...)
		objectIPs = append(objectIPs, objVal)
	}
	setVal, dgs := types.SetValue(SmbShareElemTypes, objectIPs)
	diags = append(diags, dgs...)
	smbShare.SMBShareACL = setVal
	return smbShare, diags
}

// getSMBShareDetails - method to get smb share details
func (r *resourceSMBShare) getSMBShareDetails(ctx context.Context, SMBShareID string) (*models.SMBShare, diag.Diagnostics) {
	var diags diag.Diagnostics
	// Get SMBShare Details using ID retrieved above
	SMBShareResponse, err := r.client.PStoreClient.GetSMBShare(context.Background(), SMBShareID)
	if err != nil {
		diags.AddError(
			"Error reading smb share details",
			"Could not get smb share, unexpected error: "+err.Error(),
		)
		return nil, diags
	}

	// Get SMBShare Details ACL using ID retrieved above
	SMBShareACLResponse, err := r.client.PStoreClient.GetSMBShareACL(context.Background(), SMBShareID)
	if err != nil {
		diags.AddError(
			"Error reading smb share ACL details",
			"Could not get smb share, unexpected error: "+err.Error(),
		)
		return nil, diags
	}

	// Update details to state
	state, diags := r.SMBShareStateACL(SMBShareResponse, SMBShareACLResponse)
	if diags.HasError() {
		return nil, diags
	}
	return &state, diags

}

// updateSMBShareACL - method to update smb share ACL
func (r *resourceSMBShare) updateSMBShareACL(ctx context.Context, resp *resource.CreateResponse, SMBShareID string, aclParams *gopowerstore.ModifySMBShareACL) diag.Diagnostics {
	var diags diag.Diagnostics
	var err error
	_, err = r.client.PStoreClient.SetSMBShareACL(context.Background(), SMBShareID, aclParams)
	if err != nil {
		diags.AddError(
			"Error creating smb share ACL",
			"Could not create smb share ACL, unexpected error: "+err.Error(),
		)
		// Delete SMBShare by calling API to prevent from getting tainted
		_, err := r.client.PStoreClient.DeleteSMBShare(context.Background(), SMBShareID)
		if err != nil {
			diags.AddError(
				"Error deleting SMBShare",
				"Could not delete SMBShareID "+SMBShareID+": "+err.Error(),
			)
			return diags
		}
		// delete has passed so we need to remove state and return SMBShare ACL update failure
		resp.State.RemoveResource(ctx)
		return diags
	}
	return nil
}
