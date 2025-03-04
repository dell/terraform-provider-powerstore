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

// Test to Fetch Remote Systems
func TestAccRemoteSystemDs(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	// TODO: Use resource as preq when it gets implemented
	dsPreq := `
	data powerstore_remote_system preq {
		lifecycle {
			postcondition {
				condition = length(self.remote_systems) > 0
				error_message = "For Remote Systems ds test, array should have at least one Remote Systems."
			}
		}
	}
	`
	preqItem := "data.powerstore_remote_system.preq.remote_systems[0]"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				// Get all Remote Systems
				Config: ProviderConfigForTesting + `
				data "powerstore_remote_system" "test" {
				}
				`,
			},
			{
				// Get Remote Systems by Invalid ID
				Config: ProviderConfigForTesting + `
				data "powerstore_remote_system" "test" {
					id = "invalid"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read Remote System with id.*"),
			},
			{
				// Get Remote Systems by Invalid Name
				Config: ProviderConfigForTesting + `
				data "powerstore_remote_system" "test" {
					name = "invalid"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read Remote System with name.*"),
			},
			{
				// Get Remote Systems by invalid filter expression
				Config: ProviderConfigForTesting + `
				data "powerstore_remote_system" "test" {
					filter_expression = "name=inv.terraform"
				}
				`,
				ExpectError: regexp.MustCompile(".*Could not read Remote Systems with error.*"),
			},
			// TODO: Use resource as prereq when it gets implemented
			{
				//Get Remote Systems by ID
				Config: ProviderConfigForTesting + dsPreq + `
				data "powerstore_remote_system" "test" {
					id = ` + preqItem + `.id
				}
				`,
			},
			{
				// Get Remote Systems by name
				Config: ProviderConfigForTesting + dsPreq + `
				data "powerstore_remote_system" "test" {
					name = ` + preqItem + `.name
				}
				`,
			},
			{
				// Get Remote Systems by filter expression
				Config: ProviderConfigForTesting + dsPreq + `
				data "powerstore_remote_system" "test" {
					filter_expression = "management_address=eq.${` + preqItem + `.management_address}"
				}
				`,
			},
		},
	})
}
