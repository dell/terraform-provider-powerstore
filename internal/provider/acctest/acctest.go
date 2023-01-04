package acctest

import (
	"fmt"
	"os"
	"terraform-provider-powerstore/internal/provider"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testProvider = provider.New("test")()

var testProviderFactory = map[string]func() (tfprotov6.ProviderServer, error){
	"powerstore": providerserver.NewProtocol6WithError(testProvider),
}

var (
	hostID        string
	hostGroupID   string
	volumeGroupID string
	hostName      string
	hostGroupName string

	providerConfigForTesting string
)

func init() {
	var endpoint = os.Getenv("POWERSTORE_ENDPOINT")
	var username = os.Getenv("POWERSTORE_USERNAME")
	var password = os.Getenv("POWERSTORE_PASSWORD")

	hostID = os.Getenv("HOST_ID")
	hostGroupID = os.Getenv("HOST_GROUP_ID")
	volumeGroupID = os.Getenv("VOLUME_GROUP_ID")
	hostName = os.Getenv("HOST_NAME")
	hostGroupName = os.Getenv("HOST_GROUP_NAME")

	providerConfigForTesting = fmt.Sprintf(`
		provider "powerstore" {
			username = "%s"
			password = "%s"
			endpoint = "%s"
			insecure = true
		}
	`, username, password, endpoint)
}

func testAccPreCheck(t *testing.T) {

	mustEnvVariables := []string{
		"POWERSTORE_ENDPOINT",
		"POWERSTORE_USERNAME",
		"POWERSTORE_PASSWORD",
		"HOST_ID",
		"HOST_GROUP_ID",
		"VOLUME_GROUP_ID",
		"HOST_NAME",
		"HOST_GROUP_NAME",
	}

	for _, envVar := range mustEnvVariables {

		if v := os.Getenv(envVar); v == "" {
			t.Fatal(envVar, " must be set for acceptance tests")
		}
	}
}
