package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicraft/terraform-bot-provider/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", nil),
			},
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HOST", nil),
			},
		},
    ResourcesMap: map[string]*schema.Resource{
      "hashicraft_bot": resourceBot(),
      "hashicraft_mine": resourceMine(),
    },
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
  var diags diag.Diagnostics

  host := d.Get("host").(string)
  key := d.Get("api_key").(string)

  if(host == "") {
    diags := append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create client",
				Detail:   "Please specify the host location for the bot server",
			})

      return nil, diags
  }

  if(key == "") {
    diags := append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create client",
				Detail:   "Please specify the API key for the bot server",
			})

      return nil, diags
  }

  c:= client.New(host,key)

	return c, nil
}
