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
				Config: ProviderConfigForTesting + VolumeSnapshotDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeSnapshotDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeSnapshotDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsIDAndName,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsEmptyID,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsEmptyName,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Snapshots"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeSnapshotDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Snapshots"),
			},
		},
	})
}

var VolumeSnapshotDataSourceparamsID = `
data "powerstore_volume_snapshot" "test1" {
	id = "` + volumeSnapshotID + `"
}
`

var VolumeSnapshotDataSourceparamsIDNegative = `
data "powerstore_volume_snapshot" "test1" {
	id = "invalid-id"
}
`

var VolumeSnapshotDataSourceparamsEmptyID = `
data "powerstore_volume_snapshot" "test1" {
	id = ""
}
`

var VolumeSnapshotDataSourceparamsName = `
data "powerstore_volume_snapshot" "test1" {
	name = "` + volumeSnapshotName + `"
}
`

var VolumeSnapshotDataSourceparamsNameNegative = `
data "powerstore_volume_snapshot" "test1" {
	name = "invalid-name"
}
`

var VolumeSnapshotDataSourceparamsEmptyName = `
data "powerstore_volume_snapshot" "test1" {
	name = ""
}
`

var VolumeSnapshotDataSourceparamsAll = `
data "powerstore_volume_snapshot" "test1" {
}
`

var VolumeSnapshotDataSourceparamsIDAndName = `
data "powerstore_volume_snapshot" "test1" {
	id = "` + volumeSnapshotID + `"
	name = "` + volumeSnapshotName + `"
}
`
