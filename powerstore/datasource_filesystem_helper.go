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
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// updateFileSystemState returns the list of filesystems
func updateFileSystemState(filesystems []gopowerstore.FileSystem) (response []models.FileSystemDatasource) {
	for _, filesystem := range filesystems {
		response = append(response, newFileSystem(filesystem))
	}
	return response
}

// newFileSystem returns a new filesystem
func newFileSystem(input gopowerstore.FileSystem) models.FileSystemDatasource {
	return models.FileSystemDatasource{
		AccessPolicy:               types.StringValue(input.AccessPolicy),
		AccessType:                 types.StringValue(input.AccessType),
		ConfigType:                 types.StringValue(input.ConfigType),
		Description:                types.StringValue(input.Description),
		ExpirationTimestamp:        types.StringValue(input.ExpirationTimestamp),
		FilesystemType:             types.StringValue(string(input.FilesystemType)),
		FolderRenamePolicy:         types.StringValue(input.FolderRenamePolicy),
		ID:                         types.StringValue(input.ID),
		IsAsyncMTimeEnabled:        types.BoolValue(input.IsAsyncMTimeEnabled),
		IsSmbNoNotifyEnabled:       types.BoolValue(input.IsSmbNoNotifyEnabled),
		IsSmbNotifyOnAccessEnabled: types.BoolValue(input.IsSmbNotifyOnAccessEnabled),
		IsSmbNotifyOnWriteEnabled:  types.BoolValue(input.IsSmbNotifyOnWriteEnabled),
		IsSmbOpLocksEnabled:        types.BoolValue(input.IsSmbOpLocksEnabled),
		IsSmbSyncWritesEnabled:     types.BoolValue(input.IsSmbSyncWritesEnabled),
		LockingPolicy:              types.StringValue(input.LockingPolicy),
		Name:                       types.StringValue(input.Name),
		NasServerID:                types.StringValue(input.NasServerID),
		ParentID:                   types.StringValue(input.ParentID),
		ProtectionPolicyID:         types.StringValue(input.ProtectionPolicyID),
		SizeTotal:                  types.Int64Value(int64(input.SizeTotal)),
		SizeUsed:                   types.Int64Value(int64(input.SizeUsed)),
		SmbNotifyOnChangeDirDepth:  types.Int64Value(int64(input.SmbNotifyOnChangeDirDepth)),
		FlrAttributes: models.FLRAttributesDatasource{
			DefaultRetention: types.StringValue(input.FlrCreate.DefaultRetention),
			MaximumRetention: types.StringValue(input.FlrCreate.MaximumRetention),
			MinimumRetention: types.StringValue(input.FlrCreate.MinimumRetention),
			Mode:             types.StringValue(input.FlrCreate.Mode),
		},
	}
}
