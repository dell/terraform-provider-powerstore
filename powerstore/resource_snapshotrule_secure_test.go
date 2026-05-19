/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test U-26: Create secure snapshot rule
func TestAccSnapshotRule_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "name", "ZZZ_AT_tf_secure_snapshotrule"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "interval", "Four_Hours"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "desired_retention", "24"),
				),
			},
		},
	})
}

// Test U-27: Create non-secure snapshot rule - is_secure defaults to false
func TestAccSnapshotRule_CreateNonSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + NonSecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_nonsecure", "name", "ZZZ_AT_tf_nonsecure_snapshotrule"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_nonsecure", "is_secure", "false"),
				),
			},
		},
	})
}

// Test U-28: Update snapshot rule to secure
func TestAccSnapshotRule_UpdateToSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + NonSecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_nonsecure", "is_secure", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + SnapshotRuleParamsUpdateToSecure,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_nonsecure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test U-29: Reject unsecure snapshot rule (one-way lock)
func TestAccSnapshotRule_RejectUnsecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:      ProviderConfigForTesting + SnapshotRuleParamsRevertSecure,
				ExpectError: regexp.MustCompile(".*[Oo]ne.way lock.*|.*cannot.*unlock.*|.*cannot.*change.*true.*false.*"),
			},
		},
	})
}

// Test U-30: Import secure snapshot rule
func TestAccSnapshotRule_ImportSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:       ProviderConfigForTesting + SecureSnapshotRuleParams,
				ResourceName: "powerstore_snapshotrule.test_secure",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "ZZZ_AT_tf_secure_snapshotrule", s[0].Attributes["name"])
					assert.Equal(t, "true", s[0].Attributes["is_secure"])
					return nil
				},
			},
		},
	})
}

// Test I-09: Idempotent re-apply of secure snapshot rule
func TestAccSnapshotRule_SecureIdempotent(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "is_secure", "true"),
				),
			},
			// Re-apply same config - should produce no changes
			{
				Config: ProviderConfigForTesting + SecureSnapshotRuleParams,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_secure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test E-10: Explicit is_secure=false for snapshot rule
func TestAccSnapshotRule_ExplicitFalse(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
resource "powerstore_snapshotrule" "test_explicit_false" {
  name              = "ZZZ_AT_tf_explicit_false_rule"
  interval          = "Four_Hours"
  desired_retention = 24
  is_secure         = false
  delete_snaps      = true
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.test_explicit_false", "is_secure", "false"),
				),
			},
		},
	})
}

// --- HCL Test Configs ---

var SecureSnapshotRuleParams = `
resource "powerstore_snapshotrule" "test_secure" {
  name              = "ZZZ_AT_tf_secure_snapshotrule"
  interval          = "Four_Hours"
  desired_retention = 24
  is_secure         = true
  delete_snaps      = true
}
`

var NonSecureSnapshotRuleParams = `
resource "powerstore_snapshotrule" "test_nonsecure" {
  name              = "ZZZ_AT_tf_nonsecure_snapshotrule"
  interval          = "Four_Hours"
  desired_retention = 24
  delete_snaps      = true
}
`

var SnapshotRuleParamsUpdateToSecure = `
resource "powerstore_snapshotrule" "test_nonsecure" {
  name              = "ZZZ_AT_tf_nonsecure_snapshotrule"
  interval          = "Four_Hours"
  desired_retention = 24
  is_secure         = true
  delete_snaps      = true
}
`

var SnapshotRuleParamsRevertSecure = `
resource "powerstore_snapshotrule" "test_secure" {
  name              = "ZZZ_AT_tf_secure_snapshotrule"
  interval          = "Four_Hours"
  desired_retention = 24
  is_secure         = false
  delete_snaps      = true
}
`
