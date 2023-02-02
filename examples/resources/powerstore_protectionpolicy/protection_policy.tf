terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source = "registry.terraform.io/dell/powerstore"
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
  snapshot_rule_names = ["snapshot_test_emi"]
  replication_rule_ids = ["5d45b173-9a85-473e-8ab8-e107f8b8085e"]
}

//Below example is for import operation
/*resource "powerstore_protectionpolicy" "terraform-provider-test-import"{
}*/
