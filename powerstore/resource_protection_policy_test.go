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

// Test to Create ProtectionPolicy
func TestAccProtectionPolicy_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ProtectionPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
				),
			},
			{
				Config:       ProviderConfigForTesting + ProtectionPolicyParamsCreate,
				ResourceName: "powerstore_protectionpolicy.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "protectionpolicy_acc_new", s[0].Attributes["name"])
					assert.Equal(t, "Test CreateProtectionPolicy", s[0].Attributes["description"])
					assert.Equal(t, "Protection", s[0].Attributes["type"])
					return nil
				},
			},
		},
	})
}

// Test to Update existing ProtectionPolicy Params
func TestAccProtectionPolicy_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ProtectionPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
				),
			},
			{
				Config: ProviderConfigForTesting + ProtectionPolicyParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test UpdateProtectionPolicy"),
				),
			},
		},
	})
}

// Test to Create ProtectionPolicy with Invalid values, will result in error
func TestAccProtectionPolicy_CreateWithInvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + ProtectionPoliycParamsCreateServerError,
				ExpectError: regexp.MustCompile(CreatePPDetailErrorMsg),
			},
		},
	})
}

// Test to update existing ProtectionPolicy params but will result in error
func TestAccProtectionPolicy_UpdateError(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ProtectionPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
				),
			},
			{
				Config:      ProviderConfigForTesting + ProtectionPoliycParamsCreateServerError,
				ExpectError: regexp.MustCompile(UpdatePPDetailErrorMsg),
			},
		},
	})
}

// Test to Create ProtectionPolicy with mutually exclusive params available
func TestAccProtectionPolicy_CreateWithMutuallyExclusiveParams(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	tests := []resource.TestStep{
		{
			Config:      ProviderConfigForTesting + ProtectionPolicyParamsWithSnapshotIDAndName,
			ExpectError: regexp.MustCompile(SnapshotIDSnapshotNameErroMsg),
		},
		{
			Config:      ProviderConfigForTesting + ProtectionPolicyParamsWithReplicationIDAndName,
			ExpectError: regexp.MustCompile(ReplicationIDReplicationNameErrorMsg),
		},
	}

	for i := range tests {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testProviderFactory,
			Steps:                    []resource.TestStep{tests[i]},
		})
	}
}

// Test to Create ProtectionPolicy with SnapshotRule Name
func TestAccProtectionPolicy_CreateWithSnapshotRuleName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ProtectionPolicyParamsWithSnapshotRuleName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_names.0", "tf_snapshotrule"),
				),
			},
		},
	})
}

// Test to Create ProtectionPolicy with ReplicationRule Name
func TestAccProtectionPolicy_CreateWithReplicationRuleName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + ProtectionPolicyParamsWithReplicationRuleName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
				),
			},
		},
	})
}

// Negative - Test to import protection policy
func TestAccProtectionPolicy_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + ProtectionPolicyParamsCreate,
				ResourceName:  "powerstore_protectionpolicy.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportPPDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var ProtectionPolicyParamsCreate = SnapshotRuleParamsWithTimeOfDay + ReplicationRuleParamsCreate + `
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test CreateProtectionPolicy"
	replication_rule_ids = [powerstore_replication_rule.test.id]
	snapshot_rule_ids = [powerstore_snapshotrule.test.id]
}
`

var ProtectionPolicyParamsUpdate = SnapshotRuleParamsWithTimeOfDay + ReplicationRuleParamsCreate + `
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	replication_rule_ids = [powerstore_replication_rule.test.id]
	snapshot_rule_ids = [powerstore_snapshotrule.test.id]
}
`

var ProtectionPoliycParamsCreateServerError = SnapshotRuleParamsWithTimeOfDay + ReplicationRuleParamsCreate + `
resource "powerstore_protectionpolicy" "test" {
	depends_on = [powerstore_snapshotrule.test, powerstore_replication_rule.test]
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
}
`

var ProtectionPolicyParamsWithSnapshotIDAndName = SnapshotRuleParamsWithTimeOfDay + `
resource "powerstore_protectionpolicy" "test" {
	depends_on = [powerstore_snapshotrule.test]
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	snapshot_rule_names = [powerstore_snapshotrule.test.name]
	snapshot_rule_ids = [powerstore_snapshotrule.test.id]
}
`

var ProtectionPolicyParamsWithReplicationIDAndName = ReplicationRuleParamsCreate + `
resource "powerstore_protectionpolicy" "test" {
	depends_on = [powerstore_replication_rule.test]
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	replication_rule_names = [powerstore_replication_rule.test.name]
	replication_rule_ids = [powerstore_replication_rule.test.id]
}
`

var ProtectionPolicyParamsWithSnapshotRuleName = SnapshotRuleParamsWithTimeOfDay + ReplicationRuleParamsCreate + `
resource "powerstore_protectionpolicy" "test" {
	depends_on = [powerstore_snapshotrule.test]
	name = "protectionpolicy_acc_new"
	description = "Test CreateProtectionPolicy"
	snapshot_rule_names = [powerstore_snapshotrule.test.name]
	replication_rule_ids = [powerstore_replication_rule.test.id]
}
`

var ProtectionPolicyParamsWithReplicationRuleName = SnapshotRuleParamsWithTimeOfDay + ReplicationRuleParamsCreate + `
resource "powerstore_protectionpolicy" "test" {
	depends_on = [powerstore_replication_rule.test]
	name = "protectionpolicy_acc_new"
	description = "Test CreateProtectionPolicy"
	snapshot_rule_ids = [powerstore_snapshotrule.test.id]
	replication_rule_names = [powerstore_replication_rule.test.name]
}
`
