terraform {
  cloud {
    organization = "ChrisWilding"

    workspaces {
      name = "go-base"
    }
  }

  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "5.1.8"
    }
  }
}

provider "heroku" {
}
