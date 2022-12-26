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
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "replication_rule_ids.*", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "snapshot_rule_ids.*", "153df6eb-3433-4b5e-942e-ecf90348df20"),
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
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "description", "Test CreateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "replication_rule_ids.*", "5d45b173-9a85-473e-8ab8-e107f8b8085e"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "snapshot_rule_ids.*", "153df6eb-3433-4b5e-942e-ecf90348df20"),
				),
			},
			{
				Config: ProtectionPolicyParamsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "name", "protectionpolicy_acc_new"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "description", "Test UpdateProtectionPolicy"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "replication_rule_ids.*", "5520daa7-aedb-4966-93c5-f0ae82b040ee"),
					resource.TestCheckResourceAttr("powerstore_storagecontainer.test", "snapshot_rule_ids.*", "4be81573-c0e6-4956-a32f-a0e396a9b86d"),
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

// Test to Create ProtectionPolicy with mutually exclusive params available
func TestAccProtectionPolicy_CreateWithMutuallyExclusiveParams(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	tests := []resource.TestStep{
		{
			Config:      ProtectionPolicyParamsWithSnapshotIDAndName,
			ExpectError: regexp.MustCompile("Either of Snapshot Rule ID or Snapshot Rule Name should be present"),
		},
		{
			Config:      ProtectionPolicyParamsWithReplicationIDAndName,
			ExpectError: regexp.MustCompile("Either of Replication Rule ID or Replication Rule Name should be present"),
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
	snapshot_rule_ids = ["153df6eb-3433-4b5e-942e-ecf90348df20"]
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
	replication_rule_ids = ["5520daa7-aedb-4966-93c5-f0ae82b040ee"]
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
resource "powerstore_storagecontainer" "test" {
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
resource "powerstore_storagecontainer" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	snapshot_rule_names = ["test_snapshotrule_1"]
	snapshot_rule_ids = ["26837335-f92e-4da8-8ea7-b6e1aab7c7cf"]
}
`
var ProtectionPolicyParamsWithReplicationIDAndName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_storagecontainer" "test" {
	name = "protectionpolicy_acc_new"
	description = "Test UpdateProtectionPolicy"
	replication_rule_names = ["rr-llau-csi-test2-RT-D8337-Five_Minutes"]
	replication_rule_ids = ["ac0a3611-bea0-4298-89ad-536e5a8788bd"]
}
`
