resource "powerstore_protectionpolicy" "terraform-provider-test" {
  # (resource arguments)
  description = "Creating Protection Policy"
  name = "test_protection_policy"
  snapshot_rule_ids = ["153df6eb-3433-4b5e-942e-ecf90348df20"]
  replication_rule_names = ["vsi_aut_rep_rule"]
}
