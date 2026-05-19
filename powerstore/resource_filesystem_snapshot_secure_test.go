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

// --- Config generators for secure FS snapshot tests ---
// The dedicated filesystem is created/managed by the SDK in setupSecureFSAndProbe.
// These functions accept the filesystem ID directly so Terraform only manages the snapshot.

func secureFSSnapConfig(fsID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_filesystem_snapshot" "test_secure" {
  name                 = "%s"
  description          = "Test Secure FS Snapshot"
  filesystem_id        = "%s"
  expiration_timestamp = "%s"
  access_type          = "Snapshot"
  is_secure            = true
}
`, snapName, fsID, expTS)
}

func secureFSSnapRevertConfig(fsID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_filesystem_snapshot" "test_secure" {
  name                 = "%s"
  description          = "Test Secure FS Snapshot"
  filesystem_id        = "%s"
  expiration_timestamp = "%s"
  access_type          = "Snapshot"
  is_secure            = false
}
`, snapName, fsID, expTS)
}

func nonsecureFSSnapConfig(fsID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_filesystem_snapshot" "test_nonsecure" {
  name                 = "%s"
  description          = "Non-Secure FS Snapshot"
  filesystem_id        = "%s"
  expiration_timestamp = "%s"
  access_type          = "Snapshot"
}
`, snapName, fsID, expTS)
}

func nonsecureFSSnapUpdateToSecureConfig(fsID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_filesystem_snapshot" "test_nonsecure" {
  name                 = "%s"
  description          = "Non-Secure FS Snapshot"
  filesystem_id        = "%s"
  expiration_timestamp = "%s"
  access_type          = "Snapshot"
  is_secure            = true
}
`, snapName, fsID, expTS)
}

// setupSecureFSAndProbe creates the dedicated filesystem via SDK,
// probes the array clock via an FS snapshot, and returns (fsID, expTS, cleanup).
func setupSecureFSAndProbe(t *testing.T) (string, string, func()) {
	t.Helper()
	client, err := secureTestSDKClient()
	if err != nil {
		t.Fatalf("failed to create SDK client: %v", err)
	}

	// Create the dedicated filesystem
	fsName := "ZZZ_AT_tf_secure_fs"
	var fsID string
	fs, err := client.GetFSByName(context.Background(), fsName)
	if err == nil && fs.ID != "" {
		fsID = fs.ID
	} else {
		createResp, err := client.CreateFS(context.Background(), &gopowerstore.FsCreate{
			Name:         fsName,
			NASServerID:  nasServerID,
			Size:         3221225472, // 3 GB
			ConfigType:   "General",
			AccessPolicy: "UNIX",
		})
		if err != nil {
			t.Fatalf("failed to create dedicated secure test filesystem: %v", err)
		}
		fsID = createResp.ID
	}

	// Probe the array clock via an FS snapshot
	expTS, err := probeArrayClockViaFSSnapshot(client, fsID)
	if err != nil {
		t.Fatalf("failed to probe array clock via FS snapshot: %v", err)
	}

	cleanup := func() {
		// Best-effort cleanup; may fail if secure snapshots still exist
		_, _ = client.DeleteFS(context.Background(), fsID)
	}

	return fsID, expTS, cleanup
}

// Test U-21: Create secure filesystem snapshot
func TestAccFileSystemSnapshot_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_create_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "name", snapName),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "expiration_timestamp", expTS),
				),
			},
		},
	})
}

// Test U-22: Create secure FS snapshot missing expiration - should fail
func TestAccFileSystemSnapshot_CreateSecureMissingExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
resource "powerstore_filesystem_snapshot" "test_secure_no_exp" {
  name          = "ZZZ_AT_tf_fs_snap_secure_no_exp"
  description   = "Secure FS Snapshot Without Expiration"
  filesystem_id = "fake-fs-id-for-validation"
  access_type   = "Snapshot"
  is_secure     = true
}
`,
				ExpectError: regexp.MustCompile(".*Secure snapshots require.*expiration.*"),
			},
		},
	})
}

// Test U-23: Create non-secure FS snapshot - is_secure should default to false
func TestAccFileSystemSnapshot_CreateNonSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_create_nonsec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + nonsecureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_nonsecure", "name", snapName),
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
		},
	})
}

// Test U-24: Update FS snapshot - mark as secure (one-way lock)
func TestAccFileSystemSnapshot_UpdateToSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_upd_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + nonsecureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + nonsecureFSSnapUpdateToSecureConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_nonsecure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test U-25: Reject unsecure FS snapshot (one-way lock)
func TestAccFileSystemSnapshot_RejectUnsecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_reject")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:      ProviderConfigForTesting + secureFSSnapRevertConfig(fsID, expTS, snapName),
				ExpectError: regexp.MustCompile(".*one-way lock.*|.*cannot.*changed from true to false.*"),
			},
		},
	})
}

// Test U-26: Import secure FS snapshot
func TestAccFileSystemSnapshot_ImportSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_import_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:       ProviderConfigForTesting + secureFSSnapConfig(fsID, expTS, snapName),
				ResourceName: "powerstore_filesystem_snapshot.test_secure",
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

// Test I-08: Idempotent re-apply of secure FS snapshot
func TestAccFileSystemSnapshot_SecureIdempotent(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_idempotent")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
				),
			},
			// Re-apply same config - should produce no changes
			{
				Config: ProviderConfigForTesting + secureFSSnapConfig(fsID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_secure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test E-09: Explicit is_secure=false for FS snapshot
func TestAccFileSystemSnapshot_ExplicitFalse(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	fsID, expTS, cleanup := setupSecureFSAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("fs_explicit_false")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
resource "powerstore_filesystem_snapshot" "test_explicit_false" {
  name                 = "%s"
  description          = "Explicitly non-secure FS snapshot"
  filesystem_id        = "%s"
  expiration_timestamp = "%s"
  access_type          = "Snapshot"
  is_secure            = false
}
`, snapName, fsID, expTS),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_filesystem_snapshot.test_explicit_false", "is_secure", "false"),
				),
			},
		},
	})
}
