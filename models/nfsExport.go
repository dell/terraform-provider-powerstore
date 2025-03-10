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

package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-powerstore/powerstore/customtypes/nfshostset"
)

type NFSExport struct {
	ID                 types.String            `tfsdk:"id"`
	FileSystemID       types.String            `tfsdk:"file_system_id"`
	Name               types.String            `tfsdk:"name"`
	Path               types.String            `tfsdk:"path"`
	AnonymousGID       types.Int32             `tfsdk:"anonymous_gid"`
	AnonymousUID       types.Int32             `tfsdk:"anonymous_uid"`
	Description        types.String            `tfsdk:"description"`
	IsNoSUID           types.Bool              `tfsdk:"is_no_suid"`
	MinSecurity        types.String            `tfsdk:"min_security"`
	NfsOwnerUsername   types.String            `tfsdk:"nfs_owner_username"`
	DefaultAccess      types.String            `tfsdk:"default_access"`
	NoAccessHosts      nfshostset.HostSetValue `tfsdk:"no_access_hosts"`
	ReadOnlyHosts      nfshostset.HostSetValue `tfsdk:"read_only_hosts"`
	ReadOnlyRootHosts  nfshostset.HostSetValue `tfsdk:"read_only_root_hosts"`
	ReadWriteHosts     nfshostset.HostSetValue `tfsdk:"read_write_hosts"`
	ReadWriteRootHosts nfshostset.HostSetValue `tfsdk:"read_write_root_hosts"`
}

type NFSExportDs struct {
	ID           types.String          `tfsdk:"id"`
	FileSystemID types.String          `tfsdk:"file_system_id"`
	Name         types.String          `tfsdk:"name"`
	Filters      FilterExpressionValue `tfsdk:"filter_expression"`
	Items        []NFSExportDsData     `tfsdk:"nfs_exports"`
}

type NFSExportDsData struct {
	ID               types.String `tfsdk:"id"`
	FileSystemID     types.String `tfsdk:"file_system_id"`
	Name             types.String `tfsdk:"name"`
	Path             types.String `tfsdk:"path"`
	RWHosts          []string     `tfsdk:"read_write_hosts"`
	ROHosts          []string     `tfsdk:"read_only_hosts"`
	RWRootHosts      []string     `tfsdk:"read_write_root_hosts"`
	RORootHosts      []string     `tfsdk:"read_only_root_hosts"`
	NoAccessHosts    []string     `tfsdk:"no_access_hosts"`
	AnonymousGID     types.Int32  `tfsdk:"anonymous_gid"`
	AnonymousUID     types.Int32  `tfsdk:"anonymous_uid"`
	Description      types.String `tfsdk:"description"`
	IsNoSUID         types.Bool   `tfsdk:"is_no_suid"`
	MinSecurity      types.String `tfsdk:"min_security"`
	NfsOwnerUsername types.String `tfsdk:"nfs_owner_username"`
	DefaultAccess    types.String `tfsdk:"default_access"`
}
