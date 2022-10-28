terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source = "powerstore.com/powerstoreprovider/powerstore"
    }
  }
}

variable "username" {
  type=string
}

variable "password" {
  type=string
}

variable "endpoint" {
  type=string
}

provider "powerstore" {
  username = "${var.username}"
  password = "${var.password}"
  endpoint = "${var.endpoint}"
  insecure = true
}

resource "powerstore_volume" "test" {
  name = "test_vol2"
  size = 7516192768
  description = "Creating volume"
}
