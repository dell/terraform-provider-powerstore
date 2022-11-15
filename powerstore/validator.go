package powerstore

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
func (validator emptyStringtValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	s, ok := validateString(ctx, req, resp)

	if !ok {
		return
	}

	if l := len(s); l == 0 {
		resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			req.AttributePath,
			"Invalid empty string",
			fmt.Sprintf("Attribute %s %s, got : empty string", req.AttributePath, validator.Description(ctx)),
		))

		return
	}
}

// oneOfStringtValidator validates that a string Attribute's must be one of given acceptableStringValues
type oneOfStringtValidator struct {
	acceptableStringValues []string
	isList                 bool
}

// Description describes the validation in plain text formatting.
func (validator oneOfStringtValidator) Description(_ context.Context) string {
	return fmt.Sprintf("must be one of these : %s", strings.Join(validator.acceptableStringValues, " , "))
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator oneOfStringtValidator) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("must be one of these :\n%s", strings.Join(validator.acceptableStringValues, " \n "))
}

// Validate performs the validation.
func (validator oneOfStringtValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {

	// using same validator for list too
	// once migtrated to > v0.15.0 , will make use of terraform-plugin-framework-validators
	if validator.isList {

		elems, ok := validateList(ctx, req, resp)
		if !ok {
			return
		}

		for k, elem := range elems {
			attrPath := req.AttributePath.AtListIndex(k)
			request := tfsdk.ValidateAttributeRequest{
				AttributePath:           attrPath,
				AttributePathExpression: attrPath.Expression(),
				AttributeConfig:         elem,
				Config:                  req.Config,
			}

			oneOfStringtValidator{
				acceptableStringValues: validator.acceptableStringValues,
			}.Validate(ctx, request, resp)
		}

		return
	}

	s, ok := validateString(ctx, req, resp)

	if !ok {
		return
	}

	for _, acceptableString := range validator.acceptableStringValues {
		if s == acceptableString {
			return
		}
	}

	resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
		req.AttributePath,
		"Invalid unknown value",
		fmt.Sprintf("Attribute %s %s :: got : %s", req.AttributePath, validator.Description(ctx), s),
	))
}

// https://github.com/hashicorp/terraform-plugin-framework-validators/blob/4a4c520b56aa33e071c4356e0f3ac0ad2633dc7b/stringvalidator/type_validation.go#L12
// validateString ensures that the request contains a String value.
func validateString(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) (string, bool) {
	t := req.AttributeConfig.Type(ctx)
	if t != types.StringType {
		resp.Diagnostics.Append(diag.NewAttributeErrorDiagnostic(
			req.AttributePath,
			"Invalid value",
			fmt.Sprintf("Attribute %s expected value of type string :: got : %s", req.AttributePath, t.String()),
		))
		return "", false
	}

	s := req.AttributeConfig.(types.String)

	if s.Unknown || s.Null {
		return "", false
	}

	return s.Value, true
}

// validateList ensures that the request contains a List value.
// https://github.com/hashicorp/terraform-plugin-framework-validators/blob/4a4c520b56aa33e071c4356e0f3ac0ad2633dc7b/listvalidator/type_validation.go#L12
func validateList(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) ([]attr.Value, bool) {

	var l types.List

	diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &l)

	if diags.HasError() {
		response.Diagnostics.Append(diags...)

		return nil, false
	}

	if l.Unknown || l.Null {
		return nil, false
	}

	return l.Elems, true
}
