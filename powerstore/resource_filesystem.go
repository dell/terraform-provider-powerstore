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
	"log"
	"regexp"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type fileSystemResource struct {
	client *client.Client
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &fileSystemResource{}
	_ resource.ResourceWithConfigure   = &fileSystemResource{}
	_ resource.ResourceWithImportState = &fileSystemResource{}
)

// newFileSystemResource is a helper function to simplify the provider implementation.
func newFileSystemResource() resource.Resource {
	return &fileSystemResource{}
}

func (r fileSystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystem"
}

func (r fileSystemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This resource is used to manage the file system entity of PowerStore Array. We can Create, Update and Delete the file system using this resource. We can also import an existing file system from PowerStore array.",
		MarkdownDescription: "This resource is used to manage the file system entity of PowerStore Array. We can Create, Update and Delete the file system using this resource. We can also import an existing file system from PowerStore array.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the file system.",
				MarkdownDescription: "Unique identifier of the file system.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Description:         "Name of the file system.",
				MarkdownDescription: "Name of the file system.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				Description:         "File system description.",
				MarkdownDescription: "File system description.",
				Optional:            true,
				Computed:            true,
			},
			"size": schema.Float64Attribute{
				Description:         "Size that the file system presents to the host or end user.",
				MarkdownDescription: "Size that the file system presents to the host or end user.",
				Required:            true,
			},
			"capacity_unit": schema.StringAttribute{
				Description:         "The Capacity Unit corresponding to the size.",
				MarkdownDescription: "The Capacity Unit corresponding to the size.",
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("GB"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"MB",
						"GB",
						"TB",
					}...),
				},
			},
			"nas_server_id": schema.StringAttribute{
				Description:         "Unique identifier of the NAS Server on which the file system is mounted.",
				MarkdownDescription: "Unique identifier of the NAS Server on which the file system is mounted.",
				Required:            true,
			},
			"config_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "File system security access policies.",
				MarkdownDescription: "File system security access policies.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("General"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"General",
						"VMware",
					}...),
				},
			},
			"access_policy": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "File system security access policies.",
				MarkdownDescription: "File system security access policies.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("Native"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"Native",
						"UNIX",
						"Windows",
					}...),
				},
			},

			"locking_policy": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "File system locking policies.",
				MarkdownDescription: "File system locking policies.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("Advisory"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"Advisory",
						"Mandatory",
					}...),
				},
			},
			"folder_rename_policy": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "File system folder rename policies for the file system with multiprotocol access enabled.",
				MarkdownDescription: "File system folder rename policies for the file system with multiprotocol access enabled.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("All_Forbidden"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"All_Allowed",
						"SMB_Forbidden",
						"All_Forbidden",
					}...),
				},
			},
			"is_smb_sync_writes_enabled": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Indicates whether the synchronous writes option is enabled on the file system.",
				MarkdownDescription: "Indicates whether the synchronous writes option is enabled on the file system.",
			},
			"is_smb_no_notify_enabled": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Indicates whether notifications of changes to directory file structure are enabled.",
				MarkdownDescription: "Indicates whether notifications of changes to directory file structure are enabled.",
			},
			"is_smb_op_locks_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether opportunistic file locking is enabled on the file system.",
				MarkdownDescription: "Indicates whether opportunistic file locking is enabled on the file system.",
			},
			"is_smb_notify_on_access_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether file access notifications are enabled on the file system.",
				MarkdownDescription: "Indicates whether file access notifications are enabled on the file system.",
			},
			"is_smb_notify_on_write_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Indicates whether file writes notifications are enabled on the file system.",
				MarkdownDescription: "Indicates whether file writes notifications are enabled on the file system.",
			},
			"smb_notify_on_change_dir_depth": schema.Int32Attribute{
				Computed:            true,
				Optional:            true,
				Description:         "Lowest directory level to which the enabled notifications apply, if any.",
				MarkdownDescription: "Lowest directory level to which the enabled notifications apply, if any.",
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
					int32validator.AtMost(512),
				},
			},
			"is_async_mtime_enabled": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Indicates whether asynchronous MTIME is enabled on the file system or protocol snaps that are mounted writeable.",
				MarkdownDescription: "Indicates whether asynchronous MTIME is enabled on the file system or protocol snaps that are mounted writeable.",
			},
			"protection_policy_id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Unique identifier of the protection policy applied to the file system.",
				MarkdownDescription: "Unique identifier of the protection policy applied to the file system.",
			},
			"file_events_publishing_mode": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "State of the event notification services for all file systems of the NAS server.",
				MarkdownDescription: "State of the event notification services for all file systems of the NAS server.",
				PlanModifiers: []planmodifier.String{
					DefaultAttribute("None"),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"None",
						"SMB_Only",
						"NFS_Only",
						"All",
					}...),
				},
			},
			"host_io_size": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Typical size of writes from the server or other computer using the VMware file system to the storage system.",
				MarkdownDescription: "Typical size of writes from the server or other computer using the VMware file system to the storage system.",
				Validators: []validator.String{
					stringvalidator.OneOf([]string{
						"VMware_8K",
						"VMware_16K",
						"VMware_32K",
						"Vmware_64K",
					}...),
				},
			},
			"file_system_type": schema.StringAttribute{
				Computed:            true,
				Description:         "Type of filesystem: normal or snapshot.",
				MarkdownDescription: "Type of filesystem: normal or snapshot.",
			},
			"flr_attributes": schema.SingleNestedAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Type of filesystem: normal or snapshot.",
				MarkdownDescription: "Type of filesystem: normal or snapshot.",
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description:         "The FLR type of the file system.",
						MarkdownDescription: "The FLR type of the file system.",
						Optional:            true,
						Computed:            true,
						Validators: []validator.String{
							stringvalidator.OneOf([]string{
								"None",
								"Enterprise",
								"Compliance",
							}...),
						},
						PlanModifiers: []planmodifier.String{
							DefaultAttribute("Enterprise"),
						},
					},
					"minimum_retention": schema.StringAttribute{
						Description:         "The FLR type of the file system.",
						MarkdownDescription: "The FLR type of the file system.",
						Optional:            true,
						Computed:            true,
						Validators: []validator.String{
							stringvalidator.RegexMatches(
								regexp.MustCompile(`(^\d+[DMY])|(^infinite$)`),
								"must only contain only alphanumeric characters",
							),
						},
						PlanModifiers: []planmodifier.String{
							DefaultAttribute("1D"),
						},
					},
					"default_retention": schema.StringAttribute{
						Description:         "The FLR type of the file system.",
						MarkdownDescription: "The FLR type of the file system.",
						Optional:            true,
						Computed:            true,
						Validators: []validator.String{
							stringvalidator.RegexMatches(
								regexp.MustCompile(`(^\d+[DMY])|(^infinite$)`),
								"must only contain only alphanumeric characters",
							),
						},
					},
					"maximum_retention": schema.StringAttribute{
						Description:         "The FLR type of the file system.",
						MarkdownDescription: "The FLR type of the file system.",
						Optional:            true,
						Computed:            true,
						Validators: []validator.String{
							stringvalidator.RegexMatches(
								regexp.MustCompile(`(^\d+[DMY])|(^infinite$)`),
								"must only contain only alphanumeric characters",
							),
						},
						PlanModifiers: []planmodifier.String{
							DefaultAttribute("infinite"),
						},
					},
				},
			},
			"parent_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the parent filesystem.",
				MarkdownDescription: "Unique identifier of the parent filesystem.",
			},
		},
	}
}

