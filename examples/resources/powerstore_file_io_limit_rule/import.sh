#!/bin/bash
# Copyright (c) 2024-2025 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Script to import an existing file IO limit rule into Terraform
# Usage: ./import.sh <file_io_limit_rule_id>

if [ -z "$1" ]; then
  echo "Usage: $0 <file_io_limit_rule_id>"
  echo "Example: $0 643d8f3c-7b8c-4d1e-9f2a-1b3c4d5e6f7g"
  exit 1
fi

terraform import powerstore_file_io_limit_rule.terraform-provider-test1 "$1"
