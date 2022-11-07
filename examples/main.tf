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


resource "powerstore_snapshotrule" "test" {
  name = "test_snapshotrule_1"
  interval = "Four_Hours"
  time_of_day = "21:00"
  timezone = "UTC"
  days_of_week = ["Monday"]
  desired_retention = 8
  is_read_only = true
  nas_access_type = "Snapshot"
}

resource "powerstore_volume" "test" {
  name = "test_vol"
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