// Configure - defines configuration for file system resource.
func (r *fileSystemResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create - method to create file system resource
func (r fileSystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var valInBytes int64

	log.Printf("Started Creating file system")
	var plan models.FileSystem

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	valInBytes, errmsg := convertToBytesForFileSystem(plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error creating file system",
			"Error in converting the given size into bytes"+errmsg,
		)
		return
	}

	var FlrCreate models.FlrAttributes
	if !plan.FlrAttributes.IsUnknown() {
		plan.FlrAttributes.As(ctx, &FlrCreate, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
	}

	fileSystemCreate := &gopowerstore.FsCreate{
		Name:                       plan.Name.ValueString(),
		Description:                plan.Description.ValueString(),
		Size:                       valInBytes,
		NASServerID:                plan.NASServerID.ValueString(),
		ConfigType:                 plan.ConfigType.ValueString(),
		AccessPolicy:               plan.AccessPolicy.ValueString(),
		LockingPolicy:              plan.LockingPolicy.ValueString(),
		FolderRenamePolicy:         plan.FolderRenamePolicy.ValueString(),
		IsAsyncMTimeEnabled:        plan.IsAsyncMTimeEnabled.ValueBool(),
		ProtectionPolicyID:         plan.ProtectionPolicyID.ValueString(),
		FileEventsPublishingMode:   plan.FileEventsPublishingMode.ValueString(),
		HostIOSize:                 plan.HostIOSize.ValueString(),
		IsSmbSyncWritesEnabled:     GetKnownBoolPointer(plan.IsSmbSyncWritesEnabled),
		IsSmbNoNotifyEnabled:       GetKnownBoolPointer(plan.IsSmbNoNotifyEnabled),
		IsSmbOpLocksEnabled:        GetKnownBoolPointer(plan.IsSmbOpLocksEnabled),
		IsSmbNotifyOnAccessEnabled: GetKnownBoolPointer(plan.IsSmbNotifyOnAccessEnabled),
		IsSmbNotifyOnWriteEnabled:  GetKnownBoolPointer(plan.IsSmbNotifyOnWriteEnabled),
		SmbNotifyOnChangeDirDepth:  plan.SmbNotifyOnChangeDirDepth.ValueInt32(),
		FlrCreate: gopowerstore.FlrAttributes{
			Mode:             FlrCreate.Mode.ValueString(),
			MinimumRetention: FlrCreate.MinimumRetention.ValueString(),
			DefaultRetention: FlrCreate.DefaultRetention.ValueString(),
			MaximumRetention: FlrCreate.MaximumRetention.ValueString(),
		},
	}

	// Create New FileSystem
	fsCreateResponse, err := r.client.PStoreClient.CreateFS(context.Background(), fileSystemCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating file system",
			"Could not create file system, unexpected error: "+err.Error(),
		)
		return
	}

	// Get file system Details using ID retrieved above
	fsResponse, err1 := r.client.PStoreClient.GetFS(context.Background(), fsCreateResponse.ID)
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting file system after creation",
			"Could not get file system, unexpected error: "+err1.Error(),
		)
		return
	}

	result := models.FileSystem{}
	updateFsState(&result, fsResponse)

	log.Printf("Added to result: %v", result)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Read - reads file system resource
