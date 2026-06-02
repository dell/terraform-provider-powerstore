/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

# Commands to run this tf file : terraform init && terraform plan && terraform apply
# This resource performs actions on an existing replication session.
# Supported actions: sync, pause, resume, failover, reprotect, start_failover_test, stop_failover_test
# Destroying this resource only removes it from Terraform state; it does not undo the action.

# Example: Planned failover
resource "powerstore_replication_session_action" "failover" {
  session_id = "replication-session-id-here"
  action     = "failover"
  is_planned = true
  reverse    = false
}
