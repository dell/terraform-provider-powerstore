package provider

import (
	"context"
	"terraform-provider-powerstore/internal/powerstore"
	"terraform-provider-powerstore/internal/provider/resources/protectionpolicy"
	"terraform-provider-powerstore/internal/provider/resources/snapshotrule"
	"terraform-provider-powerstore/internal/provider/resources/storagecontainer"
	"terraform-provider-powerstore/internal/provider/resources/volume"
	"terraform-provider-powerstore/internal/provider/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure PowerStore satsisfies provider.Provider interface
var _ provider.Provider = &PowerStore{}

// PowerStore defines the provider implementaion
type PowerStore struct {

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string

	// currently being used only in test case of volume resource
	Client *powerstore.Client
}

type model struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

// Metadata defines provider interface Metadata method
func (p *PowerStore) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "powerstore"
	resp.Version = p.version
}

// Schema defines provider interface Schema method
func (p *PowerStore) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider for PowerStore",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "IP or FQDN of the PowerStore host",
				Description:         "IP or FQDN of the PowerStore host",
				Required:            true,
				Validators: []validator.String{
					validators.URLString{},
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

// Configure defines provider interface Configure method
func (p *PowerStore) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	config := model{}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// initializing powerstore client
	pstoreClient, err := powerstore.NewClient(
		config.Endpoint.ValueString(),
		config.Username.ValueString(),
		config.Password.ValueString(),
		// as false is default value, so even if insecure parameter is not provided
		// value will be false
		config.Insecure.ValueBool(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create powerstore client",
			"Unable to authenticate user for authenticated powerstore client",
		)
		return
	}

	p.Client = pstoreClient
	resp.ResourceData = pstoreClient
	resp.DataSourceData = pstoreClient
}

// Resources defines provider interface Resources method
func (p *PowerStore) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		snapshotrule.NewResource,
		storagecontainer.NewResource,
		protectionpolicy.NewResource,
		volume.NewResource,
	}
}

// DataSources defines provider interface DataSources method
func (p *PowerStore) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// New returns instance of provider
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &PowerStore{
			version: version,
		}
	}
}
