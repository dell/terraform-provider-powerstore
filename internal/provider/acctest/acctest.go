package acctest

import (
	"fmt"
	"os"
	"terraform-provider-powerstore/internal/provider"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var TestProvider = provider.New("test")()

var TestProviderFactory = map[string]func() (tfprotov6.ProviderServer, error){
	"powerstore": providerserver.NewProtocol6WithError(TestProvider),
}

var (
	HostID        string
	HostGroupID   string
	VolumeGroupID string
	HostName      string
	HostGroupName string

	ProviderConfigForTesting string
)

func init() {
	var endpoint = os.Getenv("POWERSTORE_ENDPOINT")
	var username = os.Getenv("POWERSTORE_USERNAME")
	var password = os.Getenv("POWERSTORE_PASSWORD")

	HostID = os.Getenv("HOST_ID")
	HostGroupID = os.Getenv("HOST_GROUP_ID")
	VolumeGroupID = os.Getenv("VOLUME_GROUP_ID")
	HostName = os.Getenv("HOST_NAME")
	HostGroupName = os.Getenv("HOST_GROUP_NAME")

	ProviderConfigForTesting = fmt.Sprintf(`
		provider "powerstore" {
			username = "%s"
			password = "%s"
			endpoint = "%s"
			insecure = true
		}
	`, username, password, endpoint)
}

func TestAccPreCheck(t *testing.T) {

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
