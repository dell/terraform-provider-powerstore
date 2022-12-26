package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", "4be81573-c0e6-4956-a32f-a0e396a9b86d"),
				),
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
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", "4be81573-c0e6-4956-a32f-a0e396a9b86d"),
				),
			},
			{
				Config: ProtectionPolicyParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "description", "Test UpdateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", "4be81573-c0e6-4956-a32f-a0e396a9b86d"),
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
				ExpectError: regexp.MustCompile("Could not create protection policy"),
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
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", "4be81573-c0e6-4956-a32f-a0e396a9b86d"),
				),
			},
			{
				Config:      ProtectionPoliycParamsCreateServerError,
				ExpectError: regexp.MustCompile("Could not update"),
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
			ExpectError: regexp.MustCompile("either of snapshot rule id or snapshot rule name should be present"),
		},
		{
			Config:      ProtectionPolicyParamsWithReplicationIDAndName,
			ExpectError: regexp.MustCompile("either of replication rule id or replication rule name should be present"),
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
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_names.0", "test_snapshotrule_1"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_ids.0", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
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
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "snapshot_rule_ids.0", "4be81573-c0e6-4956-a32f-a0e396a9b86d"),
					resource.TestCheckResourceAttr("powerstore_protectionpolicy.test", "replication_rule_names.0", "Emalee-SRA-7416-Rep"),
				),
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
	replication_rule_ids = ["5d45b173-9a85-473e-8ab8-e107f8b8085e"]
	snapshot_rule_ids = ["4be81573-c0e6-4956-a32f-a0e396a9b86d"]
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
	replication_rule_ids = ["5d45b173-9a85-473e-8ab8-e107f8b8085e"]
	snapshot_rule_ids = ["4be81573-c0e6-4956-a32f-a0e396a9b86d"]
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
	snapshot_rule_names = ["test_snapshotrule_1"]
	snapshot_rule_ids = ["4be81573-c0e6-4956-a32f-a0e396a9b86d"]
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
	replication_rule_names = ["Emalee-SRA-7416-Rep"]
	replication_rule_ids = ["5d45b173-9a85-473e-8ab8-e107f8b8085e"]
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
	snapshot_rule_names = ["test_snapshotrule_1"]
	replication_rule_ids = ["5d45b173-9a85-473e-8ab8-e107f8b8085e"]
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
	snapshot_rule_ids = ["4be81573-c0e6-4956-a32f-a0e396a9b86d"]
	replication_rule_names = ["Emalee-SRA-7416-Rep"]
}
`
