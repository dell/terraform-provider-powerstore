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
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// --- Config generators for secure volume snapshot tests ---
// The dedicated volume is created/managed by the SDK in setupSecureVolAndProbe.
// These functions accept the volume ID directly so Terraform only manages the snapshot.

// secureSnapConfig generates an HCL config for a secure volume snapshot.
func secureSnapConfig(volID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volume_snapshot" "test_secure" {
  name                  = "%s"
  description           = "Test Secure Snapshot Resource"
  volume_id             = "%s"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "%s"
  is_secure             = true
}
`, snapName, volID, expTS)
}

// secureSnapRevertConfig generates an HCL config that tries to set
// is_secure=false on a previously secure snapshot (should be rejected).
func secureSnapRevertConfig(volID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volume_snapshot" "test_secure" {
  name                  = "%s"
  description           = "Test Secure Snapshot Resource"
  volume_id             = "%s"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "%s"
  is_secure             = false
}
`, snapName, volID, expTS)
}

// nonsecureSnapConfig generates an HCL config for a non-secure snapshot.
func nonsecureSnapConfig(volID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volume_snapshot" "test_nonsecure" {
  name                  = "%s"
  description           = "Test Non-Secure Snapshot Resource"
  volume_id             = "%s"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "%s"
}
`, snapName, volID, expTS)
}

// nonsecureSnapUpdateToSecureConfig generates an HCL config that updates
// a previously non-secure snapshot to secure.
func nonsecureSnapUpdateToSecureConfig(volID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volume_snapshot" "test_nonsecure" {
  name                  = "%s"
  description           = "Test Non-Secure Snapshot Resource"
  volume_id             = "%s"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "%s"
  is_secure             = true
}
`, snapName, volID, expTS)
}

// setupSecureVolAndProbe creates the dedicated volume via the SDK,
// probes the array clock, and returns (volumeID, expirationTS, cleanup, error).
// The caller MUST call the returned cleanup function in a defer.
func setupSecureVolAndProbe(t *testing.T) (string, string, func()) {
	t.Helper()
	client, err := secureTestSDKClient()
	if err != nil {
		t.Fatalf("failed to create SDK client: %v", err)
	}

	// Create the dedicated volume
	volName := "ZZZ_AT_tf_secure_vol"
	var volID string
	// Check if volume already exists
	vol, err := client.GetVolumeByName(context.Background(), volName)
	if err == nil && vol.ID != "" {
		volID = vol.ID
	} else {
		size := int64(1048576) // 1 GB in sectors (512-byte)
		createResp, err := client.CreateVolume(context.Background(), &gopowerstore.VolumeCreate{
			Name: &volName,
			Size: &size,
		})
		if err != nil {
			t.Fatalf("failed to create dedicated secure test volume: %v", err)
		}
		volID = createResp.ID
	}

	// Probe the array clock
	expTS, err := probeArrayClockViaSnapshot(client, volID)
	if err != nil {
		t.Fatalf("failed to probe array clock: %v", err)
	}

	cleanup := func() {
		// Best-effort cleanup of the dedicated volume
		_, _ = client.DeleteVolume(context.Background(), nil, volID)
	}

	return volID, expTS, cleanup
}

// Test U-01: Create secure volume snapshot
func TestAccVolumeSnapshot_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("create_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureSnapConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "name", snapName),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "expiration_timestamp", expTS),
				),
			},
		},
	})
}

// Test U-02: Create secure snapshot missing expiration - should fail
// This is a client-side validation so no array clock probe needed.
func TestAccVolumeSnapshot_CreateSecureMissingExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
resource "powerstore_volume_snapshot" "test_secure_no_exp" {
  name        = "ZZZ_AT_tf_snap_secure_no_exp"
  description = "Secure Snapshot Without Expiration"
  volume_id   = "fake-vol-id-for-validation"
  is_secure   = true
}
`,
				ExpectError: regexp.MustCompile(".*Secure snapshots require.*expiration.*"),
			},
		},
	})
}

// Test U-03: Create non-secure snapshot - is_secure should default to false
func TestAccVolumeSnapshot_CreateNonSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("create_nonsec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + nonsecureSnapConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "name", snapName),
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

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("upd_to_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create non-secure snapshot
			{
				Config: ProviderConfigForTesting + nonsecureSnapConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
			// Step 2: Update to secure (should succeed)
			{
				Config: ProviderConfigForTesting + nonsecureSnapUpdateToSecureConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_nonsecure", "expiration_timestamp", expTS),
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

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("reject_unsec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create secure snapshot
			{
				Config: ProviderConfigForTesting + secureSnapConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
			// Step 2: Try to unsecure - should fail
			{
				Config:      ProviderConfigForTesting + secureSnapRevertConfig(volID, expTS, snapName),
				ExpectError: regexp.MustCompile(".*one-way lock.*|.*cannot.*changed from true to false.*"),
			},
		},
	})
}

// Test U-12: Import secure snapshot
func TestAccVolumeSnapshot_ImportSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("import_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureSnapConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:       ProviderConfigForTesting + secureSnapConfig(volID, expTS, snapName),
				ResourceName: "powerstore_volume_snapshot.test_secure",
				ImportState:  true,
				ExpectError:  nil,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, snapName, s[0].Attributes["name"])
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

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("idempotent")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureSnapConfig(volID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_secure", "is_secure", "true"),
				),
			},
			// Re-apply same config - should produce no changes
			{
				Config: ProviderConfigForTesting + secureSnapConfig(volID, expTS, snapName),
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

	volID, expTS, cleanup := setupSecureVolAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("explicit_false")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
resource "powerstore_volume_snapshot" "test_explicit_false" {
  name                  = "%s"
  description           = "Explicitly non-secure snapshot"
  volume_id             = "%s"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "%s"
  is_secure             = false
}
`, snapName, volID, expTS),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volume_snapshot.test_explicit_false", "is_secure", "false"),
				),
			},
		},
	})
}
