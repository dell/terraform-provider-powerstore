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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// var testAccProviders map[string]func() tfsdk.Provider
var testProvider provider.Provider
var testProviderFactory map[string]func() (tfprotov6.ProviderServer, error)

var endpoint = setDefault(os.Getenv("POWERSTORE_ENDPOINT"), "http://localhost:3003/api/rest")
var username = setDefault(os.Getenv("POWERSTORE_USERNAME"), "test")
var password = setDefault(os.Getenv("POWERSTORE_PASSWORD"), "test")
var hostID = setDefault(os.Getenv("HOST_ID"), "tfacc_host_id")
var hostIDRead = setDefault(os.Getenv("HOST_ID_READ"), "tfacc_host_id")
var hostGroupID = setDefault(os.Getenv("HOST_GROUP_ID"), "tfacc_host_group_id")
var volumeGroupID = setDefault(os.Getenv("VOLUME_GROUP_ID"), "tfacc_volume_group_id")
var hostName = setDefault(os.Getenv("HOST_NAME"), "tfacc_host_name")
var hostNameRead = setDefault(os.Getenv("HOST_NAME_READ"), "tfacc_host_name")
var hostGroupName = setDefault(os.Getenv("HOST_GROUP_NAME"), "tfacc_host_group_name")
var volumeID = setDefault(os.Getenv("VOLUME_ID"), "tfacc_volume_id")
var volumeName = setDefault(os.Getenv("VOLUME_NAME"), "tfacc_volume_name")
var snapshotRuleID = setDefault(os.Getenv("SNAPSHOT_RULE_ID"), "tfacc_snapshot_rule_id")
var replicationRuleID = setDefault(os.Getenv("REPLICATION_RULE_ID"), "tfacc_replication_rule_id")
var snapshotRuleName = setDefault(os.Getenv("SNAPSHOT_RULE_NAME"), "tfacc_snapshot_rule_name")
var replicationRuleName = setDefault(os.Getenv("REPLICATION_RULE_NAME"), "tfacc_replication_rule_name")
var policyName = setDefault(os.Getenv("PROTECTION_POLICY_NAME"), "tfacc_policy_name")
var policyID = setDefault(os.Getenv("PROTECTION_POLICY_ID"), "tfacc_policy_id")
var volumeGroupName = setDefault(os.Getenv("VOLUME_GROUP_NAME"), "tfacc_volume_group_name")
var volumeGroupSnapshotName = setDefault(os.Getenv("VOLUME_GROUP_SNAPSHOT_NAME"), "tfacc_volume_group_snapshot_name")
var volumeGroupSnapshotID = setDefault(os.Getenv("VOLUME_GROUP_SNAPSHOT_ID"), "tfacc_volume_group_snapshot_id")
var volumeSnapshotID = setDefault(os.Getenv("VOLUME_SNAPSHOT_ID"), "tfacc_volume_snapshot_id")
var volumeSnapshotName = setDefault(os.Getenv("VOLUME_SNAPSHOT_NAME"), "tfacc_volume_snapshot_name")
var nasServerID = setDefault(os.Getenv("NAS_SERVER_ID"), "tfacc_nas_server_id")
var fileSystemID = setDefault(os.Getenv("FILESYSTEM_ID"), "tfacc_filesystem_id")
var remoteSystemID = setDefault(os.Getenv("REMOTE_SYSTEM_ID"), "")
var fileSystemSnapshotID = setDefault(os.Getenv("FILESYSTEM_SNAPSHOT_ID"), "tfacc_filesystem_snapshot_id")
var fileSystemSnapshotName = setDefault(os.Getenv("FILESYSTEM_SNAPSHOT_NAME"), "tfacc_filesystem_snapshot_name")

var ProviderConfigForTesting = ``

func init() {

	username := username
	password := password
	endpoint := endpoint
	insecure := "true"

	ProviderConfigForTesting = fmt.Sprintf(`
		provider "powerstore" {
			username = "%s"
			password = "%s"
			endpoint = "%s"
			insecure = "%s"
		}
	`, username, password, endpoint, insecure)

	testProvider = New("test")()
	testProviderFactory = map[string]func() (tfprotov6.ProviderServer, error){
		"powerstore": providerserver.NewProtocol6WithError(testProvider),
	}
}

func testAccPreCheck(t *testing.T) {
	if v := username; v == "" {
		t.Fatal("POWERSTORE_USERNAME must be set for acceptance tests")
	}

	if v := password; v == "" {
		t.Fatal("POWERSTORE_PASSWORD must be set for acceptance tests")
	}

	if v := endpoint; v == "" {
		t.Fatal("POWERSTORE_ENDPOINT must be set for acceptance tests")
	}

}

func setDefault(osInput string, defaultStr string) string {
	if osInput == "" {
		return defaultStr
	}
	return osInput
}
