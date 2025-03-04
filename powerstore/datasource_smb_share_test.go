/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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

// Test to Fetch SMB Shares
func TestAccSMBShareDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	// TODO: Use resource as preq when it gets implemented
	dsPreq := `
	data powerstore_smb_share preq {
		lifecycle {
			postcondition {
				condition = length(self.smb_shares) > 0
				error_message = "For SMB Share ds test, array should have at least one SMB share."
			}
		}
	}
	`
	preqItem := "data.powerstore_smb_share.preq.smb_shares[0]"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				// Get all SMB Shares
				Config: ProviderConfigForTesting + `
				data "powerstore_smb_share" "test" {
				}
				`,
			},
			{
				// Get SMB Share by Invalid ID
				Config: ProviderConfigForTesting + `
				data "powerstore_smb_share" "test" {
					id = "invalid"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read SMB Share with id.*"),
			},
			{
				// Get SMB Share by Invalid Name
				Config: ProviderConfigForTesting + `
				data "powerstore_smb_share" "test" {
					name = "invalid"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read SMB Share with name.*"),
			},
			{
				// Get SMB Share by invalid filter expression
				Config: ProviderConfigForTesting + `
				data "powerstore_smb_share" "test" {
					filter_expression = "name=inv.terraform"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read SMB Shares with error.*"),
			},
			// TODO: Use resource as prereq when it gets implemented
			{
				//Get SMB Share by ID
				Config: ProviderConfigForTesting + dsPreq + `
				data "powerstore_smb_share" "test" {
					id = ` + preqItem + `.id
				}
				`,
			},
			{
				// Get SMB Share by name
				Config: ProviderConfigForTesting + dsPreq + `
				data "powerstore_smb_share" "test" {
					name = ` + preqItem + `.name
				}
				`,
			},
			{
				// Get SMB Share by filesystem id
				Config: ProviderConfigForTesting + dsPreq + `
				data "powerstore_smb_share" "test" {
					file_system_id = ` + preqItem + `.file_system_id
				}
				`,
			},
		},
	})
}
