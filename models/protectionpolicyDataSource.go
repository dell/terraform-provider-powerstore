/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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

import "github.com/hashicorp/terraform-plugin-framework/types"

// ProtectionPolicyDataSource - Protection Policy DataSource properties
type ProtectionPolicyDataSource struct {
	ID               types.String       `tfsdk:"id"`
	Name             types.String       `tfsdk:"name"`
	Description      types.String       `tfsdk:"description"`
	Type             types.String       `tfsdk:"type"`
	ManagedBy        types.String       `tfsdk:"managed_by"`
	ManagedByID      types.String       `tfsdk:"managed_by_id"`
	IsReadOnly       types.Bool         `tfsdk:"is_read_only"`
	IsReplica        types.Bool         `tfsdk:"is_replica"`
	TypeL10          types.String       `tfsdk:"type_l10n"`
	ManagedByL10     types.String       `tfsdk:"managed_by_l10n"`
	VirtualMachines  []VirtualMachines  `tfsdk:"virtual_machines"`
	Volumes          []Volumes          `tfsdk:"volumes"`
	VolumeGroups     []VolumeGroups     `tfsdk:"volume_groups"`
	FileSystems      []FileSystems      `tfsdk:"file_systems"`
	PerformanceRules []PerformanceRules `tfsdk:"performance_rules"`
	SnapshotRules    []SnapshotRules    `tfsdk:"snapshot_rules"`
	ReplicationRules []ReplicationRules `tfsdk:"replication_rules"`
}

// VirtualMachines - Details of virtual machine
type VirtualMachines struct {
	ID           types.String `tfsdk:"id"`
	InstanceUUID types.String `tfsdk:"instance_uuid"`
	Name         types.String `tfsdk:"name"`
}

// Volumes - Details of volume
type Volumes struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

// VolumeGroups - Details of volume group
type VolumeGroups struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

// FileSystems - Details of file system
type FileSystems struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

// PerformanceRules - Details of performance rule
type PerformanceRules struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	IoPriority types.String `tfsdk:"io_priority"`
}

// SnapshotRules - Details of snapshot rule
type SnapshotRules struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ReplicationRules - Details of replication rule
type ReplicationRules struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
