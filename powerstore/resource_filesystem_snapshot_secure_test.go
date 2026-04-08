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

// Unit test HCL configs for secure filesystem snapshot tests

var securefsSnapMockedConfig = `
resource "powerstore_filesystem_snapshot" "secure_test" {
  name                 = "tf_secure_fs_snap_test"
  description          = "Test Secure FS Snapshot"
  filesystem_id        = "fs-parent-001"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var securefsSnapMockedDefaultConfig = `
resource "powerstore_filesystem_snapshot" "secure_test" {
  name                 = "tf_secure_fs_snap_default"
  description          = "Test Default FS Snapshot"
  filesystem_id        = "fs-parent-001"
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var securefsSnapMockedUpdateConfig = `
resource "powerstore_filesystem_snapshot" "secure_test" {
  name                 = "tf_secure_fs_snap_test"
  description          = "Updated Description"
  filesystem_id        = "fs-parent-001"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

// U-009: TestSecureFSSnapshot_Create - Mocked, is_secure=true in config, check state
func TestSecureFSSnapshot_Create(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateFsSnapshot).
						Return(gopowerstore.CreateResponse{ID: "fs-snap-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetFS).
						Return(gopowerstore.FileSystem{
							ID:                  "fs-snap-001",
							Name:                "tf_secure_fs_snap_test",
							Description:         "Test Secure FS Snapshot",
							ExpirationTimestamp: "2035-05-06T09:01:47+00:00",
							ParentID:            "fs-parent-001",
							AccessType:          "Snapshot",
							IsSecure:            true,
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteFsSnapshot).
						Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + securefsSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "name", "tf_secure_fs_snap_test"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "description", "Test Secure FS Snapshot"),
				),
			},
		},
	})
}

// U-010: TestSecureFSSnapshot_CreateDefault - Mocked, no is_secure in config, check is_secure=false
func TestSecureFSSnapshot_CreateDefault(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateFsSnapshot).
						Return(gopowerstore.CreateResponse{ID: "fs-snap-002"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetFS).
						Return(gopowerstore.FileSystem{
							ID:                  "fs-snap-002",
							Name:                "tf_secure_fs_snap_default",
							Description:         "Test Default FS Snapshot",
							ExpirationTimestamp: "2035-05-06T09:01:47+00:00",
							ParentID:            "fs-parent-001",
							AccessType:          "Snapshot",
							IsSecure:            false,
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteFsSnapshot).
						Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + securefsSnapMockedDefaultConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "is_secure", "false"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "name", "tf_secure_fs_snap_default"),
				),
			},
		},
	})
}

// U-011: TestSecureFSSnapshot_ReadState - Mocked read with IsSecure=true, verify all state attributes
func TestSecureFSSnapshot_ReadState(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateFsSnapshot).
						Return(gopowerstore.CreateResponse{ID: "fs-snap-003"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetFS).
						Return(gopowerstore.FileSystem{
							ID:                  "fs-snap-003",
							Name:                "tf_secure_fs_snap_test",
							Description:         "Test Secure FS Snapshot",
							ExpirationTimestamp: "2035-05-06T09:01:47+00:00",
							ParentID:            "fs-parent-001",
							AccessType:          "Snapshot",
							IsSecure:            true,
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteFsSnapshot).
						Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + securefsSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "id", "fs-snap-003"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "name", "tf_secure_fs_snap_test"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "description", "Test Secure FS Snapshot"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "expiration_timestamp", "2035-05-06T09:01:47Z"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "filesystem_id", "fs-parent-001"),
				),
			},
		},
	})
}

// U-012: TestSecureFSSnapshot_UpdateDescription - Mocked update of description, is_secure preserved
func TestSecureFSSnapshot_UpdateDescription(t *testing.T) {
	var createMock, getMock, modifyMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	defer func() {
		if createMock != nil {
			createMock.UnPatch()
		}
		if getMock != nil {
			getMock.UnPatch()
		}
		if modifyMock != nil {
			modifyMock.UnPatch()
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateFsSnapshot).
						Return(gopowerstore.CreateResponse{ID: "fs-snap-004"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetFS).
						Return(gopowerstore.FileSystem{
							ID:                  "fs-snap-004",
							Name:                "tf_secure_fs_snap_test",
							Description:         "Test Secure FS Snapshot",
							ExpirationTimestamp: "2035-05-06T09:01:47+00:00",
							ParentID:            "fs-parent-001",
							AccessType:          "Snapshot",
							IsSecure:            true,
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteFsSnapshot).
						Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + securefsSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "description", "Test Secure FS Snapshot"),
				),
			},
			{
				PreConfig: func() {
					if createMock != nil {
						createMock.UnPatch()
						createMock = nil
					}
					if getMock != nil {
						getMock.UnPatch()
					}
					modifyMock = mockey.Mock((*gopowerstore.ClientIMPL).ModifyFS).
						Return(nil, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetFS).
						Return(gopowerstore.FileSystem{
							ID:                  "fs-snap-004",
							Name:                "tf_secure_fs_snap_test",
							Description:         "Updated Description",
							ExpirationTimestamp: "2035-05-06T09:01:47+00:00",
							ParentID:            "fs-parent-001",
							AccessType:          "Snapshot",
							IsSecure:            true,
						}, nil).Build()
				},
				Config: ProviderConfigForTesting + securefsSnapMockedUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_test", "description", "Updated Description"),
				),
			},
		},
	})
}

// U-013: TestSecureFSSnapshot_CreateReadError - Mock create success but GetFS error
func TestSecureFSSnapshot_CreateReadError(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateFsSnapshot).
						Return(gopowerstore.CreateResponse{ID: "fs-snap-005"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetFS).
						Return(gopowerstore.FileSystem{}, fmt.Errorf("mock error: unable to read filesystem snapshot")).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteFsSnapshot).
						Return(nil, nil).Build()
				},
				Config:      ProviderConfigForTesting + securefsSnapMockedConfig,
				ExpectError: regexp.MustCompile(".*Error getting filesystem snapshot after creation.*"),
			},
		},
	})
}

// I-002: TestAccSecureFSSnapshot_Create - TF_ACC gated, full lifecycle
func TestAccSecureFSSnapshot_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureFSSnapAccConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_acc_test", "name", "tf_secure_fs_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_acc_test", "description", "Acceptance Test Secure FS Snapshot"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_acc_test", "is_secure", "true"),
				),
			},
		},
	})
}

// I-006: TestAccSecureFSSnapshot_Import - TF_ACC gated, import test
func TestAccSecureFSSnapshot_Import(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureFSSnapAccConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_acc_test", "name", "tf_secure_fs_snap_acc"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.secure_acc_test", "is_secure", "true"),
				),
			},
			// Import Testing
			{
				Config:       ProviderConfigForTesting + SecureFSSnapAccConfig,
				ResourceName: "powerstore_filesystem_snapshot.secure_acc_test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_secure_fs_snap_acc", s[0].Attributes["name"])
					return nil
				},
			},
		},
	})
}

// Acceptance test HCL config for secure filesystem snapshot
var SecureFSSnapAccConfig = FsParams + `
resource "powerstore_filesystem_snapshot" "secure_acc_test" {
  name                 = "tf_secure_fs_snap_acc"
  description          = "Acceptance Test Secure FS Snapshot"
  filesystem_id        = powerstore_filesystem.test_fs_create.id
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
  access_type          = "Snapshot"
  depends_on           = [powerstore_filesystem.test_fs_create]
}
`
