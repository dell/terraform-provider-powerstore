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

resource "powerstore_snapshotrule" "test1" {
  name = "test_snapshotrule_2"
  # interval = "Four_Hours"
  time_of_day = "21:00"
  timezone = "UTC"
  days_of_week = ["Monday"]
  desired_retention = 56
  nas_access_type = "Snapshot"
  is_read_only = false
  delete_snaps = true
}

resource "powerstore_snapshotrule" "test1" {
}