func (r fileSystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.FileSystem
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get file system details from API and then update what is in state from what the API returns
	fsID := state.ID.ValueString()
	fsResponse, err := r.client.PStoreClient.GetFS(context.Background(), fsID)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading file system",
			"Could not read file system with error "+fsID+": "+err.Error(),
		)
		return
	}

	updateFsState(&state, fsResponse)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

// Update - updates file system resource
func (r fileSystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	// Get plan values
	var plan models.FileSystem
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state models.FileSystem
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.Name.ValueString() != state.Name.ValueString() {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Name of the file system can't be updated",
		)
		return
	}

	if plan.NASServerID.ValueString() != state.NASServerID.ValueString() {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"NAS server ID can't be updated",
		)
		return
	}

	if plan.HostIOSize.ValueString() != state.HostIOSize.ValueString() {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Host IO size can't be updated",
		)
		return
	}
	if plan.ConfigType.ValueString() != state.ConfigType.ValueString() {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Config type can't be updated",
		)
		return
	}

	var FlrCreate models.FlrAttributes
	plan.FlrAttributes.As(ctx, &FlrCreate, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	var FlrCreateState models.FlrAttributes
	state.FlrAttributes.As(ctx, &FlrCreateState, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	if FlrCreate.Mode.ValueString() != FlrCreateState.Mode.ValueString() {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Mode of the flr attributes can't be updated",
		)
		return
	}

	valInBytes, errmsg := convertToBytesForFileSystem(plan)
	if errmsg != "" {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Error in converting the given size into bytes "+errmsg,
		)
		return
	}

	fsModify := &gopowerstore.FSModify{
		Description:                plan.Description.ValueString(),
		Size:                       int(valInBytes),
		AccessPolicy:               plan.AccessPolicy.ValueString(),
		LockingPolicy:              plan.LockingPolicy.ValueString(),
		FolderRenamePolicy:         plan.FolderRenamePolicy.ValueString(),
		IsAsyncMtimeEnabled:        GetKnownBoolPointer(plan.IsAsyncMTimeEnabled),
		ProtectionPolicyID:         plan.ProtectionPolicyID.ValueString(),
		FileEventsPublishingMode:   plan.FileEventsPublishingMode.ValueString(),
		IsSmbSyncWritesEnabled:     GetKnownBoolPointer(plan.IsSmbSyncWritesEnabled),
		IsSmbNoNotifyEnabled:       GetKnownBoolPointer(plan.IsSmbNoNotifyEnabled),
		IsSmbOpLocksEnabled:        GetKnownBoolPointer(plan.IsSmbOpLocksEnabled),
		IsSmbNotifyOnAccessEnabled: GetKnownBoolPointer(plan.IsSmbNotifyOnAccessEnabled),
		IsSmbNotifyOnWriteEnabled:  GetKnownBoolPointer(plan.IsSmbNotifyOnWriteEnabled),
		SmbNotifyOnChangeDirDepth:  plan.SmbNotifyOnChangeDirDepth.ValueInt32(),
		FlrCreate: gopowerstore.FlrAttributes{
			MinimumRetention: FlrCreate.MinimumRetention.ValueString(),
			DefaultRetention: FlrCreate.DefaultRetention.ValueString(),
			MaximumRetention: FlrCreate.MaximumRetention.ValueString(),
		},
	}

	_, err := r.client.PStoreClient.ModifyFS(context.Background(), fsModify, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Could not update file system "+err.Error(),
		)
		return
	}

	fsResponse, err1 := r.client.PStoreClient.GetFS(context.Background(), state.ID.ValueString())
	if err1 != nil {
		resp.Diagnostics.AddError(
			"Error getting file system after creation",
			"Could not get file system, unexpected error: "+err1.Error(),
		)
		return
	}

	updateFsState(&state, fsResponse)

	//Set State
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Done with Update")
}

