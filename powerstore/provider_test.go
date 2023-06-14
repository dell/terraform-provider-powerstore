/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

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
var endpoint = os.Getenv("POWERSTORE_ENDPOINT")
var username = os.Getenv("POWERSTORE_USERNAME")
var password = os.Getenv("POWERSTORE_PASSWORD")
var hostID = os.Getenv("HOST_ID")
var hostID1 = os.Getenv("HOST_ID1")
var hostID2 = os.Getenv("HOST_ID2")
var hostID3 = os.Getenv("HOST_ID3")
var hostGroupID = os.Getenv("HOST_GROUP_ID")
var volumeGroupID = os.Getenv("VOLUME_GROUP_ID")
var hostName = os.Getenv("HOST_NAME")
var hostGroupName = os.Getenv("HOST_GROUP_NAME")
var volumeID = os.Getenv("VOLUME_ID")
var volumeName = os.Getenv("VOLUME_NAME")
var snapshotRuleID = os.Getenv("SNAPSHOT_RULE_ID")
var replicationRuleID = os.Getenv("REPLICATION_RULE_ID")
var snapshotRuleName = os.Getenv("SNAPSHOT_RULE_NAME")
var replicationRuleName = os.Getenv("REPLICATION_RULE_NAME")
var policyName = os.Getenv("PROTECTION_POLICY_NAME")
var policyID = os.Getenv("PROTECTION_POLICY_ID")
var volumeGroupName = os.Getenv("VOLUME_GROUP_NAME")
var volumeGroupSnapshotName = os.Getenv("VOLUME_GROUP_SNAPSHOT_NAME")
var volumeGroupSnapshotID = os.Getenv("VOLUME_GROUP_SNAPSHOT_ID")
var volumeSnapshotID = os.Getenv("VOLUME_SNAPSHOT_ID")
var volumeSnapshotName = os.Getenv("VOLUME_SNAPSHOT_NAME")

var ProviderConfigForTesting = ``

func init() {
	username := os.Getenv("POWERSTORE_USERNAME")
	password := os.Getenv("POWERSTORE_PASSWORD")
	endpoint := os.Getenv("POWERSTORE_ENDPOINT")
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
	if v := os.Getenv("POWERSTORE_USERNAME"); v == "" {
		t.Fatal("POWERSTORE_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("POWERSTORE_PASSWORD"); v == "" {
		t.Fatal("POWERSTORE_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("POWERSTORE_ENDPOINT"); v == "" {
		t.Fatal("POWERSTORE_ENDPOINT must be set for acceptance tests")
	}

}
