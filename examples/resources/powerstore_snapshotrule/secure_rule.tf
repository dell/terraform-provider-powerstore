/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

# Example: Create a secure snapshot rule
# Secure snapshot rules create secure snapshots that cannot be deleted
# before their expiration time.
# Was added in PowerStore API version 3.5.0.0.

resource "powerstore_snapshotrule" "secure_rule" {
  name              = "secure_hourly_rule"
  interval          = "One_Hour"
  desired_retention = 168
  is_secure         = true
}
