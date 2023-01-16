package powerstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

// Test to Create Volume
func TestAccVolume_FetchVolume(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: VolumeDataSourceparamsName,
			},
			{
				Config: VolumeDataSourceparamsID,
			},
			{
				Config: VolumeDataSourceparamsAll,
			},
		},
	})
}

var VolumeDataSourceparamsName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume" "test1" {
	name = "tf_vol"
}
`

var VolumeDataSourceparamsID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume" "test1" {
	id = "a0b0c773-1c50-425a-89dc-aef9162ec787"
}
`

var VolumeDataSourceparamsAll = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume" "test1" {
}
`
