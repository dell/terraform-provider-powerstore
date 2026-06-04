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
    '/volume',
    '/volume/{id}',
    '/volume/{id}/snapshot',
    '/volume/{id}/configure_metro',
    '/volume/{id}/end_metro',
    '/volume_group',
    '/volume_group/{id}',
    '/volume_group/{id}/add_members',
    '/volume_group/{id}/remove_members',
    '/volume_group/{id}/snapshot',
    '/volume_group/{id}/configure_metro',
    '/volume_group/{id}/end_metro',
    '/cluster',
    '/file_system',
    '/file_system/{id}',
    '/file_system/{id}/snapshot',
    '/snapshot_rule',
    '/snapshot_rule/{id}',
    '/login_session',
    '/recycle_bin',
    '/recycle_bin/{id}',
    '/recycle_bin/{id}/recover',
    '/recycle_bin/empty',
    '/recycle_bin_config/{id}',
    '/replication_session',
    '/replication_session/{id}',
    '/replication_session/{id}/failover',
    '/replication_session/{id}/reprotect',
    '/replication_session/{id}/pause',
    '/replication_session/{id}/resume',
    '/replication_session/{id}/sync',
    '/replication_session/{id}/start_failover_test',
    '/replication_session/{id}/stop_failover_test',
    '/remote_system',
    '/remote_system/{id}',
    '/security_config/generate_temp_credentials',
    '/x509_certificate/exchange',
]