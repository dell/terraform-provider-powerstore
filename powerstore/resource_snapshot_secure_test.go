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

// Mock data for secure snapshot tests
var mockSecureVolSnap = gopowerstore.Volume{
	ID:          "snap-secure-001",
	Name:        "tf_secure_snap_test",
	Description: "Test Secure Snapshot",
	ProtectionData: gopowerstore.ProtectionData{
		ExpirationTimeStamp: "2035-05-06T09:01:47+00:00",
		ParentID:            "vol-parent-001",
		IsSecure:            true,
	},
	PerformancePolicyID: "default_medium",
}

var mockDefaultVolSnap = gopowerstore.Volume{
	ID:          "snap-default-001",
	Name:        "tf_secure_snap_default",
	Description: "Test Default Snapshot",
	ProtectionData: gopowerstore.ProtectionData{
		ExpirationTimeStamp: "2035-05-06T09:01:47+00:00",
		ParentID:            "vol-parent-001",
		IsSecure:            false,
	},
	PerformancePolicyID: "default_medium",
}

var mockSecureVolSnapUpdated = gopowerstore.Volume{
	ID:          "snap-secure-001",
	Name:        "tf_secure_snap_test_updated",
	Description: "Test Secure Snapshot Updated",
	ProtectionData: gopowerstore.ProtectionData{
		ExpirationTimeStamp: "2035-05-06T09:01:47+00:00",
		ParentID:            "vol-parent-001",
		IsSecure:            true,
	},
	PerformancePolicyID: "default_medium",
}

// Mocked HCL configs (hardcoded volume_id since API calls are mocked)
var secureSnapMockedConfig = `
resource "powerstore_volume_snapshot" "secure_test" {
  name                 = "tf_secure_snap_test"
  description          = "Test Secure Snapshot"
  volume_id            = "vol-parent-001"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var secureSnapMockedDefaultConfig = `
