/*
PowerStore REST API

Storage cluster REST API definition. ( For \"Try It Out\", use the cluster management IP address to load this swaggerui interface. )

API version: 4.1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package clientgen

// LoginSessionInstance struct for LoginSessionInstance
type LoginSessionInstance struct {
	// Unique identifier of the login session.
	Id *string `json:"id,omitempty"`
	// Fully qualified user account name being used to log in.
	User *string `json:"user,omitempty"`
	// Roles to which the logged-in user is mapped.
	RoleIds []string `json:"role_ids,omitempty"`
	// Remaining idle time until the session will expire, in seconds.
	IdleTimeout *int32 `json:"idle_timeout,omitempty"`
	// Indicates whether the logged-in user requires a password change.
	IsPasswordChangeRequired *bool `json:"is_password_change_required,omitempty"`
	// Indicates whether the logged-in user is predefined.
	IsBuiltInUser *bool `json:"is_built_in_user,omitempty"`
}
