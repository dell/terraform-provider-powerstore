/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// MigrationResourceTypeEnum Storage resource types eligible for migration. Values are: * volume * virtual_volume * volume_group * virtual_machine * replication_group  Values was added in 3.0.0.0: virtual_machine, replication_group.
type MigrationResourceTypeEnum string

// List of MigrationResourceTypeEnum
const (
	MIGRATIONRESOURCETYPEENUM_VOLUME            MigrationResourceTypeEnum = "volume"
	MIGRATIONRESOURCETYPEENUM_VIRTUAL_VOLUME    MigrationResourceTypeEnum = "virtual_volume"
	MIGRATIONRESOURCETYPEENUM_VOLUME_GROUP      MigrationResourceTypeEnum = "volume_group"
	MIGRATIONRESOURCETYPEENUM_VIRTUAL_MACHINE   MigrationResourceTypeEnum = "virtual_machine"
	MIGRATIONRESOURCETYPEENUM_REPLICATION_GROUP MigrationResourceTypeEnum = "replication_group"
)

// All allowed values of MigrationResourceTypeEnum enum
var AllowedMigrationResourceTypeEnumEnumValues = []MigrationResourceTypeEnum{
	"volume",
	"virtual_volume",
	"volume_group",
	"virtual_machine",
	"replication_group",
}

func (v *MigrationResourceTypeEnum) Value() string {
	return string(*v)
}
