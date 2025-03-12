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
	"terraform-provider-powerstore/client"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

var yesmocker, nomocker *mockey.Mocker

// Test to Create NFSExport
func TestAccNFSExport_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	defer func() {
		if yesmocker != nil {
			yesmocker.UnPatch()
		}
		if nomocker != nil {
			nomocker.UnPatch()
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Validate error with invalid host set IPs
			{
				Config:      ProviderConfigForTesting + nfsHostInvalidIP,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(".*Error parsing host values.*"),
			},
			// Validate error with invalid host ip unknown
			{
				Config:      ProviderConfigForTesting + nfsHostInvalidIPUnknown,
				ExpectError: regexp.MustCompile(".*Error parsing host values.*"),
			},
			// Validate error with invalid host set IPs unknown
			{
				Config:      ProviderConfigForTesting + nfsHostInvalidIPSetUnknown,
				ExpectError: regexp.MustCompile(".*Error parsing host values.*"),
			},
			// Create Error
			{
				Config:      ProviderConfigForTesting + nfsCreateError,
				ExpectError: regexp.MustCompile(".*Error creating nfs export.*"),
			},
			// read error after create - mocked
			{
				PreConfig: func() {
					yesmocker = mockey.Mock((*gopowerstore.ClientIMPL).CreateNFSExport).Return(gopowerstore.CreateResponse{ID: "1"}, nil).Build()
					nomocker = mockey.Mock((*gopowerstore.ClientIMPL).GetNFSExport).Return(nil, fmt.Errorf("Error reading nfs export")).Build()
				},
				Config:      ProviderConfigForTesting + nfsCreateError,
				ExpectError: regexp.MustCompile(".*Error reading nfs export.*"),
			},
			// Create Testing
			{
				PreConfig: func() {
					yesmocker.UnPatch()
					nomocker.UnPatch()
				},
				Config: ProviderConfigForTesting + nfsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "name", "terraform_nfs"),
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "description", "terraform nfs export"),
				),
			},
			// Import Testing
			{
				Config:       ProviderConfigForTesting + nfsCreate,
				ResourceName: "powerstore_nfs_export.test1",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "terraform_nfs", s[0].Attributes["name"])
					assert.Equal(t, "terraform nfs export", s[0].Attributes["description"])
					return nil
				},
			},
			// error in update - mocked
			{
				PreConfig: func() {
					yesmocker = mockey.Mock((*client.Client).ModifyNFSExport).Return(fmt.Errorf("Error reading nfs export")).Build()
				},
				Config:      ProviderConfigForTesting + nfsUpdate,
				ExpectError: regexp.MustCompile(".*Error reading nfs export.*"),
			},
			// Update Testing
			{
				PreConfig: func() {
					yesmocker.UnPatch()
				},
				Config: ProviderConfigForTesting + nfsUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "name", "terraform_nfs"),
					resource.TestCheckResourceAttr("powerstore_nfs_export.test1", "description", "terraform nfs export update"),
				),
			},
			// Update Error - cannot change name
			{
				Config:      ProviderConfigForTesting + nfsInvalidUpate,
				ExpectError: regexp.MustCompile(".*Error updating nfs export resource.*"),
			},
			// delete error
			{
				PreConfig: func() {
					nomocker = mockey.Mock((*gopowerstore.ClientIMPL).DeleteNFSExport).Return(nil, fmt.Errorf("mock error")).Build()
				},
				Config:      ProviderConfigForTesting + nfsUpdate,
				Destroy:     true,
				ExpectError: regexp.MustCompile(".*Error deleting nfsExport.*"),
			},
			// Import Error
			{
				PreConfig: func() {
					nomocker.UnPatch()
				},
				Config:        ProviderConfigForTesting + nfsCreate,
				ResourceName:  "powerstore_nfs_export.test1",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(".*Error reading nfs export.*"),
				ImportStateId: "invalid-id",
			},
		},
	})
}

var nfsCreate = FsParams + `
resource "powerstore_nfs_export" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_nfs"
  path = "/test_fs"

  anonymous_gid = -23
  anonymous_uid = -23 
  description = "terraform nfs export"
  is_no_suid = true 

  min_security = "Sys"
  default_access = "Read_Only"

  no_access_hosts = [
    "192.168.1.0/24",
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",
    "2001:db8:85a3::8a2e:370:7334",
    "2001:db8:85a3::/64",
    "2001:db8:85a3::8a2e:370:7334",
    "hostname1",
    "hostname2",
    "@netgroup1",
    "dellinv.gov.in",
    "dellinv.gov",
  ]
}
`

var nfsHostInvalidIP = `
resource "powerstore_nfs_export" "test1" {
  file_system_id = "whatever"
  name = "terraform_nfs"
  path = "/test_fs"

  no_access_hosts = [
    "192.168.1.0/24",
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",
    "192.168.1.54/255.1009.255.0",    
    "2001:db8:85a3::8a2e:370:7334/255.255.255.0",
    "2001:db8:85a3::8a2e:370:7334",
    "2001:db8:85a3::/64",
  ]
}
`

var nfsHostInvalidIPUnknown = `
resource terraform_data inv_ip {
	input = "192.168.1.54/255.1009.255.0"
}
resource "powerstore_nfs_export" "test1" {
  file_system_id = "whatever"
  name = "terraform_nfs"
  path = "/test_fs"

  no_access_hosts = [
    "192.168.1.0/24",
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",
	terraform_data.inv_ip.output, 
    "2001:db8:85a3::8a2e:370:7334/255.255.255.0",
    "2001:db8:85a3::8a2e:370:7334",
    "2001:db8:85a3::/64",
  ]
}
`

var nfsHostInvalidIPSetUnknown = `
resource terraform_data inv_ip_set {
	input = [
    "192.168.1.0/24",
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",
    "192.168.1.54/255.1009.255.0",    
    "2001:db8:85a3::8a2e:370:7334/255.255.255.0",
    "2001:db8:85a3::8a2e:370:7334",
    "2001:db8:85a3::/64",
  ]
}
resource "powerstore_nfs_export" "test1" {
  file_system_id = "whatever"
  name = "terraform_nfs"
  path = "/test_fs"

  no_access_hosts = terraform_data.inv_ip_set.output
}
`

var nfsCreateError = `
resource "powerstore_nfs_export" "test2" {
  file_system_id = "invalid"
  name = "terraform_nfs"
  path = "/test_fs"
}
`

var nfsUpdate = FsParams + `
resource "powerstore_nfs_export" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform_nfs"
  path = "/test_fs"

  anonymous_gid = 23
  anonymous_uid = 23 
  description = "terraform nfs export update"
  is_no_suid = false 

  min_security = "Kerberos"
  default_access = "Read_Write"

  no_access_hosts = [
    "192.168.1.0/24",
    "192.168.1.0/26",
    "192.168.1.54/255.255.255.0",
    "2001:db8:85a3::8a2e:370:7334",
    "2001:db8:85a3::/64",
    "2001:db8:85a3::8a2e:370:7334",
    "hostname1",
    "@netgroup1",
    "dellinv.gov",
  ]
}
`

var nfsInvalidUpate = FsParams + `
resource "powerstore_nfs_export" "test1" {
  file_system_id = powerstore_filesystem.test_fs_create.id
  name = "terraform-nfs1"
  path = "/test_fs1"
}
`
