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

import json

def _get_refs(json_obj: dict) -> set:
    """
    Gets all models referenced in an openapi spec section.
    The section could be a model or a path.

    Args:
        json_obj (dict): The JSON object.
    
    Returns:
        set: The set of referenced models
    """
    result = set()
    if isinstance(json_obj, dict):
        for key, value in json_obj.items():
            if key == "$ref":
                result.add(value.split('/')[-1])
            else:
                result = result | _get_refs(value)
    elif isinstance(json_obj, list):
        for item in json_obj:
            result = result | _get_refs(item)
    return result

def _filter_by_paths(json_obj, paths):
    ret = {}
    for key in json_obj:
        if key in paths:
            ret[key] = json_obj[key]
    return ret

def _getnextqueue(json_obj_defs, refs):
    next_queue = set()
    for ref in refs:
        next_queue = next_queue | _get_refs(json_obj_defs[ref])
    return next_queue

def _getAllRefsRec(all_models, refs):
    print("The initial queue is: ", refs)
    it = 0
    ret = set()
    while(len(refs)>0):
        print("iteration: ", str(it), " =============== queue: ", refs)
        ret = ret | refs
        nextlevel = _getnextqueue(all_models, refs)
        refs = nextlevel - ret
        it+=1
    return ret

def _get_all_required_models(all_models: dict, top_level_models: set) -> dict:
    """
    This function filters the JSON object definitions to only include the top-level references and their dependencies.

    Args:
        all_models (dict): All model definitions.
        top_level_refs (set): The top-level models.

    Returns:
        dict: The filtered JSON object definitions.
    """
    
    # Get all references recursively
    all_refs = _getAllRefsRec(all_models, top_level_models)
    print("All refs are: ", all_refs)

    # Find the redundant keys by taking the difference between all keys and the references
    redundant_keys = set(all_models.keys()) - all_refs
    
    # Delete the redundant keys from the JSON object definitions
    for key in redundant_keys:
        del all_models[key]
    
    return all_models

def _find_back_edges(definitions: dict, priority_roots: set) -> set:
    """
    Finds back-edges in the definition reference graph using DFS.
    A back-edge is an edge (source, target) where target is an ancestor
    of source in the DFS tree, indicating a cycle.

    priority_roots are DFS'd first so that their direct references become
    tree edges (preserved), and back-edges are found at the deep end of
    cycles instead.

    Returns:
        set of (source_def, target_def) tuples representing back-edges.
    """
    WHITE, GRAY, BLACK = 0, 1, 2
    color = {name: WHITE for name in definitions}
    back_edges = set()

    def dfs(node):
        color[node] = GRAY
        for neighbor in sorted(_get_refs(definitions[node])):
            if neighbor not in color:
                continue
            if color[neighbor] == GRAY:
                back_edges.add((node, neighbor))
            elif color[neighbor] == WHITE:
                dfs(neighbor)
        color[node] = BLACK

    # DFS from priority roots first to protect their direct references
    for name in sorted(priority_roots):
        if name in color and color[name] == WHITE:
            dfs(name)
    for name in definitions:
        if color[name] == WHITE:
            dfs(name)
    return back_edges

def _replace_refs(json_obj, target_ref: str):
    """
    Recursively replaces {"$ref": "#/definitions/<target_ref>"} with {"type": "object"},
    preserving "description" if present. Vendor extensions (x-ref, x-added, etc.) are
    dropped — they expand the serialized schema and cause the Java generator to OOM.
    Returns the modified object.
    """
    ref_value = "#/definitions/" + target_ref
    if isinstance(json_obj, dict):
        if json_obj.get("$ref") == ref_value:
            result = {"type": "object"}
            if "description" in json_obj:
                result["description"] = json_obj["description"]
            return result
        return {k: _replace_refs(v, target_ref) for k, v in json_obj.items()}
    elif isinstance(json_obj, list):
        return [_replace_refs(item, target_ref) for item in json_obj]
    return json_obj

def _break_circular_refs(definitions: dict, dfs_roots: set, protected_defs: set) -> dict:
    """
    Detects circular references in definitions and breaks them by replacing
    cycle-causing $ref entries with {"type": "object"}.
    This prevents the OpenAPI generator's YAML serializer and ExampleGenerator
    from exploding on circular object graphs.

    dfs_roots: all path-referenced defs (GET + non-GET); DFS'd first so their
               direct references become tree edges (preserved).
    protected_defs: definitions never modified — both GET response models and
                    non-GET request bodies. Only deeply-nested non-path-referenced
                    definitions have their circular back-refs broken.
    """
    back_edges = _find_back_edges(definitions, dfs_roots)
    safe_edges = {(s, t) for s, t in back_edges if s not in protected_defs}
    if safe_edges:
        print("Breaking", len(safe_edges), "circular references in deeply-nested definitions (skipped", len(back_edges) - len(safe_edges), "in path-referenced models)")
        for src, tgt in sorted(safe_edges):
            print("  ", src, "->", tgt)
    for source, target in safe_edges:
        definitions[source] = _replace_refs(definitions[source], target)
    return definitions

def _get_get_refs(paths_obj: dict) -> set:
    """
    Gets all model references from GET method definitions only.
    These are the response models we actually read and use.
    """
    result = set()
    for path, methods in paths_obj.items():
        if 'get' in methods:
            result |= _get_refs(methods['get'])
    return result

def ProcessOpenapiSpec(file_path, paths):
    with open(file_path, 'r') as file:
        json_obj = json.load(file)
    json_obj['paths'] = _filter_by_paths(json_obj['paths'], paths)
    
    top_level_refs = _get_refs(json_obj['paths'])
    print("The top level models are: ", top_level_refs)

    _get_all_required_models(json_obj['definitions'], top_level_refs)
    print("The number of models is: ", len(json_obj['definitions'].keys()))

    get_refs = _get_get_refs(json_obj['paths'])
    non_get_refs = top_level_refs - get_refs
    print("GET response models:", get_refs)
    print("Non-GET request models:", non_get_refs)
    _break_circular_refs(json_obj['definitions'], top_level_refs, top_level_refs)

    return json_obj
