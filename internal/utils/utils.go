package utils

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// StringListToSetValues returns setValue by converting given string list appends prefix to every element of list
func StringListToSetValues(list []string, prefix string) (basetypes.SetValue, diag.Diagnostics) {

	setValues := []attr.Value{}

	uniqueElems := map[string]struct{}{}

	for _, str := range list {
		if _, ok := uniqueElems[str]; ok {
			continue
		}
		uniqueElems[str] = struct{}{}
		setValues = append(setValues, types.StringValue(fmt.Sprintf("%s%s", prefix, str)))
	}

	return types.SetValue(types.StringType, setValues)
}

// SetValuesToStringList returns string list by converting given setValues and appends prefix to every element of list
func SetValuesToStringList(setValues basetypes.SetValue, prefix string) (list []string) {

	for _, element := range setValues.Elements() {
		list = append(list, fmt.Sprintf("%s%s", prefix, strings.Trim(element.String(), "\"")))
	}

	return

}
