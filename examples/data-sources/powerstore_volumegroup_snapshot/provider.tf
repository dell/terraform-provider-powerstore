terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source  = "registry.terraform.io/dell/powerstore"
    }
  }
}

provider "powerstore" {
  username = var.username
  password = var.password
  endpoint = var.endpoint
  insecure = true
  timeout  = var.timeout
}
