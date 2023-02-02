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

resource "powerstore_storagecontainer" "test1" {
  name = "scterraform1"
  quota = 0
  storage_protocol = "NVMe"
  high_water_mark = 100
}

//Below example is for import operation
/*resource "powerstore_storagecontainer" "terraform-provider-test-import" {
}*/
