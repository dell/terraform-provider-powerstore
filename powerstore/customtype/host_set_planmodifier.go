package customtype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// HostSetPlanModifier implements the plan modifier.
type HostSetPlanModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m HostSetPlanModifier) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m HostSetPlanModifier) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m HostSetPlanModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// Do nothing if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Do nothing if there is no known plan value.
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}

	// convert to HostSetValue
	// ignoring diags here, but this will give errors if used with any type other than HostSetType
	planSet, _ := NewHostSetType().ValueFromSet(ctx, req.PlanValue)
	stateSet, _ := NewHostSetType().ValueFromSet(ctx, req.StateValue)
	plan, state := planSet.(HostSetValue), stateSet.(HostSetValue)
	// check if sets are equal semantically
	ok, dgs := plan.SetSemanticEquals(ctx, state)
	if dgs.HasError() {
		resp.Diagnostics.Append(dgs...)
		return
	}
	if ok {
		// set plan value to state value
		resp.PlanValue = req.StateValue
	}
}
