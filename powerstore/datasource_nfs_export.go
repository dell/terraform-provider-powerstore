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
	"net/url"

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

// newNFSExportDatasource returns nfsExport new datasource instance
func newNFSExportDatasource() datasource.DataSource {
	return &datasourceNFSExport{}
}

type datasourceNFSExport struct {
	client *client.Client
}

// Metadata defines datasource interface Metadata method
func (r *datasourceNFSExport) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfs_export"
}

// Schema defines datasource interface Schema method
func (r *datasourceNFSExport) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "This resource is used to manage the nfs export entity of PowerStore Array. We can Create, Update and Delete the nfs export using this resource. We can also import an existing nfs export from PowerStore array.",
		Description:         "This resource is used to manage the nfs export entity of PowerStore Array. We can Create, Update and Delete the nfs export using this resource. We can also import an existing nfs export from PowerStore array.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of the NAS Server. Conflicts with `name`, `file_system_id` and `filter_expression`.",
				MarkdownDescription: "Unique identifier of the NAS Server. Conflicts with `name`, `file_system_id` and `filter_expression`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("name"),
						path.MatchRoot("file_system_id"),
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"name": schema.StringAttribute{
				Description:         "Name of the NFS export to fetch. Conflicts with `id` and `filter_expression`.",
				MarkdownDescription: "Name of the NFS export to fetch. Conflicts with `id` and `filter_expression`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"file_system_id": schema.StringAttribute{
				Description:         "The unique identifier of the file system on which the NFS Export is created. Conflicts with `id` and `filter_expression`.",
				MarkdownDescription: "The unique identifier of the file system on which the NFS Export is created. Conflicts with `id` and `filter_expression`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.ConflictsWith(
						path.MatchRoot("filter_expression"),
					),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter NFS exports by. Conflicts with `id`.",
				MarkdownDescription: "PowerStore filter expression to filter NFS exports by. Conflicts with `id`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"nfs_exports": schema.ListNestedAttribute{
				Description:         "List of NFS exports.",
				MarkdownDescription: "List of NFS exports.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: r.nfsExportDsSchema()},
			},
		},
	}
}

