/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

package powerstore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// this defines the operation being executed on resource via terraform
type operation uint8

const (
	operationRead operation = iota + 1
	operationCreate
	operationUpdate
	operationDelete
	operationImport
)

// DefaultAttributePlanModifier is to set default value for an attribute
type DefaultAttributePlanModifier struct {
	value interface{}
}

// DefaultAttribute is used to set the default value
func DefaultAttribute(value interface{}) DefaultAttributePlanModifier {
	return DefaultAttributePlanModifier{value: value}
}

// PlanModifyString sets string default value
func (m DefaultAttributePlanModifier) PlanModifyString(
	ctx context.Context,
	req planmodifier.StringRequest,
	resp *planmodifier.StringResponse,
) {
	// if configuration was provided, then don't use the default
	if !req.ConfigValue.IsNull() {
		return
	}
	// If the plan is known and not null (for example due to another plan modifier),
	// don't set the default value
	if !resp.PlanValue.IsUnknown() && !resp.PlanValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.StringValue(m.value.(string))
	}
}

// PlanModifyInt64 sets int64 default value
func (m DefaultAttributePlanModifier) PlanModifyInt64(
	ctx context.Context,
	req planmodifier.Int64Request,
	resp *planmodifier.Int64Response,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		resp.PlanValue = types.Int64Value(m.value.(int64))
	}
}

// Description of default parameter
func (m DefaultAttributePlanModifier) Description(ctx context.Context) string {
	return "Use a static default value for an attribute"
}

// MarkdownDescription of default parameter
func (m DefaultAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}
