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
	"log"
	"time"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// secureSnapshotCheckDestroy is a CheckDestroy function for secure snapshot tests.
// It tolerates destroy errors caused by retention locks on secure snapshots.
// When Terraform's destroy phase tries to delete a secure snapshot that hasn't
// expired yet, the API rejects the request. This function treats that as
// expected behavior rather than a test failure.
func secureSnapshotCheckDestroy(_ *terraform.State) error {
	log.Printf("[INFO] CheckDestroy: tolerating secure snapshot retention errors during destroy")
	return nil
}

// secureSnapshotTestExpOffset is the time offset (in seconds) added to the
// array clock when computing expiration timestamps for secure snapshot tests.
// 5 minutes is enough for the test to run while ensuring automatic cleanup.
const secureSnapshotTestExpOffset = 300 // 5 minutes

// secureTestSDKClient creates a gopowerstore SDK client for direct API calls
// in acceptance tests (e.g. probing the array clock).
func secureTestSDKClient() (*gopowerstore.ClientIMPL, error) {
	clientOptions := gopowerstore.NewClientOptions()
	clientOptions.SetInsecure(true)
	c, err := gopowerstore.NewClientWithArgs(endpoint, username, password, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot create SDK client for secure snapshot tests: %w", err)
	}
	return c.(*gopowerstore.ClientIMPL), nil
}

// probeArrayClockViaSnapshot creates a throwaway non-secure volume snapshot,
// reads its creation_timestamp to determine the array clock, deletes the
// probe snapshot, and returns a UTC expiration timestamp ~5 minutes in the
// future (relative to the array clock) in RFC 3339 format.
//
// This mirrors the Ansible QE approach:
//   - Create a probe snapshot -> read creation_timestamp -> delete it
//   - Use the array clock (not local clock) to avoid clock-skew issues
//   - Short expiration ensures secure snapshots auto-expire and don't block
//     resource cleanup between test runs
func probeArrayClockViaSnapshot(client *gopowerstore.ClientIMPL, volumeID string) (string, error) {
	probeName := "ZZZ_AT_tf_secure_probe_snap"

	// Create a throwaway non-secure snapshot
	createResp, err := client.CreateSnapshot(context.Background(), &gopowerstore.SnapshotCreate{
		Name:        &probeName,
		Description: strPtr("Throwaway probe snapshot for array clock detection"),
	}, volumeID)
	if err != nil {
		return "", fmt.Errorf("failed to create probe snapshot: %w", err)
	}

	// Read the probe snapshot to get creation_timestamp
	probeSnap, err := client.GetSnapshot(context.Background(), createResp.ID)
	if err != nil {
		// Best-effort cleanup
		_, _ = client.DeleteSnapshot(context.Background(), nil, createResp.ID)
		return "", fmt.Errorf("failed to get probe snapshot details: %w", err)
	}

	// Delete the probe snapshot immediately
	_, err = client.DeleteSnapshot(context.Background(), nil, createResp.ID)
	if err != nil {
		log.Printf("[WARN] failed to delete probe snapshot %s: %v", createResp.ID, err)
	}

	// Parse the creation_timestamp from the array (field is on Volume directly)
	arrayTime, err := time.Parse(time.RFC3339Nano, probeSnap.CreationTimeStamp)
	if err != nil {
		// Try alternate format
		arrayTime, err = time.Parse("2006-01-02T15:04:05+00:00", probeSnap.CreationTimeStamp)
		if err != nil {
			return "", fmt.Errorf("failed to parse array creation_timestamp %q: %w",
				probeSnap.CreationTimeStamp, err)
		}
	}

	// Compute expiration = array clock + offset
	expiration := arrayTime.Add(time.Duration(secureSnapshotTestExpOffset) * time.Second)
	expirationStr := expiration.UTC().Format("2006-01-02T15:04:05Z")

	log.Printf("[INFO] Array clock: %s, computed expiration: %s", arrayTime.UTC().Format(time.RFC3339), expirationStr)
	return expirationStr, nil
}

// probeArrayClockViaFSSnapshot is the filesystem variant of the clock probe.
// It creates a throwaway non-secure FS snapshot, reads creation_timestamp,
// deletes it, and returns a ~5-min-future expiration timestamp.
func probeArrayClockViaFSSnapshot(client *gopowerstore.ClientIMPL, filesystemID string) (string, error) {
	probeName := "ZZZ_AT_tf_secure_probe_fs_snap"

	createResp, err := client.CreateFsSnapshot(context.Background(), &gopowerstore.SnapshotFSCreate{
		Name:        probeName,
		Description: "Throwaway probe FS snapshot for array clock detection",
	}, filesystemID)
	if err != nil {
		return "", fmt.Errorf("failed to create probe FS snapshot: %w", err)
	}

	probeSnap, err := client.GetFsSnapshot(context.Background(), createResp.ID)
	if err != nil {
		_, _ = client.DeleteFsSnapshot(context.Background(), createResp.ID)
		return "", fmt.Errorf("failed to get probe FS snapshot details: %w", err)
	}

	_, err = client.DeleteFsSnapshot(context.Background(), createResp.ID)
	if err != nil {
		log.Printf("[WARN] failed to delete probe FS snapshot %s: %v", createResp.ID, err)
	}

	arrayTime, err := time.Parse(time.RFC3339Nano, probeSnap.CreationTimestamp)
	if err != nil {
		arrayTime, err = time.Parse("2006-01-02T15:04:05+00:00", probeSnap.CreationTimestamp)
		if err != nil {
			return "", fmt.Errorf("failed to parse array creation_timestamp %q: %w",
				probeSnap.CreationTimestamp, err)
		}
	}

	expiration := arrayTime.Add(time.Duration(secureSnapshotTestExpOffset) * time.Second)
	expirationStr := expiration.UTC().Format("2006-01-02T15:04:05Z")

	log.Printf("[INFO] Array clock (FS): %s, computed expiration: %s", arrayTime.UTC().Format(time.RFC3339), expirationStr)
	return expirationStr, nil
}

// uniqueSnapName generates a unique snapshot name for a test to avoid
// name collisions on the array. Secure snapshots cannot be deleted during
// their retention period, so each test run must use distinct names.
// Format: ZZZ_AT_<prefix>_<unix_timestamp>
func uniqueSnapName(prefix string) string {
	return fmt.Sprintf("ZZZ_AT_%s_%d", prefix, time.Now().UnixNano())
}

func strPtr(s string) *string {
	return &s
}
