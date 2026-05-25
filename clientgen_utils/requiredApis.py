# Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://mozilla.org/MPL/2.0/


# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

RequiredAPIs = [
    '/volume_group',
    "/volume_group/{id}",
    "/volume_group/{id}/add_members",
    "/volume_group/{id}/remove_members",
    "/io_limit_rule",
    "/io_limit_rule/{id}",
    "/file_io_limit_rule",
    "/file_io_limit_rule/{id}",
    "/policy",
    "/policy/{id}",
    "/login_session",
    "/recycle_bin",
    "/recycle_bin/{id}",
    "/recycle_bin/{id}/recover",
    "/recycle_bin/empty",
    "/recycle_bin_config/{id}",
    "/file_system",
    "/file_system/{id}",
]