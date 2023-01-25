terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source = "powerstore.com/powerstoreprovider/powerstore"
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
  quota = 10737418240
  storage_protocol = "SCSI"
  high_water_mark = 70
}

resource "powerstore_storagecontainer" "test1" {
  
}