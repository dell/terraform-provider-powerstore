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

# Example: Create a secure volume group snapshot
# Secure snapshots cannot be deleted before their expiration time,
# and the expiration time can only be extended.
# Was added in PowerStore API version 3.5.0.0.

resource "powerstore_volumegroup_snapshot" "secure_snap" {
  name                 = "secure_vg_snap"
  description          = "Secure volume group snapshot for data protection"
  volume_group_id      = "01d88dea-7d71-4a1b-abd6-be07f94aecd9"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
