/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"testing"

	"github.com/bytedance/mockey"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create SMBShare
func TestAccSMBShare_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Create Testing
			{
				Config: ProviderConfigForTesting + smbShareCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_smb_share.test1", "name", "terraform_smb"),
					resource.TestCheckResourceAttr("powerstore_smb_share.test1", "description", "terraform smb share"),
				),
			},
			// Import Testing
			{
				Config:            ProviderConfigForTesting + smbShareCreate,
				ResourceName:      "powerstore_smb_share.test1",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "terraform_smb", s[0].Attributes["name"])
					assert.Equal(t, "terraform smb share", s[0].Attributes["description"])
					return nil
				},
			},
			// Update Testing
			{
				Config: ProviderConfigForTesting + smbShareUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_smb_share.test1", "name", "terraform_smb"),
					resource.TestCheckResourceAttr("powerstore_smb_share.test1", "description", "terraform smb share update"),
				),
			},
			// Update Error
			{
				Config:      ProviderConfigForTesting + smbShareRename,
				ExpectError: regexp.MustCompile(".*Error updating smb share resource*"),
			},
			{
				Config:      ProviderConfigForTesting + smbShareCreateWithInvalidumask,
				ExpectError: regexp.MustCompile(".*Error updating smb share*"),
			},
			// Import Error
			{
				Config:        ProviderConfigForTesting + smbShareCreate,
				ResourceName:  "powerstore_smb_share.test1",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(".*Error reading smb share.*"),
				ImportStateId: "invalid-id",
			},
		},
	})

}

// Test to Create SMBShare with Invalid Values
func TestAccSMBShare_InvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + smbShareCreateWithInvalidFileSystemId,
				ExpectError: regexp.MustCompile(".*Enter a valid File System ID*"),
			},
			{
				Config:      ProviderConfigForTesting + smbShareCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
			},
			{
				Config:      ProviderConfigForTesting + smbShareCreateWithInvalidumask,
				ExpectError: regexp.MustCompile(".*Error creating smb share*"),
			},
		},
	})
}

// Test to mock SMB Share ACL error
func TestAccSMBShare_mockErrorACL(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).SetSMBShareACL).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + smbShareCreateWithACl,
				ExpectError: regexp.MustCompile("mock error"),
			},
			{
				PreConfig: func() {
					FunctionMocker.UnPatch()
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).GetSMBShareACL).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + smbShareCreateWithACl,
				ExpectError: regexp.MustCompile("mock error"),
			},
		},
	})
}

// Test to Add and remove SMB Share resource with ACl
func TestAccSMBShare_AddRemoveACL(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + smbShareCreateWithACl,
			},
			{
				Config: ProviderConfigForTesting + smbShareUpdateWithACl,
			},
			{
				Config: ProviderConfigForTesting + smbShareUpdateAddACl,
			},
			{
				Config: ProviderConfigForTesting + smbShareUpdateRemoveACl,
			},
		},
	})
}

var smbShareCreate = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"
}
`

var smbShareRename = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb1"
  path = "/test_fs"
  description = "terraform smb share"
}
`

var smbShareCreateWithInvalidFileSystemId = `
resource "powerstore_smb_share" "test1" {
  file_system_id = "invalid"
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"
}
`

var smbShareCreateWithoutName = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  path = "/test_fs"
  description = "terraform smb share"
}
`

var smbShareCreateWithInvalidumask = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"
  umask = "878"
}
`

var smbShareUpdate = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share update"
}
`

var smbShareCreateWithACl = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"   
  aces = [{"access_level":"Read","access_type":"Deny","trustee_name":"Everyone","trustee_type":"WellKnown"}]
}
`

var smbShareUpdateWithACl = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"
  aces = [{trustee_name="Everyone",trustee_type="WellKnown",access_level="Read",access_type="Allow"}]
}
`

var smbShareUpdateAddACl = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"
  aces = [{trustee_name="Everyone",trustee_type="WellKnown",access_level="Read",access_type="Allow"},{trustee_name="S-1-5-21-8-5-1-32",trustee_type="SID",access_level="Full",access_type="Allow"}]
}
`

var smbShareUpdateRemoveACl = FsParams +`
resource "powerstore_smb_share" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_smb"
  path = "/test_fs"
  description = "terraform smb share"
  aces = [{trustee_name="Everyone",trustee_type="WellKnown",access_level="Read",access_type="Allow"}]
}
`
