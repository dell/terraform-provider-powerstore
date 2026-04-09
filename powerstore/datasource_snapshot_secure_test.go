/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"terraform-provider-powerstore/client"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// HCL configs for secure snapshot datasource tests

var secureVolSnapDatasourceConfig = `
data "powerstore_volume_snapshot" "secure_test" {
  name = "tf_secure_snap_test"
}
`

var secureFSSnapDatasourceConfig = `
data "powerstore_filesystem_snapshot" "secure_test" {
  name = "tf_secure_fs_snap_test"
}
`

var secureRuleDatasourceConfig = `
data "powerstore_snapshotrule" "secure_test" {
  name = "tf_secure_rule_test"
}
`

// U-021: TestSecureVolSnapshotDatasource_Read verifies that the is_secure attribute
// is populated in the volume snapshot datasource when ProtectionData.IsSecure is true.
func TestSecureVolSnapshotDatasource_Read(t *testing.T) {
	var newClientMocker *mockey.Mocker
	var getVolumesMocker *mockey.Mocker
	var getSnapByNameMocker *mockey.Mocker
	var getHostMappingMocker *mockey.Mocker
	var getVolGroupMocker *mockey.Mocker

	defer func() {
		if newClientMocker != nil {
			newClientMocker.UnPatch()
		}
		if getVolumesMocker != nil {
			getVolumesMocker.UnPatch()
		}
		if getSnapByNameMocker != nil {
			getSnapByNameMocker.UnPatch()
		}
		if getHostMappingMocker != nil {
			getHostMappingMocker.UnPatch()
		}
		if getVolGroupMocker != nil {
			getVolGroupMocker.UnPatch()
		}
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
			FunctionMocker = nil
		}
	}()

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
						FunctionMocker = nil
					}
					newClientMocker = mockey.Mock(client.NewClient).Return(
						&client.Client{
							PStoreClient: &gopowerstore.ClientIMPL{},
						}, nil,
					).Build()
					getVolumesMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return(
						[]gopowerstore.Volume{}, nil,
					).Build()
					getSnapByNameMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotByName).Return(
						gopowerstore.Volume{
							ID:   "snap-secure-001",
							Name: "tf_secure_snap_test",
							Size: 1048576,
							ProtectionData: gopowerstore.ProtectionData{
								SourceID:    "vol-parent-001",
								CreatorType: "User",
							},
						}, nil,
					).Build()
					getHostMappingMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetHostVolumeMappingByVolumeID).Return(
						[]gopowerstore.HostVolumeMapping{}, nil,
					).Build()
					getVolGroupMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumeGroupsByVolumeID).Return(
						gopowerstore.VolumeGroups{}, nil,
					).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureVolSnapDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_volume_snapshot.secure_test", "volumes.0.protection_data.is_secure", "true"),
				),
			},
		},
	})
}

// U-022: TestSecureFSSnapshotDatasource_Read verifies that the is_secure attribute
// is populated in the filesystem snapshot datasource when IsSecure is true.
func TestSecureFSSnapshotDatasource_Read(t *testing.T) {
	var newClientMocker *mockey.Mocker
	var getVolumesMocker *mockey.Mocker
	var getFsByFilterMocker *mockey.Mocker

	defer func() {
		if newClientMocker != nil {
			newClientMocker.UnPatch()
		}
		if getVolumesMocker != nil {
			getVolumesMocker.UnPatch()
		}
		if getFsByFilterMocker != nil {
			getFsByFilterMocker.UnPatch()
		}
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
			FunctionMocker = nil
		}
	}()

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
						FunctionMocker = nil
					}
					newClientMocker = mockey.Mock(client.NewClient).Return(
						&client.Client{
							PStoreClient: &gopowerstore.ClientIMPL{},
						}, nil,
					).Build()
					getVolumesMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return(
						[]gopowerstore.Volume{}, nil,
					).Build()
					getFsByFilterMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetFsByFilter).Return(
						[]gopowerstore.FileSystem{
							{
								ID:             "fs-snap-secure-001",
								Name:           "tf_secure_fs_snap_test",
								FilesystemType: gopowerstore.FileSystemTypeEnumSnapshot,
							},
						}, nil,
					).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureFSSnapDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_filesystem_snapshot.secure_test", "filesystem_snapshots.0.is_secure", "true"),
				),
			},
		},
	})
}

// U-023: TestSecureSnapshotRuleDatasource_Read verifies that the is_secure attribute
// is populated in the snapshot rule datasource when IsSecure is true.
func TestSecureSnapshotRuleDatasource_Read(t *testing.T) {
	var newClientMocker *mockey.Mocker
	var getVolumesMocker *mockey.Mocker
	var getSnapRuleByNameMocker *mockey.Mocker

	defer func() {
		if newClientMocker != nil {
			newClientMocker.UnPatch()
		}
		if getVolumesMocker != nil {
			getVolumesMocker.UnPatch()
		}
		if getSnapRuleByNameMocker != nil {
			getSnapRuleByNameMocker.UnPatch()
		}
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
			FunctionMocker = nil
		}
	}()

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
						FunctionMocker = nil
					}
					newClientMocker = mockey.Mock(client.NewClient).Return(
						&client.Client{
							PStoreClient: &gopowerstore.ClientIMPL{},
						}, nil,
					).Build()
					getVolumesMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return(
						[]gopowerstore.Volume{}, nil,
					).Build()
					getSnapRuleByNameMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotRuleByName).Return(
						gopowerstore.SnapshotRule{
							ID:               "rule-secure-001",
							Name:             "tf_secure_rule_test",
							DesiredRetention: 24,
							IsReplica:        false,
							IsReadOnly:       false,
						}, nil,
					).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureRuleDatasourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.powerstore_snapshotrule.secure_test", "snapshot_rules.0.is_secure", "true"),
				),
			},
		},
	})
}
