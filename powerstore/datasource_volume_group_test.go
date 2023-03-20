package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Volume Groups
func TestAccVolumeGroup_FetchVolumeGroup(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
		},
	})
}

var VolumeGroupDataSourceparamsName = `
data "powerstore_volumegroup" "test1" {
	name = "` + volumeGroupName + `"
}
`
var VolumeGroupDataSourceparamsNameNegative = `
data "powerstore_volumegroup" "test1" {
	name = "invalid-name"
}
`

var VolumeGroupDataSourceparamsID = `
data "powerstore_volumegroup" "test1" {
	id = "` + volumeGroupID + `"
}
`

var VolumeGroupDataSourceparamsAll = `
data "powerstore_volumegroup" "test1" {
}
`
