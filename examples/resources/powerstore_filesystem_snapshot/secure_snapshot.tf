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

# Example: Create a secure filesystem snapshot
# Secure snapshots cannot be deleted before their expiration time,
# and the expiration time can only be extended.
# Was added in PowerStore API version 4.1.0.0.

resource "powerstore_filesystem_snapshot" "secure_snap" {
  name                 = "secure_fs_snap"
  description          = "Secure filesystem snapshot for data protection"
  filesystem_id        = "61d68815-1ac2-fc68-7263-96e2d715e865"
  is_secure            = true
  expiration_timestamp = "2035-05-06T09:01:47Z"
}
