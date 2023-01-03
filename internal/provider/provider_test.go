package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testProviderFactory = map[string]func() (tfprotov6.ProviderServer, error){
	"powerstore": providerserver.NewProtocol6WithError(New("test")()),
}

var endpoint = os.Getenv("POWERSTORE_ENDPOINT")
var username = os.Getenv("POWERSTORE_USERNAME")
var password = os.Getenv("POWERSTORE_PASSWORD")
var hostID = os.Getenv("HOST_ID")
var hostGroupID = os.Getenv("HOST_GROUP_ID")
var volumeGroupID = os.Getenv("VOLUME_GROUP_ID")
var hostName = os.Getenv("HOST_NAME")
var hostGroupName = os.Getenv("HOST_GROUP_NAME")

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
