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
	"fmt"
	"os"
	"regexp"
	"terraform-provider-powerstore/client"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Unit test HCL config for secure volume group snapshot tests

var secureVGSnapMockedConfig = `
resource "powerstore_volumegroup_snapshot" "secure_test" {
  name                 = "tf_secure_vg_snap_test"
  description          = "Test Secure VG Snapshot"
  volume_group_id      = "vg-parent-001"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

// U-014: TestSecureVGSnapshot_Create - Mocked, is_secure=true
func TestSecureVGSnapshot_Create(t *testing.T) {
	var createMock, getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	defer func() {
		if createMock != nil {
			createMock.UnPatch()
		}
		if getMock != nil {
			getMock.UnPatch()
		}
		if deleteMock != nil {
			deleteMock.UnPatch()
		}
		if newClientMock != nil {
			newClientMock.UnPatch()
		}
		if getVolumesMock != nil {
			getVolumesMock.UnPatch()
		}
	}()

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateVolumeGroupSnapshot).
						Return(gopowerstore.CreateResponse{ID: "vg-snap-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumeGroupSnapshot).
						Return(gopowerstore.VolumeGroup{
							ID:          "vg-snap-001",
							Name:        "tf_secure_vg_snap_test",
							Description: "Test Secure VG Snapshot",
							ProtectionData: gopowerstore.ProtectionData{
								ExpirationTimeStamp: "2035-05-06T09:01:47+00:00",
								ParentID:            "vg-parent-001",
							},
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteVolumeGroup).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureVGSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "name", "tf_secure_vg_snap_test"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "description", "Test Secure VG Snapshot"),
				),
			},
		},
	})
}

// U-015: TestSecureVGSnapshot_ReadState - Mocked read with IsSecure=true
func TestSecureVGSnapshot_ReadState(t *testing.T) {
	var createMock, getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	defer func() {
		if createMock != nil {
			createMock.UnPatch()
		}
		if getMock != nil {
			getMock.UnPatch()
		}
		if deleteMock != nil {
			deleteMock.UnPatch()
		}
		if newClientMock != nil {
			newClientMock.UnPatch()
		}
		if getVolumesMock != nil {
			getVolumesMock.UnPatch()
		}
	}()

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateVolumeGroupSnapshot).
						Return(gopowerstore.CreateResponse{ID: "vg-snap-002"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumeGroupSnapshot).
						Return(gopowerstore.VolumeGroup{
							ID:          "vg-snap-002",
							Name:        "tf_secure_vg_snap_test",
							Description: "Test Secure VG Snapshot",
							ProtectionData: gopowerstore.ProtectionData{
								ExpirationTimeStamp: "2035-05-06T09:01:47+00:00",
								ParentID:            "vg-parent-001",
							},
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteVolumeGroup).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureVGSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "id", "vg-snap-002"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "name", "tf_secure_vg_snap_test"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "description", "Test Secure VG Snapshot"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "expiration_timestamp", "2035-05-06T09:01:47Z"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_test", "volume_group_id", "vg-parent-001"),
				),
			},
		},
	})
}

// U-016: TestSecureVGSnapshot_CreateError - Mocked create returning error
func TestSecureVGSnapshot_CreateError(t *testing.T) {
	var createMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	defer func() {
		if createMock != nil {
			createMock.UnPatch()
		}
		if deleteMock != nil {
			deleteMock.UnPatch()
		}
		if newClientMock != nil {
			newClientMock.UnPatch()
		}
		if getVolumesMock != nil {
			getVolumesMock.UnPatch()
		}
	}()

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateVolumeGroupSnapshot).
						Return(gopowerstore.CreateResponse{}, fmt.Errorf("mock error: failed to create volume group snapshot")).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteVolumeGroup).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config:      ProviderConfigForTesting + secureVGSnapMockedConfig,
				ExpectError: regexp.MustCompile(".*Error creating volume group snapshot.*"),
			},
		},
	})
}

// I-003: TestAccSecureVGSnapshot_Create - TF_ACC gated
func TestAccSecureVGSnapshot_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureVGSnapAccConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_acc_test", "name", "tf_secure_vg_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_acc_test", "description", "Acceptance Test Secure VG Snapshot"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.secure_acc_test", "is_secure", "true"),
				),
			},
			// Import Testing
			{
				Config:       ProviderConfigForTesting + SecureVGSnapAccConfig,
				ResourceName: "powerstore_volumegroup_snapshot.secure_acc_test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_secure_vg_snap_acc", s[0].Attributes["name"])
					return nil
				},
			},
		},
	})
}

// Acceptance test HCL config for secure volume group snapshot
var SecureVGSnapAccConfig = VolumeGroupParamsWithVolumeName + `
resource "powerstore_volumegroup_snapshot" "secure_acc_test" {
  name                 = "tf_secure_vg_snap_acc"
  description          = "Acceptance Test Secure VG Snapshot"
  volume_group_id      = powerstore_volumegroup.test.id
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
  depends_on           = [powerstore_volumegroup.test]
}
`
