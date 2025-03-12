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

	"github.com/bytedance/mockey"
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
var nasServerID = setDefault(os.Getenv("NAS_SERVER_ID"), "tfacc_nas_server_id")
var nasServerName = setDefault(os.Getenv("NAS_SERVER_NAME"), "tfacc_nas")
var remoteSystemID = setDefault(os.Getenv("REMOTE_SYSTEM_ID"), "db11abb3-789e-47f9-96b5-84b5374cbcd2")
var FunctionMocker *mockey.Mocker

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

	// Before each test clear out the mocker
	if FunctionMocker != nil {
		FunctionMocker.UnPatch()
	}
}

func setDefault(osInput string, defaultStr string) string {
	if osInput == "" {
		return defaultStr
	}
	return osInput
}
