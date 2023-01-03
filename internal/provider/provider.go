package provider

import (
	"context"
	"log"
	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure PowerStorPprovider satsisfies various provider framework
var _ provider.Provider = &PowerStore{}

// PowerStore defines the provider implementaion
type PowerStore struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
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

// ProviderData describes the provider data model.
type ProviderData struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func (p *PowerStore) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	log.Printf("mayank metadata")
	resp.TypeName = "powerstore"
	resp.Version = p.version
}

func (p *PowerStore) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	log.Printf("mayank schema")
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider for PowerStore",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "IP or FQDN of the PowerStore host",
				Description:         "IP or FQDN of the PowerStore host",
				Required:            true,
				Validators: []validator.String{
					validators.UrlString{},
				},
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
		},
	}
}

func (p *PowerStore) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	log.Printf("mayank configure")

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
	)
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

func (p *PowerStore) Resources(ctx context.Context) []func() resource.Resource {
	log.Printf("mayank Resources")
	return []func() resource.Resource{}
}

func (p *PowerStore) DataSources(ctx context.Context) []func() datasource.DataSource {
	log.Printf("mayank DataSources")
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	log.Printf("mayank New")
	return func() provider.Provider {
		return &PowerStore{
			version: version,
		}
	}
}
