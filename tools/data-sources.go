//go:build tools

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

package main

var datasourceFacts = map[string]Facts{
	"host": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_host.test1.attribute_name",
	},
	"hostgroup": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_hostgroup.test1.attribute_name",
	},
	"protectionpolicy": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_protectionpolicy.test1.attribute_name",
	},
	"replication_rule": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_replication_rule.test1.attribute_name",
	},
	"snapshotrule": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_replication_rule.test1.attribute_name",
	},
	"volume_snapshot": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_volume_snapshot.test1.attribute_name",
	},
	"volume": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_volume.test1.attribute_name",
	},
	"volumegroup_snapshot": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_volumegroup_snapshot.test1.attribute_name",
	},
	"volumegroup": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_volumegroup.test1.attribute_name",
	},
	"filesystem": {
		ExampleVar: "data.powerstore_filesystem.test1",
	},
	"filesystem_snapshot": {
		Note:       "> **Note:** Only one of `name` or `id` can be provided at a time.",
		ExampleVar: "data.powerstore_filesystem_snapshot.test1.attribute_name",
	},
	"nas_server": {
		ExampleVar: "data.powerstore_nas_server.test1.attribute_name",
	},
	"nfs_export": {
		Note:       "> **Note:** `id` and `filter_expression` cannot be used with any other attribute. `name` and `file_system_id` can be used together.",
		ExampleVar: "data.powerstore_nfs_export.nfs_export_by_name_regex.attribute_name",
	},
	"smb_share": {
		Note:       "> **Note:** Only one of `name`, `id`, `file_system_id` or `filter_expression` can be provided at a time.",
		ExampleVar: "data.powerstore_smb_share.smb_share_by_filters.attribute_name",
	},
	"remote_system": {
		Note:       "> **Note:** Only one of `name`, `id` or `filter_expression` can be provided at a time.",
		ExampleVar: "data.powerstore_remote_system.remote_system_by_filters.attribute_name",
	},
}
