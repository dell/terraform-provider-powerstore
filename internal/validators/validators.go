package validators

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure UrlString satsisfies validator.String
var _ validator.String = UrlString{}

type UrlString struct{}

func (u UrlString) Description(context.Context) string {
	return "string must be valid uri"
}

func (u UrlString) MarkdownDescription(context.Context) string {
	return "string must be valid uri"
}

func (u UrlString) ValidateString(ctx context.Context, req validator.StringRequest, res *validator.StringResponse) {
	log.Printf("mayank %+v", req)
	fmt.Println(req)
}
