package powerstore

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// var testAccProviders map[string]func() tfsdk.Provider
var testProvider tfsdk.Provider
var testProviderFactory map[string]func() (tfprotov6.ProviderServer, error)
var endpoint = os.Getenv("POWERSTORE_ENDPOINT")
var username = os.Getenv("POWERSTORE_USERNAME")
var password = os.Getenv("POWERSTORE_PASSWORD")

func init() {
	testProvider = New("test")()
	testProvider.Configure(context.Background(), tfsdk.ConfigureProviderRequest{}, &tfsdk.ConfigureProviderResponse{})
	testProviderFactory = map[string]func() (tfprotov6.ProviderServer, error){
		"powerstore": providerserver.NewProtocol6WithError(testProvider),
	}
}

//lint:ignore U1000 used by the internal provider, to be checked
func testAccProvider(t *testing.T, p tfsdk.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providers := p.(*Pstoreprovider)
		if !providers.configured {
			return fmt.Errorf("provider not configured")
		}

		if providers.client.PStoreClient == nil {
			return fmt.Errorf("provider not configured")
		}
		return nil
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

const EmptyEndpointConfig = `
provider "powerstore" {
	username = "username"
	password = "password"
}
`
