package powerstore

import (
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

func init() {
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
