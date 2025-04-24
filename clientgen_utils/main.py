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

import argparse
import json

from requiredApis import RequiredAPIs
from commonUtils import ProcessOpenapiSpec
from powerStoreUtils import AddPowerStoreOpIds, AddPowerStoreFlexibleQuery

parser = argparse.ArgumentParser(description='Process PowerStore OpenAPI spec.')
parser.add_argument('--input', help='Input PowerStore OpenAPI spec file path.', required=True)
parser.add_argument('--output', help='Output filtered PowerStore OpenAPI spec file path.', required=True)

args = parser.parse_args()

# common processing of OpenAPI spec
filtered_json = ProcessOpenapiSpec(args.input, RequiredAPIs)

# powerstore specific processing
filtered_json = AddPowerStoreOpIds(filtered_json)
filtered_json = AddPowerStoreFlexibleQuery(filtered_json)

# write to file
with open(args.output, 'w') as outfile:
    json.dump(filtered_json, outfile, indent="\t")

