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

resource "powerstore_volume" "test1" {
  name = "test_vol1"
  size = 7516192768
  description = "Creating volume"
  host_id=""
  host_group_id=""
  appliance_id="A1"
  volume_group_id=""
  min_size=1048576
  sector_size=512
  protection_policy_id=""
  performance_policy_id="default_medium"
  app_type=""
  app_type_other=""
}

resource "powerstore_volume" "test1" {
}