// Delete - method to delete file system resource
func (r fileSystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with Delete")

	var state models.FileSystem
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get vg ID from state
	fsID := state.ID.ValueString()

	// Delete file system  by calling API
	_, err := r.client.PStoreClient.DeleteFS(context.Background(), fsID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting file system",
			"Could not delete file system "+fsID+": "+err.Error(),
		)
		return
	}

	// Remove resource from state
	resp.State.RemoveResource(ctx)
	log.Printf("Done with Delete")
}

// ImportState import state for existing file system
func (r fileSystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func convertToBytesForFileSystem(plan models.FileSystem) (int64, string) {
	var valInBytes float64
	switch plan.CapacityUnit.ValueString() {
	case "MB":
		valInBytes = plan.Size.ValueFloat64() * MiB
	case "TB":
		valInBytes = plan.Size.ValueFloat64() * TiB
	case "GB":
		valInBytes = plan.Size.ValueFloat64() * GiB
	default:
		return 0, "Invalid Capacity unit"
	}
	return int64(valInBytes), ""
}

func convertFromBytesForFileSystem(bytes int64) (float64, string) {
	var newSize float64
	var unit int
	var units = []string{"KB", "MB", "GB", "TB"}
	for newSize = float64(bytes); newSize >= 1024 && unit < len(units); unit++ {
		newSize = newSize / 1024
	}
	if unit > 0 {
		return newSize, units[unit-1]
	}
	return newSize, units[unit]
}

func updateFsState(fsState *models.FileSystem, fsResponse gopowerstore.FileSystem) {
	// Update value from file system Response to State
	fsState.ID = types.StringValue(fsResponse.ID)
	fsState.Name = types.StringValue(fsResponse.Name)
	fsState.Description = types.StringValue(fsResponse.Description)
	fsState.NASServerID = types.StringValue(fsResponse.NasServerID)
	size, unit := convertFromBytesForFileSystem(fsResponse.SizeTotal)
	fsState.Size = types.Float64Value(size)
	fsState.CapacityUnit = types.StringValue(unit)
	fsState.ConfigType = types.StringValue(fsResponse.ConfigType)
	fsState.AccessPolicy = types.StringValue(fsResponse.AccessPolicy)
	fsState.LockingPolicy = types.StringValue(fsResponse.LockingPolicy)
	fsState.FolderRenamePolicy = types.StringValue(fsResponse.FolderRenamePolicy)
	fsState.IsAsyncMTimeEnabled = types.BoolValue(fsResponse.IsAsyncMTimeEnabled)
	fsState.ProtectionPolicyID = types.StringValue(fsResponse.ProtectionPolicyID)
	fsState.FileEventsPublishingMode = types.StringValue(fsResponse.FileEventsPublishingMode)
	fsState.HostIOSize = types.StringValue(fsResponse.HostIOSize)
	fsState.IsSmbSyncWritesEnabled = types.BoolValue(fsResponse.IsSmbSyncWritesEnabled)
	fsState.IsSmbNoNotifyEnabled = types.BoolValue(fsResponse.IsSmbNoNotifyEnabled)
	fsState.IsSmbOpLocksEnabled = types.BoolValue(fsResponse.IsSmbOpLocksEnabled)
	fsState.IsSmbNotifyOnAccessEnabled = types.BoolValue(fsResponse.IsSmbNotifyOnAccessEnabled)
	fsState.IsSmbNotifyOnWriteEnabled = types.BoolValue(fsResponse.IsSmbNotifyOnWriteEnabled)
	fsState.SmbNotifyOnChangeDirDepth = types.Int32Value(fsResponse.SmbNotifyOnChangeDirDepth)
	fsState.ParentID = types.StringValue(fsResponse.ParentID)
	fsState.FilesystemType = types.StringValue(string(fsResponse.FilesystemType))
	fsState.FlrAttributes, _ = types.ObjectValue(map[string]attr.Type{
		"mode":              types.StringType,
		"minimum_retention": types.StringType,
		"default_retention": types.StringType,
		"maximum_retention": types.StringType,
	}, map[string]attr.Value{
		"mode":              types.StringValue(fsResponse.FlrCreate.Mode),
		"minimum_retention": types.StringValue(fsResponse.FlrCreate.MinimumRetention),
		"default_retention": types.StringValue(fsResponse.FlrCreate.DefaultRetention),
		"maximum_retention": types.StringValue(fsResponse.FlrCreate.MaximumRetention),
	})

}

func GetKnownBoolPointer(in types.Bool) *bool {
	if in.IsUnknown() {
		return nil
	}
	return in.ValueBoolPointer()
}
