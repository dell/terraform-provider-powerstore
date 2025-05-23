/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// PolicyTypeEnum Supported policy types.  * Protection - A protection policy for associating with volumes or volume groups, consisting of snapshot and replication rules.  * Performance - A performance policy for associating with volumes, consisting of performance rules.  * QoS - A performance policy for associating with volumes or volume groups, consisting of quality of service rules.  * File_Performance - A performance policy for associating with nas servers or file systems, consisting of file_io_limit_rules.
type PolicyTypeEnum string

// List of PolicyTypeEnum
const (
	POLICYTYPEENUM_PROTECTION       PolicyTypeEnum = "Protection"
	POLICYTYPEENUM_PERFORMANCE      PolicyTypeEnum = "Performance"
	POLICYTYPEENUM_QO_S             PolicyTypeEnum = "QoS"
	POLICYTYPEENUM_FILE_PERFORMANCE PolicyTypeEnum = "File_Performance"
)

// All allowed values of PolicyTypeEnum enum
var AllowedPolicyTypeEnumEnumValues = []PolicyTypeEnum{
	"Protection",
	"Performance",
	"QoS",
	"File_Performance",
}

func (v *PolicyTypeEnum) Value() string {
	return string(*v)
}
