package powerstore

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Fetch Volume Group snapshots
func TestAccVolumeGroupSnapshot_FetchVolumeGroupSnapshot(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsName,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsID,
			},
			{
				Config: ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsAll,
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsIDAndNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsEmptyIDNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsEmptyNameNegative,
				ExpectError: regexp.MustCompile("Invalid Attribute Value Length"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsNameNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
			{
				Config:      ProviderConfigForTesting + VolumeGroupSnapshotDataSourceparamsIDNegative,
				ExpectError: regexp.MustCompile("Unable to Read PowerStore Volume Group"),
			},
		},
	})
}

var VolumeGroupSnapshotDataSourceparamsName = `
data "powerstore_volumegroup_snapshot" "test1" {
	name = "` + volumeGroupSnapshotName + `"
}
`

var VolumeGroupSnapshotDataSourceparamsNameNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	name = "invalid-name"
}
`

var VolumeGroupSnapshotDataSourceparamsID = `
data "powerstore_volumegroup_snapshot" "test1" {
	id = "` + volumeGroupSnapshotID + `"
}
`

var VolumeGroupSnapshotDataSourceparamsAll = `
data "powerstore_volumegroup_snapshot" "test1" {
}
`

var VolumeGroupSnapshotDataSourceparamsIDNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	id = "invalid-id"
}
`

var VolumeGroupSnapshotDataSourceparamsIDAndNameNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	id = "` + volumeGroupSnapshotID + `"
	name = "` + volumeGroupSnapshotName + `"
}
`

var VolumeGroupSnapshotDataSourceparamsEmptyIDNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	id = ""
}
`

var VolumeGroupSnapshotDataSourceparamsEmptyNameNegative = `
data "powerstore_volumegroup_snapshot" "test1" {
	name = ""
}
`
