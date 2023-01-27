terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source = "dell/powerstore"
    }
  }
}

provider "powerstore" {
  username = "${var.username}"
  password = "${var.password}"
  endpoint = "${var.endpoint}"
  insecure = true
}

resource "powerstore_protectionpolicy" "terraform-provider-test1" {
  # (resource arguments)
  description = "Creating Protection Policy"
  name = "test_protection_policy1"
  snapshot_rule_names = ["vsi_aut_snaprule","snapshot_test_emi","test_snapshotrule_1","snap-use-for-nfs-test"]
  replication_rule_names = ["Emalee-SRA-7416-Rep"]
}

resource "powerstore_protectionpolicy" "terraform-provider-test-import"{

}