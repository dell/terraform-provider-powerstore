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

// Unit test HCL configs for secure snapshot rule tests

var secureRuleMockedConfig = `
resource "powerstore_snapshotrule" "secure_test" {
  name              = "tf_secure_rule_test"
  interval          = "One_Hour"
  desired_retention = 168
  is_secure         = true
}
`

var secureRuleMockedDefaultConfig = `
resource "powerstore_snapshotrule" "secure_test" {
  name              = "tf_secure_rule_default"
  interval          = "One_Hour"
  desired_retention = 168
}
`

var secureRuleMockedUpdateConfig = `
resource "powerstore_snapshotrule" "secure_test" {
  name              = "tf_secure_rule_test"
  interval          = "One_Hour"
  desired_retention = 336
  is_secure         = true
}
`

// U-017: TestSecureSnapshotRule_Create - Mocked, is_secure=true
func TestSecureSnapshotRule_Create(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshotRule).
						Return(gopowerstore.CreateResponse{ID: "sr-001"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotRule).
						Return(gopowerstore.SnapshotRule{
							ID:               "sr-001",
							Name:             "tf_secure_rule_test",
							Interval:         gopowerstore.SnapshotRuleIntervalEnum("One_Hour"),
							DesiredRetention: int32(168),
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshotRule).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureRuleMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "name", "tf_secure_rule_test"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "interval", "One_Hour"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "desired_retention", "168"),
				),
			},
		},
	})
}

// U-018: TestSecureSnapshotRule_ReadState - Mocked read with IsSecure=true
func TestSecureSnapshotRule_ReadState(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshotRule).
						Return(gopowerstore.CreateResponse{ID: "sr-002"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotRule).
						Return(gopowerstore.SnapshotRule{
							ID:               "sr-002",
							Name:             "tf_secure_rule_test",
							Interval:         gopowerstore.SnapshotRuleIntervalEnum("One_Hour"),
							DesiredRetention: int32(168),
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshotRule).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureRuleMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "id", "sr-002"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "name", "tf_secure_rule_test"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "interval", "One_Hour"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "desired_retention", "168"),
				),
			},
		},
	})
}

// U-019: TestSecureSnapshotRule_Update - Mocked update setting is_secure
func TestSecureSnapshotRule_Update(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshotRule).
						Return(gopowerstore.CreateResponse{ID: "sr-003"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotRule).
						Return(gopowerstore.SnapshotRule{
							ID:               "sr-003",
							Name:             "tf_secure_rule_test",
							Interval:         gopowerstore.SnapshotRuleIntervalEnum("One_Hour"),
							DesiredRetention: int32(168),
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshotRule).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(true).Build()
				},
				Config: ProviderConfigForTesting + secureRuleMockedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "desired_retention", "168"),
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
					modifyMock = mockey.Mock((*gopowerstore.ClientIMPL).ModifySnapshotRule).
						Return(nil, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotRule).
						Return(gopowerstore.SnapshotRule{
							ID:               "sr-003",
							Name:             "tf_secure_rule_test",
							Interval:         gopowerstore.SnapshotRuleIntervalEnum("One_Hour"),
							DesiredRetention: int32(336),
						}, nil).Build()
				},
				Config: ProviderConfigForTesting + secureRuleMockedUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "desired_retention", "336"),
				),
			},
		},
	})
}

// U-020: TestSecureSnapshotRule_CreateDefault - Mocked, no is_secure, check false
func TestSecureSnapshotRule_CreateDefault(t *testing.T) {
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
					createMock = mockey.Mock((*gopowerstore.ClientIMPL).CreateSnapshotRule).
						Return(gopowerstore.CreateResponse{ID: "sr-004"}, nil).Build()
					getMock = mockey.Mock((*gopowerstore.ClientIMPL).GetSnapshotRule).
						Return(gopowerstore.SnapshotRule{
							ID:               "sr-004",
							Name:             "tf_secure_rule_default",
							Interval:         gopowerstore.SnapshotRuleIntervalEnum("One_Hour"),
							DesiredRetention: int32(168),
						}, nil).Build()
					deleteMock = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshotRule).
						Return(nil, nil).Build()
					mockey.Mock((*client.Client).FetchIsSecure).Return(false).Build()
				},
				Config: ProviderConfigForTesting + secureRuleMockedDefaultConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "is_secure", "false"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_test", "name", "tf_secure_rule_default"),
				),
			},
		},
	})
}

// I-004: TestAccSecureSnapshotRule_Create - TF_ACC gated
func TestAccSecureSnapshotRule_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureRuleAccConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_acc_test", "name", "tf_secure_rule_acc"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_acc_test", "interval", "One_Hour"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_acc_test", "desired_retention", "168"),
					resource.TestCheckResourceAttr("powerstore_snapshotrule.secure_acc_test", "is_secure", "true"),
				),
			},
			// Import Testing
			{
				Config:            ProviderConfigForTesting + SecureRuleAccConfig,
				ResourceName:      "powerstore_snapshotrule.secure_acc_test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: false,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_secure_rule_acc", s[0].Attributes["name"])
					assert.Equal(t, "One_Hour", s[0].Attributes["interval"])
					assert.Equal(t, "168", s[0].Attributes["desired_retention"])
					return nil
				},
			},
			// Import Negative Testing
			{
				Config:        ProviderConfigForTesting + SecureRuleAccConfig,
				ResourceName:  "powerstore_snapshotrule.secure_acc_test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportSRDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

// Acceptance test HCL config for secure snapshot rule
var SecureRuleAccConfig = `
resource "powerstore_snapshotrule" "secure_acc_test" {
  name              = "tf_secure_rule_acc"
  interval          = "One_Hour"
  desired_retention = 168
  is_secure         = true
}
`
