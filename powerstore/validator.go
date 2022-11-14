package powerstore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// once we migrate to > v0.15.0 , we can use
// https://github.com/hashicorp/terraform-plugin-framework-validators/blob/v0.5.0/stringvalidator/length_at_least.go
// till then a custom implemented

// emptyStringtValidator validates that a string Attribute's must not be a empty string.
type emptyStringtValidator struct{}

// Description describes the validation in plain text formatting.
func (validator emptyStringtValidator) Description(_ context.Context) string {
	return "must not be an empty string"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator emptyStringtValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator emptyStringtValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, request, response)

	if !ok {
		return
	}

	if l := len(s); l == 0 {
		response.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			request.AttributePath,
			"Invalid empty string",
			fmt.Sprintf("Attribute %s %s, got : empty string", request.AttributePath, validator.Description(ctx)),
		))

		return
	}
}

// https://github.com/hashicorp/terraform-plugin-framework-validators/blob/4a4c520b56aa33e071c4356e0f3ac0ad2633dc7b/stringvalidator/type_validation.go#L12
// validateString ensures that the request contains a String value.
func validateString(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) (string, bool) {
	t := request.AttributeConfig.Type(ctx)
	if t != types.StringType {
		response.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			request.AttributePath,
			"Invalid empty string",
			fmt.Sprintf("Attribute %s expected value of type string, got : %s", request.AttributePath, t.String()),
		))
		return "", false
	}

	s := request.AttributeConfig.(types.String)

	if s.Unknown || s.Null {
		return "", false
	}

	return s.Value, true
}
