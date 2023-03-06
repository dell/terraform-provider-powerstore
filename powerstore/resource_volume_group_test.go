package powerstore

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create VolumeGroup
func TestAccVolumeGroup_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: VolumeGroupParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "name", "test_volume_group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "description", "Creating Volume Group"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "is_write_order_consistent", "false"),
					resource.TestCheckResourceAttr("powerstore_volumegroup.test", "protection_policy_id", protectionPolicyID),
				),
			},
		},
	})
}

var VolumeGroupParamsCreate = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}
resource "powerstore_volumegroup" "test" {
  name = "test_volume_group"
  description = "Creating Volume Group"
  is_write_order_consistent = false
  protection_policy_id = "` + protectionPolicyID + `"
}
`
