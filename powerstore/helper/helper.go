package helper

import "github.com/hashicorp/terraform-plugin-framework/types"

func GetKnownBoolPointer(in types.Bool) *bool {
	if in.IsUnknown() {
		return nil
	}
	return in.ValueBoolPointer()
}
