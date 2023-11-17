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
				Config: ProtectionPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", replicationRuleID),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", snapshotRuleID),
				),
			},
			{
				Config:       ProtectionPolicyParamsCreate,
				ResourceName: "powerstore_protectionpolicy.test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "protectionpolicy_acc_new", s[0].Attributes["name"])
					assert.Equal(t, "Test CreateProtectionPolicy", s[0].Attributes["description"])
					assert.Equal(t, "Protection", s[0].Attributes["type"])
					assert.Equal(t, replicationRuleID, s[0].Attributes["replication_rule_ids.0"])
					assert.Equal(t, snapshotRuleID, s[0].Attributes["snapshot_rule_ids.0"])
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
				Config: ProtectionPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", replicationRuleID),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", snapshotRuleID),
				),
			},
			{
				Config: ProtectionPolicyParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test UpdateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", replicationRuleID),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", snapshotRuleID),
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
				Config:      ProtectionPoliycParamsCreateServerError,
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
				Config: ProtectionPolicyParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", replicationRuleID),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", snapshotRuleID),
				),
			},
			{
				Config:      ProtectionPoliycParamsCreateServerError,
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
			Config:      ProtectionPolicyParamsWithSnapshotIDAndName,
			ExpectError: regexp.MustCompile(SnapshotIDSnapshotNameErroMsg),
		},
		{
			Config:      ProtectionPolicyParamsWithReplicationIDAndName,
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
				Config: ProtectionPolicyParamsWithSnapshotRuleName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_names.0", snapshotRuleName),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", replicationRuleID),
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
				Config: ProtectionPolicyParamsWithReplicationRuleName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", snapshotRuleID),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_names.0", replicationRuleName),
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
				Config:        ProtectionPolicyParamsCreate,
				ResourceName:  "powerstore_protectionpolicy.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportPPDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var ProtectionPolicyParamsCreate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test CreateProtectionPolicy"
	replication_rule_ids = ["` + replicationRuleID + `"]
	snapshot_rule_ids = ["` + snapshotRuleID + `"]
}
`

var ProtectionPolicyParamsUpdate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	replication_rule_ids = ["` + replicationRuleID + `"]
	snapshot_rule_ids = ["` + snapshotRuleID + `"]
}
`

var ProtectionPoliycParamsCreateServerError = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
}
`

var ProtectionPolicyParamsWithSnapshotIDAndName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	snapshot_rule_names = ["` + snapshotRuleName + `"]
	snapshot_rule_ids = ["` + snapshotRuleID + `"]
}
`

var ProtectionPolicyParamsWithReplicationIDAndName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	replication_rule_names = ["` + replicationRuleName + `"]
	replication_rule_ids = ["` + replicationRuleID + `"]
}
`

var ProtectionPolicyParamsWithSnapshotRuleName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test CreateProtectionPolicy"
	snapshot_rule_names = ["` + snapshotRuleName + `"]
	replication_rule_ids = ["` + replicationRuleID + `"]
}
`

var ProtectionPolicyParamsWithReplicationRuleName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_protectionpolicy" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test CreateProtectionPolicy"
	snapshot_rule_ids = ["` + snapshotRuleID + `"]
	replication_rule_names = ["` + replicationRuleName + `"]
}
`
