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
      version = "5.0.1"
    }
  }
}

provider "heroku" {
}
