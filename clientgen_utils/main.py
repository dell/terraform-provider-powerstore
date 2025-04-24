import argparse
import json

from requiredApis import RequiredAPIs
from commonUtils import get_openapi
from powerStoreUtils import AddPowerStoreOpIds, AddPowerStoreFlexibleQuery

parser = argparse.ArgumentParser(description='Process PowerStore OpenAPI spec.')
parser.add_argument('--input', help='Input PowerStore OpenAPI spec file path.', required=True)
parser.add_argument('--output', help='Output filtered PowerStore OpenAPI spec file path.', required=True)

args = parser.parse_args()

# common processing of OpenAPI spec
filtered_json = get_openapi(args.input, RequiredAPIs)

# powerstore specific processing
filtered_json = AddPowerStoreOpIds(filtered_json)
filtered_json = AddPowerStoreFlexibleQuery(filtered_json)

# write to file
with open(args.output, 'w') as outfile:
    json.dump(filtered_json, outfile, indent="\t")

