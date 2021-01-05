terraform {
  required_providers {
    hashicraft = {
      source = "local/hashicraft/hashicraft"
      version = "0.1.0"
    }
  }
}

provider "hashicraft" {}

resource "hashicraft_bot" "nic" {
  username = var.minecraft_username
  password = var.minecraft_password
  server = var.minecraft_server
}

resource "hashicraft_mine" "mine" {
  bot = hashicraft_bot.nic.id

  start {
    x = 300
    y = 5
    z = -137
  }
  
  end {
    x = 255
    y = 5
    z = -144
  }
  
 tools {
    x = 291
    y = 8
    z = -107
  }
 
  drop {
    x = 295
    y = 9
    z = -104
  }
}
