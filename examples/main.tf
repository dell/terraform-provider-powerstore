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
  # interval = "Four_Hours"
  time_of_day = "21:00"
  timezone = "UTC"
  days_of_week = ["Monday"]
  desired_retention = 56
  nas_access_type = "Snapshot"
  is_read_only = false
}

resource "powerstore_volume" "test" {
  name = "test_vol"
  size =  3
  capacity_unit= "GB"
}

resource "powerstore_volume" "test1" {
  name = "test_vol1"
  size =  3
  capacity_unit= "GB"
  description = "Creating volume"
  host_id="022c3fbc-4e92-48b6-928b-18565c803d0e"
  appliance_id="A1"
  volume_group_id="069b594c-6f68-4485-ab56-1c10b6230d71"
  min_size=1048576
  sector_size=512
  protection_policy_id="ea88-9c6e-a549-4281-b346-762451758e43"
  performance_policy_id="default_medium"
  app_type="Relational_Databases_Other"
  app_type_other="db1"
}

resource "powerstore_volume" "test2" {
  name = "test_vol2"
  size =  3
  capacity_unit= "GB"
  description = "Creating volume"
  host_group_id="80c4c618-cf91-4b67-9df3-b2c0f0d6564c"
  appliance_id="A1"
  volume_group_id="069b594c-6f68-4485-ab56-1c10b6230d71"
  min_size=1048576
  sector_size=512
  protection_policy_id="ea889c6e-a549-4281-b346-762451758e43"
  performance_policy_id="default_medium"
  app_type="Relational_Databases_Other"
  app_type_other="db1"
}

resource "powerstore_storagecontainer" "test" {
  name = "scterraform"
  quota = 10737418240
  storage_protocol = "SCSI"
  high_water_mark = 70
}

resource "powerstore_protectionpolicy" "terraform-provider-test" {
  # (resource arguments)
  description = "Creating Protection Policy"
  name = "test_protection_policy"
  replication_rule_ids = ["5d45b173-9a85-473e-8ab8-e107f8b8085e"]
  snapshot_rule_ids = ["4be81573-c0e6-4956-a32f-a0e396a9b86d","cafb7709-8368-4363-bb1e-a60765b05e1e","2023295c-7098-4794-87ff-413d9c8386a0","153df6eb-3433-4b5e-942e-ecf90348df20"]
}