package curl

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		// Inputs for authenticating the provider
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CURL_CLIENT_ID", nil),
			},
			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CURL_CLIENT_SECRET", nil),
			},
			"resource": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CURL_RESOURCE", nil),
			},
			"tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CURL_TENANT_ID", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"curl":     dataSource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	clientId := d.Get("client_id").(string)
	secret := d.Get("secret").(string)
	tenantId := d.Get("tenant_id").(string)
	resource := d.Get("resource").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	opts := ApiClientOpts{
		insecure:    false,
		timeout:     0,
	}

	if (clientId != "") && (secret != "") && (tenantId != "") && (resource != "") {
		log.Printf("[INFO] ******* Creating OAuthOpts")
		oauthOpts := OAuthOpts{
			ClientId:     clientId,
			ClientSecret: secret,
			Resource:     resource,
			TenantId:     tenantId,
		}
		opts.oauthConfig = oauthOpts

		c, err := NewClient(opts)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := NewClient(opts)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}

