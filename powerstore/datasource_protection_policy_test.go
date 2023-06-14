/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Protection Policy
func TestAccProtectionPolicy_FetchPolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ProtectionPolicyDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + ProtectionPolicyDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + ProtectionPolicyDataSourceparamsAll,
			},
		},
	})
}

// Test to Fetch Protection Policy- Negative
func TestAccProtectionPolicy_FetchPolicyNegative(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + ProtectionPolicyDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Protection"),
			},
			{
				Config:      ProviderConfigForTesting + ProtectionPolicyDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Protection"),
			},
			{
				Config:      ProviderConfigForTesting + ProtectionPolicyDataSourceparamsEmptyID,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + ProtectionPolicyDataSourceparamsEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
		},
	})
}

var ProtectionPolicyDataSourceparamsName = `
data "powerstore_protectionpolicy" "test1" {
	name = "` + policyName + `"
}
`
var ProtectionPolicyDataSourceparamsNameNegative = `
data "powerstore_protectionpolicy" "test1" {
	name = "invalid-name"
}
`

var ProtectionPolicyDataSourceparamsIDNegative = `
data "powerstore_protectionpolicy" "test1" {
	id = "invalid-id"
}
`

var ProtectionPolicyDataSourceparamsID = `
data "powerstore_protectionpolicy" "test1" {
	id = "` + policyID + `"
}
`

var ProtectionPolicyDataSourceparamsEmptyName = `
data "powerstore_protectionpolicy" "test1" {
	name = ""
}
`

var ProtectionPolicyDataSourceparamsEmptyID = `
data "powerstore_protectionpolicy" "test1" {
	id = ""
}
`

var ProtectionPolicyDataSourceparamsAll = `
data "powerstore_protectionpolicy" "test1" {
}
`
