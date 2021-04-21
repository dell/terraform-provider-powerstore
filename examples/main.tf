terraform {
  required_providers {
    powerstore = {
      versions = ["0.0.1"]
      source = "dell.com/incubation/powerstore"
    }
  }
}

provider "powerstore" {
  username = "UserName"
  password = "Password"
  host = "https://10.10.10.10"
}


resource "powerstore_volume" "test" {
  # (resource arguments)
      name = "test_vol"
      size = 9663676416
      appliance_id = "A1"
      creation_timestamp = "2020-11-17T01:29:56.356792+00:00"
      is_replication_destination = false
      node_affinity = "System_Selected_Node_A"
      node_affinity_l10n = "System Selected Node A"
      nsid = 0
      state  = "Ready"
      state_l10n = "Ready"
      type = "Primary"
      type_l10n   = "Primary"
      wwn = "naa.68ccf09800f9ecd28abfa5b729f25466"
      is_app_consistent = false
      creator_type = "User"
      creator_type_l10n = "User"
      family_id = "97036821-b960-452b-9655-dbcb2761c818"

}

resource "powerstore_snapshot_rule" "snapshot1" {
  # (resource arguments)
      name = "test_snapshot_rule"
      interval= "Three_Hours"
      time_of_day= null
      days_of_week= ["Sunday","Monday","Tuesday","Wednesday", "Thursday", "Friday"]
      desired_retention= 168
}
resource "powerstore_protection_policy" "terraform-provider-test" {
  # (resource arguments)
  description = "update v5"
  name = "test_protection_policy"

  snapshot_rules {
    id = "a50a222f-3ac6-4168-b3d7-6c14b5874c61"
  }
}

  resource "powerstore_storage_container" "test-sc" {
  # (resource arguments)
      name = "Test-SC_new"
      quota = 109951162777
}