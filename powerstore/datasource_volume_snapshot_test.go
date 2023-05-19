package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Volume Snapshot
func TestAccVolume_FetchVolumeSnapshot(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: VolumeSnapshotDataSourceparamsID,
			},
			{
				Config: VolumeSnapshotDataSourceparamsName,
			},
			{
				Config: VolumeSnapshotDataSourceparamsAll,
			},
			{
				Config:      VolumeSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Snapshots"),
			},
			{
				Config:      VolumeSnapshotDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Snapshots"),
			},
		},
	})
}

var VolumeSnapshotDataSourceparamsID = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume_snapshot" "test1" {
	id = "` + volumeSnapshotID + `"
}
`

var VolumeSnapshotDataSourceparamsIDNegative = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume_snapshot" "test1" {
	id = "invalid-id"
}
`

var VolumeSnapshotDataSourceparamsName = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume_snapshot" "test1" {
	name = "` + volumeSnapshotName + `"
}
`
var VolumeSnapshotDataSourceparamsNameNegative = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume_snapshot" "test1" {
	name = "invalid-name"
}
`

var VolumeSnapshotDataSourceparamsAll = `
provider "powerstore" {
	username = "` + username + `"
	password = "` + password + `"
	endpoint = "` + endpoint + `"
	insecure = true
}

data "powerstore_volume_snapshot" "test1" {
}
`
