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
	"context"
	client "terraform-provider-powerstore/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ provider.Provider = &Pstoreprovider{}

// Pstoreprovider satisfies the tfsdk.Provider interface and usually is included
// // with all Resource and DataSource implementations.
type Pstoreprovider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	client *client.Client

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ProviderData can be used to store data from the Terraform configuration.
type ProviderData struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
	Timeout  types.Int64  `tfsdk:"timeout"`
}

// Metadata defines provider interface Metadata method
func (p *Pstoreprovider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "powerstore"
	resp.Version = p.version
}

// Schema defines provider interface Schema method
func (p *Pstoreprovider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider for PowerStore",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "IP or FQDN of the PowerStore host",
				Description:         "IP or FQDN of the PowerStore host",
				Required:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Boolean variable to specify whether to validate SSL certificate or not.",
				Description:         "Boolean variable to specify whether to validate SSL certificate or not.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The password of the PowerStore host.",
				Description:         "The password of the PowerStore host.",
				Required:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "The username of the PowerStore host.",
				Description:         "The username of the PowerStore host.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: "The default timeout value for the Powerstore host.",
				Description:         "The default timeout value for the Powerstore host.",
				Optional:            true,
			},
		},
	}
}

// Configure defines provider interface Configure method
func (p *Pstoreprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	config := ProviderData{}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initializing powerstore client
	pstoreClient, err := client.NewClient(
		config.Endpoint.ValueString(),
		config.Username.ValueString(),
		config.Password.ValueString(),
		// as false is default value, so even if insecure parameter is not provided
		// value will be false
		config.Insecure.ValueBool(),
		config.Timeout.ValueInt64(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create powerstore client",
			"Unable to authenticate user for authenticated powerstore client",
		)
		return
	}

	p.client = pstoreClient
	resp.ResourceData = pstoreClient
	resp.DataSourceData = pstoreClient
}

// Resources defines provider interface Resources method
func (p *Pstoreprovider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newVolumeResource,
		newSnapshotRuleResource,
		newStorageContainerResource,
		newProtectionPolicyResource,
		newVolumeGroupResource,
		newHostResource,
		newHostGroupResource,
		newVGSnapshotResource,
		newVolumeSnapshotResource,
	}
}

// DataSources defines provider interface DataSources method
func (p *Pstoreprovider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newVolumeDataSource,
		newSnapshotRuleDataSource,
		newProtectionPolicyDataSource,
		newVolumeGroupDataSource,
		newHostDataSource,
		newHostGroupDataSource,
		newVolumeGroupSnapshotDataSource,
		newVolumeSnapshotDataSource,
	}
}

// New returns instance of provider
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Pstoreprovider{
			version: version,
		}
	}
}
