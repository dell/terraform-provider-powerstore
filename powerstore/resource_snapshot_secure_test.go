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
	"testing"

	"github.com/bytedance/mockey"
	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test U-01: Create secure volume snapshot
func TestAccVolumeSnapshot_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "name", "tf_snap_secure_acc"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "expiration_timestamp", "2035-05-06T09:01:47Z"),
				),
			},
		},
	})
}

// Test U-02: Create secure snapshot missing expiration - should fail
func TestAccVolumeSnapshot_CreateSecureMissingExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + SecureSnapParamsCreateNoExpiration,
				ExpectError: regexp.MustCompile(".*[Ss]ecure snapshots require.*expiration.*"),
			},
		},
	})
}

// Test U-03: Create non-secure snapshot - is_secure should default to false
func TestAccVolumeSnapshot_CreateNonSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + NonSecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "name", "tf_snap_nonsecure_acc"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
		},
	})
}

// Test U-07/U-08: Update - mark existing snapshot as secure (one-way lock)
func TestAccVolumeSnapshot_UpdateToSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create non-secure snapshot
			{
				Config: ProviderConfigForTesting + NonSecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
			// Step 2: Update to secure (should succeed)
			{
				Config: ProviderConfigForTesting + SnapParamsUpdateToSecure,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "expiration_timestamp", "2035-05-06T09:01:47Z"),
				),
			},
		},
	})
}

// Test U-08: Update - reject unsecure attempt (one-way lock)
func TestAccVolumeSnapshot_RejectUnsecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			// Step 1: Create secure snapshot
			{
				Config: ProviderConfigForTesting + SecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
			// Step 2: Try to unsecure - should fail
			{
				Config:      ProviderConfigForTesting + SnapParamsRevertSecure,
				ExpectError: regexp.MustCompile(".*[Oo]ne.way lock.*|.*cannot.*unlock.*|.*cannot.*change.*true.*false.*"),
			},
		},
	})
}

// Test U-11: Delete secure snapshot - should get API error
func TestAccVolumeSnapshot_DeleteSecureMocked(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapParamsCreate,
			},
			// Mock the delete to return 422
			{
				PreConfig: func() {
					FunctionMocker = mockey.Mock((*gopowerstore.ClientIMPL).DeleteSnapshot).Return(gopowerstore.CreateResponse{}, fmt.Errorf("secure snapshot cannot be deleted before expiration")).Build()
				},
				Config:      ProviderConfigForTesting + SecureSnapParamsCreate,
				Destroy:     true,
				ExpectError: regexp.MustCompile(".*[Ss]ecure snapshot.*|.*cannot.*delete.*expiration.*"),
			},
		},
	})
}

// Test U-12: Import secure snapshot
func TestAccVolumeSnapshot_ImportSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:       ProviderConfigForTesting + SecureSnapParamsCreate,
				ResourceName: "powerstore_volume_snapshot.test_secure",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "tf_snap_secure_acc", s[0].Attributes["name"])
					assert.Equal(t, "true", s[0].Attributes["is_secure"])
					return nil
				},
			},
		},
	})
}

// Test I-06: Idempotent re-apply of secure snapshot
func TestAccVolumeSnapshot_SecureIdempotent(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
			// Re-apply same config - should produce no changes
			{
				Config: ProviderConfigForTesting + SecureSnapParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test E-07: Explicit is_secure=false
func TestAccVolumeSnapshot_ExplicitFalse(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + SnapParamsExplicitFalse,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_explicit_false", "is_secure", "false"),
				),
			},
		},
	})
}

// --- HCL Test Configs ---

var SecureSnapParamsCreate = VolumeParams + `
resource "powerstore_volume_snapshot" "test_secure" {
  name = "tf_snap_secure_acc"
  description = "Test Secure Snapshot Resource"
  volume_id = powerstore_volume.volume_create_test.id
  performance_policy_id = "default_medium"
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = true
}
`

var SecureSnapParamsCreateNoExpiration = VolumeParams + `
resource "powerstore_volume_snapshot" "test_secure_no_exp" {
  name = "tf_snap_secure_no_exp"
  description = "Secure Snapshot Without Expiration"
  volume_id = powerstore_volume.volume_create_test.id
  is_secure = true
}
`

var NonSecureSnapParamsCreate = VolumeParams + `
resource "powerstore_volume_snapshot" "test_nonsecure" {
  name = "tf_snap_nonsecure_acc"
  description = "Test Non-Secure Snapshot Resource"
  volume_id = powerstore_volume.volume_create_test.id
  performance_policy_id = "default_medium"
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
`

var SnapParamsUpdateToSecure = VolumeParams + `
resource "powerstore_volume_snapshot" "test_nonsecure" {
  name = "tf_snap_nonsecure_acc"
  description = "Test Non-Secure Snapshot Resource"
  volume_id = powerstore_volume.volume_create_test.id
  performance_policy_id = "default_medium"
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = true
}
`

var SnapParamsRevertSecure = VolumeParams + `
resource "powerstore_volume_snapshot" "test_secure" {
  name = "tf_snap_secure_acc"
  description = "Test Secure Snapshot Resource"
  volume_id = powerstore_volume.volume_create_test.id
  performance_policy_id = "default_medium"
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = false
}
`

var SnapParamsExplicitFalse = VolumeParams + `
resource "powerstore_volume_snapshot" "test_explicit_false" {
  name = "tf_snap_explicit_false"
  description = "Explicitly non-secure snapshot"
  volume_id = powerstore_volume.volume_create_test.id
  performance_policy_id = "default_medium"
  expiration_timestamp = "2035-05-06T09:01:47Z"
  is_secure = false
}
`
