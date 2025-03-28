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

var resourceFacts = map[string]Facts{
	// Host Access Management
	"host": {
		Note: "~> **Note:** `name`, `os_type` and `initiators` are the required attributes to create." +
			"\n~> **Note:** `os_type` cannot be updated." +
			"\n~> **Note:** `port_name` is the required attribute for `initiators`." +
			"\n~> **Note:** `chap_single_password` must be present when `chap_single_username` is given and vice-versa." +
			"\n~> **Note:** `chap_mutual_password` must be present when `chap_mutual_username` is given and vice-versa." +
			"\n~> **Note:** `chap_mutual_username` and `chap_mutual_password` can be used only when `chap_single_username` and `chap_single_password` are present.",
		ExampleVar:  "host",
		SubCategory: "Host Access Management",
	},
	"hostgroup": {
		Note: "~> **Note:** Exactly one of `host_ids` and `host_names` is required." +
			"\n~> **Note:** `host_connectivity` cannot be used while creating host group resource but it can be used while updating the host group resource.",
		ExampleVar:  "host group",
		SubCategory: "Host Access Management",
	},
	// Data Protection Management
	"protectionpolicy": {
		ExampleVar:  "protection policy",
		SubCategory: "Data Protection Management",
	},
	"replication_rule": {
		ExampleVar:  "replication rule",
		SubCategory: "Data Protection Management",
	},
	"snapshotrule": {
		ExampleVar:  "snapshot rule",
		SubCategory: "Data Protection Management",
	},
	"volume_snapshot": {
		Note: "> **Note:** `volume_id`/`volume_name` is the required attribute to create volume snapshot." +
			"\n> **Note:** if `name` is present in the config it cannot be blank(\"\"). if absent, default value is allocated to it." +
			"\n> **Note:** During create operation, if `expiration_timestamp` is not specified or set to blank(\"\"), snapshot will be created with infinite retention." +
			"\n> **Note:** During modify operation, to set infinite retention, `expiration_timestamp` can be set to blank(\"\")." +
			"\n> **Note:** Volume DataSource can be used to fetch volume ID/Name for volume snapshot creation." +
			"\n> **Note:** Exactly one of `volume_id` and `volume_name` should be provided.",
		ExampleVar:  "volume snapshot",
		SubCategory: "Data Protection Management",
	},
	"volumegroup_snapshot": {
		Note: "~> **Note:** `volume_group_id`/`volume_group_name` is the required attribute to create volume group snapshot." +
			"\n~> **Note:** `expiration_timestamp` if present in config cannot be blank(\"\"). if absent, default value is allocated to it." +
			"\n~> **Note:** During create operation, if `expiration_timestamp` is not specified or set to blank(\"\"), snapshot will be created with infinite retention." +
			"\n~> **Note:** During modify operation, to set infinite retention, `expiration_timestamp` can be set to blank(\"\")." +
			"\n~> **Note:** Volume group DataSource can be used to fetch volume group ID/Name." +
			"\n~> **Note:** Exactly one of `volume_group_id` and `volume_group_name` should be provided.",
		ExampleVar:  "volume group snapshot",
		SubCategory: "Data Protection Management",
	},
	"filesystem_snapshot": {
		ExampleVar:  "filesystem snapshot",
		SubCategory: "Data Protection Management",
	},
	// Block Storage Management
	"volume": {
		ExampleVar:  "volume",
		SubCategory: "Block Storage Management",
	},
	"volumegroup": {
		Note: "> **Note:** Exactly one of `volume_ids` and `volume_names` is required." +
			"\n> **Note:** Exactly one of `protection_policy_id` and `protection_policy_name` is required.",
		ExampleVar:  "volume group",
		SubCategory: "Block Storage Management",
	},
	"storagecontainer": {
		ExampleVar:  "storage container",
		SubCategory: "Block Storage Management",
	},
	// File Storage Management
	"filesystem": {
		ExampleVar:  "filesystem",
		SubCategory: "File Storage Management",
	},
	"nas_server": {
		ExampleVar:  "NAS Server",
		SubCategory: "File Storage Management",
	},
	"nfs_export": {
		ExampleVar:  "NFS Export",
		SubCategory: "File Storage Management",
	},
	"smb_share": {
		SubCategory: "File Storage Management",
	},
}
