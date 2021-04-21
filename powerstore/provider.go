package powerstore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("POWERSTORE_HOST", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("POWERSTORE_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("POWERSTORE_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"powerstore_volume":            resourceVolume(),
			"powerstore_snapshot_rule":     resourceSnapshotRule(),
			"powerstore_storage_container": resourceStorageContainer(),
			"powerstore_protection_policy": resourceProtectionPolicy(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//getting username , password and host url from terraform scripts (.tf files)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	hostURL := d.Get("host").(string)

	if (username != "") && (password != "") && (hostURL != "") {
		//calling new client to create Client struct for powerstore appliance
		c, err := NewClient(&hostURL, &username, &password)
		if err != nil {

			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create powerstore client",
				Detail:   "Unable to authenticate user for authenticated powerstore client",
			})
			return nil, diags
		}

		return c, diags

	}
	return nil, diags
}
