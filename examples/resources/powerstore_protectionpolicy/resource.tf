resource "powerstore_protectionpolicy" "terraform-provider-test" {
  # (resource arguments)
  description = "Creating Protection Policy"
  name = "test_protection_policy"
  snapshot_rule_names = ["vsi_aut_snaprule","snapshot_test_emi","test_snapshotrule_1","snap-use-for-nfs-test"]
  replication_rule_names = ["Emalee-SRA-7416-Rep"]
}