resource "powerstore_volume_snapshot" "secure_test" {
  name                 = "tf_secure_snap_default"
  description          = "Test Default Snapshot"
  volume_id            = "vol-parent-001"
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var secureSnapMockedUpdateConfig = `
resource "powerstore_volume_snapshot" "secure_test" {
  name                 = "tf_secure_snap_test_updated"
  description          = "Test Secure Snapshot Updated"
  volume_id            = "vol-parent-001"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

// Acceptance test HCL configs
var SecureSnapCreate = VolumeParams + `
resource "powerstore_volume_snapshot" "secure_test" {
  name                 = "tf_secure_snap_test"
  description          = "Test Secure Snapshot"
  volume_id            = powerstore_volume.volume_create_test.id
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var SecureSnapCreateDefault = VolumeParams + `
resource "powerstore_volume_snapshot" "secure_test" {
  name                 = "tf_secure_snap_default"
  description          = "Test Default Snapshot"
  volume_id            = powerstore_volume.volume_create_test.id
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

// --- Unit Tests (mocked, no TF_ACC check) ---

// U-001: Test secure snapshot creation with is_secure = true
func TestSecureVolumeSnapshot_Create(t *testing.T) {
	var getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
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
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "is_secure", "true"),
				),
			},
		},
	})
}

// U-002: Test snapshot creation with default is_secure (false when not specified)
func TestSecureVolumeSnapshot_CreateDefault(t *testing.T) {
	var getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
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
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-default-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockDefaultVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedDefaultConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "is_secure", "false"),
				),
			},
		},
	})
}

// U-003: Test that Read correctly maps is_secure from API response
func TestSecureVolumeSnapshot_ReadState(t *testing.T) {
	var getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
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
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "name", "tf_secure_snap_test"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "description", "Test Secure Snapshot"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "volume_id", "vol-parent-001"),
				),
			},
		},
	})
}

// U-004: Test is_secure persists through update
func TestSecureVolumeSnapshot_Update(t *testing.T) {
	var getMock, deleteMock, modifyMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
		}
		if getMock != nil {
			getMock.UnPatch()
		}
		if deleteMock != nil {
			deleteMock.UnPatch()
		}
		if modifyMock != nil {
			modifyMock.UnPatch()
		}
		if newClientMock != nil {
			newClientMock.UnPatch()
		}
		if getVolumesMock != nil {
			getVolumesMock.UnPatch()
		}
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "is_secure", "true"),
				),
			},
			// Step 2: Update name and description, verify is_secure persists
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					if getMock != nil {
						getMock.UnPatch()
					}
					if deleteMock != nil {
						deleteMock.UnPatch()
					}
					modifyMock = mockey.Mock((*gopowerstore.ClientIMPL).ModifyVolume).Return(nil, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnapUpdated, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "name", "tf_secure_snap_test_updated"),
				),
			},
		},
	})
}

// U-005: Test error when GetSnapshot fails after CreateSnapshot succeeds
func TestSecureVolumeSnapshot_CreateReadError(t *testing.T) {
	var getMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
		}
		if getMock != nil {
			getMock.UnPatch()
		}
		if newClientMock != nil {
			newClientMock.UnPatch()
		}
		if getVolumesMock != nil {
			getVolumesMock.UnPatch()
		}
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(gopowerstore.Volume{}, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + secureSnapMockedConfig,
				ExpectError: regexp.MustCompile(".*Error getting volume snapshot after creation.*"),
			},
		},
	})
}

// U-006: Test Read error when GetSnapshot fails during refresh
func TestSecureVolumeSnapshot_ReadError(t *testing.T) {
	var getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
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
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create successfully
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
			},
			// Step 2: Mock GetSnapshot to return error during refresh
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					if getMock != nil {
						getMock.UnPatch()
					}
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(gopowerstore.Volume{}, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + secureSnapMockedConfig,
				ExpectError: regexp.MustCompile(".*Error reading snapshot.*"),
			},
			// Step 3: Restore GetSnapshot mock for auto-destroy cleanup
			{
				PreConfig: func() {
					if getMock != nil {
						getMock.UnPatch()
					}
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
			},
		},
	})
}

// U-007: Test error when deleting a secure snapshot is blocked
func TestSecureVolumeSnapshot_DeleteSecureError(t *testing.T) {
	var getMock, deleteMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
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
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create successfully
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
			},
			// Step 2: Mock DeleteSnapshot to return error (simulating secure snapshot deletion blocked)
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					if deleteMock != nil {
						deleteMock.UnPatch()
					}
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, fmt.Errorf("secure snapshot cannot be deleted before expiration")).Build()
				},
				Config:      ProviderConfigForTesting + secureSnapMockedConfig,
				Destroy:     true,
				ExpectError: regexp.MustCompile(".*Error deleting snapshot.*"),
			},
			// Step 3: Restore DeleteSnapshot mock for auto-destroy cleanup
			{
				PreConfig: func() {
					if deleteMock != nil {
						deleteMock.UnPatch()
					}
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
			},
		},
	})
}

// U-008: Test error when ModifyVolume fails during update
func TestSecureVolumeSnapshot_UpdateModifyError(t *testing.T) {
	var getMock, deleteMock, modifyMock, newClientMock, getVolumesMock *mockey.Mocker
	t.Cleanup(func() {
		if FunctionMocker != nil {
			FunctionMocker.UnPatch()
		}
		if getMock != nil {
			getMock.UnPatch()
		}
		if deleteMock != nil {
			deleteMock.UnPatch()
		}
		if modifyMock != nil {
			modifyMock.UnPatch()
		}
		if newClientMock != nil {
			newClientMock.UnPatch()
		}
		if getVolumesMock != nil {
			getVolumesMock.UnPatch()
		}
	})

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create successfully
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					newClientMock = mockey.Mock(client.NewClient).Return(&client.Client{PStoreClient: &gopowerstore.ClientIMPL{}}, nil).Build()
					getVolumesMock = mockey.Mock((*gopowerstore.ClientIMPL).GetVolumes).Return([]gopowerstore.Volume{}, nil).Build()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshot).Return(gopowerstore.CreateResponse{ID: "snap-secure-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshot).Return(mockSecureVolSnap, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(nil, nil).Build()
				},
				Config: ProviderConfigForTesting + secureSnapMockedConfig,
			},
			// Step 2: Mock ModifyVolume to return error
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					modifyMock = mockey.Mock((*gopowerstore.ClientIMPL).ModifyVolume).Return(nil, fmt.Errorf("mock modify error")).Build()
				},
				Config:      ProviderConfigForTesting + secureSnapMockedUpdateConfig,
				ExpectError: regexp.MustCompile(".*Error updating volume snapshot resource.*"),
			},
		},
	})
}

// --- Acceptance Tests (require TF_ACC) ---

// I-001: Acceptance test for secure snapshot creation with full lifecycle
func TestAccSecureVolumeSnapshot_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "name", "tf_secure_snap_test"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "description", "Test Secure Snapshot"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.secure_test", "is_secure", "true"),
				),
			},
		},
	})
}

// I-005: Acceptance test for importing an existing secure snapshot
func TestAccSecureVolumeSnapshot_Import(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapCreate,
			},
			{
				Config:       ProviderConfigForTesting + SecureSnapCreate,
				ResourceName: "powerstore_volume_snapshot.secure_test",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_secure_snap_test", s[0].Attributes["name"])
					assert.Equal(t, "true", s[0].Attributes["is_secure"])
					return nil
				},
			},
		},
	})
}
