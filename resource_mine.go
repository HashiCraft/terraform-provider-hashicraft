package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicraft/terraform-bot-provider/client"
)

var coord = map[string]*schema.Schema{
	"x": {
		Description: "The X position",
		Type:        schema.TypeInt,
		Required: true,
	},
	"y": {
		Description: "The Y position",
		Type:        schema.TypeInt,
		Required: true,
	},
	"z": {
		Description: "The Z position",
		Type:        schema.TypeInt,
		Required: true,
	},
}

func resourceMine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMineCreate,
		ReadContext:   resourceMineRead,
		UpdateContext: resourceMineUpdate,
		DeleteContext: resourceMineDelete,
		Schema: map[string]*schema.Schema{
	    "bot": {
        Description: "The id of the bot resource",
        Type:        schema.TypeString,
        Required: true,
	    },
			"start": {
				Description: "The start position of the mining area.",
				Elem: &schema.Resource{
					Schema: coord,
				},
        Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
			},
			"end": {
				Description: "The end position for the mining area.",
				Elem: &schema.Resource{
					Schema: coord,
				},
        Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
			},
			"tools": {
				Description: "The location of the chest containing the mining tools",
				Elem: &schema.Resource{
					Schema: coord,
				},
        Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
			},
			"drop": {
				Description: "The location of the chest to drop mined resources.",
				Elem: &schema.Resource{
					Schema: coord,
				},
        Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
			},
		},
	}
}

func resourceMineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  config := client.ConfigRequest{}

  start := d.Get("start").([]interface{})
  config.MineStart = fmt.Sprintf(
    "%d,%d,%d",
    start[0].(map[string]interface{})["x"].(int),
    start[0].(map[string]interface{})["y"].(int),
    start[0].(map[string]interface{})["z"].(int),
  )

  end := d.Get("end").([]interface{})
  config.MineEnd = fmt.Sprintf(
    "%d,%d,%d",
    end[0].(map[string]interface{})["x"].(int),
    end[0].(map[string]interface{})["y"].(int),
    end[0].(map[string]interface{})["z"].(int),
  )

  tools := d.Get("tools").([]interface{})
  config.ToolChest = fmt.Sprintf(
    "%d,%d,%d",
    tools[0].(map[string]interface{})["x"].(int),
    tools[0].(map[string]interface{})["y"].(int),
    tools[0].(map[string]interface{})["z"].(int),
  )

  drop := d.Get("drop").([]interface{})
  config.DropChest = fmt.Sprintf(
    "%d,%d,%d",
    drop[0].(map[string]interface{})["x"].(int),
    drop[0].(map[string]interface{})["y"].(int),
    drop[0].(map[string]interface{})["z"].(int),
  )

  id := d.Get("bot").(string)

  c := m.(*client.Bot)
  err := c.Configure(id, config)
  if err != nil {
    diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to configure bot",
				Detail:   err.Error(),
			})

      return diags
  }

  err = c.Start(id)
  if err != nil {
    diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to start bot",
				Detail:   err.Error(),
			})

      return diags
  }

  d.SetId(id)

	return diags
}

func resourceMineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
	return diags
}

func resourceMineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
	return diags
}

func resourceMineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
	return diags
}
