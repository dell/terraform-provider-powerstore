terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source  = "registry.terraform.io/dell/powerstore"
    }
  }
}

variable "username" {
  type = string
}

variable "password" {
  type = string
}

variable "endpoint" {
  type = string
}

provider "powerstore" {
  username = var.username
  password = var.password
  endpoint = var.endpoint
  insecure = true
}


resource "powerstore_snapshotrule" "test" {
  name = "test_snapshotrule_1"
  # interval = "Four_Hours"
  time_of_day       = "21:00"
  timezone          = "UTC"
  days_of_week      = ["Monday"]
  desired_retention = 56
  nas_access_type   = "Snapshot"
  is_read_only      = false
}

resource "powerstore_volume" "test" {
  name          = "test_vol"
  size          = 3
  capacity_unit = "GB"
}

resource "powerstore_volume" "test1" {
  name                  = "test_vol1"
  size                  = 3
  capacity_unit         = "GB"
  description           = "Creating volume"
  host_id               = "022c3fbc-4e92-48b6-928b-18565c803d0e"
  appliance_id          = "A1"
  volume_group_id       = "069b594c-6f68-4485-ab56-1c10b6230d71"
  min_size              = 1048576
  sector_size           = 512
  protection_policy_id  = "ea88-9c6e-a549-4281-b346-762451758e43"
  performance_policy_id = "default_medium"
  app_type              = "Relational_Databases_Other"
  app_type_other        = "db1"
}

resource "powerstore_volume" "test2" {
  name                  = "test_vol2"
  size                  = 3
  capacity_unit         = "GB"
  description           = "Creating volume"
  host_group_id         = "80c4c618-cf91-4b67-9df3-b2c0f0d6564c"
  appliance_id          = "A1"
  volume_group_id       = "069b594c-6f68-4485-ab56-1c10b6230d71"
  min_size              = 1048576
  sector_size           = 512
  protection_policy_id  = "ea889c6e-a549-4281-b346-762451758e43"
  performance_policy_id = "default_medium"
  app_type              = "Relational_Databases_Other"
  app_type_other        = "db1"
}

resource "powerstore_volume" "test3" {
  name                   = "test_vol3"
  size                   = 3
  capacity_unit          = "GB"
  description            = "Creating volume"
  host_group_name        = "tf_hostgroup1"
  appliance_name         = "server"
  volume_group_name      = "tf_volumegroup"
  min_size               = 1048576
  sector_size            = 512
  protection_policy_name = "snap_policy"
  performance_policy_id  = "default_medium"
  app_type               = "Relational_Databases_Other"
  app_type_other         = "db1"
}

resource "powerstore_storagecontainer" "test" {
  name             = "scterraform"
  quota            = 10737418240
  storage_protocol = "SCSI"
  high_water_mark  = 70
}

resource "powerstore_protectionpolicy" "terraform-provider-test" {
  # (resource arguments)
  description            = "Creating Protection Policy"
  name                   = "test_protection_policy"
  snapshot_rule_names    = ["vsi_aut_snaprule", "snapshot_test_emi", "test_snapshotrule_1", "snap-use-for-nfs-test"]
  replication_rule_names = ["Emalee-SRA-7416-Rep"]
}

data "powerstore_volume" "test1" {
  name = "tf_vol"
}

data "powerstore_snapshotrule" "test1" {
  name = "tf_snapshotRule"
}

resource "powerstore_host" "test" {
  name       = "new-host1"
  os_type    = "Linux"
  initiators = [{ port_name = "iqn.1994-05.com.redhat:88cb605", port_type = "NVMe" }]
}

data "powerstore_host" "test1" {
  name = "tf_host"
}

resource "powerstore_volumegroup_snapshot" "test" {
  name                 = "test_snap"
  volume_group_id      = "410943f5-0033-45e5-b900-6e48305ea007"
  expiration_timestamp = "2023-05-06T09:01:47Z"
}

resource "powerstore_volume_snapshot" "test" {
  name                  = "test_snap"
  volume_id             = "01d88dea-7d71-4a1b-abd6-be07f94aecd9"
  performance_policy_id = "default_medium"
  expiration_timestamp  = "2023-05-06T09:01:47Z"
  creator_type          = "User"
}
