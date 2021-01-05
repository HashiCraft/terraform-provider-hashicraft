package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicraft/terraform-bot-provider/client"
)

func resourceBot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBotCreate,
		ReadContext:   resourceBotRead,
		UpdateContext: resourceBotUpdate,
		DeleteContext: resourceBotDelete,
		Schema: map[string]*schema.Schema{
			"server": {
		    Description: "The location of the minecraft server",
		    Type:        schema.TypeString,
		    Required: true,
      },
			"port": {
		    Description: "The port for the minecraft server",
		    Type:        schema.TypeInt,
        Optional: true,
      },
			"username": {
		    Description: "The username of the mojang account the bot uses",
		    Type:        schema.TypeString,
		    Required: true,
      },
			"password": {
		    Description: "The password of the mojang account the bot uses",
		    Type:        schema.TypeString,
		    Required: true,
      },
    },
  }
}

func resourceBotCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  host := d.Get("server").(string)
  username := d.Get("username").(string)
  password := d.Get("password").(string)
  port := d.Get("port").(int)

  if port == 0 {
    port = 25565
  }

  config := client.NewRequest{
    Host:host,
    Port: port,
    Username: username,
    Password:password,
  }

  c := m.(*client.Bot)
  id, err := c.New(config)
  if err != nil {
    diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to configure bot",
				Detail:   err.Error(),
			})
  }

  d.SetId(id)

	return diags
}

func resourceBotRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
	return diags
}

func resourceBotUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
	return diags
}

func resourceBotDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  id := d.Id()
  c := m.(*client.Bot)
  err := c.Delete(id)
  if err != nil {
    diags = append(diags, diag.Diagnostic{
      Severity: diag.Error,
      Summary:  "Unable to delete bot",
      Detail:   err.Error(),
    })
  }

	return diags
}
