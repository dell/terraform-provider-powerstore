/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// RemoteSnapshotSessionStateEnum State of the remote snapshot session:  * Initializing - The session is being created. The status changes to Idle after the session is created.  * Idle - No other operation. The session is operating normally.  * Prepare - Indicates that the session is preparing for the data copy.  * In_Progress - Indicates that the data copy between the source and remote system has started.  * IO_Forwarding - Indicates that the session is forwarding the host IO's. Only applies to instant_access type sessions.  * Completed - Indicates that all operations completed successfully. Only applies to retrieve type sessions.  * Failed - Indicates that there was an error during the session operation.  * Cancelling - Session is being cancelled.  * Cancelled - Indicates that a user forcefully cancelled the session. Sessions which are in Prepare, In_progress & Paused state can be cancelled.  * Deleting - Session is being deleted.  * System_Paused -  A non-disruptive upgrade or Migration paused the session.  * Paused - Indicates that the session cannot be used for further backups. Sessions which are in Idle, Prepare, In_progress and Failed state can be paused.  * Rollback_In_Progress- Indicates that there was an error during the session operation and changes are getting reverted.  * Failed_Cleanup_Required- Indicates that there was an error during the session operation and while reverting the changes there is another failure. This state will be cleaned up automatically by the cleanup service that runs periodically. Session will be moved to failed state once cleanup is successful. In case of a backup session, no further backup is allowed while the session is in this state.  * Cancel_Cleanup_Required- Indicates that there was an error during the session cancel operation. This state will be cleaned up automatically by the cleanup service that runs periodically. Session will be moved to cancelled state once cleanup is successful. In case of a backup session, no further backup is allowed while the session is in this state.  * Cleanup_Required- Indicates that operation has completely successfully but while cleaning the resources locally there was an error. This state will be cleaned up automatically by the cleanup service that runs periodically. Session will be moved to idle/completed state once the cleanup is successful. In case of a backup session, no further backup is allowed while the session is in this state.  * Cleanup_In_Progress- Indicates that a cleanup is in progress.  Was added in version 3.5.0.0.
type RemoteSnapshotSessionStateEnum string

// List of RemoteSnapshotSessionStateEnum
const (
	REMOTESNAPSHOTSESSIONSTATEENUM_INITIALIZING            RemoteSnapshotSessionStateEnum = "Initializing"
	REMOTESNAPSHOTSESSIONSTATEENUM_IDLE                    RemoteSnapshotSessionStateEnum = "Idle"
	REMOTESNAPSHOTSESSIONSTATEENUM_PREPARE                 RemoteSnapshotSessionStateEnum = "Prepare"
	REMOTESNAPSHOTSESSIONSTATEENUM_IN_PROGRESS             RemoteSnapshotSessionStateEnum = "In_Progress"
	REMOTESNAPSHOTSESSIONSTATEENUM_IO_FORWARDING           RemoteSnapshotSessionStateEnum = "IO_Forwarding"
	REMOTESNAPSHOTSESSIONSTATEENUM_COMPLETED               RemoteSnapshotSessionStateEnum = "Completed"
	REMOTESNAPSHOTSESSIONSTATEENUM_FAILED                  RemoteSnapshotSessionStateEnum = "Failed"
	REMOTESNAPSHOTSESSIONSTATEENUM_CANCELLING              RemoteSnapshotSessionStateEnum = "Cancelling"
	REMOTESNAPSHOTSESSIONSTATEENUM_CANCELLED               RemoteSnapshotSessionStateEnum = "Cancelled"
	REMOTESNAPSHOTSESSIONSTATEENUM_DELETING                RemoteSnapshotSessionStateEnum = "Deleting"
	REMOTESNAPSHOTSESSIONSTATEENUM_SYSTEM_PAUSED           RemoteSnapshotSessionStateEnum = "System_Paused"
	REMOTESNAPSHOTSESSIONSTATEENUM_PAUSED                  RemoteSnapshotSessionStateEnum = "Paused"
	REMOTESNAPSHOTSESSIONSTATEENUM_ROLLBACK_IN_PROGRESS    RemoteSnapshotSessionStateEnum = "Rollback_In_Progress"
	REMOTESNAPSHOTSESSIONSTATEENUM_FAILED_CLEANUP_REQUIRED RemoteSnapshotSessionStateEnum = "Failed_Cleanup_Required"
	REMOTESNAPSHOTSESSIONSTATEENUM_CANCEL_CLEANUP_REQUIRED RemoteSnapshotSessionStateEnum = "Cancel_Cleanup_Required"
	REMOTESNAPSHOTSESSIONSTATEENUM_CLEANUP_REQUIRED        RemoteSnapshotSessionStateEnum = "Cleanup_Required"
	REMOTESNAPSHOTSESSIONSTATEENUM_CLEANUP_IN_PROGRESS     RemoteSnapshotSessionStateEnum = "Cleanup_In_Progress"
)

// All allowed values of RemoteSnapshotSessionStateEnum enum
var AllowedRemoteSnapshotSessionStateEnumEnumValues = []RemoteSnapshotSessionStateEnum{
	"Initializing",
	"Idle",
	"Prepare",
	"In_Progress",
	"IO_Forwarding",
	"Completed",
	"Failed",
	"Cancelling",
	"Cancelled",
	"Deleting",
	"System_Paused",
	"Paused",
	"Rollback_In_Progress",
	"Failed_Cleanup_Required",
	"Cancel_Cleanup_Required",
	"Cleanup_Required",
	"Cleanup_In_Progress",
}

func (v *RemoteSnapshotSessionStateEnum) Value() string {
	return string(*v)
}
