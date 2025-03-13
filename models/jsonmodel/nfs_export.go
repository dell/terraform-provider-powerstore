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

package jsonmodel

// NFSExportModify - json model for modifying nfs export
// TODO: Move to using auto-generated client library
type NFSExportModify struct {
	// An optional description for the host.
	// The description should not be more than 256 UTF-8 characters long and should not have any unprintable characters.
	// Empty string is a valid description value, so omitempty should not be used.
	Description string `json:"description"`

	// Read-Write
	// Hosts to add to the current read_write_hosts list. Hosts can be entered by Hostname, IP addresses
	AddRWHosts []string `json:"add_read_write_hosts,omitempty"`
	// Hosts to remove from the current read_write_hosts list. Hosts can be entered by Hostname, IP addresses.
	RemoveRWHosts []string `json:"remove_read_write_hosts,omitempty"`

	// Read-Only
	// Hosts to add to the current read_only_hosts list. Hosts can be entered by Hostname, IP addresses
	AddROHosts []string `json:"add_read_only_hosts,omitempty"`
	// Hosts to remove from the current read_only_hosts list. Hosts can be entered by Hostname, IP addresses.
	RemoveROHosts []string `json:"remove_read_only_hosts,omitempty"`

	// Read-Write, allow Root
	// Hosts to add to the current read_write_root_hosts list. Hosts can be entered by Hostname, IP addresses
	AddRWRootHosts []string `json:"add_read_write_root_hosts,omitempty"`
	// Hosts to remove from the current read_write_root_hosts list. Hosts can be entered by Hostname, IP addresses.
	RemoveRWRootHosts []string `json:"remove_read_write_root_hosts,omitempty"`

	// Read-Only, allow Roots
	// Hosts to add to the current read_only_hosts list. Hosts can be entered by Hostname, IP addresses
	AddRORootHosts []string `json:"add_read_only_root_hosts,omitempty"`
	// Hosts to remove from the current read_only_hosts list. Hosts can be entered by Hostname, IP addresses.
	RemoveRORootHosts []string `json:"remove_read_only_root_hosts,omitempty"`

	// No-Access
	// Hosts to add to the current no_access_hosts list. Hosts can be entered by Hostname, IP addresses
	AddNoAccessHosts []string `json:"add_no_access_hosts,omitempty"`
	// Hosts to remove from the current no_access_hosts list. Hosts can be entered by Hostname, IP addresses
	RemoveNoAccessHosts []string `json:"remove_no_access_hosts,omitempty"`

	// Default access level for all hosts that can access the Export.
	DefaultAccess string `json:"default_access,omitempty"`
	// NFS enforced security type for users accessing an NFS Export.
	MinSecurity string `json:"min_security,omitempty"`
	// Hosts with no access to the NFS export or its snapshots.
	// Empty list is a valid value, so omitempty should not be used.
	NoAccessHosts []string `json:"no_access_hosts"`
	// Hosts with read-only access to the NFS export and its snapshots.
	// Empty list is a valid value, so omitempty should not be used.
	ReadOnlyHosts []string `json:"read_only_hosts"`
	// Hosts with read-only and read-only for root user access to the NFS Export and its snapshots.
	// Empty list is a valid value, so omitempty should not be used.
	ReadOnlyRootHosts []string `json:"read_only_root_hosts"`
	// Hosts with read and write access to the NFS export and its snapshots.
	// Empty list is a valid value, so omitempty should not be used.
	ReadWriteHosts []string `json:"read_write_hosts"`
	// Hosts with read and write and read and write for root user access to the NFS Export and its snapshots.
	// Empty list is a valid value, so omitempty should not be used.
	ReadWriteRootHosts []string `json:"read_write_root_hosts"`
	// Specifies the user ID of the anonymous account.
	// Zero ID is a valid value, so omitempty should not be used.
	AnonymousUID int32 `json:"anonymous_UID"`
	// Specifies the group ID of the anonymous account.
	// Zero ID is a valid value, so omitempty should not be used.
	AnonymousGID int32 `json:"anonymous_GID"`
	// If set, do not allow access to set SUID. Otherwise, allow access.
	IsNoSUID bool `json:"is_no_SUID"`
	// (*Applies to NFS shares of VMware NFS storage resources.*) Default owner of the NFS Export associated with the datastore. Required if secure NFS enabled. For NFSv3 or NFSv4 without Kerberos, the default owner is root. Was added in version 3.0.0.0.
	NFSOwnerUsername string `json:"nfs_owner_username,omitempty"`
}
