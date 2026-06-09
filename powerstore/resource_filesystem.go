/*
Copyright (c) 2024-2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"net/url"
	"regexp"

	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"
	"terraform-provider-powerstore/powerstore/helper"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Description:         "File system description.",
				MarkdownDescription: "File system description.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"size": schema.Float64Attribute{
				Description:         "Size that the file system presents to the host or end user.",
				MarkdownDescription: "Size that the file system presents to the host or end user.",
				Required:            true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
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
						"GB",
						"TB",
					}...),
				},
			},
			"nas_server_id": schema.StringAttribute{
				Description:         "Unique identifier of the NAS Server on which the file system is mounted.",
				MarkdownDescription: "Unique identifier of the NAS Server on which the file system is mounted.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"performance_policy_id": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Description:         "Unique identifier of the File_Performance type QoS policy applied to the file system.",
				MarkdownDescription: "Unique identifier of the File_Performance type QoS policy applied to the file system.",
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
						"VMware_64K",
					}...),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
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
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
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
							stringplanmodifier.UseStateForUnknown(),
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
					},
					"auto_lock": schema.BoolAttribute{
						Description:         "Indicates whether to automatically lock files in an FLR-enabled file system.",
						MarkdownDescription: "Indicates whether to automatically lock files in an FLR-enabled file system.",
						Optional:            true,
						Computed:            true,
					},
					"auto_delete": schema.BoolAttribute{
						Description:         "Indicates whether locked files will be automatically delete from an FLR-enabled file system once their retention periods have expired.",
						MarkdownDescription: "Indicates whether locked files will be automatically delete from an FLR-enabled file system once their retention periods have expired.",
						Optional:            true,
						Computed:            true,
					},
					"policy_interval": schema.Int32Attribute{
						Description:         "Indicates how long to wait (in seconds) after files are modified before the files are automatically locked.",
						MarkdownDescription: "Indicates how long to wait (in seconds) after files are modified before the files are automatically locked.",
						Optional:            true,
						Computed:            true,
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
	if plan.CapacityUnit.ValueString() == "GB" && plan.Size.ValueFloat64() > 1023 {
		resp.Diagnostics.AddError(
			"Error creating file system",
			"Use capacity unit TB instead of GB if the size is greater than 1023 GB",
		)
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

	fileSystemCreate := clientgen.FileSystemCreate{
		Name:                       plan.Name.ValueString(),
		Description:                helper.ValueToPointer[string](plan.Description),
		SizeTotal:                  valInBytes,
		NasServerId:                plan.NASServerID.ValueString(),
		ConfigType:                 helper.ValueToEnumPointer[string, clientgen.FileSystemConfigTypeEnum](plan.ConfigType),
		AccessPolicy:               helper.ValueToEnumPointer[string, clientgen.FileSystemAccessPolicyEnum](plan.AccessPolicy),
		LockingPolicy:              helper.ValueToEnumPointer[string, clientgen.FileSystemLockingPolicyEnum](plan.LockingPolicy),
		FolderRenamePolicy:         helper.ValueToEnumPointer[string, clientgen.FileSystemFolderRenamePolicyEnum](plan.FolderRenamePolicy),
		IsAsyncMTimeEnabled:        helper.ValueToPointer[bool](plan.IsAsyncMTimeEnabled),
		ProtectionPolicyId:         helper.ValueToPointer[string](plan.ProtectionPolicyID),
		PerformancePolicyId:        helper.ValueToPointer[string](plan.PerformancePolicyID),
		FileEventsPublishingMode:   helper.ValueToEnumPointer[string, clientgen.FileEventsPublishingModeEnum](plan.FileEventsPublishingMode),
		HostIoSize:                 helper.ValueToEnumPointer[string, clientgen.FileSystemHostIoSizeEnum](plan.HostIOSize),
		IsSmbSyncWritesEnabled:     helper.ValueToPointer[bool](plan.IsSmbSyncWritesEnabled),
		IsSmbNoNotifyEnabled:       helper.ValueToPointer[bool](plan.IsSmbNoNotifyEnabled),
		IsSmbOpLocksEnabled:        helper.ValueToPointer[bool](plan.IsSmbOpLocksEnabled),
		IsSmbNotifyOnAccessEnabled: helper.ValueToPointer[bool](plan.IsSmbNotifyOnAccessEnabled),
		IsSmbNotifyOnWriteEnabled:  helper.ValueToPointer[bool](plan.IsSmbNotifyOnWriteEnabled),
		SmbNotifyOnChangeDirDepth:  helper.ValueToPointer[int32](plan.SmbNotifyOnChangeDirDepth),
	}

	var FlrCreate models.FlrAttributes
	if !plan.FlrAttributes.IsUnknown() {
		plan.FlrAttributes.As(ctx, &FlrCreate, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if !(FlrCreate.AutoLock.IsUnknown()) || !(FlrCreate.AutoDelete.IsUnknown()) || !(FlrCreate.PolicyInterval.IsUnknown()) {
			resp.Diagnostics.AddError(
				"Error Creating file system",
				"auto_lock or auto_delete or policy_interval mustn't be provided during creation",
			)
			return
		}
		fileSystemCreate.FlrAttributes = &clientgen.FlrCreate{
			Mode:             helper.ValueToEnumPointer[string, clientgen.FileSystemFLRModeEnum](FlrCreate.Mode),
			MinimumRetention: helper.ValueToPointer[string](FlrCreate.MinimumRetention),
			DefaultRetention: helper.ValueToPointer[string](FlrCreate.DefaultRetention),
			MaximumRetention: helper.ValueToPointer[string](FlrCreate.MaximumRetention),
		}
	}

	// Create New FileSystem
	fsCreateResponse, _, err := r.client.GenClient.FileSystemApi.PostAllFileSystems(ctx).Body(fileSystemCreate).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating file system",
			"Could not create file system, unexpected error: "+err.Error(),
		)
		return
	}

	// Get file system Details using ID retrieved above
	fsResponse, err1 := r.readFileSystemAPI(ctx, *fsCreateResponse.Id)
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
	fsResponse, err := r.readFileSystemAPI(ctx, fsID)

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

	if plan.CapacityUnit.ValueString() == "GB" && plan.Size.ValueFloat64() > 1023 {
		resp.Diagnostics.AddError(
			"Error creating file system",
			"Use capacity unit TB instead of GB if the size is greater than 1023 GB",
		)
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
	if state.ConfigType.ValueString() == "VMware" && (FlrCreate.MinimumRetention.ValueString() != FlrCreateState.MinimumRetention.ValueString() || FlrCreate.DefaultRetention.ValueString() != FlrCreateState.DefaultRetention.ValueString() || FlrCreate.MaximumRetention.ValueString() != FlrCreateState.MaximumRetention.ValueString() || FlrCreate.AutoLock.ValueBool() != FlrCreateState.AutoLock.ValueBool() || FlrCreate.AutoDelete.ValueBool() != FlrCreateState.AutoDelete.ValueBool() || FlrCreate.PolicyInterval.ValueInt32() != FlrCreateState.PolicyInterval.ValueInt32()) {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"flr attributes can't be updated when config type is VMware",
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
	fsModify := clientgen.FileSystemModify{
		Description:                helper.ValueToPointer[string](plan.Description),
		SizeTotal:                  &valInBytes,
		AccessPolicy:               helper.ValueToEnumPointer[string, clientgen.FileSystemAccessPolicyEnum](plan.AccessPolicy),
		LockingPolicy:              helper.ValueToEnumPointer[string, clientgen.FileSystemLockingPolicyEnum](plan.LockingPolicy),
		FolderRenamePolicy:         helper.ValueToEnumPointer[string, clientgen.FileSystemFolderRenamePolicyEnum](plan.FolderRenamePolicy),
		IsAsyncMTimeEnabled:        helper.ValueToPointer[bool](plan.IsAsyncMTimeEnabled),
		ProtectionPolicyId:         helper.ValueToPointer[string](plan.ProtectionPolicyID),
		PerformancePolicyId:        helper.ValueToPointer[string](plan.PerformancePolicyID),
		FileEventsPublishingMode:   helper.ValueToEnumPointer[string, clientgen.FileEventsPublishingModeEnum](plan.FileEventsPublishingMode),
		IsSmbSyncWritesEnabled:     helper.ValueToPointer[bool](plan.IsSmbSyncWritesEnabled),
		IsSmbNoNotifyEnabled:       helper.ValueToPointer[bool](plan.IsSmbNoNotifyEnabled),
		IsSmbOpLocksEnabled:        helper.ValueToPointer[bool](plan.IsSmbOpLocksEnabled),
		IsSmbNotifyOnAccessEnabled: helper.GetKnownBoolPointer(plan.IsSmbNotifyOnAccessEnabled),
		IsSmbNotifyOnWriteEnabled:  helper.GetKnownBoolPointer(plan.IsSmbNotifyOnWriteEnabled),
		SmbNotifyOnChangeDirDepth:  helper.ValueToPointer[int32](plan.SmbNotifyOnChangeDirDepth),
	}

	if state.ConfigType.ValueString() == "General" && (FlrCreate.MinimumRetention.ValueString() != FlrCreateState.MinimumRetention.ValueString() || FlrCreate.DefaultRetention.ValueString() != FlrCreateState.DefaultRetention.ValueString() || FlrCreate.MaximumRetention.ValueString() != FlrCreateState.MaximumRetention.ValueString() || FlrCreate.AutoLock.ValueBool() != FlrCreateState.AutoLock.ValueBool() || FlrCreate.AutoDelete.ValueBool() != FlrCreateState.AutoDelete.ValueBool() || FlrCreate.PolicyInterval.ValueInt32() != FlrCreateState.PolicyInterval.ValueInt32()) {
		fsModify.FlrAttributes = &clientgen.FlrModify{
			MinimumRetention: helper.ValueToPointer[string](FlrCreate.MinimumRetention),
			DefaultRetention: helper.ValueToPointer[string](FlrCreate.DefaultRetention),
			MaximumRetention: helper.ValueToPointer[string](FlrCreate.MaximumRetention),
			AutoLock:         helper.GetKnownBoolPointer(FlrCreate.AutoLock),
			AutoDelete:       helper.GetKnownBoolPointer(FlrCreate.AutoDelete),
			PolicyInterval:   helper.ValueToPointer[int32](FlrCreate.PolicyInterval),
		}
	}

	_, err := r.client.GenClient.FileSystemApi.PatchFileSystemById(ctx, state.ID.ValueString()).Body(fsModify).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating file system",
			"Could not update file system "+err.Error(),
		)
		return
	}

	fsResponse, err1 := r.readFileSystemAPI(ctx, state.ID.ValueString())
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

	// Get ID from state
	fsID := state.ID.ValueString()

	// Delete file system  by calling API
	_, err := r.client.GenClient.FileSystemApi.DeleteFileSystemById(ctx, fsID).Execute()
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

func (r fileSystemResource) readFileSystemAPI(ctx context.Context, id string) (*clientgen.FileSystemInstance, error) {
	queries := make(url.Values)
	queries.Set("select", "id,name,description,nas_server_id,parent_id,filesystem_type,size_total,config_type,protection_policy_id,performance_policy_id,access_policy,locking_policy,folder_rename_policy,is_smb_sync_writes_enabled,is_smb_op_locks_enabled,is_smb_no_notify_enabled,is_smb_notify_on_access_enabled,is_smb_notify_on_write_enabled,smb_notify_on_change_dir_depth,is_async_MTime_enabled,file_events_publishing_mode,host_io_size,flr_attributes")
	response, _, err := r.client.GenClient.FileSystemApi.GetFileSystemById(ctx, id).Queries(queries).Execute()
	return response, err
}

func updateFsState(fsState *models.FileSystem, fsResponse *clientgen.FileSystemInstance) {
	// Update value from file system Response to State
	fsState.ID = helper.TfString(fsResponse.Id)
	fsState.Name = helper.TfString(fsResponse.Name)
	fsState.Description = helper.TfStringNN(fsResponse.Description)
	fsState.NASServerID = helper.TfString(fsResponse.NasServerId)
	if fsResponse.SizeTotal != nil {
		size, unit := convertFromBytesForFileSystem(*fsResponse.SizeTotal)
		fsState.Size = types.Float64Value(size)
		fsState.CapacityUnit = types.StringValue(unit)
	}
	fsState.ConfigType = helper.TfString(fsResponse.ConfigType)
	fsState.AccessPolicy = helper.TfString(fsResponse.AccessPolicy)
	fsState.LockingPolicy = helper.TfString(fsResponse.LockingPolicy)
	fsState.FolderRenamePolicy = helper.TfString(fsResponse.FolderRenamePolicy)
	fsState.IsAsyncMTimeEnabled = helper.TfBool(fsResponse.IsAsyncMTimeEnabled)
	fsState.ProtectionPolicyID = helper.TfStringNN(fsResponse.ProtectionPolicyId)
	fsState.PerformancePolicyID = helper.TfStringNN(fsResponse.PerformancePolicyId)
	fsState.FileEventsPublishingMode = helper.TfString(fsResponse.FileEventsPublishingMode)
	fsState.HostIOSize = helper.TfString(fsResponse.HostIoSize)
	fsState.IsSmbSyncWritesEnabled = helper.TfBool(fsResponse.IsSmbSyncWritesEnabled)
	fsState.IsSmbNoNotifyEnabled = helper.TfBool(fsResponse.IsSmbNoNotifyEnabled)
	fsState.IsSmbOpLocksEnabled = helper.TfBool(fsResponse.IsSmbOpLocksEnabled)
	fsState.IsSmbNotifyOnAccessEnabled = helper.TfBool(fsResponse.IsSmbNotifyOnAccessEnabled)
	fsState.IsSmbNotifyOnWriteEnabled = helper.TfBool(fsResponse.IsSmbNotifyOnWriteEnabled)
	fsState.SmbNotifyOnChangeDirDepth = helper.TfInt32(fsResponse.SmbNotifyOnChangeDirDepth)
	fsState.ParentID = helper.TfString(fsResponse.ParentId)
	fsState.FilesystemType = helper.TfString(fsResponse.FilesystemType)

	if fsResponse.FlrAttributes != nil {
		flr := fsResponse.FlrAttributes
		fsState.FlrAttributes, _ = types.ObjectValue(map[string]attr.Type{
			"mode":              types.StringType,
			"minimum_retention": types.StringType,
			"default_retention": types.StringType,
			"maximum_retention": types.StringType,
			"auto_lock":         types.BoolType,
			"auto_delete":       types.BoolType,
			"policy_interval":   types.Int32Type,
		}, map[string]attr.Value{
			"mode":              helper.TfString(flr.Mode),
			"minimum_retention": helper.TfString(flr.MinimumRetention),
			"default_retention": helper.TfString(flr.DefaultRetention),
			"maximum_retention": helper.TfString(flr.MaximumRetention),
			"auto_lock":         helper.TfBool(flr.AutoLock),
			"auto_delete":       helper.TfBool(flr.AutoDelete),
			"policy_interval":   helper.TfInt32(flr.PolicyInterval),
		})
	}
}
