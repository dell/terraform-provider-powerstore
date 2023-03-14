package powerstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Test to Create Host Resource
func TestAccHost_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_host.test", "name", hostName),
					resource.TestCheckResourceAttr("powerstore_host.test", "description", "Test Host Resource"),
				),
			},
		},
	})
}

// Test to Create Host Resource
func TestAccHost_CreateWithoutPolicy(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostParamsCreateWithoutPolicy,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
			},
		},
	})
}

// Test to Create Host Resource Without Name
func TestAccHost_CreateWithoutName(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + HostParamsCreateWithoutName,
				ExpectError: regexp.MustCompile(CreateResourceMissingErrorMsg),
			},
		},
	})
}

// Negative test case for import
func TestAccHost_ImportFailure(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:        ProviderConfigForTesting + HostParamsCreate,
				ResourceName:  "powerstore_host.test",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(ImportHostDetailErrorMsg),
				ImportStateId: "invalid-id",
			},
		},
	})
}

// Test to import successfully
func TestAccHost_ImportSuccess(t *testing.T) {

	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + HostParamsCreate,
			},
			{
				Config:            ProviderConfigForTesting + HostParamsCreate,
				ResourceName:      "powerstore_host.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: true,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					//assert.Equal(t, hostName, s[0].Attributes["name"])
					assert.Equal(t, "Linux", s[0].Attributes["os_type"])
					return nil
				},
			},
		},
	})

}

var HostParamsCreate = `
resource "powerstore_host" "test" {
	name = "` + hostName + `"
	description = "Test Host Resource"
	os_type = "Linux"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb605eb13", port_type="NVMe"}]
}
`

var HostParamsCreateWithoutPolicy = `
resource "powerstore_host" "test" {
	name = "` + hostName + `"
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb605eb13", port_type="NVMe"}]
	description = "Test Host Resource"
}
`
var HostParamsCreateWithoutName = `
resource "powerstore_host" "test" {
	initiators = [{port_name= "iqn.1994-05.com.redhat:88cb605eb13", port_type="NVMe"}]
	description = "Test Host Resource"
os_type = "Linux"
}
`
