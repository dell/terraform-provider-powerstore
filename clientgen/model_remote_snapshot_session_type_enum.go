/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// RemoteSnapshotSessionTypeEnum Type of the remote snapshot session:  * Backup - Remote snapshot session to backup the snapshot to remote.  * Retrieve - Remote snapshot session to retrieve the snapshot from the remote.  * Instant_Access - Provides instant access to remote backup snapshot through proxy volume/volume_group without bringing the remote copy on to the PoweStore.  Was added in version 3.5.0.0.
type RemoteSnapshotSessionTypeEnum string

// List of RemoteSnapshotSessionTypeEnum
const (
	REMOTESNAPSHOTSESSIONTYPEENUM_BACKUP         RemoteSnapshotSessionTypeEnum = "Backup"
	REMOTESNAPSHOTSESSIONTYPEENUM_RETRIEVE       RemoteSnapshotSessionTypeEnum = "Retrieve"
	REMOTESNAPSHOTSESSIONTYPEENUM_INSTANT_ACCESS RemoteSnapshotSessionTypeEnum = "Instant_Access"
)

// All allowed values of RemoteSnapshotSessionTypeEnum enum
var AllowedRemoteSnapshotSessionTypeEnumEnumValues = []RemoteSnapshotSessionTypeEnum{
	"Backup",
	"Retrieve",
	"Instant_Access",
}

func (v *RemoteSnapshotSessionTypeEnum) Value() string {
	return string(*v)
}
