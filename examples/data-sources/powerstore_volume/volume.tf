terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source = "registry.terraform.io/dell/powerstore"
    }
  }
}

provider "powerstore" {
  username = "${var.username}"
  password = "${var.password}"
  endpoint = "${var.endpoint}"
  insecure = true
}

data "powerstore_volume" "test1" {
        name = "tf_vol"
}
