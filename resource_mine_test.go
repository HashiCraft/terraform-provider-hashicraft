package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMine_basic(t *testing.T) {
    rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)

    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders(),
        CheckDestroy: testAccCheckMineResourceDestroy,
        Steps: []resource.TestStep{
          {
            Config: testAccMineBasicResource(rName),
            Check: resource.ComposeTestCheckFunc(
              testAccCheckBotResourceExists("hashicraft_bot." + rName),
              testAccCheckBotResourceExists("hashicraft_mine." + rName),
            ),
          },
        },
    })
}

func testAccCheckMineResourceDestroy(s *terraform.State) error {
  return nil
}

func testAccCheckMineResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
    return nil
  }
}

func testAccMineBasicResource(id string) string {
	return fmt.Sprintf(`
%s

resource "hashicraft_mine" "%s" {
  bot = hashicraft_bot.%s.id

  start {
    x = 123
    y = 87
    z = 98
  }
  
  end {
    x = 123
    y = 87
    z = 98
  }
  
 tools {
    x = 123
    y = 87
    z = 98
  }
 
  drop {
    x = 123
    y = 87
    z = 98
  }
}
  `,
  testAccBotBasicResource(id),
  id,
  id,
  )
}

