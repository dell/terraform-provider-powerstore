package validators

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure UrlString satsisfies validator.String
var _ validator.String = UrlString{}

// UrlString validates if given string is url
type UrlString struct{}

// Description satisfies validator.String interface
func (u UrlString) Description(context.Context) string {
	return "string must be valid uri"
}

// MarkdownDescription satisfies validator.String interface
func (u UrlString) MarkdownDescription(context.Context) string {
	return "string must be valid uri"
}

// ValidateString validates if string is url
func (u UrlString) ValidateString(ctx context.Context, req validator.StringRequest, res *validator.StringResponse) {

	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	_, err := url.ParseRequestURI(req.ConfigValue.ValueString())
	if err != nil {
		res.Diagnostics.AddError(
			fmt.Sprintf("%s: invalid url", req.PathExpression),
			fmt.Sprintf("%s: invalid uri, %s", req.PathExpression, err.Error()),
		)
	}
}
