package validators

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure UrlString satsisfies validator.String
var _ validator.String = URLString{}

// URLString validates if given string is url
type URLString struct{}

// Description satisfies validator.String interface
func (u URLString) Description(context.Context) string {
	return "string must be valid uri"
}

// MarkdownDescription satisfies validator.String interface
func (u URLString) MarkdownDescription(context.Context) string {
	return "string must be valid uri"
}

// ValidateString validates if string is url
func (u URLString) ValidateString(ctx context.Context, req validator.StringRequest, res *validator.StringResponse) {

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
