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
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// updateNasServerState returns the list of NAS servers
func updateNasServerState(nasServers []gopowerstore.NAS) (response []models.NasServerDataSource) {
	for _, nasServer := range nasServers {
		response = append(response, newNasServer(nasServer))
	}
	return response
}

// newNasServer returns a new NAS server
func newNasServer(input gopowerstore.NAS) models.NasServerDataSource {
	return models.NasServerDataSource{
		Id:                              types.StringValue(input.ID),
		Name:                            types.StringValue(input.Name),
		CurrentNodeId:                   types.StringValue(input.CurrentNodeID),
		CurrentPreferredIPv4InterfaceId: types.StringValue(input.CurrentPreferredIPv4InterfaceID),
		CurrentPreferredIPv6InterfaceId: types.StringValue(input.CurrentPreferredIPv6InterfaceID),
		DefaultUnixUser:                 types.StringValue(input.DefaultUnixUser),
		DefaultWindowsUser:              types.StringValue(input.DefaultWindowsUser),
		ProductionIPv4InterfaceId:       types.StringValue(input.BackupIPv4InterfaceID),
		ProductionIPv6InterfaceId:       types.StringValue(input.ProductionIPv6InterfaceID),
		ProtectionPolicyId:              types.StringValue(input.ProtectionPolicyID),
		CurrentUnixDirectoryService:     types.StringValue(input.CurrentUnixDirectoryService),
		CurrentUnixDirectoryServiceL10n: types.StringValue(input.CurrentUnixDirectoryServiceL10n),
		Description:                     types.StringValue(input.Description),
		BackupIPv4InterfaceId:           types.StringValue(input.BackupIPv4InterfaceID),
		BackupIPv6InterfaceId:           types.StringValue(input.BackupIPv6InterfaceID),
		FileEventsPublishingMode:        types.StringValue(input.FileEventsPublishingMode),
		FileEventsPublishingModeL10n:    types.StringValue(input.FileEventsPublishingModeL10n),
		IsAutoUserMappingEnabled:        types.BoolValue(input.IsAutoUserMappingEnabled),
		IsDrTest:                        types.BoolValue(input.IsDRTest),
		IsProductionModeEnabled:         types.BoolValue(input.IsProductionModeEnabled),
		IsReplicationDestination:        types.BoolValue(input.IsReplicationDestination),
		IsUsernameTranslationEnabled:    types.BoolValue(input.IsUsernameTranslationEnabled),
		OperationalStatus:               types.StringValue(string(input.OperationalStatus)),
		OperationalStatusL10n:           types.StringValue(input.OperationalStatusL10n),
		PreferredNodeId:                 types.StringValue(input.PreferredNodeID),
	}
}
