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

// --- Config generators for secure VG snapshot tests ---
// The dedicated volume + VG are created/managed by the SDK in setupSecureVGAndProbe.
// These functions accept the VG ID directly so Terraform only manages the snapshot.

func secureVGSnapConfig(vgID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volumegroup_snapshot" "test_secure" {
  name                 = "%s"
  description          = "Secure VG Snapshot"
  volume_group_id      = "%s"
  expiration_timestamp = "%s"
  is_secure            = true
}
`, snapName, vgID, expTS)
}

func secureVGSnapRevertConfig(vgID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volumegroup_snapshot" "test_secure" {
  name                 = "%s"
  description          = "Secure VG Snapshot"
  volume_group_id      = "%s"
  expiration_timestamp = "%s"
  is_secure            = false
}
`, snapName, vgID, expTS)
}

func nonsecureVGSnapConfig(vgID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volumegroup_snapshot" "test_nonsecure" {
  name                 = "%s"
  description          = "Non-Secure VG Snapshot"
  volume_group_id      = "%s"
  expiration_timestamp = "%s"
}
`, snapName, vgID, expTS)
}

func nonsecureVGSnapUpdateToSecureConfig(vgID, expTS, snapName string) string {
	return fmt.Sprintf(`
resource "powerstore_volumegroup_snapshot" "test_nonsecure" {
  name                 = "%s"
  description          = "Non-Secure VG Snapshot"
  volume_group_id      = "%s"
  expiration_timestamp = "%s"
  is_secure            = true
}
`, snapName, vgID, expTS)
}

// setupSecureVGAndProbe creates the dedicated volume + VG via SDK, probes
// the array clock via a volume snapshot, and returns (vgID, expTS, cleanup).
func setupSecureVGAndProbe(t *testing.T) (string, string, func()) {
	t.Helper()
	client, err := secureTestSDKClient()
	if err != nil {
		t.Fatalf("failed to create SDK client: %v", err)
	}

	// Create the dedicated volume
	volName := "ZZZ_AT_tf_secure_vg_vol"
	var volID string
	vol, err := client.GetVolumeByName(context.Background(), volName)
	if err == nil && vol.ID != "" {
		volID = vol.ID
	} else {
		size := int64(1048576)
		createResp, err := client.CreateVolume(context.Background(), &gopowerstore.VolumeCreate{
			Name: &volName,
			Size: &size,
		})
		if err != nil {
			t.Fatalf("failed to create dedicated secure VG test volume: %v", err)
		}
		volID = createResp.ID
	}

	// Create the dedicated volume group
	vgName := "ZZZ_AT_tf_secure_vg"
	var vgID string
	vg, err := client.GetVolumeGroupByName(context.Background(), vgName)
	if err == nil && vg.ID != "" {
		vgID = vg.ID
	} else {
		createVGResp, err := client.CreateVolumeGroup(context.Background(), &gopowerstore.VolumeGroupCreate{
			Name:      vgName,
			VolumeIDs: []string{volID},
		})
		if err != nil {
			t.Fatalf("failed to create dedicated secure VG: %v", err)
		}
		vgID = createVGResp.ID
	}

	// Probe the array clock via a volume snapshot
	expTS, err := probeArrayClockViaSnapshot(client, volID)
	if err != nil {
		t.Fatalf("failed to probe array clock: %v", err)
	}

	cleanup := func() {
		vg, err := client.GetVolumeGroupByName(context.Background(), vgName)
		if err == nil && vg.ID != "" {
			_, _ = client.DeleteVolumeGroup(context.Background(), vg.ID)
		}
		_, _ = client.DeleteVolume(context.Background(), nil, volID)
	}

	return vgID, expTS, cleanup
}

// Test U-16: Create secure volume group snapshot
func TestAccVolumeGroupSnapshot_CreateSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_create_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "name", snapName),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "expiration_timestamp", expTS),
				),
			},
		},
	})
}

// Test U-17: Create secure VG snapshot without expiration - should fail
func TestAccVolumeGroupSnapshot_CreateSecureMissingExpiration(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + `
resource "powerstore_volumegroup_snapshot" "test_secure_no_exp" {
  name            = "ZZZ_AT_tf_vg_snap_secure_no_exp"
  description     = "Secure VG Snapshot Without Expiration"
  volume_group_id = "fake-vg-id-for-validation"
  is_secure       = true
}
`,
				ExpectError: regexp.MustCompile(".*Secure snapshots require.*expiration.*"),
			},
		},
	})
}

// Test U-18: Create non-secure VG snapshot - is_secure should default to false
func TestAccVolumeGroupSnapshot_CreateNonSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_create_nonsec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + nonsecureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_nonsecure", "name", snapName),
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
		},
	})
}

// Test U-19: Update VG snapshot - mark as secure (one-way lock)
func TestAccVolumeGroupSnapshot_UpdateToSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_upd_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + nonsecureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_nonsecure", "is_secure", "false"),
				),
			},
			{
				Config: ProviderConfigForTesting + nonsecureVGSnapUpdateToSecureConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_nonsecure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test U-20: Reject unsecure VG snapshot (one-way lock)
func TestAccVolumeGroupSnapshot_RejectUnsecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_reject")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:      ProviderConfigForTesting + secureVGSnapRevertConfig(vgID, expTS, snapName),
				ExpectError: regexp.MustCompile(".*one-way lock.*|.*cannot.*changed from true to false.*"),
			},
		},
	})
}

// Test U-21: Import secure VG snapshot
func TestAccVolumeGroupSnapshot_ImportSecure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_import_sec")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
				),
			},
			{
				Config:       ProviderConfigForTesting + secureVGSnapConfig(vgID, expTS, snapName),
				ResourceName: "powerstore_volumegroup_snapshot.test_secure",
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

// Test I-07: Idempotent re-apply of secure VG snapshot
func TestAccVolumeGroupSnapshot_SecureIdempotent(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_idempotent")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + secureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
				),
			},
			// Re-apply same config - should produce no changes
			{
				Config: ProviderConfigForTesting + secureVGSnapConfig(vgID, expTS, snapName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_secure", "is_secure", "true"),
				),
			},
		},
	})
}

// Test E-08: Explicit is_secure=false for VG snapshot
func TestAccVolumeGroupSnapshot_ExplicitFalse(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	vgID, expTS, cleanup := setupSecureVGAndProbe(t)
	defer cleanup()

	snapName := uniqueSnapName("vg_explicit_false")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		CheckDestroy:             secureSnapshotCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + fmt.Sprintf(`
resource "powerstore_volumegroup_snapshot" "test_explicit_false" {
  name                 = "%s"
  description          = "Explicitly non-secure VG snapshot"
  volume_group_id      = "%s"
  expiration_timestamp = "%s"
  is_secure            = false
}
`, snapName, vgID, expTS),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_volumegroup_snapshot.test_explicit_false", "is_secure", "false"),
				),
			},
		},
	})
}
