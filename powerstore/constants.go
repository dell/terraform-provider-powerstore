package powerstore

const (
	//CreatePPDetailErrorMsg specifies error details occured while creating protection policy
	CreatePPDetailErrorMsg = "Could not create protection policy"

	//ReadPPDetailErrorMsg specifies error details occured while reading protection policy
	ReadPPDetailErrorMsg = "Could not read protection policy"

	//UpdatePPDetailErrorMsg specifies error details occured while updating protection policy
	UpdatePPDetailErrorMsg = "Could not update protection"

	//DeletePPDetailErrorMsg specifies error details occured while deleting protection policy
	DeletePPDetailErrorMsg = "Could not delete protection policy"

	//ImportPPDetailErrorMsg specifies error details occured while importing protection policy
	ImportPPDetailErrorMsg = "Could not read protection policy with error invalid-id"

	//CreateSRDetailErrorMsg specifies error details occured while creating snapshot rule
	CreateSRDetailErrorMsg = "Could not create snapshot rule"

	//ReadSRDetailErrorMsg specifies error details occured while reading snapshot rule
	ReadSRDetailErrorMsg = "Could not read snapshot rule"

	//UpdateSRDetailErrorMsg specifies error details occured while updating snapshot rule
	UpdateSRDetailErrorMsg = "Could not update snapshot rule"

	//DeleteSRDetailErrorMsg specifies error details occured while deleting snapshot rule
	DeleteSRDetailErrorMsg = "Could not delete snapshot rule"

	//ImportSRDetailErrorMsg specifies error details occured while importing snapshot rule
	ImportSRDetailErrorMsg = "Could not read snapshot rule with error invalid-id"

	//CreateSCDetailErrorMsg specifies error details occured while creating storage container
	CreateSCDetailErrorMsg = "Could not create Storage Container"

	//ReadSCDetailErrorMsg specifies error details occured while reading storage container
	ReadSCDetailErrorMsg = "Could not read storage container"

	//UpdateSCDetailErrorMsg specifies error details occured while updating storage container
	UpdateSCDetailErrorMsg = "Could not update storageContainer"

	//DeleteSCDetailErrorMsg specifies error details occured while deleting storage container
	DeleteSCDetailErrorMsg = "Could not delete storage container"

	//ImportSCDetailErrorMsg specifies error details occured while importing storage container
	ImportSCDetailErrorMsg = "Could not read storageContainerID with error invalid-id"

	//CreateVolumeDetailErrorMsg specifies error details occured while creating volume
	CreateVolumeDetailErrorMsg = "Could not create volume"

	//ReadVolumeDetailErrorMsg specifies error details occured while reading storage volume
	ReadVolumeDetailErrorMsg = "Could not read volume"

	//UpdateVolumeDetailErrorMsg specifies error details occured while updating volume
	UpdateVolumeDetailErrorMsg = "Could not update volume"

	//DeleteVolumeDetailErrorMsg specifies error details occured while deleting volume
	DeleteVolumeDetailErrorMsg = "Could not delete volume"

	//ImportVolumeDetailErrorMsg specifies error details occured while importing volume
	ImportVolumeDetailErrorMsg = "Could not read volume with error invalid-id"

	//InvalidAttributeErrorMsg specifies error details for invalid attributes used while creating volume
	InvalidAttributeErrorMsg = "Invalid Attribute Value Match"

	//InvalidSizeErrorMsg specifies error caused due to invalid size while updating volume
	InvalidSizeErrorMsg = "Failed to update all parameters of Volume"

	//InvalidAppliaceErrorMsg specifies error caused due to invalid appliance while creating/updating volume
	InvalidAppliaceErrorMsg = "Unable to find an appliance"

	//HostIDHostGroupErrorMsg sepecifies error caused by specifying both host id and hostgroup id while updating volume
	HostIDHostGroupErrorMsg = "Either of HostID and Host GroupID should be present."

	//CreateVolumeHostIDErrorMsg sepecifies error caused by specifying both host id and hostgroup id while creating volume
	CreateVolumeHostIDErrorMsg = "Could not create volume Either HostID or HostGroupID can be present"

	//InvalidStorageProtocolErrorMsg specifies error caused by specifying invalid storage_protocol while creating snapshot rule
	InvalidStorageProtocolErrorMsg = "Attribute storage_protocol value must be one of"

	//InvalidIntervalErrorMsg specifies error caused by specifying invalid interval while creating snapshot rule
	InvalidIntervalErrorMsg = "Attribute interval value must be one of"

	//InvalidTimezoneErrorMsg - specifies error caused by specifying invalid timezone while creating snapshot rule
	InvalidTimezoneErrorMsg = "Attribute timezone value must be one of"

	//InvalidDaysOfWeekErrorMsg - specifies error caused by specifying invalid days_of_week while creating snapshot rule
	InvalidDaysOfWeekErrorMsg = "Attribute days_of_week[^ ]* value must be one of"

	//InvalidNasAccessTypeErrorMsg - specifies error caused by specifying invalid nas_access_type while creating snapshot rule
	InvalidNasAccessTypeErrorMsg = "Attribute nas_access_type value must be one of"

	//SnapshotIDSnapshotNameErroMsg specifies error caused if both snapshot_rule_id and snapshot_rule_name are specified while creating protection policy
	SnapshotIDSnapshotNameErroMsg = "Invalid Attribute Combination"

	//ReplicationIDReplicationNameErrorMsg specifies error caused if both replication_rule_id and replication_rule_name are specified while creating protection policy
	ReplicationIDReplicationNameErrorMsg = "Invalid Attribute Combination"

	//CreateVolumeGroupDetailErrorMsg specifies error caused when invalid attribute values are provided
	CreateVolumeGroupDetailErrorMsg = "Error creating volume group"

	//CreateResourceMissingErrorMsg specifies error caused when required attribute value is not provided
	CreateResourceMissingErrorMsg = "Missing required argument"

	//ImportHostDetailErrorMsg specifies error details while importing host
	ImportHostDetailErrorMsg = "Could not read hostID with error invalid-id"

	//CreateVolumeGroupMissingErrorMsg specifies error caused when required attribute value is not provided
	CreateVolumeGroupMissingErrorMsg = "Missing required argument"

	//UpdateVolumeGroupDetailErrorMsg specifies error caused when invalid attribute values are provided
	UpdateVolumeGroupDetailErrorMsg = "Error updating volume group"

	//CreateVolumeGroupInvalidPolicyErrorMsg specifies error caused when invalid protection policy value is provided
	CreateVolumeGroupInvalidPolicyErrorMsg = "Error getting protection policy"

	//VolumeGroupInvalidAttributeCombinationErrorMsg specifies error caused if mutually exclusive paramaters are provided
	VolumeGroupInvalidAttributeCombinationErrorMsg = "Invalid Attribute Combination"

	//InvalidLengthErrorMsg specifies error details for empty name/ID
	InvalidLengthErrorMsg = "Invalid Attribute Value Length"

	//CreateHGDetailErrorMsg specifies error caused when required fields are not provided
	CreateHGDetailErrorMsg = "Missing required argument"

	//CreateHGInvalidHostErrorMsg specifies error caused when Invalid attribute value is provided
	CreateHGInvalidHostErrorMsg = "Error creating host group"

	//CreateHGWithBlankName specifies error when blank name is provided
	CreateHGWithBlankName = "Invalid Attribute Value Length"

	//ImportVGDetailErrorMsg specifies error details occured while importing volume group
	ImportVGDetailErrorMsg = "Could not read volume group with error invalid-id"

	//ImportHGDetailErrorMsg specifies error details occured while importing host group
	ImportHGDetailErrorMsg = "Could not read host group with error invalid-id"

	//CreateSnapshotErrorMsg specifies error caused when invalid attribute values are provided during snapshot creation
	CreateSnapshotErrorMsg = "Error creating volume snapshot"

	//CreateHostGroupInvalidHostErrorMsg specifies error caused when invalid host value is provided
	CreateHostGroupInvalidHostErrorMsg = "Error getting host"

	//InvalidAttributeCombinationErrorMsg specifies error caused if mutually exclusive paramaters are provided
	InvalidAttributeCombinationErrorMsg = "Invalid Attribute Combination"
)
