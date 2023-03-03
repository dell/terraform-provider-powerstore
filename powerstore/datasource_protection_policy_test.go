package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Volume
func TestAccProtectionPolicy_FetchPolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProtectionPolicyDataSourceparamsName,
			},
			{
				Config: ProtectionPolicyDataSourceparamsID,
			},
			{
				Config: ProtectionPolicyDataSourceparamsAll,
			},
			{
				Config:      ProtectionPolicyDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("nable to Read PowerStore Protection"),
			},
		},
	})
}

var ProtectionPolicyDataSourceparamsName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_protectionpolicy" "test1" {
	name = "` + policyName + `"
}
`
var ProtectionPolicyDataSourceparamsNameNegative = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_protectionpolicy" "test1" {
	name = "invalid-name"
}
`

var ProtectionPolicyDataSourceparamsID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_protectionpolicy" "test1" {
	id = "` + policyID + `"
}
`

var ProtectionPolicyDataSourceparamsAll = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_protectionpolicy" "test1" {
}
`
