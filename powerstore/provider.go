package powerstore

import (
	"context"
	"fmt"
	"os"
	"strconv"
	client "terraform-provider-powerstore/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.Provider = &Pstoreprovider{}

// Provider -
type Pstoreprovider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	//
	client client.Client

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

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
}

func (p *Pstoreprovider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// If the upstream provider SDK or HTTP client requires configuration, such
	// as authentication or logging, this is a great opportunity to do so.
	// TODO: Implement client using schema
	config := ProviderData{}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)

	// User must provide a user to the provider
	if config.Username.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.Value == "" {
		config.Username.Value = os.Getenv("UNISPHERE_USERNAME")
	}

	if config.Username.Value == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a user to the provider
	if config.Password.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.Value == "" {
		config.Password.Value = os.Getenv("UNISPHERE_PASSWORD")
	}

	if config.Password.Value == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a user to the provider
	if config.Endpoint.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Endpoint.Value == "" {
		config.Endpoint.Value = os.Getenv("UNISPHERE_HOST")
	}

	if config.Endpoint.Value == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find endpoint",
			"Endpoint cannot be an empty string",
		)
		return
	}

	// User must provide a user to the provider
	if config.Insecure.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure flag",
		)
		return
	}

	if !config.Insecure.Value {
		// config.Insecure.Value
		insecure, _ := strconv.ParseBool(os.Getenv("UNISPHERE_INSECURE"))
		config.Insecure.Value = insecure
	}

	pstoreClient, err := client.NewClient(config.Endpoint.Value, config.Username.Value, config.Password.Value, config.Insecure.Value)
	if err != nil {
		p.configured = false
		resp.Diagnostics.AddError(
			"Unable to create powerstore client",
			"Unable to authenticate user for authenticated powerstore client",
		)
		return
	}

	p.client = *pstoreClient
	p.configured = true
}

func (p *Pstoreprovider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"powerstore_volume": resourceVolumeType{}}, nil
}

func (p *Pstoreprovider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}

func (p *Pstoreprovider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{

		MarkdownDescription: "Provider for PowerMax",
		Attributes: map[string]tfsdk.Attribute{
			"endpoint": {
				MarkdownDescription: "IP or FQDN of the Unisphere host",
				Description:         "IP or FQDN of the Unisphere host",
				Type:                types.StringType,
				Required:            true,
			},
			"insecure": {
				MarkdownDescription: "Boolean variable to specify whether to validate SSL certificate or not.",
				Description:         "Boolean variable to specify whether to validate SSL certificate or not.",
				Type:                types.BoolType,
				Optional:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					DefaultAttribute(types.Bool{Value: false}),
				},
			},
			"password": {
				MarkdownDescription: "The password of the Unisphere host.",
				Description:         "The password of the Unisphere host.",
				Type:                types.StringType,
				Required:            true,
				Sensitive:           true,
			},
			"username": {
				MarkdownDescription: "The username of the Unisphere host.",
				Description:         "The username of the Unisphere host.",
				Type:                types.StringType,
				Required:            true,
			},
		},
	}, nil
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
//lint:ignore U1000 used by the internal provider, to be checked
func convertProviderType(in tfsdk.Provider) (Pstoreprovider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*Pstoreprovider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return Pstoreprovider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return Pstoreprovider{}, diags
	}

	return *p, diags
}

// New accepts version as parameter and returns tfsdk provider
func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &Pstoreprovider{
			version: version,
		}
	}
}

