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

def _add_op_id(path: str, opType: str) -> str:
    """
    This function generates a unique operationId based on the provided path and operation type.

    Args:
        path (str): The path of the operation.
        opType (str): The type of the operation (e.g., get, post, put, delete).

    Returns:
        str: A unique operationId.
    """
    components = path.split('/')
    if len(components) == 3:
        # these have format /<resource>/{id}
        # these are operation on a specific resource
        op = opType + '_' + components[1] + '_by_id'
    elif len(components) == 4:
        # these have format /<resource>/{id}/<action>
        # these are actions on specific resources. Usually will be either a post of patch call only.
        op = components[1] + '_' + components[3]
    else :
        # these have format /<resource>
        # these are operation on the resource collection
        op = opType + '_all_' + components[1] + 's'
    return op


def AddPowerStoreOpIds(json_obj: dict) -> dict:
    """
    Adds operation IDs to every API in a PowerStore OpenAPI spec JSON object.

    Args:
        json_obj (dict): The JSON OpenAPI spec of PowerStore.

    Returns:
        dict: The modified JSON OpenAPI spec.

    """
    for key in json_obj['paths']:
        for op in json_obj['paths'][key]:
            if op not in ['get', 'put', 'post', 'delete', 'options', 'head', 'patch']:
                continue
            json_obj['paths'][key][op]['operationId'] = _add_op_id(key, op)
    return json_obj

def AddPowerStoreFlexibleQuery(json_obj: dict) -> dict:
    """
    Adds 'x-flexible-query' = "true" to GET APIs of PowerStore OpenAPI spec.

    Args:
        json_obj (dict): The JSON OpenAPI spec of PowerStore.

    Returns:
        dict: The modified JSON OpenAPI spec with 'x-flexible-query' = "true" added to GET APIs.
    """
    for key in json_obj['paths']:
        if "get" in json_obj['paths'][key]:
            json_obj['paths'][key]['get']['x-flexible-query'] = "true"
    return json_obj
