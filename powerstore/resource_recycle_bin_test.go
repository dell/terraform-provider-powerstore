/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

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
	"os"
	"regexp"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"testing"

	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
)

// Test to Create (apply) Recycle Bin Config
func TestAccRecycleBinConfig_Create(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinConfigParams7Days,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_recycle_bin_config.test", "id", "0"),
					resource.TestCheckResourceAttr("powerstore_recycle_bin_config.test", "expiration_duration", "7"),
				),
			},
		},
	})
}

// Test to Update Recycle Bin Config
func TestAccRecycleBinConfig_Update(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinConfigParams7Days,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_recycle_bin_config.test", "expiration_duration", "7"),
				),
			},
			{
				Config: ProviderConfigForTesting + RecycleBinConfigParams14Days,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_recycle_bin_config.test", "expiration_duration", "14"),
				),
			},
		},
	})
}

// Test with expiration_duration = 0 (boundary value)
func TestAccRecycleBinConfig_ZeroDays(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinConfigParams0Days,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_recycle_bin_config.test", "expiration_duration", "0"),
				),
			},
		},
	})
}

// Test with invalid values (out of range)
func TestAccRecycleBinConfig_InvalidValues(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config:      ProviderConfigForTesting + RecycleBinConfigParamsInvalidDuration,
				ExpectError: regexp.MustCompile(`must be between 0 and 30`),
			},
		},
	})
}

// Test to import Recycle Bin Config
func TestAccRecycleBinConfig_ImportSuccess(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Dont run with units tests because it will try to create the context")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testProviderFactory,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfigForTesting + RecycleBinConfigParams7Days,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("powerstore_recycle_bin_config.test", "expiration_duration", "7"),
				),
			},
			{
				Config:            ProviderConfigForTesting + RecycleBinConfigParams7Days,
				ResourceName:      "powerstore_recycle_bin_config.test",
				ImportState:       true,
				ExpectError:       nil,
				ImportStateVerify: false,
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					assert.Equal(t, "0", s[0].Attributes["id"])
					return nil
				},
			},
		},
	})
}

var RecycleBinConfigParams7Days = `
resource "powerstore_recycle_bin_config" "test" {
	expiration_duration = 7
}
`

var RecycleBinConfigParams14Days = `
resource "powerstore_recycle_bin_config" "test" {
	expiration_duration = 14
}
`

var RecycleBinConfigParams0Days = `
resource "powerstore_recycle_bin_config" "test" {
	expiration_duration = 0
}
`

var RecycleBinConfigParamsInvalidDuration = `
resource "powerstore_recycle_bin_config" "test" {
	expiration_duration = 31
}
`

// Unit tests for Configure method
func TestResourceRecycleBin_Configure_InvalidType(t *testing.T) {
	r := &resourceRecycleBin{}
	req := fwresource.ConfigureRequest{
		ProviderData: "invalid_type",
	}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.True(t, resp.Diagnostics.HasError())
	assert.NotEmpty(t, resp.Diagnostics.Errors()[0].Summary)
}

func TestResourceRecycleBin_Configure_Nil(t *testing.T) {
	r := &resourceRecycleBin{}
	req := fwresource.ConfigureRequest{
		ProviderData: nil,
	}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Nil(t, r.client)
}

func TestResourceRecycleBin_Configure_Success(t *testing.T) {
	r := &resourceRecycleBin{}
	c := &client.Client{GenClient: &clientgen.APIClient{}}
	req := fwresource.ConfigureRequest{
		ProviderData: c,
	}
	resp := &fwresource.ConfigureResponse{}

	r.Configure(context.Background(), req, resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.NotNil(t, r.client)
}

// Test recycleBinConfigToState helper function
func TestRecycleBinConfigToState(t *testing.T) {
	duration := int32(7)
	id := "0"
	cfg := &clientgen.RecycleBinConfigInstance{
		Id:                 &id,
		ExpirationDuration: &duration,
	}

	state := recycleBinConfigToState(cfg)
	assert.Equal(t, "0", state.ID.ValueString())
	assert.Equal(t, int32(7), state.ExpirationDuration.ValueInt32())
}

func TestRecycleBinConfigToState_Nil(t *testing.T) {
	cfg := &clientgen.RecycleBinConfigInstance{}

	state := recycleBinConfigToState(cfg)
	assert.Equal(t, "", state.ID.ValueString())
	assert.Equal(t, int32(0), state.ExpirationDuration.ValueInt32())
}
