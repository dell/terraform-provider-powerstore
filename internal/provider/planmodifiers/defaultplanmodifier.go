package planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.String = DefaultAttributePlanModifier{}
var _ planmodifier.Int64 = DefaultAttributePlanModifier{}

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
