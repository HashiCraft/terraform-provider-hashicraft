package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBot_basic(t *testing.T) {
    rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders(),
        CheckDestroy: testAccCheckBotResourceDestroy,
        Steps: []resource.TestStep{
          {
            Config: testAccBotBasicResource(rName),
            Check: resource.ComposeTestCheckFunc(
              testAccCheckBotResourceExists("hashicraft_bot." + rName),
            ),
          },
        },
    })
}

func testAccProviders() map[string]*schema.Provider{
return map[string]*schema.Provider {
			"hashicraft": Provider(),
		}
}

func testAccPreCheck(t *testing.T) {
  if os.Getenv("MINECRAFT_USER") == "" ||
    os.Getenv("MINECRAFT_PASSWORD") == "" ||
    os.Getenv("MINECRAFT_HOST") == "" {
    t.Fatal("Environment variables for the minecraft server need to be set")
  }
}

func testAccCheckBotResourceDestroy(s *terraform.State) error {
  return nil
}

func testAccCheckBotResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
    return nil
  }
}

func testAccBotBasicResource(id string) string {
	return fmt.Sprintf(`
provider "hashicraft" {}

resource "hashicraft_bot" "%s" {
  username = "%s"
  password = "%s"
  server = "%s"
}
  `,
  id,
  os.Getenv("MINECRAFT_USER"),
  os.Getenv("MINECRAFT_PASSWORD"),
  os.Getenv("MINECRAFT_HOST"),
  )
}