// nfsExportDsSchema defines datasource interface Schema method
func (r *datasourceNFSExport) nfsExportDsSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			Description:         "The unique identifier of the NFS Export.",
			MarkdownDescription: "The unique identifier of the NFS Export.",
		},
		"file_system_id": schema.StringAttribute{
			MarkdownDescription: "The unique identifier of the file	system on which the NFS Export will be created.",
			Description:         "The unique identifier of the file system on which the NFS Export will be created.",
			Computed:            true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "The name of the NFS Export.",
			Description:         "The name of the NFS Export.",
			Computed:            true,
		},
		"path": schema.StringAttribute{
			MarkdownDescription: "The local path to export relative to the nfs export root directory. With NFS, each export of a file_system or file_nfs must have a unique local path. Before you can create additional Exports within an NFS shared folder, you must create directories within it from a Linux/Unix host that is connected to the nfs export. After a directory has been created from a mounted host, you can create a corresponding Export and Set access permissions accordingly.",
			Description:         "The local path to export relative to the nfs export root directory. With NFS, each export of a file_system or file_nfs must have a unique local path. Before you can create additional Exports within an NFS shared folder, you must create directories within it from a Linux/Unix host that is connected to the nfs export. After a directory has been created from a mounted host, you can create a corresponding Export and Set access permissions accordingly.",
			Computed:            true,
		},
		"read_write_hosts": schema.ListAttribute{
			MarkdownDescription: "List of Read-Write hosts",
			Description:         "List of Read-Write hosts",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"read_only_hosts": schema.ListAttribute{
			MarkdownDescription: "List of Read-Only hosts",
			Description:         "List of Read-Only hosts",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"no_access_hosts": schema.ListAttribute{
			MarkdownDescription: "List of hosts with no access to the NFS export or its snapshots.",
			Description:         "List of hosts with no access to the NFS export or its snapshots.",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"read_write_root_hosts": schema.ListAttribute{
			MarkdownDescription: "List of Read-Write, allow Root hosts",
			Description:         "List of Read-Write, allow Root hosts",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"read_only_root_hosts": schema.ListAttribute{
			MarkdownDescription: "List of Read-Only, allow Roots hosts",
			Description:         "List of Read-Only, allow Roots hosts",
			Computed:            true,
			ElementType:         types.StringType,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "A user-defined description of the NFS Export.",
			Description:         "A user-defined description of the NFS Export.",
			Computed:            true,
		},
		"min_security": schema.StringAttribute{
			MarkdownDescription: "The NFS enforced security type for users accessing the NFS Export. Valid values are: 'Sys', 'Kerberos', 'Kerberos_With_Integrity', 'Kerberos_With_Encryption'.",
			Description:         "The NFS enforced security type for users accessing the NFS Export. Valid values are: 'Sys', 'Kerberos', 'Kerberos_With_Integrity', 'Kerberos_With_Encryption'.",
			Computed:            true,
		},
		"anonymous_gid": schema.Int32Attribute{
			MarkdownDescription: "The GID (Group ID) of the anonymous user. This is the group ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Description:         "The GID (Group ID) of the anonymous user. This is the group ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Computed:            true,
		},
		"anonymous_uid": schema.Int32Attribute{
			MarkdownDescription: "The UID (User ID) of the anonymous user. This is the user ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Description:         "The UID (User ID) of the anonymous user. This is the user ID of the anonymous user. The anonymous user is the user ID (UID) that is used when the true user's identity cannot be determined.",
			Computed:            true,
		},
		"is_no_suid": schema.BoolAttribute{
			MarkdownDescription: "If Set, do not allow access to Set SUID. Otherwise, allow access.",
			Description:         "If Set, do not allow access to Set SUID. Otherwise, allow access.",
			Computed:            true,
		},
		"nfs_owner_username": schema.StringAttribute{
			MarkdownDescription: "The default owner of the NFS Export associated with the datastore. Required if secure NFS enabled. For NFSv3 or NFSv4 without Kerberos, the default owner is root. Was added in version 3.0.0.0.",
			Description:         "The default owner of the NFS Export associated with the datastore. Required if secure NFS enabled. For NFSv3 or NFSv4 without Kerberos, the default owner is root. Was added in version 3.0.0.0.",
			Computed:            true,
		},
		"default_access": schema.StringAttribute{
			MarkdownDescription: "The default access level for all hosts that can access the NFS Export. The default access level is the access level that is assigned to a host that is not explicitly Seted in the 'no_access_hosts', 'read_only_hosts', 'read_only_root_hosts', 'read_write_hosts', or 'read_write_root_hosts' Sets. Valid values are: 'No_Access', 'Read_Only', 'Read_Write', 'Root', 'Read_Only_Root'.",
			Description:         "The default access level for all hosts that can access the NFS Export. The default access level is the access level that is assigned to a host that is not explicitly Seted in the 'no_access_hosts', 'read_only_hosts', 'read_only_root_hosts', 'read_write_hosts', or 'read_write_root_hosts' Sets. Valid values are: 'No_Access', 'Read_Only', 'Read_Write', 'Root', 'Read_Only_Root'.",
			Computed:            true,
		},
	}
}

// Configure - defines configuration for nfs export datasource
func (r *datasourceNFSExport) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read - reads nfs export datasource information
func (r *datasourceNFSExport) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	var conf models.NFSExportDs
	diags := req.Config.Get(ctx, &conf)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var nfsExports []gopowerstore.NFSExport

	if !conf.ID.IsNull() && !conf.ID.IsUnknown() {
		nfsExportID := conf.ID.ValueString()

		// Get nfsExport details from API and then update what is in state from what the API returns
		nfsExportResponse, err := r.client.PStoreClient.GetNFSExport(context.Background(), nfsExportID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading nfs export",
				"Could not read nfs export with id "+nfsExportID+": "+err.Error(),
			)
			return
		}

		nfsExports = append(nfsExports, nfsExportResponse)
	} else {
		filters := make(map[string]string)
		if !conf.Filters.IsNull() {
			filters = convertQueriesToMap(conf.Filters.ValueQueries())
		}
		if !conf.Name.IsNull() {
			filters["name"] = "eq." + conf.Name.ValueString()
		}
		if !conf.FileSystemID.IsNull() {
			filters["file_system_id"] = "eq." + conf.FileSystemID.ValueString()
		}

		nfsExportResponse, err := r.client.PStoreClient.GetNFSExportByFilter(ctx, filters)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading nfs exports",
				"Could not read nfs exports with error "+err.Error(),
			)
			return
		}
		nfsExports = append(nfsExports, nfsExportResponse...)
		conf.ID = types.StringValue("dummy id")
	}

	// Set state
	diags = resp.State.Set(ctx, r.getState(conf, nfsExports))
	resp.Diagnostics.Append(diags...)
}

// getState - method to update terraform state
func (r *datasourceNFSExport) getState(cfg models.NFSExportDs, input []gopowerstore.NFSExport) *models.NFSExportDs {
	ret := models.NFSExportDs{
		ID:           cfg.ID,
		Name:         cfg.Name,
		FileSystemID: cfg.FileSystemID,
		Filters:      cfg.Filters,
		Items:        make([]models.NFSExportDsData, 0, len(input)),
	}
	for _, nfsExport := range input {
		ret.Items = append(ret.Items, r.getItemState(nfsExport))
	}
	return &ret
}

// nfsExportDsState - method to update terraform state
func (r *datasourceNFSExport) getItemState(input gopowerstore.NFSExport) models.NFSExportDsData {
	return models.NFSExportDsData{
		ID:               types.StringValue(input.ID),
		FileSystemID:     types.StringValue(input.FileSystemID),
		Name:             types.StringValue(input.Name),
		Description:      types.StringValue(input.Description),
		DefaultAccess:    types.StringValue(string(input.DefaultAccess)),
		Path:             types.StringValue(input.Path),
		MinSecurity:      types.StringValue(input.MinSecurity),
		NfsOwnerUsername: types.StringValue(input.NFSOwnerUsername),
		AnonymousGID:     types.Int32Value(input.AnonymousGID),
		AnonymousUID:     types.Int32Value(input.AnonymousUID),
		IsNoSUID:         types.BoolValue(input.IsNoSUID),
		RWHosts:          input.RWHosts,
		ROHosts:          input.ROHosts,
		RWRootHosts:      input.RWRootHosts,
		RORootHosts:      input.RORootHosts,
		NoAccessHosts:    input.NoAccessHosts,
	}
}

func convertQueriesToMap(in url.Values) map[string]string {
	out := make(map[string]string)
	for k := range in {
		out[k] = in.Get(k)
	}
	return out
}
