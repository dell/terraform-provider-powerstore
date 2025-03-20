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

type FSModify struct {
	//Size, in bytes, presented to the host or end user. This can be used for both expand and shrink on a file system.
	Size                       int           `json:"size_total,omitempty"`
	Description                string        `json:"description"` // empty to unassign
	AccessPolicy               string        `json:"access_policy,omitempty"`
	LockingPolicy              string        `json:"locking_policy,omitempty"`
	FolderRenamePolicy         string        `json:"folder_rename_policy,omitempty"`
	IsSmbSyncWritesEnabled     *bool         `json:"is_smb_sync_writes_enabled,omitempty"`
	IsSmbOpLocksEnabled        *bool         `json:"is_smb_op_locks_enabled,omitempty"`
	IsSmbNotifyOnAccessEnabled *bool         `json:"is_smb_notify_on_access_enabled,omitempty"`
	IsSmbNotifyOnWriteEnabled  *bool         `json:"is_smb_notify_on_write_enabled,omitempty"`
	SmbNotifyOnChangeDirDepth  int32         `json:"smb_notify_on_change_dir_depth,omitempty"`
	IsSmbNoNotifyEnabled       *bool         `json:"is_smb_no_notify_enabled,omitempty"`
	IsAsyncMtimeEnabled        *bool         `json:"is_async_MTime_enabled,omitempty"`
	ProtectionPolicyID         string        `json:"protection_policy_id"` // empty to unassign
	FileEventsPublishingMode   string        `json:"file_events_publishing_mode,omitempty"`
	FlrCreate                  FlrAttributes `json:"flr_attributes,omitempty"`
	ExpirationTimestamp        string        `json:"expiration_timestamp,omitempty"`
}

type FlrAttributes struct {
	Mode                 string `json:"mode,omitempty"`
	MinimumRetention     string `json:"minimum_retention,omitempty"`
	DefaultRetention     string `json:"default_retention,omitempty"`
	MaximumRetention     string `json:"maximum_retention,omitempty"`
	AutoLock             *bool  `json:"auto_lock,omitempty"`
	AutoDelete           *bool  `json:"auto_delete,omitempty"`
	PolicyInterval       int32  `json:"policy_interval,omitempty"`
	HasProtectedFiles    bool   `json:"has_protected_files,omitempty"`
	ClockTime            string `json:"clock_time,omitempty"`
	MaximumRetentionDate string `json:"maximum_retention_date,omitempty"`
}
