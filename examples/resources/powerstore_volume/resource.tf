resource "powerstore_volume" "test" {
  name = "test_vol"
  size = 1
  capacity_unit = "TB"
  description = "Creating volume"
  host_id=""
  host_group_name = "tf_hostgroup1"
  appliance_id="A1"
  volume_group_id=""
  min_size=1048576
  sector_size=512
  protection_policy_id=""
  performance_policy_id="default_medium"
  app_type="Relational_Databases_Other"
  app_type_other=""
}
