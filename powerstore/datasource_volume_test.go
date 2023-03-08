package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Volume
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
			{
				Config:      VolumeDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volumes"),
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
	name = "` + volumeName + `"
}
`
var VolumeDataSourceparamsNameNegative = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume" "test1" {
	name = "invalid-name"
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
	id = "` + volumeID + `"
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
